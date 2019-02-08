package crontab

import (
	"log"
	"time"

	"github.com/robfig/cron"
	"goweb/pkg/logging"
)

func CronHandel() {
	log.Println("Starting...")

	c := cron.New()

	c.AddFunc("@every 1m", func() {
		logging.Info("corn is running...")
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
