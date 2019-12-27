package urlstore

import (
	"context"

	"github.com/hotstone-seo/hotstone-server/app/repository"
	"github.com/hotstone-seo/hotstone-server/app/service"
	"go.uber.org/dig"
)

// URLStoreServer
type URLStoreServer interface {
	FullSync() error
	Sync() error

	Match(url string) (int, VarMap)
}

func NewURLStoreServer(svc service.URLStoreSyncService) URLStoreServer {
	return &URLStoreServerImpl{
		URLStoreSyncService: svc,
		urlStore:            InitURLStore(),
		latestVersion:       0,
	}
}

type URLStoreServerImpl struct {
	dig.In
	URLStoreSyncService service.URLStoreSyncService

	urlStore      URLStore
	latestVersion int
}

func (s *URLStoreServerImpl) FullSync() error {

	list, err := s.URLStoreSyncService.List(context.Background())
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

	s.urlStore = newURLStore
	s.latestVersion = int(oldestURLStoreSync.Version)

	return nil
}

func (s *URLStoreServerImpl) Sync() error {
	ctx := context.Background()

	latestVersionSync, err := s.URLStoreSyncService.GetLatestVersion(ctx)
	if err != nil {
		return err
	}

	if s.latestVersion == int(latestVersionSync) {
		return nil
	}

	if s.latestVersion != 0 && latestVersionSync == 0 {
		s.urlStore = InitURLStore()
		s.latestVersion = int(latestVersionSync)
		return nil
	}

	if s.latestVersion > int(latestVersionSync) {
		return s.FullSync()
	}

	if s.latestVersion < int(latestVersionSync) {
		listDiffURLStoreSync, err := s.URLStoreSyncService.GetListDiff(ctx, int64(s.latestVersion))
		if err != nil {
			return err
		}
		if err = s.buildURLStore(s.urlStore, listDiffURLStoreSync); err != nil {
			return err
		}

		oldestURLStoreSync := listDiffURLStoreSync[len(listDiffURLStoreSync)-1]
		s.latestVersion = int(oldestURLStoreSync.Version)
	}

	return nil
}

func (s *URLStoreServerImpl) Match(url string) (int, VarMap) {
	return s.urlStore.Get(url)
}

func (s *URLStoreServerImpl) buildURLStore(urlStore URLStore, listURLStoreSync []*repository.URLStoreSync) error {

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
