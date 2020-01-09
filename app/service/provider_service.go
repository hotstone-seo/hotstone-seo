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

	log "github.com/sirupsen/logrus"

	"github.com/hotstone-seo/hotstone-seo/app/metric"
	"github.com/hotstone-seo/hotstone-seo/app/urlstore"

	"github.com/hotstone-seo/hotstone-seo/app/repository"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"go.uber.org/dig"
)

// ProviderService contain logic for ProviderController
type ProviderService interface {
	MatchRule(context.Context, MatchRuleRequest) (*MatchRuleResponse, error)
	RetrieveData(context.Context, RetrieveDataRequest) (*http.Response, error)
	Tags(context.Context, ProvideTagsRequest) ([]*InterpolatedTag, error)
}

// ProviderServiceImpl is implementation of ProviderService
type ProviderServiceImpl struct {
	dig.In
	MetricsUnmatchedService
	MetricsRuleMatchingService
	repository.DataSourceRepo
	repository.RuleRepo
	repository.TagRepo
	urlstore.URLStoreServer
}

// NewProviderService return new instance of ProviderService
func NewProviderService(impl ProviderServiceImpl) ProviderService {
	return &impl
}

// MatchRule to match rule
func (p *ProviderServiceImpl) MatchRule(ctx context.Context, req MatchRuleRequest) (resp *MatchRuleResponse, err error) {
	ctx = metric.InitializeLatencyTracking(ctx)
	mtx := &repository.MetricsRuleMatching{}
	defer func() {
		metric.RecordLatency(ctx)
		if errInsert := p.MetricsRuleMatchingService.Insert(ctx, *mtx); errInsert != nil {
			log.Warnf("Failed to record rule matching metric: %+v", errInsert)
		}
	}()

	url, err := url.Parse(req.Path)
	if err != nil {
		return
	}

	ruleID, pathParam := p.URLStoreServer.Match(url.Path)
	if ruleID == -1 {
		// mismatched

		if errRecord := p.MetricsUnmatchedService.Record(ctx, url.Path); errRecord != nil {
			log.Warnf("Failed to record unmatched metrics: %+v", errRecord)
		}

		ctx = metric.SetMismatched(ctx, url.Path)
		p.MetricsRuleMatchingService.SetMismatched(mtx, url.Path)

		return nil, fmt.Errorf("No rule match: %s", url.Path)
	} else {
		// matched
		ctx = metric.SetMatched(ctx)
		p.MetricsRuleMatchingService.SetMatched(mtx)
	}

	resp = &MatchRuleResponse{RuleID: int64(ruleID), PathParam: pathParam}
	return resp, nil
}

func (p *ProviderServiceImpl) RetrieveData(ctx context.Context, req RetrieveDataRequest) (resp *http.Response, err error) {
	var dataSource *repository.DataSource
	if dataSource, err = p.DataSourceRepo.FindOne(ctx, req.DataSourceID); err != nil {
		return
	}
	return http.Get(dataSource.Url)
}

func (p *ProviderServiceImpl) Tags(ctx context.Context, req ProvideTagsRequest) (interpolatedTags []*InterpolatedTag, err error) {
	var (
		tags []*repository.Tag
		data = req.Data
	)
	if tags, err = p.TagRepo.FindByRuleAndLocale(ctx, req.RuleID, req.LocaleID); err != nil {
		return
	}
	if data == nil {
		// NOTE: We can omit another call to the repository here by including the whole rule object in
		// ProvideTagsRequest. Will be done later since the change impact is not local.
		var (
			rule *repository.Rule
			resp *http.Response
			body []byte
		)
		if rule, err = p.RuleRepo.FindOne(ctx, req.RuleID); err != nil {
			return
		}
		if resp, err = p.RetrieveData(ctx, RetrieveDataRequest{DataSourceID: *rule.DataSourceID}); err != nil {
			return
		}
		if body, err = ioutil.ReadAll(resp.Body); err != nil {
			return
		}
		defer resp.Body.Close()
		if err = json.Unmarshal(body, &data); err != nil {
			return
		}
	}
	for _, tag := range tags {
		var (
			attribute dbkit.JSON
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
			LocaleID:   tag.LocaleID,
			Type:       tag.Type,
			Attributes: attribute,
			Value:      value,
			UpdatedAt:  tag.UpdatedAt,
			CreatedAt:  tag.CreatedAt,
		})
	}
	return
}

func interpolateAttribute(ori dbkit.JSON, data interface{}) (interpolated dbkit.JSON, err error) {
	var (
		tmpl *template.Template
		buf  bytes.Buffer
	)
	if tmpl, err = template.New("tmpl").Parse(string(ori)); err != nil {
		return
	}
	if err = tmpl.Execute(&buf, data); err != nil {
		return
	}
	return buf.Bytes(), nil
}

func interpolateValue(ori string, data interface{}) (s string, err error) {
	var (
		tmpl *template.Template
		buf  bytes.Buffer
	)
	if tmpl, err = template.New("tmpl").Parse(ori); err != nil {
		return
	}
	if err = tmpl.Execute(&buf, data); err != nil {
		return
	}
	return buf.String(), nil
}

type InterpolatedTag repository.Tag
