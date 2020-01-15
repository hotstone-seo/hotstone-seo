package repository

import (
	"context"
	"time"
)

// MetricsRuleMatching represented  metrics_rule_matching entity
type MetricsRuleMatching struct {
	Time      time.Time
	IsMatched int
	URL       *string
}

// MetricsRuleMatchingRepo to handle metrics_rule_matching entity [mock]
type MetricsRuleMatchingRepo interface {
	Insert(context.Context, MetricsRuleMatching) (err error)
	ListMismatchedCount(ctx context.Context) (list []*MetricsMismatchedCount, err error)
	CountMatched(ctx context.Context) (count int64, err error)
	CountUniquePage(ctx context.Context) (count int64, err error)
	ListCountHitPerDay(ctx context.Context, startDate string, endDate string) (list []*MetricsCountHitPerDay, err error)
}

// NewMetricsRuleMatchingRepo return new instance of MetricsRuleMatchingRepo [autowire]
func NewMetricsRuleMatchingRepo(impl MetricsRuleMatchingRepoImpl) MetricsRuleMatchingRepo {
	return &impl
}
