package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"text/template"
	"time"

	"github.com/go-redis/redis"
	"github.com/imantung/mario"
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-rest-server/pkg/dbtype"

	"github.com/hotstone-seo/hotstone-seo/server/urlstore"

	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"go.uber.org/dig"
)

var (
	DataCacheExpire time.Duration = 60 * time.Second
)

// ProviderService contain logic for ProviderController [mock]
type ProviderService interface {
	MatchRule(context.Context, MatchRuleRequest) (*MatchRuleResponse, error)
	RetrieveData(context.Context, RetrieveDataRequest, bool) ([]byte, error)
	Tags(context.Context, ProvideTagsRequest, bool) ([]*InterpolatedTag, error)
	DumpRuleTree(context.Context) (string, error)
}

// ProviderServiceImpl is implementation of ProviderService
type ProviderServiceImpl struct {
	dig.In
	MetricsRuleMatchingService
	repository.DataSourceRepo
	repository.RuleRepo
	repository.TagRepo
	urlstore.URLService
	Redis *redis.Client
}

// InterpolatedTag is tag after interpolated with data
type InterpolatedTag repository.Tag

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
func (p *ProviderServiceImpl) RetrieveData(ctx context.Context, req RetrieveDataRequest, useCache bool) (data []byte, err error) {
	var (
		dataSource *repository.DataSource
		tmpl       *template.Template
		buf        bytes.Buffer
	)
	if dataSource, err = p.DataSourceRepo.FindOne(ctx, req.DataSourceID); err != nil {
		return
	}
	if tmpl, err = template.New("tmpl").Parse(dataSource.Url); err != nil {
		return
	}
	if err = tmpl.Execute(&buf, req.PathParam); err != nil {
		return
	}

	log.Warnf("DS_buf: %s", buf.String())

	if useCache {
		data, err = p.Redis.Get(buf.String()).Bytes()
		if err == redis.Nil {
			// data not exist in cache
			return p.getDataThenSetCache(buf.String())
		} else if err != nil {
			// err when getting data in cache
			return
		} else {
			// data exist in cache
			return data, err
		}
	} else {
		return p.getDataThenSetCache(buf.String())
	}
}

func (p *ProviderServiceImpl) getData(dsURL string) (data []byte, err error) {
	resp, err := http.Get(dsURL)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}

	return data, err
}

func (p *ProviderServiceImpl) getDataThenSetCache(url string) (data []byte, err error) {
	if data, err = p.getData(url); err != nil {
		return
	}
	if err = p.Redis.Set(url, data, DataCacheExpire).Err(); err != nil {
		return
	}
	return
}

// Tags to return interpolated tag
func (p *ProviderServiceImpl) Tags(ctx context.Context, req ProvideTagsRequest, useCache bool) (interpolatedTags []*InterpolatedTag, err error) {
	var (
		tags       []*repository.Tag
		data       = req.Data
		dataFromDS []byte
	)
	if tags, err = p.TagRepo.Find(ctx, repository.TagFilter{RuleID: req.RuleID, Locale: req.Locale}); err != nil {
		return
	}
	if data == nil {
		// NOTE: We can omit another call to the repository here by including the whole rule object in
		// ProvideTagsRequest. Will be done later since the change impact is not local.
		var (
			rule *repository.Rule
		)
		if rule, err = p.RuleRepo.FindOne(ctx, req.RuleID); err != nil {
			return
		}
		if rule.DataSourceID != nil {
			if dataFromDS, err = p.RetrieveData(
				ctx,
				RetrieveDataRequest{
					DataSourceID: *rule.DataSourceID,
					PathParam:    req.PathParam,
				}, useCache,
			); err != nil {
				return
			}
			if err = json.Unmarshal(dataFromDS, &data); err != nil {
				return
			}
		}
	}
	interpolatedTags = make([]*InterpolatedTag, 0)
	for _, tag := range tags {
		var (
			attribute dbtype.JSON
			value     string
		)
		if attribute, err = interpolateAttribute(tag.Attributes, data); err != nil {
			return
		}
		if value, err = interpolateValue(tag.Value, data); err != nil {
			return
		}
		interpolatedTags = append(interpolatedTags, &InterpolatedTag{
			ID:         tag.ID,
			RuleID:     tag.RuleID,
			Locale:     tag.Locale,
			Type:       tag.Type,
			Attributes: attribute,
			Value:      value,
			UpdatedAt:  tag.UpdatedAt,
			CreatedAt:  tag.CreatedAt,
		})
	}
	return
}

// DumpRuleTree to dump the rule tree
func (p *ProviderServiceImpl) DumpRuleTree(ctx context.Context) (dump string, err error) {
	return p.URLService.DumpTree(), nil
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
