package service

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/hotstone-seo/hotstone-seo/pkg/errkit"
	"github.com/hotstone-seo/hotstone-seo/server/metric"
)

const (
	insertMetricTimeout = 30 * time.Second
	pathParam           = "_path"
)

// MatchResponse is response of match rule
type MatchResponse struct {
	RuleID    int64             `json:"rule_id"`
	PathParam map[string]string `json:"path_param"`
}

// Match url with its rule
func (p *ProviderServiceImpl) Match(ctx context.Context, vals url.Values) (resp *MatchResponse, err error) {
	var (
		ruleID    int64
		pathParam map[string]string
	)

	path := vals.Get("_path")

	if path == "" {
		return nil, errkit.ValidationErr("_path can't empty")
	}

	if _, err = url.Parse(path); err != nil {
		return nil, fmt.Errorf("URL: %w", err)
	}

	data, param := p.Store.Get(path)

	if data == nil {
		go p.onNotMatched(path)
		return &MatchResponse{}, nil
	}

	if ruleID, err = convertToRuleID(data); err != nil {
		return
	}

	if param != nil {
		pathParam = param.Map()
	}

	// matched
	go p.onMatched(path, ruleID)
	return &MatchResponse{
		RuleID:    ruleID,
		PathParam: pathParam,
	}, nil
}

func convertToRuleID(data interface{}) (ruleID int64, err error) {
	var (
		s  string
		ok bool
	)

	if s, ok = data.(string); !ok {
		err = fmt.Errorf("Failed to cast data to string. data=%+v", data)
		return
	}

	if ruleID, err = strconv.ParseInt(s, 10, 64); err != nil {
		err = fmt.Errorf("Failed to convert string data to int. idStr=%+s", s)
		return
	}
	return
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
