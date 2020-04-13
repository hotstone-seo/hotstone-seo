package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"go.uber.org/dig"
)

// RuleRepoImpl is implementation rule repository
type RuleRepoImpl struct {
	dig.In
	*typpostgres.DB
}

// FindOne rule
func (r *RuleRepoImpl) FindOne(ctx context.Context, id int64) (rule *Rule, err error) {
	var rows *sql.Rows
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.Select("id", "name", "url_pattern", "data_source_id", "updated_at", "created_at", "status", "change_status_at").
		From("rules").
		Where(sq.Eq{"id": id})
	if rows, err = builder.RunWith(dbtxn.BaseRunner(ctx, r)).QueryContext(ctx); err != nil {
		return
	}
	defer rows.Close()
	if rows.Next() {
		if rule, err = scanRule(rows); err != nil {
			dbtxn.SetError(ctx, err)
			return
		}
	}
	return
}

// Find rule
func (r *RuleRepoImpl) Find(ctx context.Context, paginationParam PaginationParam) (list []*Rule, err error) {
	var rows *sql.Rows
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.Select("id", "name", "url_pattern", "data_source_id", "updated_at", "created_at", "status", "change_status_at").
		From("rules")
	if rows, err = composePagination(builder, paginationParam).RunWith(dbtxn.BaseRunner(ctx, r)).QueryContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	defer rows.Close()
	list = make([]*Rule, 0)
	for rows.Next() {
		var rule *Rule
		if rule, err = scanRule(rows); err != nil {
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
		Columns("data_source_id", "name", "url_pattern").
		Values(rule.DataSourceID, rule.Name, rule.UrlPattern).
		Suffix("RETURNING \"id\"").
		RunWith(dbtxn.BaseRunner(ctx, r)).
		PlaceholderFormat(sq.Dollar)
	if err = query.QueryRowContext(ctx).Scan(&rule.ID); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	lastInsertID = rule.ID
	return
}

// Delete rule
func (r *RuleRepoImpl) Delete(ctx context.Context, id int64) (err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.Delete("rules").Where(sq.Eq{"id": id})
	if _, err = builder.RunWith(dbtxn.BaseRunner(ctx, r)).ExecContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	return
}

// Update rule
func (r *RuleRepoImpl) Update(ctx context.Context, rule Rule) (err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	builder := psql.Update("rules").
		Set("data_source_id", rule.DataSourceID).
		Set("name", rule.Name).
		Set("status", rule.Status).
		Set("url_pattern", rule.UrlPattern).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": rule.ID}).RunWith(dbtxn.BaseRunner(ctx, r))

	if rule.Status != "" {
		builder = builder.Set("change_status_at", time.Now())
	}

	if _, err = builder.ExecContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
	}
	return
}

func scanRule(rows *sql.Rows) (*Rule, error) {
	var (
		rule Rule
		err  error
	)
	if err = rows.Scan(&rule.ID, &rule.Name, &rule.UrlPattern, &rule.DataSourceID, &rule.UpdatedAt, &rule.CreatedAt, &rule.Status, &rule.ChangeStatusAt); err != nil {
		return nil, err
	}
	return &rule, nil
}
