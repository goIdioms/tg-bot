package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func InitDB() {
	_ = godotenv.Load()
	dsn := os.Getenv("DSN")

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	log.Println("Успешное подключение к PostgreSQL!")
}
