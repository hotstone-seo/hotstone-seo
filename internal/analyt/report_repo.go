package analyt

import (
	"context"
	"database/sql"
	"net/url"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"go.uber.org/dig"
)

type (
	// ReportRepo responsble to generate report
	// @mock
	ReportRepo interface {
		MismatchReports(ctx context.Context, paginationParam repository.PaginationParam) (list []*MismatchReport, err error)
		DailyReports(ctx context.Context, startDate, endDate, ruleID string) (list []*DailyReport, err error)

		CountMatched(ctx context.Context, whereParams url.Values) (count int64, err error)
		CountUniquePage(ctx context.Context, whereParams url.Values) (count int64, err error)
	}

	// MismatchReport contain sumary of total count of not-matching rule
	MismatchReport struct {
		URL       string    `json:"url"`
		Count     int64     `json:"count"`
		FirstSeen time.Time `json:"first_seen"`
		LastSeen  time.Time `json:"last_seen"`
	}

	// DailyReport contain sumary of daily hit
	DailyReport struct {
		Date     time.Time `json:"date"`
		HitCount int64     `json:"count"`
	}

	// ReportRepoImpl is implementation of ReportRepo
	ReportRepoImpl struct {
		dig.In
		*sql.DB `name:"analyt"`
	}
)

// NewReportRepo return new instance of MetricsRuleMatchingRepo
// @constructor
func NewReportRepo(impl ReportRepoImpl) ReportRepo {
	return &impl
}

// MismatchReports return list of not-matching report
func (r *ReportRepoImpl) MismatchReports(ctx context.Context, paginationParam repository.PaginationParam) (list []*MismatchReport, err error) {
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

	builder = repository.ComposePagination(builder, paginationParam).RunWith(dbtxn.DB(ctx, r))

	if rows, err = builder.QueryContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	defer rows.Close()
	list = make([]*MismatchReport, 0)
	for rows.Next() {
		var report MismatchReport
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

// CountMatched is count matched
func (r *ReportRepoImpl) CountMatched(ctx context.Context, whereParams url.Values) (count int64, err error) {

	builder := sq.Select().
		Column("count(is_matched)").
		From("metrics_rule_matching").
		Where(sq.Eq{"is_matched": 1}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.DB(ctx, r))

	builder = buildWhereQuery(builder, whereParams, []string{"rule_id"})

	if err = builder.QueryRowContext(ctx).Scan(&count); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}

	return
}

// CountUniquePage is count unique page
func (r *ReportRepoImpl) CountUniquePage(ctx context.Context, whereParams url.Values) (count int64, err error) {

	builder := sq.Select().
		Column("count(distinct(url))").
		From("metrics_rule_matching").
		Where(sq.Eq{"is_matched": 1})

	builder = buildWhereQuery(builder, whereParams, []string{"rule_id"}).
		PlaceholderFormat(sq.Dollar).RunWith(dbtxn.DB(ctx, r))

	if err = builder.QueryRowContext(ctx).Scan(&count); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}

	return
}

// DailyReports return list of daily report
func (r *ReportRepoImpl) DailyReports(ctx context.Context, startDate, endDate, ruleID string) (list []*DailyReport, err error) {
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
