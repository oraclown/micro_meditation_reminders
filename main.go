package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
	"github.com/robfig/cron"
	"github.com/sfreiberg/gotwilio"
)

// Ready environment variables.
func loadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

// Send meditation reminder text using Twilio.
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

// Send meditation reminder texts every 50th min from 9a-8p using cronjobs.
func scheduleReminders() {
	c := cron.New()

	jobTime := "0 50 9-20 * * *"

	c.AddFunc(jobTime, func() { sendSMS() })
	c.AddFunc(jobTime, func() { fmt.Println("sending text") })

	c.Start()
}

// Route for app homepage.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles("./ui/html/home.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// Schedule reminder texts and serve a simple single page app.
func main() {
	go scheduleReminders()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
