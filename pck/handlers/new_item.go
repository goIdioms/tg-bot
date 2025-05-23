package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func SendNewItemMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	message := `
Команды:

/add_task - Добавить новую задачу
/cancel_task - Отменить добавление задачи

/add_teory - Добавить новую теорию (в разработке)
/cancel_teory - Отменить добавление теории (в разработке)

/add_soft_skill - Добавить новую популярную тему на собеседовании (в разработке)
/cancel_soft_skill - Отменить добавление популярной темы (в разработке)

/menu - главное меню`

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	bot.Send(msg)
}
