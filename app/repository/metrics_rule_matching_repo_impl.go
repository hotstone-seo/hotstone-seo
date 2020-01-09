package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"go.uber.org/dig"
)

// MetricsRuleMatchingRepoImpl is implementation metrics_rule_matching repository
type MetricsRuleMatchingRepoImpl struct {
	dig.In
	*sql.DB
}

// Insert metrics_rule_matching
func (r *MetricsRuleMatchingRepoImpl) Insert(ctx context.Context, e MetricsRuleMatching) (err error) {
	builder := sq.
		Insert("metrics_rule_matching").
		Columns("is_matched", "url_mismatched").
		Values(e.IsMatched, e.URLMismatched).
		PlaceholderFormat(sq.Dollar).RunWith(dbkit.TxCtx(ctx, r))

	if _, err = builder.ExecContext(ctx); err != nil {
		return
	}
	return
}

func (r *MetricsRuleMatchingRepoImpl) ListMismatchedCount(ctx context.Context) (list []*MetricsMismatchedCount, err error) {
	var rows *sql.Rows

	subQuery := sq.
		Select("url_mismatched").
		Column("count(url_mismatched)").
		Column(sq.Alias(sq.Expr("max(time)"), "since")).
		From("metrics_rule_matching").
		Where(sq.Eq{"is_matched": 0}).
		GroupBy("url_mismatched").
		PlaceholderFormat(sq.Dollar)

	builder := sq.
		Select("url_mismatched", "count", "since").
		FromSelect(subQuery, "u").
		OrderBy("u.since desc", "u.count desc").
		PlaceholderFormat(sq.Dollar).RunWith(dbkit.TxCtx(ctx, r))

	if rows, err = builder.QueryContext(ctx); err != nil {
		return
	}
	list = make([]*MetricsMismatchedCount, 0)
	for rows.Next() {
		var e0 MetricsMismatchedCount
		if err = rows.Scan(&e0.RequestPath, &e0.Count, &e0.Since); err != nil {
			return
		}
		list = append(list, &e0)
	}
	return
}
