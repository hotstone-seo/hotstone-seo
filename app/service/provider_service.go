package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/hotstone-seo/hotstone-server/app/repository"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"go.uber.org/dig"
)

// ProviderService contain logic for ProviderController
type ProviderService interface {
	MatchRule(Matcher, MatchRuleRequest) (*MatchRuleResponse, error)
	RetrieveData(context.Context, RetrieveDataRequest) (*http.Response, error)
	Tags(context.Context, ProvideTagsRequest) ([]*InterpolatedTag, error)
}

// ProviderServiceImpl is implementation of ProviderService
type ProviderServiceImpl struct {
	dig.In
	repository.DataSourceRepo
	repository.TagRepo
}

type Matcher interface {
	Match(url string) (int, map[string]string)
}

// NewProviderService return new instance of ProviderService
func NewProviderService() ProviderService {
	return &ProviderServiceImpl{}
}

// MatchRule to match rule
func (*ProviderServiceImpl) MatchRule(matcher Matcher, req MatchRuleRequest) (resp *MatchRuleResponse, err error) {
	ruleID, pathParam := matcher.Match(req.Path)
	if ruleID == -1 {
		return nil, errors.New("No rule match")
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
		interpolatedTags = append(interpolatedTags, &InterpolatedTag{
			ID:         tag.ID,
			RuleID:     tag.RuleID,
			LocaleID:   tag.LocaleID,
			Type:       tag.Type,
			Attributes: interpolateAttribute(tag.Attributes, req.Data),
			Value:      interpolateValue(tag.Value, req.Data),
			UpdatedAt:  tag.UpdatedAt,
			CreatedAt:  tag.CreatedAt,
		})
	}
	return
}

func interpolateAttribute(ori dbkit.JSON, data interface{}) dbkit.JSON {
	return ori
}

func interpolateValue(ori string, data interface{}) string {
	return ori
}

type InterpolatedTag repository.Tag
