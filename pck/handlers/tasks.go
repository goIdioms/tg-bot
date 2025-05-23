package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func SendTasksMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	message := `
	Команды:
	/easy - Получить случайную легкую задачу
	/medium - Получить случайную задачу средней сложности
	/hard - Получить случайную сложную задачу

	/menu - главное меню
`
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	bot.Send(msg)
}
