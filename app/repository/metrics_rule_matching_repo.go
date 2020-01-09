package repository

import (
	"context"
	"time"
)

// MetricsRuleMatching represented  metrics_rule_matching entity
type MetricsRuleMatching struct {
	Time          time.Time
	IsMatched     int
	URLMismatched *string
}

// MetricsRuleMatchingRepo to handle metrics_rule_matching entity
type MetricsRuleMatchingRepo interface {
	Insert(context.Context, MetricsRuleMatching) (err error)
}

// NewMetricsRuleMatchingRepo return new instance of MetricsRuleMatchingRepo
func NewMetricsRuleMatchingRepo(impl MetricsRuleMatchingRepoImpl) MetricsRuleMatchingRepo {
	return &impl
}
