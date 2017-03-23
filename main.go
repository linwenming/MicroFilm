package main

import (
	"github.com/Sirupsen/logrus"
	"MicroFilm/route"
	"github.com/bamzi/jobrunner"
	"fmt"
)

func init() {

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {

	jobrunner.Start() // optional: jobrunner.Start(pool int, concurrent int) (10, 1)
	jobrunner.Schedule("@every 5s", ReminderEmails{})

	router := route.Init()
	router.Start(":8888")

}

// Job Specific Functions
type ReminderEmails struct {
	// filtered
}

// ReminderEmails.Run() will get triggered automatically.
func (e ReminderEmails) Run() {
	// Queries the DB
	// Sends some email
	fmt.Printf("Every 5 sec send reminder emails \n")
}

