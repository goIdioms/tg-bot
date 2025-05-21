package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func InitDB() (*sql.DB, error) {
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		if err := godotenv.Load(); err != nil {
			log.Println("Не удалось загрузить .env (продакшен):", err)
		}
	}

	dsn := os.Getenv("DATABASE_URL")
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	log.Println("Успешное подключение к PostgreSQL!")
	return db, err
}
