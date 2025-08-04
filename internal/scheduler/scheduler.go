package scheduler

import (
	"fmt"
	"time"

	"github.com/Raziur306/kon-dol/internal/worker"
	"github.com/robfig/cron/v3"
)

func StartScheduler() {
	c := cron.New()
	//for now it will run every minute
	c.AddFunc("@every 5m", func() {
		fmt.Println("Running scheduled task..." + time.Now().Format(time.RFC3339))
		worker.PullNewsList()
	})
	c.Start()

}
