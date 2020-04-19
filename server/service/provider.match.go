package service

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/hotstone-seo/hotstone-seo/server/metric"
)

const (
	insertMetricTimeout = 30 * time.Second
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
	var (
		ruleID    int64
		pathParam map[string]string
	)

	if _, err = url.Parse(req.Path); err != nil {
		return nil, fmt.Errorf("URL: %w", err)
	}

	path := req.Path

	if ruleID, pathParam = p.URLService.Match(path); ruleID == -1 {
		p.onNotMatched(path)
		return nil, fmt.Errorf("No rule match: %s", path)
	}

	// matched
	p.onMatched(path, ruleID)
	return &MatchResponse{
		RuleID:    ruleID,
		PathParam: pathParam,
	}, nil
}

func (p *ProviderServiceImpl) onMatched(url string, ruleID int64) {
	ctx, cancel := context.WithTimeout(context.Background(), insertMetricTimeout)
	defer cancel()

	p.RuleMatchingRepo.Insert(ctx, metric.RuleMatching{
		IsMatched: 1,
		URL:       &url,
		RuleID:    &ruleID,
	})
}

func (p *ProviderServiceImpl) onNotMatched(url string) {
	ctx, cancel := context.WithTimeout(context.Background(), insertMetricTimeout)
	defer cancel()

	p.RuleMatchingRepo.Insert(ctx, metric.RuleMatching{
		IsMatched: 0,
		URL:       &url,
	})
}
