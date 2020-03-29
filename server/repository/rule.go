package repository

import (
	"context"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

// Rule Entity
type Rule struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name" validate:"required"`
	UrlPattern     string    `json:"url_pattern" validate:"required,uri"`
	DataSourceID   *int64    `json:"data_source_id"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedAt      time.Time `json:"created_at"`
	Status         string    `json:"status"`
	ChangeStatusAt time.Time `json:"change_status_at"`
}

// RuleRepo is rule repository [mock]
type RuleRepo interface {
	FindOne(ctx context.Context, id int64) (*Rule, error)
	Find(ctx context.Context, paginationParam PaginationParam) ([]*Rule, error)
	Insert(ctx context.Context, rule Rule) (lastInsertID int64, err error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, rule Rule) error
}

// NewRuleRepo return new instance of RuleRepo [constructor]
func NewRuleRepo(impl RuleRepoImpl) RuleRepo {
	return &impl
}

func (rule Rule) Validate() error {
	return validator.New().Struct(rule)
}
