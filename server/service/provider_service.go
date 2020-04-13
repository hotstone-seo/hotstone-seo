package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtype"
	"github.com/imantung/mario"
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"

	"github.com/hotstone-seo/hotstone-seo/pkg/cachekit"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"go.uber.org/dig"
)

// ProviderService contain logic for ProviderController [mock]
type ProviderService interface {
	MatchRule(context.Context, MatchRuleRequest) (*MatchRuleResponse, error)
	RetrieveData(context.Context, RetrieveDataRequest, *cachekit.Pragma) (*RetrieveDataResponse, error)
	Tags(context.Context, ProvideTagsRequest, *cachekit.Pragma) ([]*InterpolatedTag, error)
	DumpRuleTree(context.Context) (string, error)

	FetchTags(
		ctx context.Context,
		id int64,
		locale string,
	) ([]*repository.Tag, error)
}

// ProviderServiceImpl is implementation of ProviderService
type ProviderServiceImpl struct {
	dig.In
	MetricsRuleMatchingService
	repository.DataSourceRepo
	repository.RuleRepo
	repository.TagRepo
	URLService

	Redis *redis.Client
}

// InterpolatedTag is tag after interpolate with data
type InterpolatedTag repository.Tag

// InterpolatedDataSource is datasource after interpolate with data
type InterpolatedDataSource repository.DataSource

// NewProviderService return new instance of ProviderService [constructor]
func NewProviderService(impl ProviderServiceImpl) ProviderService {
	return &impl
}

// MatchRule to match rule
func (p *ProviderServiceImpl) MatchRule(ctx context.Context, req MatchRuleRequest) (resp *MatchRuleResponse, err error) {
	mtx := &repository.MetricsRuleMatching{}
	defer func() {
		if errInsert := p.MetricsRuleMatchingService.Insert(ctx, *mtx); errInsert != nil {
			log.Warnf("Failed to record rule matching metric: %+v", errInsert)
		}
	}()

	url, err := url.Parse(req.Path)
	if err != nil {
		return
	}

	ruleID, pathParam := p.URLService.Match(url.Path)
	if ruleID == -1 {
		// mismatched
		p.MetricsRuleMatchingService.SetMismatched(mtx, url.Path)

		return nil, fmt.Errorf("No rule match: %s", url.Path)
	}

	// matched
	p.MetricsRuleMatchingService.SetMatched(mtx, url.Path, int64(ruleID))
	return &MatchRuleResponse{
		RuleID:    int64(ruleID),
		PathParam: pathParam,
	}, nil
}

// RetrieveData to retrieve the data from data provider
func (p *ProviderServiceImpl) RetrieveData(ctx context.Context, req RetrieveDataRequest, pragma *cachekit.Pragma) (resp *RetrieveDataResponse, err error) {
	var (
		ds           *repository.DataSource
		interpolated *InterpolatedDataSource
	)

	if ds, err = p.DataSourceRepo.FindOne(ctx, req.DataSourceID); err != nil {
		return nil, err
	}
	if ds == nil {
		return nil, fmt.Errorf("Data-Source#%d not found", req.DataSourceID)
	}

	if interpolated, err = interpolateDataSource(ds, req.PathParam); err != nil {
		return
	}

	cache := cachekit.New(interpolated.Url, callDatasoure(interpolated))

	resp = new(RetrieveDataResponse)
	if err = cache.Execute(p.Redis.WithContext(ctx), resp, pragma); err != nil {
		return nil, err
	}

	return resp, nil
}

func callDatasoure(ds *InterpolatedDataSource) cachekit.RefreshFn {
	var (
		resp *http.Response
		data []byte
		err  error
	)

	return func() (interface{}, error) {
		if resp, err = http.Get(ds.Url); err != nil {
			return data, err
		}
		defer resp.Body.Close()

		if data, err = ioutil.ReadAll(resp.Body); err != nil {
			return data, err
		}

		return &RetrieveDataResponse{
			Data: data,
		}, nil
	}
}

