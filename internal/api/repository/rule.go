package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/typical-go/typical-rest-server/pkg/dbtxn"
	"go.uber.org/dig"
	"gopkg.in/go-playground/validator.v9"
)

// Rule Entity
type Rule struct {
	ID             int64     `json:"id" structs:"id"`
	Name           string    `json:"name" structs:"name" validate:"required"`
	URLPattern     string    `json:"url_pattern" structs:"url_pattern" validate:"required,uri"`
	DataSourceIDs  []int64   `json:"data_source_ids" structs:"data_source_ids"`
	UpdatedAt      time.Time `json:"updated_at" structs:"updated_at"`
	CreatedAt      time.Time `json:"created_at" structs:"created_at"`
	Status         string    `json:"status" structs:"status"`
	ChangeStatusAt time.Time `json:"change_status_at" structs:"change_status_at"`
}

// RuleRepo is rule repository
// @mock
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
	*sql.DB
}

// NewRuleRepo return new instance of RuleRepo
// @ctor
func NewRuleRepo(impl RuleRepoImpl) RuleRepo {
	return &impl
}

// Validate rule
func (rule Rule) Validate() error {
	return validator.New().Struct(rule)
}

// FindOne rule
func (r *RuleRepoImpl) FindOne(ctx context.Context, id int64) (rule *Rule, err error) {
	var rows *sql.Rows

	row := sq.StatementBuilder.
		Select(
			"id",
			"name",
			"url_pattern",
			"updated_at",
			"created_at",
			"status",
			"change_status_at",
		).
		From("rules").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(r).
		QueryRowContext(ctx)

	rule = &Rule{DataSourceIDs: make([]int64, 0)}
	if err = row.Scan(
		&rule.ID,
		&rule.Name,
		&rule.URLPattern,
		&rule.UpdatedAt,
		&rule.CreatedAt,
		&rule.Status,
		&rule.ChangeStatusAt,
	); err != nil {
		return nil, err
	}

	dsBuilder := sq.StatementBuilder.
		Select("data_source_id").
		From("rule_data_sources").
		Where(sq.Eq{"rule_id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(r)

	if rows, err = dsBuilder.QueryContext(ctx); err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var dataSourceID int64
		if err = rows.Scan(&dataSourceID); err != nil {
			return nil, err
		}
		rule.DataSourceIDs = append(rule.DataSourceIDs, dataSourceID)
	}

	return
}

// Find rule
func (r *RuleRepoImpl) Find(ctx context.Context, paginationParam PaginationParam) (list []*Rule, err error) {

	builder := sq.StatementBuilder.
		Select(
			"id",
			"name",
			"url_pattern",
			"updated_at",
			"created_at",
			"status",
			"change_status_at",
		).
		From("rules").
		PlaceholderFormat(sq.Dollar).
		RunWith(r)

	builder = ComposePagination(builder, paginationParam)

	rows, err := builder.QueryContext(ctx)
	if err != nil {
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
			&rule.UpdatedAt,
			&rule.CreatedAt,
			&rule.Status,
			&rule.ChangeStatusAt,
		); err != nil {
			return
		}
		list = append(list, rule)
	}
	return
}

// Insert rule
func (r *RuleRepoImpl) Insert(ctx context.Context, rule Rule) (lastInsertID int64, err error) {
	query := sq.
		Insert("rules").
		Columns(
			"name",
			"url_pattern",
		).
		Values(
			rule.Name,
			rule.URLPattern,
		).
		Suffix("RETURNING \"id\"").
		RunWith(r).
		PlaceholderFormat(sq.Dollar).
		QueryRowContext(ctx)

	if err = query.Scan(&lastInsertID); err != nil {
		return
	}

	if len(rule.DataSourceIDs) > 0 {
		insertDataSource := sq.Insert("rule_data_sources").
			Columns(
				"rule_id",
				"data_source_id",
			).
			RunWith(r).
			PlaceholderFormat(sq.Dollar)

		for _, dataSourceID := range rule.DataSourceIDs {
			insertDataSource = insertDataSource.Values(lastInsertID, dataSourceID)
		}
		if _, err = insertDataSource.ExecContext(ctx); err != nil {
			return
		}
	}

	return
}

// Delete rule
func (r *RuleRepoImpl) Delete(ctx context.Context, id int64) (err error) {
	builder := sq.StatementBuilder.
		Delete("rules").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(r)

	if _, err = builder.ExecContext(ctx); err != nil {
		return
	}
	return
}

// Update rule
func (r *RuleRepoImpl) Update(ctx context.Context, rule Rule) (err error) {
	txn, err := dbtxn.Use(ctx, r.DB)
	if err != nil {
		return
	}
	builder := sq.StatementBuilder.
		Update("rules").
		Set("name", rule.Name).
		Set("status", rule.Status).
		Set("url_pattern", rule.URLPattern).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": rule.ID}).
		PlaceholderFormat(sq.Dollar).
		RunWith(txn.DB)

	if rule.Status != "" {
		builder = builder.Set("change_status_at", time.Now())
	}

	if _, err = builder.ExecContext(ctx); err != nil {
		return
	}

	deletePrevDataSource := sq.StatementBuilder.
		Delete("rule_data_sources").
		Where(sq.Eq{"rule_id": rule.ID}).
		PlaceholderFormat(sq.Dollar).
		RunWith(txn.DB)

	if _, err = deletePrevDataSource.ExecContext(ctx); err != nil {
		return
	}

	if len(rule.DataSourceIDs) > 0 {
		insertDataSource := sq.Insert("rule_data_sources").
			Columns(
				"rule_id",
				"data_source_id",
			).
			RunWith(txn.DB).
			PlaceholderFormat(sq.Dollar)

		for _, dataSourceID := range rule.DataSourceIDs {
			insertDataSource = insertDataSource.Values(rule.ID, dataSourceID)
		}
		if _, err = insertDataSource.ExecContext(ctx); err != nil {
			return
		}
	}
	return
}
