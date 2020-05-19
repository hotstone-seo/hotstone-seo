package urlstore

import (
	"context"
	"strconv"

	"go.uber.org/dig"
)

type (
	// Service contain logic of url
	Service struct {
		dig.In
		SyncRepo
		Store
		LatestVersion int
	}
)

// NewService return new instance of Service
// @constructor
func NewService(svc SyncRepo, store Store) *Service {
	return &Service{
		SyncRepo:      svc,
		Store:         store,
		LatestVersion: 0,
	}
}

// FullSync to sync from url-sync data to in-memory url-store from beginning
func (s *Service) FullSync(ctx context.Context) error {

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
	s.LatestVersion = int(oldestSync.Version)

	return nil
}

// Sync to  from url-sync data to in-memory url-store based on diff
func (s *Service) Sync(ctx context.Context) error {

	LatestVersionSync, err := s.GetLatestVersion(ctx)
	if err != nil {
		return err
	}

	if s.LatestVersion == int(LatestVersionSync) {
		return nil
	}

	if s.LatestVersion != 0 && LatestVersionSync == 0 {
		s.Store = NewStore()
		s.LatestVersion = int(LatestVersionSync)
		return nil
	}

	if s.LatestVersion > int(LatestVersionSync) {
		return s.FullSync(ctx)
	}

	if s.LatestVersion < int(LatestVersionSync) {
		listDiffSync, err := s.GetListDiff(ctx, int64(s.LatestVersion))
		if err != nil {
			return err
		}
		s.setStore(listDiffSync)

		oldestSync := listDiffSync[len(listDiffSync)-1]
		s.LatestVersion = int(oldestSync.Version)
	}

	return nil
}

// Insert to store
func (s *Service) Insert(id int64, key string) {
	data := strconv.FormatInt(id, 10)
	s.Store.Add(id, key, data)
}

// Update store
func (s *Service) Update(id int64, key string) {
	s.Delete(id)
	s.Insert(id, key)
}

func (s *Service) setStore(listSync []*Sync) {
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
