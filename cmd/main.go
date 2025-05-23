package main

import (
	"log"
	"net/http"
	"os"
	tgbot "telegram-golang-tasks-bot/pck/bot"
	"telegram-golang-tasks-bot/pck/database/repository"
	"telegram-golang-tasks-bot/pck/handlers"
)

// for Railway
func startHTTPServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	go func() {
		log.Printf("Starting HTTP server on port %s", port)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Telegram bot is running!"))
		})
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Printf("HTTP server error: %v", err)
		}
	}()
}

func main() {
	startHTTPServer()

	bot, err := tgbot.InitBot()
	if err != nil {
		log.Fatal(err)
	}

	router := tgbot.NewRouter()
	router.Handle("start", handlers.SendStartMessage)
	router.Handle("menu", handlers.SendMenuMessage)

	router.Handle("soft_skills", handlers.SendSoftSkillsMessage)

	router.Handle("tasks", handlers.SendTasksMessage)
	router.Handle("easy", handlers.RandomEasyTask)
	router.Handle("medium", handlers.RandomMediumTask)
	router.Handle("hard", handlers.RandomHardTask)

	router.Handle("new_item ", handlers.SendNewItemMessage)
	router.Handle("add_task", handlers.StartTaskAddition)
	router.Handle("cancel_task", handlers.CancelTaskAddition)

	router.Handle("theory", handlers.SendTheoryMessage)
	router.Handle("theory_tasks", handlers.SendTheoryTasksMessage)
	router.Handle("theory_skills", handlers.SendTheorySkillsMessage)

	for update := range bot.UpdatesChannel {

		if update.CallbackQuery != nil && bot != nil && bot.API != nil {
			handlers.CallbackQuery(bot.API, update.CallbackQuery)
			continue
		}

		if update.Message == nil {
			continue
		}

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
	}
}
