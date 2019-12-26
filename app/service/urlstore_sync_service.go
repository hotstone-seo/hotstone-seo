package service

import (
	"github.com/hotstone-seo/hotstone-server/app/repository"
	"go.uber.org/dig"
)

// URLStoreSyncService contain logic for URLStoreSync Controller
type URLStoreSyncService interface {
	repository.URLStoreSyncRepo
}

// URLStoreSyncServiceImpl is implementation of URLStoreSyncService
type URLStoreSyncServiceImpl struct {
	dig.In
	repository.URLStoreSyncRepo
}

// NewURLStoreSyncService return new instance of URLStoreSyncService
func NewURLStoreSyncService(impl URLStoreSyncServiceImpl) URLStoreSyncService {
	return &impl
}
