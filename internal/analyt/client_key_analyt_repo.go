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
	// ClientKeyAnalytRepo responsble to retrieve metric of ClientKey
	// @mock
	ClientKeyAnalytRepo interface {
		ClientKeyLastUsed(ctx context.Context, clientKeyID string) (lastUsed *time.Time, err error)
	}

	// ClientKeyAnalytRepoImpl is implementation of ClientKeyAnalytRepo
	ClientKeyAnalytRepoImpl struct {
		dig.In
		*sql.DB `name:"analyt"`
	}
)

// NewClientKeyAnalytRepo return new instance of MetricsRuleMatchingRepo
// @constructor
func NewClientKeyAnalytRepo(impl ClientKeyAnalytRepoImpl) ClientKeyAnalytRepo {
	return &impl
}

// ClientKeyLastUsed get last time client key is used
func (r *ClientKeyAnalytRepoImpl) ClientKeyLastUsed(ctx context.Context, clientKeyID string) (lastTimeUsed *time.Time, err error) {
	builder := sq.Select().
		Column(sq.Alias(sq.Expr("max(time)"), "time")).
		From("metrics_client_key").
		Where(sq.Eq{"client_key_id": clientKeyID}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbtxn.BaseRunner(ctx, r))

	if err = builder.QueryRowContext(ctx).Scan(&lastTimeUsed); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	return
}
