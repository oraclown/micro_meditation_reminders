package main

import (
	"fmt"

	"github.com/robfig/cron"
)

func scheduleReminders() {
	c := cron.New()

	jobStr := "0 * * * * *"
	c.AddFunc(jobStr, func() { sendSMS() })
	c.AddFunc(jobStr, func() { fmt.Println("sending text") })
	// reminderIntervals := formatWorkTimes()

	// for _, hourInterval := range reminderIntervals {
	// 	fmt.Println("Scheduling reminders @ min 45 of each hour from " + hourInterval)
	// 	jobStr := "45 " + hourInterval + " * * *"
	// 	c.AddFunc(jobStr, func() { sendSMS() })
	// }

	c.Start()

	// c.Stop()
}
