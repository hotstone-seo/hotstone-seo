package urlstore

import (
	"context"

	"go.uber.org/dig"
)

// URLStoreSyncService contain logic for URLStoreSync Controller
type URLStoreSyncService interface {
	FindOne(ctx context.Context, id int64) (*URLStoreSync, error)
	Find(ctx context.Context) ([]*URLStoreSync, error)
	Insert(ctx context.Context, URLStoreSync URLStoreSync) (lastInsertID int64, err error)
	GetLatestVersion(ctx context.Context) (latestVersion int64, err error)
	GetListDiff(ctx context.Context, offsetVersion int64) ([]*URLStoreSync, error)
}

// URLStoreSyncServiceImpl is implementation of URLStoreSyncService
type URLStoreSyncServiceImpl struct {
	dig.In
	URLStoreSyncRepo
}

// NewURLStoreSyncService return new instance of URLStoreSyncService
func NewURLStoreSyncService(impl URLStoreSyncServiceImpl) URLStoreSyncService {
	return &impl
}
