package server

import (
	"github.com/hotstone-seo/hotstone-seo/server/controller"
	"go.uber.org/dig"
)

type provider struct {
	dig.In
	controller.ProviderCntrl
}

func (p *provider) route(s server) {
	// TODO: remove api prefix
	s.POST("api/provider/match-rule", p.MatchRule)
	s.POST("api/provider/retrieve-data", p.RetrieveData)
	s.POST("api/provider/tags", p.Tags)
	s.GET("api/provider/rule-tree", p.DumpRuleTree)
}
