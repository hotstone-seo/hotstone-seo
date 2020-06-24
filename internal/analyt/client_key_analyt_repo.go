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
		Insert(context.Context, int64) (err error)
	}

	// ClientKeyAnalytRepoImpl is implementation of ClientKeyAnalytRepo
	ClientKeyAnalytRepoImpl struct {
		dig.In
		*sql.DB `name:"analyt"`
	}
)

// NewClientKeyAnalytRepo return new instance of MetricsRuleMatchingRepo
// @ctor
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
		RunWith(r)

	if err = builder.QueryRowContext(ctx).Scan(&lastTimeUsed); err != nil {
		return
	}
	return
}

// Insert metrics_client_key
func (r *ClientKeyAnalytRepoImpl) Insert(ctx context.Context, clientKeyID int64) (err error) {
	txn, err := dbtxn.Use(ctx, r.DB)
	if err != nil {
		return err
	}
	builder := sq.
		Insert("metrics_client_key").
		Columns("client_key_id").
		Values(clientKeyID).
		PlaceholderFormat(sq.Dollar).
		RunWith(txn.DB())

	if _, err = builder.ExecContext(ctx); err != nil {
		txn.SetError(err)
		return
	}
	return
}
