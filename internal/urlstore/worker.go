package urlstore

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/robfig/cron/v3"
	"go.uber.org/dig"
)

// Worker responsible manage the task
type Worker struct {
	dig.In
	Service
}

// Start the task
func (m Worker) Start() (err error) {
	log.Info("Start the worker")
	c := cron.New()

	ctx := context.Background()

	if err = m.Service.Sync(ctx); err != nil {
		log.Warn(err.Error())
		return
	}

	if _, err = c.AddFunc("* * * * *", func() {
		if err := m.Service.Sync(ctx); err != nil {
			log.Warn(err.Error())
		}
	}); err != nil {
		return
	}
	c.Start()
	return
}
