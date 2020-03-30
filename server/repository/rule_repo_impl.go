package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
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
	if rows, err = builder.RunWith(dbkit.TxCtx(ctx, r)).QueryContext(ctx); err != nil {
		return
	}
	defer rows.Close()
	if rows.Next() {
		rule, err = scanRule(rows)
	}
	return
}

// Find rule
func (r *RuleRepoImpl) Find(ctx context.Context, paginationParam PaginationParam) (list []*Rule, err error) {
	var rows *sql.Rows
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.Select("id", "name", "url_pattern", "data_source_id", "updated_at", "created_at", "status", "change_status_at").
		From("rules")
	if rows, err = composePagination(builder, paginationParam).RunWith(dbkit.TxCtx(ctx, r)).QueryContext(ctx); err != nil {
		return
	}
	defer rows.Close()
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

	query := sq.Insert("rules").
		Columns("data_source_id", "name", "url_pattern").
		Values(rule.DataSourceID, rule.Name, rule.UrlPattern).
		Suffix("RETURNING \"id\"").
		RunWith(dbkit.TxCtx(ctx, r)).
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
	_, err = builder.RunWith(dbkit.TxCtx(ctx, r)).ExecContext(ctx)
	return
}

// Update rule
func (r *RuleRepoImpl) Update(ctx context.Context, rule Rule) (err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	if rule.Status != "" {
		builder := psql.Update("rules").
			Set("status", rule.Status).
			Set("change_status_at", time.Now()).
			Where(sq.Eq{"id": rule.ID})
		_, err = builder.RunWith(dbkit.TxCtx(ctx, r)).ExecContext(ctx)
	} else {
		builder := psql.Update("rules").
			Set("data_source_id", rule.DataSourceID).
			Set("name", rule.Name).
			Set("url_pattern", rule.UrlPattern).
			Set("updated_at", time.Now()).
			Where(sq.Eq{"id": rule.ID})
		_, err = builder.RunWith(dbkit.TxCtx(ctx, r)).ExecContext(ctx)
	}
	return
}

func scanRule(rows *sql.Rows) (*Rule, error) {
	var rule Rule
	var err error
	if err = rows.Scan(&rule.ID, &rule.Name, &rule.UrlPattern, &rule.DataSourceID, &rule.UpdatedAt, &rule.CreatedAt, &rule.Status, &rule.ChangeStatusAt); err != nil {
		return nil, err
	}
	return &rule, nil
}
