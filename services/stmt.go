package services

import (
	"MicroFilm/db"
	"github.com/gocraft/dbr"
)

func Init() {
	session := db.Init();
	tx,_ := session.Begin();


	tx.Rollback()

	tx.Commit()
}

func startup(){

}

func loadTask(tx *dbr.Tx) {


}

func clearing() {

}

