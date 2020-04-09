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
	s.POST("p/match", p.MatchRule)
	s.POST("p/retrieve-data", p.RetrieveData)
	s.POST("p/tags", p.Tags)

	// TODO: should hide in production or require some secret
	s.GET("p/dump-rule-tree", p.DumpRuleTree)
}
