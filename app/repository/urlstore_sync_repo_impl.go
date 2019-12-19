package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/dig"
)

// URLStoreSyncRepoImpl is implementation urlStoreSync repository
type URLStoreSyncRepoImpl struct {
	dig.In
	SqlDB *sql.DB
}

func (r *URLStoreSyncRepoImpl) DB() *sql.DB {
	return r.SqlDB
}

// Find urlStoreSync
func (r *URLStoreSyncRepoImpl) Find(ctx context.Context, tx *sql.Tx, version int64) (urlStoreSync *URLStoreSync, err error) {
	var rows *sql.Rows
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.Select("version", "operation", "rule_id", "latest_url_pattern", "created_at").
		From("urlstore_sync").
		Where(sq.Eq{"version": version})
	if rows, err = builder.RunWith(tx).QueryContext(ctx); err != nil {
		return
	}
	if rows.Next() {
		urlStoreSync, err = scanURLStoreSync(rows)
	}
	return
}

// List urlStoreSync
func (r *URLStoreSyncRepoImpl) List(ctx context.Context, tx *sql.Tx) (list []*URLStoreSync, err error) {
	var rows *sql.Rows
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.Select("version", "operation", "rule_id", "latest_url_pattern", "created_at").
		From("urlstore_sync")
	if rows, err = builder.RunWith(tx).QueryContext(ctx); err != nil {
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
func (r *URLStoreSyncRepoImpl) Insert(ctx context.Context, tx *sql.Tx, urlStoreSync URLStoreSync) (lastInsertID int64, err error) {

	query := sq.Insert("urlstore_sync").
		Columns("operation", "rule_id", "latest_url_pattern").
		Values(urlStoreSync.Operation, urlStoreSync.RuleID, urlStoreSync.LatestURLPattern).
		Suffix("RETURNING \"version\"").
		RunWith(tx).
		PlaceholderFormat(sq.Dollar)
	if err = query.QueryRowContext(ctx).Scan(&urlStoreSync.Version); err != nil {
		return
	}
	lastInsertID = urlStoreSync.Version
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
