package urlstore

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"go.uber.org/dig"
)

// URLStoreSyncRepoImpl is implementation urlStoreSync repository
type URLStoreSyncRepoImpl struct {
	dig.In
	*sql.DB
}

// FindOne urlStoreSync
func (r *URLStoreSyncRepoImpl) FindOne(ctx context.Context, version int64) (urlStoreSync *URLStoreSync, err error) {
	var rows *sql.Rows
	builder := sq.
		Select("version", "operation", "rule_id", "latest_url_pattern", "created_at").
		From("urlstore_sync").
		Where(sq.Eq{"version": version}).
		PlaceholderFormat(sq.Dollar).RunWith(dbkit.TxCtx(ctx, r))
	if rows, err = builder.QueryContext(ctx); err != nil {
		return
	}
	if rows.Next() {
		urlStoreSync, err = scanURLStoreSync(rows)
	}
	return
}

// Find urlStoreSync
func (r *URLStoreSyncRepoImpl) Find(ctx context.Context) (list []*URLStoreSync, err error) {
	var rows *sql.Rows
	builder := sq.
		Select("version", "operation", "rule_id", "latest_url_pattern", "created_at").
		From("urlstore_sync").
		OrderBy("version").
		PlaceholderFormat(sq.Dollar).RunWith(dbkit.TxCtx(ctx, r))
	if rows, err = builder.QueryContext(ctx); err != nil {
		return
	}
	list = make([]*URLStoreSync, 0)
	for rows.Next() {
		var urlStoreSync *URLStoreSync
		if urlStoreSync, err = scanURLStoreSync(rows); err != nil {
			return
		}
		list = append(list, urlStoreSync)
	}
	return
}

// Insert urlStoreSync
func (r *URLStoreSyncRepoImpl) Insert(ctx context.Context, urlStoreSync URLStoreSync) (lastInsertID int64, err error) {
	query := sq.
		Insert("urlstore_sync").
		Columns("operation", "rule_id", "latest_url_pattern").
		Values(urlStoreSync.Operation, urlStoreSync.RuleID, urlStoreSync.LatestURLPattern).
		Suffix("RETURNING \"version\"").
		PlaceholderFormat(sq.Dollar).RunWith(dbkit.TxCtx(ctx, r))
	if err = query.QueryRowContext(ctx).Scan(&urlStoreSync.Version); err != nil {
		return
	}
	lastInsertID = urlStoreSync.Version
	return
}

// GetLatestVersion of url store
func (r *URLStoreSyncRepoImpl) GetLatestVersion(ctx context.Context) (latestVersion int64, err error) {
	builder := sq.
		Select("version").
		From("urlstore_sync").
		OrderBy("version DESC").
		Limit(1).
		PlaceholderFormat(sq.Dollar).RunWith(dbkit.TxCtx(ctx, r))
	if err = builder.QueryRowContext(ctx).Scan(&latestVersion); err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return
	}

	return
}

func (r *URLStoreSyncRepoImpl) GetListDiff(ctx context.Context, offsetVersion int64) (list []*URLStoreSync, err error) {
	var rows *sql.Rows
	builder := sq.
		Select("version", "operation", "rule_id", "latest_url_pattern", "created_at").
		From("urlstore_sync").
		Where(sq.Gt{"version": offsetVersion}).
		OrderBy("version").
		PlaceholderFormat(sq.Dollar).RunWith(dbkit.TxCtx(ctx, r))
	if rows, err = builder.QueryContext(ctx); err != nil {
		return
	}
	for rows.Next() {
		var urlStoreSync *URLStoreSync
		if urlStoreSync, err = scanURLStoreSync(rows); err != nil {
			return
		}
		list = append(list, urlStoreSync)
	}
	return
}

func scanURLStoreSync(rows *sql.Rows) (*URLStoreSync, error) {
	var urlStoreSync URLStoreSync
	var err error
	if err = rows.Scan(&urlStoreSync.Version, &urlStoreSync.Operation, &urlStoreSync.RuleID, &urlStoreSync.LatestURLPattern, &urlStoreSync.CreatedAt); err != nil {
		return nil, err
	}
	return &urlStoreSync, nil
}
