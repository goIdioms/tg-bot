package main

import (
	"log"
	tgbot "telegram-golang-tasks-bot/pck/bot"
	"telegram-golang-tasks-bot/pck/database/repository"
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
	router.Handle("cancel", handlers.CancelTaskAddition)

	for update := range bot.UpdatesChannel {
		chatID := update.Message.Chat.ID
		userState, exists := repository.GetUserState(chatID)
		if exists && userState.Step > 0 {
			if update.Message.Text == "/cancel" {
				handlers.CancelTaskAddition(bot.API, update)
				continue
			}
			handlers.HandleTaskAdditionProcess(bot.API, update, userState)
			continue
		}

		if update.Message.IsCommand() {
			router.Route(bot.API, update)
		}
		if update.CallbackQuery != nil {
			// handleCallbackQuery(bot, update.CallbackQuery)
		}

	}
}
