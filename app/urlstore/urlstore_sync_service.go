package urlstore

import (
	"go.uber.org/dig"
)

// URLStoreSyncService contain logic for URLStoreSync Controller
type URLStoreSyncService interface {
	URLStoreSyncRepo
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
