package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/dig"
)

func scanRule(rows *sql.Rows) (*Rule, error) {
	var rule Rule
	var err error
	if err = rows.Scan(&rule.ID, &rule.Name, &rule.UrlPattern, &rule.UpdatedAt, &rule.CreatedAt); err != nil {
		return nil, err
	}
	return &rule, nil
}

// RuleRepoImpl is implementation rule repository
type RuleRepoImpl struct {
	dig.In
	*sql.DB
}

// Find rule
func (r *RuleRepoImpl) Find(ctx context.Context, id int64) (rule *Rule, err error) {
	var rows *sql.Rows
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.Select("id", "name", "url_pattern", "updated_at", "created_at").
		From("rule").
		Where(sq.Eq{"id": id})
	if rows, err = builder.RunWith(r.DB).QueryContext(ctx); err != nil {
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
		From("rule")
	if rows, err = builder.RunWith(r.DB).QueryContext(ctx); err != nil {
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
