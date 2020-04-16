package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"go.uber.org/dig"
	"gopkg.in/go-playground/validator.v9"
)

// Rule Entity
type Rule struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name" validate:"required"`
	URLPattern     string    `json:"url_pattern" validate:"required,uri"`
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

// RuleRepoImpl is implementation rule repository
type RuleRepoImpl struct {
	dig.In
	*typpostgres.DB
}

// NewRuleRepo return new instance of RuleRepo [constructor]
func NewRuleRepo(impl RuleRepoImpl) RuleRepo {
	return &impl
}

// Validate rule
func (rule Rule) Validate() error {
	return validator.New().Struct(rule)
}

// FindOne rule
func (r *RuleRepoImpl) FindOne(ctx context.Context, id int64) (*Rule, error) {
	row := sq.StatementBuilder.
		Select(
			"id",
			"name",
			"url_pattern",
			"data_source_id",
			"updated_at",
			"created_at",
			"status",
			"change_status_at",
		).
		From("rules").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.BaseRunner(ctx, r)).
		QueryRowContext(ctx)

	rule := new(Rule)
	if err := row.Scan(
		&rule.ID,
		&rule.Name,
		&rule.URLPattern,
		&rule.DataSourceID,
		&rule.UpdatedAt,
		&rule.CreatedAt,
		&rule.Status,
		&rule.ChangeStatusAt,
	); err != nil {
		dbtxn.SetError(ctx, err)
		return nil, err
	}

	return rule, nil
}

// Find rule
func (r *RuleRepoImpl) Find(ctx context.Context, paginationParam PaginationParam) (list []*Rule, err error) {
	var (
		rows *sql.Rows
	)

	builder := sq.StatementBuilder.
		Select(
			"id",
			"name",
			"url_pattern",
			"data_source_id",
			"updated_at",
			"created_at",
			"status",
			"change_status_at",
		).
		From("rules").
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.BaseRunner(ctx, r))

	builder = composePagination(builder, paginationParam)

	if rows, err = builder.QueryContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	defer rows.Close()

	list = make([]*Rule, 0)

	for rows.Next() {
		rule := new(Rule)
		if err = rows.Scan(
			&rule.ID,
			&rule.Name,
			&rule.URLPattern,
			&rule.DataSourceID,
			&rule.UpdatedAt,
			&rule.CreatedAt,
			&rule.Status,
			&rule.ChangeStatusAt,
		); err != nil {
			dbtxn.SetError(ctx, err)
			return
		}
		list = append(list, rule)
	}
	return
}

// Insert rule
func (r *RuleRepoImpl) Insert(ctx context.Context, rule Rule) (lastInsertID int64, err error) {

	query := sq.Insert("rules").
		Columns(
			"data_source_id",
			"name",
			"url_pattern",
		).
		Values(
			rule.DataSourceID,
			rule.Name,
			rule.URLPattern,
		).
		Suffix("RETURNING \"id\"").
		RunWith(dbtxn.BaseRunner(ctx, r)).
		PlaceholderFormat(sq.Dollar).
		QueryRowContext(ctx)

	if err = query.Scan(&rule.ID); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}

	lastInsertID = rule.ID
	return
}

// Delete rule
func (r *RuleRepoImpl) Delete(ctx context.Context, id int64) (err error) {
	builder := sq.StatementBuilder.
		Delete("rules").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.BaseRunner(ctx, r))

	if _, err = builder.ExecContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	return
}

// Update rule
func (r *RuleRepoImpl) Update(ctx context.Context, rule Rule) (err error) {
	builder := sq.StatementBuilder.
		Update("rules").
		Set("data_source_id", rule.DataSourceID).
		Set("name", rule.Name).
		Set("status", rule.Status).
		Set("url_pattern", rule.URLPattern).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": rule.ID}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.BaseRunner(ctx, r))

	if rule.Status != "" {
		builder = builder.Set("change_status_at", time.Now())
	}

	if _, err = builder.ExecContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
	}
	return
}
