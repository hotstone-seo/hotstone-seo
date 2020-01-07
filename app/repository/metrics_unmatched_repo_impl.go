package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"go.uber.org/dig"
)

// MetricsUnmatchedRepoImpl is implementation metrics_unmatched repository
type MetricsUnmatchedRepoImpl struct {
	dig.In
	*sql.DB
}

// Find metrics_unmatched
func (r *MetricsUnmatchedRepoImpl) Find(ctx context.Context, id int64) (e *MetricsUnmatched, err error) {
	var rows *sql.Rows
	builder := sq.
		Select("id", "request_path", "created_at", "updated_at").
		From("metrics_unmatched").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).RunWith(dbkit.TxCtx(ctx, r))
	if rows, err = builder.QueryContext(ctx); err != nil {
		return
	}
	if rows.Next() {
		e = new(MetricsUnmatched)
		if err = rows.Scan(&e.ID, &e.RequestPath, &e.CreatedAt, &e.UpdatedAt); err != nil {
			return nil, err
		}
	}
	return
}

// List metrics_unmatched
func (r *MetricsUnmatchedRepoImpl) List(ctx context.Context) (list []*MetricsUnmatched, err error) {
	var rows *sql.Rows
	builder := sq.
		Select("id", "request_path", "created_at", "updated_at").
		From("metrics_unmatched").
		PlaceholderFormat(sq.Dollar).RunWith(dbkit.TxCtx(ctx, r))
	if rows, err = builder.QueryContext(ctx); err != nil {
		return
	}
	list = make([]*MetricsUnmatched, 0)
	for rows.Next() {
		var e0 MetricsUnmatched
		if err = rows.Scan(&e0.ID, &e0.RequestPath, &e0.CreatedAt, &e0.UpdatedAt); err != nil {
			return
		}
		list = append(list, &e0)
	}
	return
}

func (r *MetricsUnmatchedRepoImpl) ListCount(ctx context.Context) (list []*MetricsUnmatchedCount, err error) {
	var rows *sql.Rows

	subQuery := sq.
		Select("request_path").
		Column("count(request_path)").
		Column(sq.Alias(sq.Expr("max(created_at)"), "since")).
		From("metrics_unmatched").
		GroupBy("request_path").
		PlaceholderFormat(sq.Dollar)

	builder := sq.
		Select("request_path", "count", "since").
		FromSelect(subQuery, "u").
		OrderBy("u.since desc", "u.count desc").
		PlaceholderFormat(sq.Dollar).RunWith(dbkit.TxCtx(ctx, r))

	if rows, err = builder.QueryContext(ctx); err != nil {
		return
	}
	list = make([]*MetricsUnmatchedCount, 0)
	for rows.Next() {
		var e0 MetricsUnmatchedCount
		if err = rows.Scan(&e0.RequestPath, &e0.Count, &e0.Since); err != nil {
			return
		}
		list = append(list, &e0)
	}
	return
}

// Insert metrics_unmatched
func (r *MetricsUnmatchedRepoImpl) Insert(ctx context.Context, e MetricsUnmatched) (lastInsertID int64, err error) {
	builder := sq.
		Insert("metrics_unmatched").
		Columns("request_path").
		Values(e.RequestPath).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(sq.Dollar).RunWith(dbkit.TxCtx(ctx, r))
	if err = builder.QueryRowContext(ctx).Scan(&e.ID); err != nil {
		return
	}
	lastInsertID = e.ID
	return
}

// Delete metrics_unmatched
func (r *MetricsUnmatchedRepoImpl) Delete(ctx context.Context, id int64) (err error) {
	builder := sq.
		Delete("metrics_unmatched").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).RunWith(dbkit.TxCtx(ctx, r))
	_, err = builder.ExecContext(ctx)
	return
}

// Update metrics_unmatched
func (r *MetricsUnmatchedRepoImpl) Update(ctx context.Context, e MetricsUnmatched) (err error) {
	builder := sq.
		Update("metrics_unmatched").
		Set("request_path", e.RequestPath).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": e.ID}).
		PlaceholderFormat(sq.Dollar).RunWith(dbkit.TxCtx(ctx, r))
	_, err = builder.ExecContext(ctx)
	return
}