// Tags to return interpolated tag
func (p *ProviderServiceImpl) Tags(ctx context.Context, req ProvideTagsRequest, pragma *cachekit.Pragma) (interpolatedTags []*InterpolatedTag, err error) {
	var (
		tags         []*repository.Tag
		data         = req.Data
		dataResp     *RetrieveDataResponse
		rule         *repository.Rule
		interpolated *InterpolatedTag
	)

	if tags, err = p.TagRepo.Find(ctx,
		dbkit.Equal("rule_id", strconv.FormatInt(req.RuleID, 10)),
		dbkit.Equal("locale", req.Locale),
	); err != nil {
		err = fmt.Errorf("Provider: Tags: Find: %s", err.Error())
		return
	}

	if data == nil {
		// NOTE: We can omit another call to the repository here by including the whole rule object in
		// ProvideTagsRequest. Will be done later since the change impact is not local.
		if rule, err = p.RuleRepo.FindOne(ctx, req.RuleID); err != nil {
			return
		}
		if rule == nil {
			err = fmt.Errorf("Rule#%d not found", req.RuleID)
			return
		}
		if rule.DataSourceID != nil {
			if dataResp, err = p.RetrieveData(ctx,
				RetrieveDataRequest{
					DataSourceID: *rule.DataSourceID,
					PathParam:    req.PathParam,
				}, pragma,
			); err != nil {
				return
			}
			if err = json.Unmarshal(dataResp.Data, &data); err != nil {
				return
			}
		}
	}

	interpolatedTags = make([]*InterpolatedTag, 0)
	for _, tag := range tags {
		if interpolated, err = interpolateTag(tag, data); err != nil {
			return
		}
		interpolatedTags = append(interpolatedTags, interpolated)
	}
	return
}

// FetchTag handle logic for fetching tag
func (p *ProviderServiceImpl) FetchTags(
	ctx context.Context,
	ruleID int64,
	locale string,
) (tags []*repository.Tag, err error) {

	if tags, err = p.TagRepo.Find(ctx,
		dbkit.Equal("rule_id", strconv.FormatInt(ruleID, 10)),
		dbkit.Equal("locale", locale),
	); err != nil {
		err = fmt.Errorf("Provider: FetchTags: Find: %s", err.Error())
		return
	}
	return
}

// DumpRuleTree to dump the rule tree
func (p *ProviderServiceImpl) DumpRuleTree(ctx context.Context) (dump string, err error) {
	return p.URLService.DumpTree(), nil
}

func interpolateDataSource(ds *repository.DataSource, data interface{}) (*InterpolatedDataSource, error) {
	var (
		buf  bytes.Buffer
		tmpl *mario.Template
		err  error
	)

	if tmpl, err = mario.New().Parse(ds.Url); err != nil {
		return nil, err
	}
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}

	return &InterpolatedDataSource{
		ID:        ds.ID,
		Name:      ds.Name,
		Url:       buf.String(),
		UpdatedAt: ds.UpdatedAt,
		CreatedAt: ds.CreatedAt,
	}, nil
}

func interpolateTag(tag *repository.Tag, data interface{}) (*InterpolatedTag, error) {
	var (
		attribute dbtype.JSON
		value     string
		err       error
	)
	if attribute, err = interpolateAttribute(tag.Attributes, data); err != nil {
		return nil, err
	}
	if value, err = interpolateValue(tag.Value, data); err != nil {
		return nil, err
	}
	return &InterpolatedTag{
		ID:         tag.ID,
		RuleID:     tag.RuleID,
		Locale:     tag.Locale,
		Type:       tag.Type,
		Attributes: attribute,
		Value:      value,
		UpdatedAt:  tag.UpdatedAt,
		CreatedAt:  tag.CreatedAt,
	}, nil

}

func interpolateAttribute(ori dbtype.JSON, data interface{}) (interpolated dbtype.JSON, err error) {
	var (
		tmpl *mario.Template
		buf  bytes.Buffer
	)
	if tmpl, err = mario.New().Parse(string(ori)); err != nil {
		return
	}
	if err = tmpl.Execute(&buf, data); err != nil {
		return
	}
	return buf.Bytes(), nil
}

func interpolateValue(ori string, data interface{}) (s string, err error) {
	var (
		tmpl *mario.Template
		buf  bytes.Buffer
	)
	if tmpl, err = mario.New().Parse(ori); err != nil {
		return
	}
	if err = tmpl.Execute(&buf, data); err != nil {
		return
	}
	return buf.String(), nil
}
