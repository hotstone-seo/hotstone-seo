package repository

import (
	"context"
	"time"
)

// Rule Entity
type Rule struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name" validate:"required"`
	UrlPattern   string    `json:"url_pattern" validate:"required"`
	DataSourceID *int64    `json:"data_source_id"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
}

// RuleRepo is rule repository [mock]
type RuleRepo interface {
	FindOne(ctx context.Context, id int64) (*Rule, error)
	Find(ctx context.Context) ([]*Rule, error)
	Insert(ctx context.Context, rule Rule) (lastInsertID int64, err error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, rule Rule) error
}

// NewRuleRepo return new instance of RuleRepo [autowire]
func NewRuleRepo(impl RuleRepoImpl) RuleRepo {
	return &impl
}