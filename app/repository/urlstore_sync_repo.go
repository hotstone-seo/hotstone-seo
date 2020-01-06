package repository

import (
	"context"
	"time"
)

type URLStoreSync struct {
	Version          int64     `json:"version"`
	Operation        string    `json:"operation" validate:"required"`
	RuleID           int64     `json:"rule_id"`
	LatestURLPattern *string   `json:"latest_url_pattern"`
	CreatedAt        time.Time `json:"-"`
}

// URLStoreSyncRepo
type URLStoreSyncRepo interface {
	FindOne(ctx context.Context, id int64) (*URLStoreSync, error)
	Find(ctx context.Context) ([]*URLStoreSync, error)
	Insert(ctx context.Context, URLStoreSync URLStoreSync) (lastInsertID int64, err error)
	GetLatestVersion(ctx context.Context) (latestVersion int64, err error)
	GetListDiff(ctx context.Context, offsetVersion int64) ([]*URLStoreSync, error)
}

// NewURLStoreSyncRepo return new instance of URLStoreSyncRepo
func NewURLStoreSyncRepo(impl URLStoreSyncRepoImpl) URLStoreSyncRepo {
	return &impl
}
