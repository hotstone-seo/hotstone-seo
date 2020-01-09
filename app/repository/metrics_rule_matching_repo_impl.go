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
