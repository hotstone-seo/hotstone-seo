package service

import (
	"context"
	"fmt"
	"net/url"

	"github.com/hotstone-seo/hotstone-seo/server/repository"
	log "github.com/sirupsen/logrus"
)

// MatchRequest is request for match rule
type MatchRequest struct {
	Path string `json:"path"`
}

// MatchResponse is response of match rule
type MatchResponse struct {
	RuleID    int64             `json:"rule_id"`
	PathParam map[string]string `json:"path_param"`
}

// Match url with its rule
func (p *ProviderServiceImpl) Match(ctx context.Context, req MatchRequest) (resp *MatchResponse, err error) {
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
	return &MatchResponse{
		RuleID:    int64(ruleID),
		PathParam: pathParam,
	}, nil
}
