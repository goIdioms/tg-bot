package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func InitDB() (*sql.DB, error) {
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		if err := godotenv.Load(); err != nil {
			log.Println("Не удалось загрузить .env (локальная среда):", err)
		}
	} else {
		log.Println("Запуск в среде Railway")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("Переменная окружения DATABASE_URL не установлена")
	}

	log.Printf("Исходная строка подключения: %s", dsn)

	if os.Getenv("RAILWAY_ENVIRONMENT") != "" {
		if !strings.Contains(dsn, "sslmode=") {
			if strings.Contains(dsn, "?") {
				dsn += "&sslmode=require"
			} else {
				dsn += "?sslmode=require"
			}
		}

		if !strings.Contains(dsn, "host=") && strings.Contains(dsn, "postgresql://") {
			log.Println("Внимание: строка подключения может быть неполной. Проверьте настройки PostgreSQL в Railway.")
		}
	}

	if os.Getenv("RAILWAY_ENVIRONMENT") != "" {
		pgHost := os.Getenv("PGHOST")
		pgPort := os.Getenv("PGPORT")
		pgUser := os.Getenv("PGUSER")
		pgPassword := os.Getenv("PGPASSWORD")
		pgDatabase := os.Getenv("PGDATABASE")

		if pgHost != "" && pgUser != "" && pgDatabase != "" {
			log.Println("Используем переменные окружения Railway для подключения к PostgreSQL")

			if pgPort == "" {
				pgPort = "5432"
			}

			dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=require",
				pgUser, pgPassword, pgHost, pgPort, pgDatabase)
		} else if strings.HasPrefix(dsn, "postgresql://") {
			log.Println("Преобразование строки подключения для Railway...")
			dsn = strings.Replace(dsn, "postgresql://", "postgres://", 1)

			if !strings.Contains(dsn, "sslmode=") {
				if strings.Contains(dsn, "?") {
					dsn += "&sslmode=require"
				} else {
					dsn += "?sslmode=require"
				}
			}
		}
	}

	log.Printf("Итоговая строка подключения: %s", dsn)

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
