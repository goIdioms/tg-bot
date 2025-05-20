package bot

import (
	"telegram-golang-tasks-bot/pck/database/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type HandlerFunc func(bot *tgbotapi.BotAPI, update tgbotapi.Update)

type Router struct {
	handlers map[string]HandlerFunc
}

func (r *Router) Handle(command string, handler HandlerFunc) {
	r.handlers[command] = handler
}

func (r *Router) Route(bot *tgbotapi.BotAPI, update tgbotapi.Update, repo ...*repository.Repository) {
	command := update.Message.Command()
	if handler, ok := r.handlers[command]; ok {
		handler(bot, update)
	} else {
		// Обработка неизвестных команд
	}

}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]HandlerFunc),
	}
}
