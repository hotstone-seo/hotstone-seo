package analyt

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"go.uber.org/dig"
)

type (
	// RuleMatching represented  metrics_rule_matching entity
	RuleMatching struct {
		Time      time.Time
		IsMatched int
		RuleID    *int64
		URL       *string
	}

	// RuleMatchingRepo to handle metrics_rule_matching entity
	// @mock
	RuleMatchingRepo interface {
		Insert(context.Context, RuleMatching) (err error)
	}

	// RuleMatchingRepoImpl is implementation metrics_rule_matching repository
	RuleMatchingRepoImpl struct {
		dig.In
		*sql.DB `name:"analyt"`
	}
)

// NewRuleMatchingRepo return new instance of MetricsRuleMatchingRepo
// @ctor
func NewRuleMatchingRepo(impl RuleMatchingRepoImpl) RuleMatchingRepo {
	return &impl
}

// Insert metrics_rule_matching
func (r *RuleMatchingRepoImpl) Insert(ctx context.Context, e RuleMatching) (err error) {
	txn, err := dbtxn.Use(ctx, r.DB)
	if err != nil {
		return err
	}
	builder := sq.
		Insert("metrics_rule_matching").
		Columns("is_matched", "url", "rule_id").
		Values(e.IsMatched, e.URL, e.RuleID).
		PlaceholderFormat(sq.Dollar).
		RunWith(r)

	if _, err = builder.ExecContext(ctx); err != nil {
		txn.SetError(err)
		return
	}
	return
}
