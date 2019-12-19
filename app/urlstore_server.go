package app

import (
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

func NewURLStoreServer(impl URLStoreServerImpl) URLStoreServer {
	return &impl
}

type URLStoreServerImpl struct {
	dig.In
	URLStoreSyncService service.URLStoreSyncService

	URLStoreTree  repository.URLStoreTree
	latestVersion int64
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
	// TODO: FullSync implementation
	// list, err := s.URLStoreSyncService.List()
	return nil
}

func (s *URLStoreServerImpl) Sync() error {
	return nil
}
