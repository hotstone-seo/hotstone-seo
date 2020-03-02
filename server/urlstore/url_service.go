package urlstore

import (
	"context"

	"go.uber.org/dig"
)

// URLService contain logic of url [mock]
type URLService interface {
	FullSync() error
	Sync() error

	Match(url string) (int, map[string]string)
	DumpTree() string
}

// NewURLService return new instance of URLService [constructor]
func NewURLService(svc URLSyncRepo) URLService {
	return &URLServiceImpl{
		URLSyncRepo:   svc,
		URLStore:      InitURLStore(),
		LatestVersion: 0,
	}
}

// URLServiceImpl is implementation of URLService
type URLServiceImpl struct {
	dig.In
	URLSyncRepo
	URLStore
	LatestVersion int
}

func (s *URLServiceImpl) FullSync() error {

	list, err := s.Find(context.Background())
	if err != nil {
		return err
	}

	if len(list) == 0 {
		return nil
	}

	newURLStore := InitURLStore()
	if err = s.buildURLStore(newURLStore, list); err != nil {
		return err
	}

	oldestURLSync := list[len(list)-1]

	s.URLStore = newURLStore
	s.LatestVersion = int(oldestURLSync.Version)

	return nil
}

func (s *URLServiceImpl) Sync() error {
	ctx := context.Background()

	LatestVersionSync, err := s.GetLatestVersion(ctx)
	if err != nil {
		return err
	}

	if s.LatestVersion == int(LatestVersionSync) {
		return nil
	}

	if s.LatestVersion != 0 && LatestVersionSync == 0 {
		s.URLStore = InitURLStore()
		s.LatestVersion = int(LatestVersionSync)
		return nil
	}

	if s.LatestVersion > int(LatestVersionSync) {
		return s.FullSync()
	}

	if s.LatestVersion < int(LatestVersionSync) {
		listDiffURLSync, err := s.GetListDiff(ctx, int64(s.LatestVersion))
		if err != nil {
			return err
		}
		if err = s.buildURLStore(s.URLStore, listDiffURLSync); err != nil {
			return err
		}

		oldestURLSync := listDiffURLSync[len(listDiffURLSync)-1]
		s.LatestVersion = int(oldestURLSync.Version)
	}

	return nil
}

func (s *URLServiceImpl) Match(url string) (int, map[string]string) {
	return s.URLStore.Get(url)
}

func (s *URLServiceImpl) DumpTree() string {
	return s.URLStore.String()
}

func (s *URLServiceImpl) buildURLStore(urlStore URLStore, listURLSync []*URLSync) error {

	for _, URLSync := range listURLSync {
		switch URLSync.Operation {
		case "INSERT":
			urlStore.Add(int(URLSync.RuleID), *URLSync.LatestURLPattern)
		case "UPDATE":
			urlStore.Update(int(URLSync.RuleID), *URLSync.LatestURLPattern)
		case "DELETE":
			urlStore.Delete(int(URLSync.RuleID))
		}
	}

	return nil
}
