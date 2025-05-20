package main

import (
	"log"
	tgbot "telegram-golang-tasks-bot/pck/bot"
	"telegram-golang-tasks-bot/pck/handlers"
)

func main() {

	bot, err := tgbot.InitBot()
	if err != nil {
		log.Fatal(err)
	}

	router := tgbot.NewRouter()
	router.Handle("start", handlers.SendStartMessage)
	router.Handle("help", handlers.SendHelpMessage)
	router.Handle("add", handlers.StartTaskAddition)

	for update := range bot.UpdatesChannel {
		if update.Message != nil && update.Message.IsCommand() {
			router.Route(bot.API, update)
		} else {
			if update.CallbackQuery != nil {
				// handleCallbackQuery(bot, update.CallbackQuery)
			}
		}
	}
}
