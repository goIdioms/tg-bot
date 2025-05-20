package repository

import (
	"database/sql"
	"log"
	"telegram-golang-tasks-bot/pck/database"
)

var DB *sql.DB

func init() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных")
		panic(err)
	}
	DB = db
}
