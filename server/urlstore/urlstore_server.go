package urlstore

import (
	"context"

	"go.uber.org/dig"
)

// URLStoreServer
type URLStoreServer interface {
	FullSync() error
	Sync() error

	Match(url string) (int, map[string]string)
	DumpTree() string
}

// NewURLStoreServer return new instance of URLStoreServer [constructor]
func NewURLStoreServer(svc URLStoreSyncService) URLStoreServer {
	return &URLStoreServerImpl{
		URLStoreSyncService: svc,
		URLStore:            InitURLStore(),
		LatestVersion:       0,
	}
}

type URLStoreServerImpl struct {
	dig.In
	URLStoreSyncService
	URLStore
	LatestVersion int
}

func (s *URLStoreServerImpl) FullSync() error {

	list, err := s.URLStoreSyncService.Find(context.Background())
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

	oldestURLStoreSync := list[len(list)-1]

	s.URLStore = newURLStore
	s.LatestVersion = int(oldestURLStoreSync.Version)

	return nil
}

func (s *URLStoreServerImpl) Sync() error {
	ctx := context.Background()

	LatestVersionSync, err := s.URLStoreSyncService.GetLatestVersion(ctx)
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
		listDiffURLStoreSync, err := s.URLStoreSyncService.GetListDiff(ctx, int64(s.LatestVersion))
		if err != nil {
			return err
		}
		if err = s.buildURLStore(s.URLStore, listDiffURLStoreSync); err != nil {
			return err
		}

		oldestURLStoreSync := listDiffURLStoreSync[len(listDiffURLStoreSync)-1]
		s.LatestVersion = int(oldestURLStoreSync.Version)
	}

	return nil
}

func (s *URLStoreServerImpl) Match(url string) (int, map[string]string) {
	return s.URLStore.Get(url)
}

func (s *URLStoreServerImpl) DumpTree() string {
	return s.URLStore.String()
}

func (s *URLStoreServerImpl) buildURLStore(urlStore URLStore, listURLStoreSync []*URLStoreSync) error {

	for _, urlStoreSync := range listURLStoreSync {
		switch urlStoreSync.Operation {
		case "INSERT":
			urlStore.Add(int(urlStoreSync.RuleID), *urlStoreSync.LatestURLPattern)
		case "UPDATE":
			urlStore.Update(int(urlStoreSync.RuleID), *urlStoreSync.LatestURLPattern)
		case "DELETE":
			urlStore.Delete(int(urlStoreSync.RuleID))
		}
	}

	return nil
}
