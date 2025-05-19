package main

import (
	"log"
	tgbot "telegram-golang-tasks-bot/pck/bot"
	"telegram-golang-tasks-bot/pck/database"
	"telegram-golang-tasks-bot/pck/handlers"
)

func main() {
	database.InitDB()

	bot, err := tgbot.InitBot()
	if err != nil {
		log.Fatal(err)
	}

	router := tgbot.NewRouter()
	router.Handle("start", handlers.StartHandler)
	// router.Handle("help", handlers.HelpHandler)

	for update := range bot.UpdatesChannel {
		if update.Message != nil && update.Message.IsCommand() {
			router.Route(bot.API, update)
		} else {
			// обработать обычные сообщения или callback-запросы
		}
	}
}
