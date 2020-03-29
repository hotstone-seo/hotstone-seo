package repository

import (
	"context"
	"database/sql"
	"net/url"

	sq "github.com/Masterminds/squirrel"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"go.uber.org/dig"
)

// MetricsRuleMatchingRepoImpl is implementation metrics_rule_matching repository
type MetricsRuleMatchingRepoImpl struct {
	dig.In
	*typpostgres.DB
}

// Insert metrics_rule_matching
func (r *MetricsRuleMatchingRepoImpl) Insert(ctx context.Context, e MetricsRuleMatching) (err error) {
	builder := sq.
		Insert("metrics_rule_matching").
		Columns("is_matched", "url", "rule_id").
		Values(e.IsMatched, e.URL, e.RuleID).
		PlaceholderFormat(sq.Dollar).RunWith(dbkit.TxCtx(ctx, r))

	if _, err = builder.ExecContext(ctx); err != nil {
		return
	}
	return
}

// ListMismatchedCount list mistached count
func (r *MetricsRuleMatchingRepoImpl) ListMismatchedCount(ctx context.Context, paginationParam PaginationParam) (list []*MetricsMismatchedCount, err error) {
	var rows *sql.Rows

	subQuery := sq.
		Select("url").
		Column("count(url)").
		Column(sq.Alias(sq.Expr("min(time)"), "first_seen")).
		Column(sq.Alias(sq.Expr("max(time)"), "last_seen")).
		From("metrics_rule_matching").
		Where(sq.Eq{"is_matched": 0}).
		GroupBy("url")

	builder := sq.
		Select("url", "count", "first_seen", "last_seen").
		FromSelect(subQuery, "u").
		PlaceholderFormat(sq.Dollar)

	builder = composePagination(builder, paginationParam).RunWith(dbkit.TxCtx(ctx, r))

	if rows, err = builder.QueryContext(ctx); err != nil {
		return
	}
	defer rows.Close()
	list = make([]*MetricsMismatchedCount, 0)
	for rows.Next() {
		var e0 MetricsMismatchedCount
		if err = rows.Scan(&e0.URL, &e0.Count, &e0.FirstSeen, &e0.LastSeen); err != nil {
			return
		}
		list = append(list, &e0)
	}
	return
}

func (r *MetricsRuleMatchingRepoImpl) CountMatched(ctx context.Context, whereParams url.Values) (count int64, err error) {

	builder := sq.Select().
		Column("count(is_matched)").
		From("metrics_rule_matching").
		Where(sq.Eq{"is_matched": 1}).
		PlaceholderFormat(sq.Dollar).RunWith(dbkit.TxCtx(ctx, r))

	builder = buildWhereQuery(builder, whereParams, []string{"rule_id"}).
		PlaceholderFormat(sq.Dollar).RunWith(dbkit.TxCtx(ctx, r))

	if err = builder.QueryRowContext(ctx).Scan(&count); err != nil {
		return
	}

	return
}

func (r *MetricsRuleMatchingRepoImpl) CountUniquePage(ctx context.Context, whereParams url.Values) (count int64, err error) {

	builder := sq.Select().
		Column("count(distinct(url))").
		From("metrics_rule_matching").
		Where(sq.Eq{"is_matched": 1})

	builder = buildWhereQuery(builder, whereParams, []string{"rule_id"}).
		PlaceholderFormat(sq.Dollar).RunWith(dbkit.TxCtx(ctx, r))

	if err = builder.QueryRowContext(ctx).Scan(&count); err != nil {
		return
	}

	return
}

func (r *MetricsRuleMatchingRepoImpl) ListCountHitPerDay(ctx context.Context, startDate, endDate, ruleID string) (list []*MetricsCountHitPerDay, err error) {
	var rows *sql.Rows

	query := `
	WITH range_date AS (
		SELECT generate_series($1::date, $2::date, '1 day') as date
	),
	count_per_day AS (
		select date(time), count(is_matched) from metrics_rule_matching
		WHERE is_matched = 1 AND (
			CASE WHEN $3 != '' THEN
				rule_id = $3::int
			ELSE true
			END
		)
		GROUP BY date(time)
	)
	select rd.date, coalesce(cpd.count, 0) as count from range_date rd
	left join count_per_day cpd on rd.date = cpd.date
	`

	if rows, err = r.DB.Query(query, startDate, endDate, ruleID); err != nil {
		return
	}
	defer rows.Close()

	list = make([]*MetricsCountHitPerDay, 0)
	for rows.Next() {
		var e0 MetricsCountHitPerDay
		if err = rows.Scan(&e0.Date, &e0.Count); err != nil {
			return
		}
		list = append(list, &e0)
	}
	return
}

func buildWhereQuery(builder sq.SelectBuilder, whereParams url.Values, validColumns []string) sq.SelectBuilder {
	for key, val := range whereParams {
		for _, col := range validColumns {
			if col == key {
				builder = builder.Where(sq.Eq{key: val})
			}
		}
	}

	return builder
}
