package service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/hotstone-seo/hotstone-server/app/repository"
	"go.uber.org/dig"
)

// ProviderService contain logic for ProviderController
type ProviderService interface {
	MatchRule(Matcher, MatchRuleRequest) (*MatchRuleResponse, error)
	RetrieveData(context.Context, RetrieveDataRequest) (*http.Response, error)
	Tags(ruleID string, data interface{}) ([]*repository.Tag, error)
}

// ProviderServiceImpl is implementation of ProviderService
type ProviderServiceImpl struct {
	dig.In
	repository.DataSourceRepo
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

func (*ProviderServiceImpl) Tags(ruleID string, data interface{}) (tags []*repository.Tag, err error) {
	attr := []map[string]string{
		map[string]string{
			"type":    "description",
			"content": "Page Description",
		},
	}
	attrByte, _ := json.Marshal(attr)
	tags = []*repository.Tag{
		&repository.Tag{ID: 9999, Type: "title", Value: "Page Title"},
		&repository.Tag{ID: 8888, Type: "meta", Attributes: attrByte},
	}
	return
}
