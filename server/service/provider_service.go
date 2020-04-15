package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/go-redis/redis"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtype"
	"github.com/imantung/mario"
	log "github.com/sirupsen/logrus"

	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"go.uber.org/dig"
)

// ProviderService contain logic for ProviderController [mock]
type ProviderService interface {
	MatchRule(context.Context, MatchRuleRequest) (*MatchRuleResponse, error)
	FetchTags(ctx context.Context, id int64, locale string) ([]*ITag, error)
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

// ITag is tag after interpolate with data
type ITag repository.Tag

// IDataSource is datasource after interpolate with data
type IDataSource repository.DataSource

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

// FetchTags handle logic for fetching tag
func (p *ProviderServiceImpl) FetchTags(ctx context.Context, ruleID int64, locale string) (itags []*ITag, err error) {
	var (
		rule *repository.Rule
	)

	if rule, err = p.RuleRepo.FindOne(ctx, ruleID); err != nil {
		return
	}

	return p.fetchTags(ctx, rule, locale)
}

func (p *ProviderServiceImpl) fetchTags(
	ctx context.Context,
	rule *repository.Rule,
	locale string,
) (itags []*ITag, err error) {

	var (
		tags  []*repository.Tag
		ds    *IDataSource
		b     []byte
		param map[string]interface{}
		itag  *ITag
	)

	if tags, err = p.TagRepo.FindByRuleAndLocale(ctx, rule.ID, locale); err != nil {
		return
	}

	if rule.DataSourceID != nil {
		if ds, err = p.findAndInterpolateDataSource(ctx, *rule.DataSourceID, map[string]interface{}{
			"id": rule.ID,
		}); err != nil {
			return nil, err
		}

		if b, err = call(ds); err != nil {
			return nil, fmt.Errorf("Call: %w", err)
		}

		if err = json.Unmarshal(b, &param); err != nil {
			return nil, fmt.Errorf("JSON: %w", err)
		}

		for _, tag := range tags {
			if itag, err = interpolateTag(tag, param); err != nil {
				return nil, fmt.Errorf("Interpolate-Tag: %w", err)
			}
			itags = append(itags, itag)
		}

	} else {
		for _, tag := range tags {
			itag := ITag(*tag)
			itags = append(itags, &itag)
		}
	}

	return
}

func (p *ProviderServiceImpl) findAndInterpolateDataSource(ctx context.Context, dataSourceID int64, param interface{}) (interpolated *IDataSource, err error) {
	var (
		ds *repository.DataSource
	)

	if ds, err = p.DataSourceRepo.FindOne(ctx, dataSourceID); err != nil {
		return nil, fmt.Errorf("DataSource: %w", err)
	}

	if interpolated, err = interpolateDataSource(ds, param); err != nil {
		return nil, fmt.Errorf("Interpolate-DataSource: %w", err)
	}

	return
}

func interpolateDataSource(ds *repository.DataSource, data interface{}) (*IDataSource, error) {
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

	return &IDataSource{
		ID:        ds.ID,
		Name:      ds.Name,
		Url:       buf.String(),
		UpdatedAt: ds.UpdatedAt,
		CreatedAt: ds.CreatedAt,
	}, nil
}

func interpolateTag(tag *repository.Tag, data interface{}) (*ITag, error) {
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
	return &ITag{
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

func interpolateAttribute(ori dbtype.JSON, param interface{}) (interpolated dbtype.JSON, err error) {
	var (
		tmpl *mario.Template
		buf  bytes.Buffer
	)
	if tmpl, err = mario.New().Parse(string(ori)); err != nil {
		return
	}
	if err = tmpl.Execute(&buf, param); err != nil {
		return
	}
	return buf.Bytes(), nil
}

func interpolateValue(ori string, param interface{}) (s string, err error) {
	var (
		tmpl *mario.Template
		buf  bytes.Buffer
	)
	if tmpl, err = mario.New().Parse(ori); err != nil {
		return
	}
	if err = tmpl.Execute(&buf, param); err != nil {
		return
	}
	return buf.String(), nil
}

func call(ds *IDataSource) (data []byte, err error) {
	var (
		resp *http.Response
	)

	if resp, err = http.Get(ds.Url); err != nil {
		return data, err
	}
	defer resp.Body.Close()

	if data, err = ioutil.ReadAll(resp.Body); err != nil {
		return data, err
	}

	return data, nil
}
