package repository

import (
	"context"
	"database/sql"
	"time"
)

type Rule struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name" validate:"required"`
	UrlPattern   string    `json:"url_pattern" validate:"required"`
	DataSourceID *int64    `json:"data_source_id"`
	UpdatedAt    time.Time `json:"-"`
	CreatedAt    time.Time `json:"-"`
}

type RuleRepo interface {
	Find(ctx context.Context, id int64) (*Rule, error)
	List(ctx context.Context) ([]*Rule, error)
	Insert(ctx context.Context, rule Rule) (lastInsertID int64, err error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, rule Rule) error
	DB() *sql.DB
}

// NewRuleRepo return new instance of RuleRepo
func NewRuleRepo(impl RuleRepoImpl) RuleRepo {
	return &impl
}
