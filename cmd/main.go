package main

import (
	"log"
	"math/rand"
	"os"
	"telegram-golang-tasks-bot/pck/database"
	"telegram-golang-tasks-bot/pck/handlers"
	"telegram-golang-tasks-bot/pck/storage"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file")
	}

	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("TOKEN environment variable is not set")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Ошибка при создании бота: %v", err)
	}

	database.InitDB()

	bot.Debug = true
	log.Printf("Бот авторизован как %s", bot.Self.UserName)

	storage := storage.NewStorage()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("Ошибка при получении обновлений: %v", err)
	}

	handlers.HandleUpdates(bot, updates, storage)
}
