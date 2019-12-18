package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/dig"
)

// RuleRepoImpl is implementation rule repository
type RuleRepoImpl struct {
	dig.In
	SqlDB *sql.DB
}

func (r *RuleRepoImpl) DB() *sql.DB {
	return r.SqlDB
}

// Find rule
func (r *RuleRepoImpl) Find(ctx context.Context, id int64) (rule *Rule, err error) {
	var rows *sql.Rows
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.Select("id", "name", "url_pattern", "updated_at", "created_at").
		From("rules").
		Where(sq.Eq{"id": id})
	if rows, err = builder.RunWith(r.DB()).QueryContext(ctx); err != nil {
		return
	}
	if rows.Next() {
		rule, err = scanRule(rows)
	}
	return
}

// List rule
func (r *RuleRepoImpl) List(ctx context.Context) (list []*Rule, err error) {
	var rows *sql.Rows
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.Select("id", "name", "url_pattern", "updated_at", "created_at").
		From("rules")
	if rows, err = builder.RunWith(r.DB()).QueryContext(ctx); err != nil {
		return
	}
	list = make([]*Rule, 0)
	for rows.Next() {
		var rule *Rule
		if rule, err = scanRule(rows); err != nil {
			return
		}
		list = append(list, rule)
	}
	return
}

// Insert rule
func (r *RuleRepoImpl) Insert(ctx context.Context, rule Rule) (lastInsertID int64, err error) {
	// tx, err := NewTxIfNotExist(ctx, r.DB())
	tx, err := GetTx(ctx)
	if err != nil {
		return
	}

	query := sq.Insert("rules").
		Columns("data_source_id", "name", "url_pattern").
		Values(rule.DataSourceID, rule.Name, rule.UrlPattern).
		Suffix("RETURNING \"id\"").
		RunWith(tx).
		PlaceholderFormat(sq.Dollar)
	if err = query.QueryRowContext(ctx).Scan(&rule.ID); err != nil {
		return
	}
	lastInsertID = rule.ID
	return
}

// Delete rule
func (r *RuleRepoImpl) Delete(ctx context.Context, id int64) (err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.Delete("rules").Where(sq.Eq{"id": id})
	_, err = builder.RunWith(r.DB()).ExecContext(ctx)
	return
}

// Update rule
func (r *RuleRepoImpl) Update(ctx context.Context, rule Rule) (err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.Update("rules").
		Set("data_source_id", rule.DataSourceID).
		Set("name", rule.Name).
		Set("url_pattern", rule.UrlPattern).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": rule.ID})
	_, err = builder.RunWith(r.DB()).ExecContext(ctx)
	return
}

func scanRule(rows *sql.Rows) (*Rule, error) {
	var rule Rule
	var err error
	if err = rows.Scan(&rule.ID, &rule.Name, &rule.UrlPattern, &rule.UpdatedAt, &rule.CreatedAt); err != nil {
		return nil, err
	}
	return &rule, nil
}
