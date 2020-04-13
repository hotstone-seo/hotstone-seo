package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"go.uber.org/dig"
)

// URLSyncRepoImpl is implementation urlStoreSync repository
type URLSyncRepoImpl struct {
	dig.In
	*typpostgres.DB
}

// FindOne urlStoreSync
func (r *URLSyncRepoImpl) FindOne(ctx context.Context, version int64) (urlStoreSync *URLSync, err error) {
	var rows *sql.Rows
	builder := sq.
		Select("version", "operation", "rule_id", "latest_url_pattern", "created_at").
		From("url_sync").
		Where(sq.Eq{"version": version}).
		PlaceholderFormat(sq.Dollar).RunWith(dbtxn.BaseRunner(ctx, r))
	if rows, err = builder.QueryContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	defer rows.Close()
	if rows.Next() {
		if urlStoreSync, err = scanURLSync(rows); err != nil {
			dbtxn.SetError(ctx, err)
		}
	}
	return
}

// Find urlStoreSync
func (r *URLSyncRepoImpl) Find(ctx context.Context) (list []*URLSync, err error) {
	var rows *sql.Rows
	builder := sq.
		Select("version", "operation", "rule_id", "latest_url_pattern", "created_at").
		From("url_sync").
		OrderBy("version").
		PlaceholderFormat(sq.Dollar).RunWith(dbtxn.BaseRunner(ctx, r))
	if rows, err = builder.QueryContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	defer rows.Close()
	list = make([]*URLSync, 0)
	for rows.Next() {
		var urlStoreSync *URLSync
		if urlStoreSync, err = scanURLSync(rows); err != nil {
			dbtxn.SetError(ctx, err)
			return
		}
		list = append(list, urlStoreSync)
	}
	return
}

// Insert urlStoreSync
func (r *URLSyncRepoImpl) Insert(ctx context.Context, urlStoreSync URLSync) (lastInsertID int64, err error) {
	query := sq.
		Insert("url_sync").
		Columns("operation", "rule_id", "latest_url_pattern").
		Values(urlStoreSync.Operation, urlStoreSync.RuleID, urlStoreSync.LatestURLPattern).
		Suffix("RETURNING \"version\"").
		PlaceholderFormat(sq.Dollar).RunWith(dbtxn.BaseRunner(ctx, r))
	if err = query.QueryRowContext(ctx).Scan(&urlStoreSync.Version); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	lastInsertID = urlStoreSync.Version
	return
}

// GetLatestVersion of url store
func (r *URLSyncRepoImpl) GetLatestVersion(ctx context.Context) (latestVersion int64, err error) {
	builder := sq.
		Select("version").
		From("url_sync").
		OrderBy("version DESC").
		Limit(1).
		PlaceholderFormat(sq.Dollar).RunWith(dbtxn.BaseRunner(ctx, r))
	if err = builder.QueryRowContext(ctx).Scan(&latestVersion); err != nil {
		if err == sql.ErrNoRows {
			dbtxn.SetError(ctx, err)
			return 0, nil
		}
		return
	}

	return
}

func (r *URLSyncRepoImpl) GetListDiff(ctx context.Context, offsetVersion int64) (list []*URLSync, err error) {
	var rows *sql.Rows
	builder := sq.
		Select("version", "operation", "rule_id", "latest_url_pattern", "created_at").
		From("url_sync").
		Where(sq.Gt{"version": offsetVersion}).
		OrderBy("version").
		PlaceholderFormat(sq.Dollar).RunWith(dbtxn.BaseRunner(ctx, r))
	if rows, err = builder.QueryContext(ctx); err != nil {
		dbtxn.SetError(ctx, err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var urlStoreSync *URLSync
		if urlStoreSync, err = scanURLSync(rows); err != nil {
			dbtxn.SetError(ctx, err)
			return
		}
		list = append(list, urlStoreSync)
	}
	return
}

func scanURLSync(rows *sql.Rows) (*URLSync, error) {
	var urlStoreSync URLSync
	var err error
	if err = rows.Scan(&urlStoreSync.Version, &urlStoreSync.Operation, &urlStoreSync.RuleID, &urlStoreSync.LatestURLPattern, &urlStoreSync.CreatedAt); err != nil {
		return nil, err
	}
	return &urlStoreSync, nil
}
