package task

import (
	log "github.com/sirupsen/logrus"

	"github.com/robfig/cron/v3"
)

// import "go.uber.org/dig"

func StartScheduler(sync func() error) error {
	c := cron.New()
	_, err := c.AddFunc("* * * * *", func() {
		err := sync()
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
