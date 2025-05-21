package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func InitDB() (*sql.DB, error) {
	// Загружаем .env только в локальной среде
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		if err := godotenv.Load(); err != nil {
			log.Println("Не удалось загрузить .env (локальная среда):", err)
		}
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("Переменная окружения DATABASE_URL не установлена")
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	for i := 0; i < 5; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		log.Printf("Не удалось подключиться к базе данных: %v. Попытка %d из 5", err, i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		db.Close()
		log.Fatal("Не удалось подключиться к базе данных после нескольких попыток:", err)
	}

	log.Println("Успешное подключение к PostgreSQL!")
	return db, nil
}
