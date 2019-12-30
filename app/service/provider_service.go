package service

import (
	"encoding/json"

	"github.com/hotstone-seo/hotstone-server/app/repository"
	"go.uber.org/dig"
)

// ProviderService contain logic for ProviderController
type ProviderService interface {
	MatchRule(MatchRuleRequest) (*repository.Rule, error)
	RetrieveData(RetrieveDataRequest) (interface{}, error)
	Tags(string) ([]*repository.Tag, error)
}

// ProviderServiceImpl is implementation of ProviderService
type ProviderServiceImpl struct {
	dig.In
}

// NewProviderService return new instance of ProviderService
func NewProviderService() ProviderService {
	return &ProviderServiceImpl{}
}

// MatchRule to match rule
func (*ProviderServiceImpl) MatchRule(req MatchRuleRequest) (rule *repository.Rule, err error) {
	dataSourceID := int64(88888)
	rule = &repository.Rule{
		ID:           999999,
		Name:         "sample-rule",
		UrlPattern:   "some-url-pattern",
		DataSourceID: &dataSourceID,
	}
	return
}

func (*ProviderServiceImpl) RetrieveData(req RetrieveDataRequest) (data interface{}, err error) {
	data = struct {
		Name     string `json:"name"`
		Province string `json:"province"`
	}{
		Name:     "CGK",
		Province: "Banten",
	}
	return
}

func (*ProviderServiceImpl) Tags(ruleID string) (tags []*repository.Tag, err error) {
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
