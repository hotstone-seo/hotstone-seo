package repository

import (
	"context"
	"time"
)

type Rule struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name" validate:"required"`
	UrlPattern   string    `json:"url_pattern" validate:"required"`
	DataSourceID *int64    `validate:"required"`
	UpdatedAt    time.Time `json:"-"`
	CreatedAt    time.Time `json:"-"`
}

type RuleRepo interface {
	Find(ctx context.Context, id int64) (*Rule, error)
	List(ctx context.Context) ([]*Rule, error)
	Insert(ctx context.Context, rule Rule) (lastInsertID int64, err error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, rule Rule) error
}
