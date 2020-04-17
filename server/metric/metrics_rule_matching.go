package metric

import (
	"context"
	"database/sql"
	"net/url"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"go.uber.org/dig"
)

// MetricsRuleMatching represented  metrics_rule_matching entity
type MetricsRuleMatching struct {
	Time      time.Time
	IsMatched int
	RuleID    *int64
	URL       *string
}

// MetricsRuleMatchingRepo to handle metrics_rule_matching entity [mock]
type MetricsRuleMatchingRepo interface {
	Insert(context.Context, MetricsRuleMatching) (err error)

	NotMatchedReports(ctx context.Context, paginationParam repository.PaginationParam) (list []*NotMatchedReport, err error)
	DailyReports(ctx context.Context, startDate, endDate, ruleID string) (list []*DailyReport, err error)

	CountMatched(ctx context.Context, whereParams url.Values) (count int64, err error)
	CountUniquePage(ctx context.Context, whereParams url.Values) (count int64, err error)
}

// NewMetricsRuleMatchingRepo return new instance of MetricsRuleMatchingRepo [constructor]
func NewMetricsRuleMatchingRepo(impl MetricsRuleMatchingRepoImpl) MetricsRuleMatchingRepo {
	return &impl
}

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
		PlaceholderFormat(sq.Dollar).RunWith(dbtxn.BaseRunner(ctx, r))

	if _, err = builder.ExecContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	return
}

// NotMatchedReports return list of not-matching report
func (r *MetricsRuleMatchingRepoImpl) NotMatchedReports(ctx context.Context, paginationParam repository.PaginationParam) (list []*NotMatchedReport, err error) {
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
		Where("not exists(select url from metrics_rule_matching mrm where mrm.url = u.url and is_matched=1)").
		PlaceholderFormat(sq.Dollar)

	builder = repository.ComposePagination(builder, paginationParam).RunWith(dbtxn.BaseRunner(ctx, r))

	if rows, err = builder.QueryContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	defer rows.Close()
	list = make([]*NotMatchedReport, 0)
	for rows.Next() {
		var report NotMatchedReport
		if err = rows.Scan(
			&report.URL,
			&report.Count,
			&report.FirstSeen,
			&report.LastSeen,
		); err != nil {
			dbtxn.SetError(ctx, err)
			return
		}
		list = append(list, &report)
	}
	return
}

func (r *MetricsRuleMatchingRepoImpl) CountMatched(ctx context.Context, whereParams url.Values) (count int64, err error) {

	builder := sq.Select().
		Column("count(is_matched)").
		From("metrics_rule_matching").
		Where(sq.Eq{"is_matched": 1}).
		PlaceholderFormat(sq.Dollar).RunWith(dbtxn.BaseRunner(ctx, r))

	builder = buildWhereQuery(builder, whereParams, []string{"rule_id"})

	if err = builder.QueryRowContext(ctx).Scan(&count); err != nil {
		dbtxn.SetError(ctx, err)
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
		PlaceholderFormat(sq.Dollar).RunWith(dbtxn.BaseRunner(ctx, r))

	if err = builder.QueryRowContext(ctx).Scan(&count); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}

	return
}

// DailyReports return list of daily report
func (r *MetricsRuleMatchingRepoImpl) DailyReports(ctx context.Context, startDate, endDate, ruleID string) (list []*DailyReport, err error) {
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

	list = make([]*DailyReport, 0)
	for rows.Next() {
		var report DailyReport
		if err = rows.Scan(&report.Date, &report.HitCount); err != nil {
			return
		}
		list = append(list, &report)
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
