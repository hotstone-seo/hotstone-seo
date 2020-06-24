package urlstore

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/dig"

	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
)

type (
	// Sync is database entity of sync information between server and url-store
	Sync struct {
		Version          int64     `json:"version"`
		Operation        string    `json:"operation" validate:"required"`
		RuleID           int64     `json:"rule_id"`
		LatestURLPattern *string   `json:"latest_url_pattern"`
		CreatedAt        time.Time `json:"-"`
	}

	// SyncRepo is repository of Sync entity
	// @mock
	SyncRepo interface {
		FindOne(ctx context.Context, id int64) (*Sync, error)
		Find(ctx context.Context) ([]*Sync, error)
		Insert(ctx context.Context, Sync Sync) (lastInsertID int64, err error)
		GetLatestVersion(ctx context.Context) (latestVersion int64, err error)
		GetListDiff(ctx context.Context, offsetVersion int64) ([]*Sync, error)
		FindRule(ctx context.Context, ruleID int64) (*Sync, error)
	}

	// SyncRepoImpl is implementation urlStoreSync repository
	SyncRepoImpl struct {
		dig.In
		*sql.DB
	}
)

// NewSyncRepo return new instance of SyncRepo
// @ctor
func NewSyncRepo(impl SyncRepoImpl) SyncRepo {
	return &impl
}

// FindOne urlStoreSync
func (r *SyncRepoImpl) FindOne(ctx context.Context, version int64) (urlStoreSync *Sync, err error) {
	var rows *sql.Rows
	builder := sq.
		Select(
			"version",
			"operation",
			"rule_id",
			"latest_url_pattern",
			"created_at",
		).
		From("url_sync").
		Where(sq.Eq{"version": version}).
		PlaceholderFormat(sq.Dollar).
		RunWith(r)
	if rows, err = builder.QueryContext(ctx); err != nil {
		return
	}
	defer rows.Close()
	if rows.Next() {
		return scanSync(rows)
	}
	return
}

// Find urlStoreSync
func (r *SyncRepoImpl) Find(ctx context.Context) (list []*Sync, err error) {
	var rows *sql.Rows
	builder := sq.
		Select("version", "operation", "rule_id", "latest_url_pattern", "created_at").
		From("url_sync").
		OrderBy("version").
		PlaceholderFormat(sq.Dollar).
		RunWith(r)
	if rows, err = builder.QueryContext(ctx); err != nil {
		return
	}
	defer rows.Close()
	list = make([]*Sync, 0)
	for rows.Next() {
		var urlStoreSync *Sync
		if urlStoreSync, err = scanSync(rows); err != nil {
			return
		}
		list = append(list, urlStoreSync)
	}
	return
}

// Insert urlStoreSync
func (r *SyncRepoImpl) Insert(ctx context.Context, urlStoreSync Sync) (lastInsertID int64, err error) {
	txn, err := dbtxn.Use(ctx, r.DB)
	if err != nil {
		return
	}
	query := sq.
		Insert("url_sync").
		Columns("operation", "rule_id", "latest_url_pattern").
		Values(urlStoreSync.Operation, urlStoreSync.RuleID, urlStoreSync.LatestURLPattern).
		Suffix("RETURNING \"version\"").
		PlaceholderFormat(sq.Dollar).
		RunWith(txn.DB())
	if err = query.QueryRowContext(ctx).Scan(&urlStoreSync.Version); err != nil {
		txn.SetError(err)
		return
	}
	lastInsertID = urlStoreSync.Version
	return
}

// GetLatestVersion of url store
func (r *SyncRepoImpl) GetLatestVersion(ctx context.Context) (int64, error) {
	var latestVersion int64
	builder := sq.
		Select("version").
		From("url_sync").
		OrderBy("version DESC").
		Limit(1).
		PlaceholderFormat(sq.Dollar).
		RunWith(r)

	err := builder.QueryRowContext(ctx).Scan(&latestVersion)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return -1, err
	}

	return latestVersion, nil
}

func (r *SyncRepoImpl) GetListDiff(ctx context.Context, offsetVersion int64) (list []*Sync, err error) {

	builder := sq.
		Select("version", "operation", "rule_id", "latest_url_pattern", "created_at").
		From("url_sync").
		Where(sq.Gt{"version": offsetVersion}).
		OrderBy("version").
		PlaceholderFormat(sq.Dollar).
		RunWith(r)

	rows, err := builder.QueryContext(ctx)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var urlStoreSync *Sync
		if urlStoreSync, err = scanSync(rows); err != nil {
			return
		}
		list = append(list, urlStoreSync)
	}
	return
}

func (r *SyncRepoImpl) FindRule(ctx context.Context, ruleID int64) (urlStoreSync *Sync, err error) {
	var rows *sql.Rows
	builder := sq.
		Select("version", "operation", "rule_id", "latest_url_pattern", "created_at").
		From("url_sync").
		Where(sq.Eq{"rule_id": ruleID}).
		OrderBy("version DESC").
		PlaceholderFormat(sq.Dollar).
		RunWith(r)
	if rows, err = builder.QueryContext(ctx); err != nil {
		return
	}
	defer rows.Close()

	if rows.Next() {
		return scanSync(rows)
	}
	return nil, nil
}

func scanSync(rows *sql.Rows) (*Sync, error) {
	var urlStoreSync Sync
	var err error
	if err = rows.Scan(&urlStoreSync.Version, &urlStoreSync.Operation, &urlStoreSync.RuleID, &urlStoreSync.LatestURLPattern, &urlStoreSync.CreatedAt); err != nil {
		return nil, err
	}
	return &urlStoreSync, nil
}
