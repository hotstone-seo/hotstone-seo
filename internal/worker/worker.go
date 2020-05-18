package worker

import (
	"context"
	"log"

	"github.com/hotstone-seo/hotstone-seo/server/service"
	"github.com/robfig/cron/v3"
	"go.uber.org/dig"
)

// Worker responsible manage the task
type Worker struct {
	dig.In
	service.URLService
}

// Start the task
func (m Worker) Start() (err error) {
	c := cron.New()

	ctx := context.Background()

	if err = m.URLService.Sync(ctx); err != nil {
		// FIXME:
		log.Fatal(err.Error())
		return
	}

	if _, err = c.AddFunc("* * * * *", func() {
		if err := m.URLService.Sync(ctx); err != nil {
			// FIXME:
			log.Fatal(err.Error())
		}
	}); err != nil {
		return
	}
	c.Start()
	return
}
