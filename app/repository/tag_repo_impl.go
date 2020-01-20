package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"go.uber.org/dig"
)

// TagRepoImpl is implementation tag repository
type TagRepoImpl struct {
	dig.In
	*sql.DB
}

// FindOne tag
func (r *TagRepoImpl) FindOne(ctx context.Context, id int64) (e *Tag, err error) {
	var rows *sql.Rows
	builder := sq.
		Select("id", "rule_id", "locale", "type", "attributes", "value", "updated_at", "created_at").
		From("tags").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).RunWith(dbkit.TxCtx(ctx, r))
	if rows, err = builder.QueryContext(ctx); err != nil {
		return
	}
	if rows.Next() {
		e = new(Tag)
		if err = rows.Scan(&e.ID, &e.RuleID, &e.Locale, &e.Type, &e.Attributes, &e.Value, &e.UpdatedAt, &e.CreatedAt); err != nil {
			return nil, err
		}
	}
	return
}

// Find tags
func (r *TagRepoImpl) Find(ctx context.Context, filter TagFilter) (list []*Tag, err error) {
	var rows *sql.Rows
	builder := sq.
		Select("id", "rule_id", "locale", "type", "attributes", "value", "updated_at", "created_at").
		From("tags")
	if filter.RuleID > -1 {
		builder = builder.Where(sq.Eq{"rule_id": filter.RuleID})
	}
	if filter.Locale != "" {
		builder = builder.Where(sq.Eq{"locale": filter.Locale})
	}
	builder = builder.PlaceholderFormat(sq.Dollar).RunWith(dbkit.TxCtx(ctx, r))
	if rows, err = builder.QueryContext(ctx); err != nil {
		return
	}
	list = make([]*Tag, 0)
	for rows.Next() {
		var e Tag
		if err = rows.Scan(&e.ID, &e.RuleID, &e.Locale, &e.Type, &e.Attributes, &e.Value, &e.UpdatedAt, &e.CreatedAt); err != nil {
			return
		}
		list = append(list, &e)
	}
	return
}

// Insert tag
func (r *TagRepoImpl) Insert(ctx context.Context, e Tag) (lastInsertID int64, err error) {
	builder := sq.
		Insert("tags").
		Columns("rule_id", "locale", "type", "attributes", "value").
		Values(e.RuleID, e.Locale, e.Type, e.Attributes, e.Value).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(sq.Dollar).RunWith(dbkit.TxCtx(ctx, r))
	if err = builder.QueryRowContext(ctx).Scan(&e.ID); err != nil {
		return
	}
	lastInsertID = e.ID
	return
}

// Delete tag
func (r *TagRepoImpl) Delete(ctx context.Context, id int64) (err error) {
	builder := sq.
		Delete("tags").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).RunWith(dbkit.TxCtx(ctx, r))
	_, err = builder.ExecContext(ctx)
	return
}

// Update tag
func (r *TagRepoImpl) Update(ctx context.Context, e Tag) (err error) {
	builder := sq.
		Update("tags").
		Set("rule_id", e.RuleID).
		Set("locale", e.Locale).
		Set("type", e.Type).
		Set("attributes", e.Attributes).
		Set("value", e.Value).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": e.ID}).
		PlaceholderFormat(sq.Dollar).RunWith(dbkit.TxCtx(ctx, r))
	_, err = builder.ExecContext(ctx)
	return
}
