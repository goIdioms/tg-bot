package database

import (
	"database/sql"
	"log"
	"os"
	"telegram-golang-tasks-bot/pck/database/repository"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func NewRepository() *repository.Repository {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки файла .env:", err)
	}

	dsn := os.Getenv("DSN")
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}
	defer db.Close()

	repo := repository.NewRepository(db)

	err = db.Ping()
	if err != nil {
		db.Close()
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	log.Println("Успешное подключение к PostgreSQL!")

	return repo
}
