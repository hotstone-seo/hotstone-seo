package server

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/hotstone-seo/hotstone-seo/server/service"

	"github.com/robfig/cron/v3"
	"go.uber.org/dig"
)

// TaskManager responsible manage the task
type taskManager struct {
	dig.In
	service.URLService
}

// Start the task
func startTaskManager(m taskManager) (err error) {
	c := cron.New()

	if _, err = c.AddFunc("* * * * *", func() {
		if err := m.URLService.Sync(context.Background()); err != nil {
			log.Fatal(err.Error())
		}
	}); err != nil {
		return
	}
	c.Start()
	return
}
