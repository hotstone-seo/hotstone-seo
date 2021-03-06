package urlstore

import (
	"context"
	"strconv"

	"go.uber.org/dig"
)

type (

	// Service contain logic for url store api
	// @mock
	Service interface {
		FullSync(ctx context.Context) error
		Sync(ctx context.Context) error
		Insert(id int64, key string)
		Update(id int64, key string)
	}

	// ServiceImpl contain logic of url
	ServiceImpl struct {
		dig.In
		SyncRepo
		Store
		Version int64
	}
)

// NewService return new instance of ServiceImpl
// @ctor
func NewService(svc SyncRepo, store Store) Service {
	return &ServiceImpl{
		SyncRepo: svc,
		Store:    store,
		Version:  0,
	}
}

// FullSync to sync from url-sync data to in-memory url-store from beginning
func (s *ServiceImpl) FullSync(ctx context.Context) error {

	list, err := s.Find(ctx)
	if err != nil {
		return err
	}

	if len(list) == 0 {
		return nil
	}

	s.Reset()
	s.setStore(list)

	oldestSync := list[len(list)-1]
	s.Version = oldestSync.Version

	return nil
}

// Sync to  from url-sync data to in-memory url-store based on diff
func (s *ServiceImpl) Sync(ctx context.Context) error {

	latestVersion, err := s.GetLatestVersion(ctx)
	if err != nil {
		return err
	}

	if s.Version == latestVersion {
		return nil
	}

	if s.Version != 0 && latestVersion == 0 {
		s.Store.Reset()
		s.Version = latestVersion
		return nil
	}

	if s.Version > latestVersion {
		return s.FullSync(ctx)
	}

	if s.Version < latestVersion {
		listDiffSync, err := s.GetListDiff(ctx, s.Version)
		if err != nil {
			return err
		}
		s.setStore(listDiffSync)

		oldestSync := listDiffSync[len(listDiffSync)-1]
		s.Version = oldestSync.Version
	}

	return nil
}

// Insert to store
func (s *ServiceImpl) Insert(id int64, key string) {
	data := strconv.FormatInt(id, 10)
	s.Store.Add(id, key, data)
}

// Update store
func (s *ServiceImpl) Update(id int64, key string) {
	s.Delete(id)
	s.Insert(id, key)
}

func (s *ServiceImpl) setStore(listSync []*Sync) {
	for _, sync := range listSync {
		switch sync.Operation {
		case "INSERT":
			s.Insert(sync.RuleID, *sync.LatestURLPattern)
		case "UPDATE":
			s.Update(sync.RuleID, *sync.LatestURLPattern)
		case "DELETE":
			s.Store.Delete(sync.RuleID)
		}
	}
}
