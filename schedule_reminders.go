package main

import (
	"github.com/robfig/cron"
)

func scheduleReminders() {
	c := cron.New()
	reminderIntervals := formatWorkTimes()

	for _, hourInterval := range reminderIntervals {
		jobStr := "45 " + hourInterval + " * * *"
		c.AddFunc(jobStr, func() { sendSMS() })
	}

	c.Start()
}
