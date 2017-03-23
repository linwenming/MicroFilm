package services

import (
	"MicroFilm/db"
	"github.com/gocraft/dbr"
	"github.com/bamzi/jobrunner"
	"fmt"
)

func Init() {
	session := db.Init();
	tx,_ := session.Begin();


	tx.Rollback()

	tx.Commit()

	startup();
}

func startup(){
	jobrunner.Start() // optional: jobrunner.Start(pool int, concurrent int) (10, 1)
	jobrunner.Schedule("@every 5s", ReminderEmails{})
}

func loadTask(tx *dbr.Tx) {


}

func clearing() {

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