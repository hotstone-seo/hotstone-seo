package service

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"text/template"

	log "github.com/sirupsen/logrus"

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
	repository.DataSourceRepo
	repository.TagRepo
	urlstore.URLStoreServer
}

// NewProviderService return new instance of ProviderService
func NewProviderService(impl ProviderServiceImpl) ProviderService {
	return &impl
}

// MatchRule to match rule
func (p *ProviderServiceImpl) MatchRule(ctx context.Context, req MatchRuleRequest) (resp *MatchRuleResponse, err error) {
	url, err := url.Parse(req.Path)
	if err != nil {
		return
	}

	ruleID, pathParam := p.URLStoreServer.Match(url.Path)
	if ruleID == -1 {
		if errRecord := p.MetricsUnmatchedService.Record(ctx, url.Path); errRecord != nil {
			log.Warnf("Failed to record unmatched metrics: %+v", errRecord)
		}
		return nil, fmt.Errorf("No rule match: %s", url.Path)
	} else {

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
	)
	if tags, err = p.TagRepo.FindByRuleAndLocale(ctx, req.RuleID, req.LocaleID); err != nil {
		return
	}
	for _, tag := range tags {
		var (
			attribute dbkit.JSON
			value     string
		)
		if attribute, err = interpolateAttribute(tag.Attributes, req.Data); err != nil {
			return
		}
		if value, err = interpolateValue(tag.Value, req.Data); err != nil {
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
