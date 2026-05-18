package cron

import (
	"permit-license/internal/service"

	"github.com/robfig/cron/v3"
)

func StartCron() {

	c := cron.New()

	schedulerService := service.SchedulerService{}

	c.AddFunc(
		"0 8 * * *", // every day at 08.00
		func() {
			schedulerService.SendExpiredReminder()
			schedulerService.AutoMarkExpiredDocuments()
		},
	)

	c.Start()
}
