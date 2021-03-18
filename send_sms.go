package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sfreiberg/gotwilio"
)

func loadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
func sendSMS() {
	loadEnvVars()

	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := os.Getenv("NUMBER_FROM")
	to := os.Getenv("NUMBER_TO")
	message := "Time to meditate, son."

	twilio.SendSMS(from, to, message, "", "")
}
