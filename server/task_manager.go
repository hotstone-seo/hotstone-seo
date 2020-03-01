package server

import (
	"github.com/hotstone-seo/hotstone-seo/server/urlstore"
	log "github.com/sirupsen/logrus"

	"github.com/robfig/cron/v3"
	"go.uber.org/dig"
)

// TaskManager responsible manage the task
type TaskManager struct {
	dig.In
	urlstore.URLService
}

// Start the task
func (m *TaskManager) Start() (err error) {
	c := cron.New()
	if _, err = c.AddFunc("* * * * *", task(m.URLService.Sync)); err != nil {
		return
	}
	c.Start()
	return
}

func task(fn func() error) func() {
	return func() {
		if err := fn(); err != nil {
			log.Warnf("Failed to sync url store: %+v", err)
		}
	}
}
