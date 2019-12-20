package app

import (
	"context"

	"github.com/hotstone-seo/hotstone-server/app/repository"
	"github.com/hotstone-seo/hotstone-server/app/service"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type URLStoreServer interface {
	Start() error
	FullSync() error
	Sync() error
}

func NewURLStoreServer(svc service.URLStoreSyncService) URLStoreServer {
	return &URLStoreServerImpl{
		URLStoreSyncService: svc,
		urlStore:            repository.InitURLStore(),
		latestVersion:       -1,
	}
}

type URLStoreServerImpl struct {
	dig.In
	URLStoreSyncService service.URLStoreSyncService

	urlStore      repository.URLStore
	latestVersion int
}

func (s *URLStoreServerImpl) Start() error {
	if err := s.FullSync(); err != nil {
		return err
	}

	c := cron.New()
	_, err := c.AddFunc("* * * * *", func() {
		err := s.Sync()
		if err != nil {
			log.Warnf("Failed to sync url store: %+v", err)
		}
	})

	if err != nil {
		return err
	}

	c.Start()

	return nil
}

func (s *URLStoreServerImpl) FullSync() error {

	list, err := s.URLStoreSyncService.List(context.Background())
	if err != nil {
		return err
	}

	if len(list) == 0 {
		return nil
	}

	newURLStore := repository.InitURLStore()
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

	if s.latestVersion > int(latestVersionSync) {
		return s.FullSync()
	}

	if s.latestVersion < int(latestVersionSync) {
		// TODO: Sync implementation
		// diffChanges, err := s.URLStoreSyncService.GetListDiff(ctx, s.latestVersion)
	}

	return nil
}

func (s *URLStoreServerImpl) buildURLStore(urlStore repository.URLStore, listURLStoreSync []*repository.URLStoreSync) error {

	for _, urlStoreSync := range listURLStoreSync {
		switch urlStoreSync.Operation {
		case "INSERT":
			urlStore.AddURL(int(urlStoreSync.RuleID), urlStoreSync.LatestURLPattern)
		case "UPDATE":
			urlStore.UpdateURL(int(urlStoreSync.RuleID), urlStoreSync.LatestURLPattern)
		case "DELETE":
			urlStore.DeleteURL(int(urlStoreSync.RuleID))
		}
	}

	return nil
}
