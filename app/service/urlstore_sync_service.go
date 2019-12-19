package service

import (
	"context"
	"database/sql"

	"github.com/hotstone-seo/hotstone-server/app/repository"
	"go.uber.org/dig"
)

// URLStoreSyncService contain logic for URLStoreSync Controller
type URLStoreSyncService interface {
	// repository.URLStoreSyncRepo
	Find(ctx context.Context, id int64) (*repository.URLStoreSync, error)
	List(ctx context.Context) ([]*repository.URLStoreSync, error)
	Insert(ctx context.Context, urlStoreSync repository.URLStoreSync) (lastInsertID int64, err error)
}

// URLStoreSyncServiceImpl is implementation of URLStoreSyncService
type URLStoreSyncServiceImpl struct {
	dig.In
	URLStoreSyncRepo repository.URLStoreSyncRepo
}

// NewURLStoreSyncService return new instance of URLStoreSyncService
func NewURLStoreSyncService(impl URLStoreSyncServiceImpl) URLStoreSyncService {
	return &impl
}

// Find urlStoreSync
func (r *URLStoreSyncServiceImpl) Find(ctx context.Context, id int64) (urlStoreSync *repository.URLStoreSync, err error) {

	err = repository.WithTransaction(r.URLStoreSyncRepo.DB(), func(tx *sql.Tx) error {
		urlStoreSync, err = r.URLStoreSyncRepo.Find(ctx, tx, id)
		if err != nil {
			return err
		}

		return nil
	})

	return
}

// List urlStoreSync
func (r *URLStoreSyncServiceImpl) List(ctx context.Context) (list []*repository.URLStoreSync, err error) {
	err = repository.WithTransaction(r.URLStoreSyncRepo.DB(), func(tx *sql.Tx) error {
		list, err = r.URLStoreSyncRepo.List(ctx, tx)
		if err != nil {
			return err
		}

		return nil
	})

	return
}

// Insert urlStoreSync
func (r *URLStoreSyncServiceImpl) Insert(ctx context.Context, urlStoreSync repository.URLStoreSync) (lastInsertID int64, err error) {
	err = repository.WithTransaction(r.URLStoreSyncRepo.DB(), func(tx *sql.Tx) error {
		lastInsertID, err = r.URLStoreSyncRepo.Insert(ctx, tx, urlStoreSync)
		if err != nil {
			return err
		}

		return nil
	})

	return
}
