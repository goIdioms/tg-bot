package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func SendStartMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	message := `Привет!
	Я бот созданный для помощи подготовки к собеседованию по бэкэнд разработке.
Команды:

/tasks - решать задачи
/theory - теоретические вопросы
/soft_skills - популярные вопросы на собеседовании

/notify - создать напоминание о задаче (в разработке)
/statistics - статистика по решенным задачам (в разработке)

/new_item - для разработчиков
`

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	bot.Send(msg)
}
