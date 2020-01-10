package repository

import (
	"context"
	"time"
)

// MetricsUnmatched represented  metrics_unmatched entity
type MetricsUnmatched struct {
	ID          int64     `json:"id"`
	RequestPath string    `json:"request_path"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MetricsUnmatchedRepo to handle metrics_unmatched entity [mock]
type MetricsUnmatchedRepo interface {
	Find(context.Context, int64) (*MetricsUnmatched, error)
	List(context.Context) ([]*MetricsUnmatched, error)
	ListCount(context.Context) ([]*MetricsMismatchedCount, error)
	Insert(context.Context, MetricsUnmatched) (lastInsertID int64, err error)
	Delete(context.Context, int64) error
	Update(context.Context, MetricsUnmatched) error
}

// NewMetricsUnmatchedRepo return new instance of MetricsUnmatchedRepo [autowire]
func NewMetricsUnmatchedRepo(impl CachedMetricsUnmatchedRepoImpl) MetricsUnmatchedRepo {
	return &impl
}
