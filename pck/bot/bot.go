package bot

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

type Bot struct {
	API            *tgbotapi.BotAPI
	UpdatesChannel tgbotapi.UpdatesChannel
}

func InitBot() (*Bot, error) {
	_ = godotenv.Load()
	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("TOKEN environment variable is not set")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Ошибка при создании бота: %v", err)
	}

	bot.Debug = true
	log.Printf("Бот авторизован как %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("Ошибка при получении обновлений: %v", err)
	}

	return &Bot{API: bot, UpdatesChannel: updates}, err
}
