package metric

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"go.uber.org/dig"
)

// RuleMatching represented  metrics_rule_matching entity
type RuleMatching struct {
	Time      time.Time
	IsMatched int
	RuleID    *int64
	URL       *string
}

// RuleMatchingRepo to handle metrics_rule_matching entity [mock]
type RuleMatchingRepo interface {
	Insert(context.Context, RuleMatching) (err error)
}

// NewRuleMatchingRepo return new instance of MetricsRuleMatchingRepo [constructor]
func NewRuleMatchingRepo(impl RuleMatchingRepoImpl) RuleMatchingRepo {
	return &impl
}

// RuleMatchingRepoImpl is implementation metrics_rule_matching repository
type RuleMatchingRepoImpl struct {
	dig.In
	*typpostgres.DB
}

// Insert metrics_rule_matching
func (r *RuleMatchingRepoImpl) Insert(ctx context.Context, e RuleMatching) (err error) {
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
