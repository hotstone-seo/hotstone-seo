package repository

import (
	"context"
	"database/sql"
	"time"
)

type URLStoreSync struct {
	Version          int64     `json:"version"`
	Operation        string    `json:"operation" validate:"required"`
	RuleID           int64     `json:"rule_id"`
	LatestURLPattern string    `json:"latest_url_pattern"`
	CreatedAt        time.Time `json:"-"`
}

type URLStoreSyncRepo interface {
	Find(ctx context.Context, tx *sql.Tx, id int64) (*URLStoreSync, error)
	List(ctx context.Context, tx *sql.Tx) ([]*URLStoreSync, error)
	Insert(ctx context.Context, tx *sql.Tx, URLStoreSync URLStoreSync) (lastInsertID int64, err error)
	GetLatestVersion(ctx context.Context, tx *sql.Tx) (latestVersion int64, err error)
	GetListDiff(ctx context.Context, tx *sql.Tx, offsetVersion int64) ([]*URLStoreSync, error)
	DB() *sql.DB
}

// NewURLStoreSyncRepo return new instance of URLStoreSyncRepo
func NewURLStoreSyncRepo(impl URLStoreSyncRepoImpl) URLStoreSyncRepo {
	return &impl
}
