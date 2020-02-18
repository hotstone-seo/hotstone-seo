package repository

import (
	"context"
	"net/url"
	"time"
)

// MetricsRuleMatching represented  metrics_rule_matching entity
type MetricsRuleMatching struct {
	Time      time.Time
	IsMatched int
	RuleID    *int64
	URL       *string
}

// MetricsRuleMatchingRepo to handle metrics_rule_matching entity [mock]
type MetricsRuleMatchingRepo interface {
	Insert(context.Context, MetricsRuleMatching) (err error)
	ListMismatchedCount(ctx context.Context, paginationParam PaginationParam) (list []*MetricsMismatchedCount, err error)
	CountMatched(ctx context.Context, whereParams url.Values) (count int64, err error)
	CountUniquePage(ctx context.Context, whereParams url.Values) (count int64, err error)
	ListCountHitPerDay(ctx context.Context, startDate string, endDate string) (list []*MetricsCountHitPerDay, err error)
}

// NewMetricsRuleMatchingRepo return new instance of MetricsRuleMatchingRepo [constructor]
func NewMetricsRuleMatchingRepo(impl MetricsRuleMatchingRepoImpl) MetricsRuleMatchingRepo {
	return &impl
}
