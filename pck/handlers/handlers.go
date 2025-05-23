package handlers

import (
	"log"
	"strconv"
	"strings"
	"telegram-golang-tasks-bot/pck/database/repository"
	"telegram-golang-tasks-bot/pck/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendMenuMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	message := `
Главное меню:
/tasks - решать задачи
/teory - теоретические вопросы
/soft_skills - популярные вопросы на собеседовании

/new_item - для разработчиков`
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	bot.Send(msg)
}

func StartTaskAddition(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	state := models.UserState{Step: 1, Task: models.Task{}}

	err := repository.SetUserState(update.Message.Chat.ID, state)
	if err != nil {
		log.Printf("Ошибка при сохранении состояния пользователя: %v", err)
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите текст задачи:")
	bot.Send(msg)
}

func CancelTaskAddition(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	repository.ClearUserState(update.Message.Chat.ID)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Добавление задачи отменено.")

	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.Send(msg)
}

func RandomEasyTask(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	task, ok := repository.GetEasyTask()
	if !ok {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Нет задач для выбранного уровня сложности.")
		bot.Send(msg)
		return
	}

	text := "Задача:\n```go\n" + task.Question + "\n```\nСложность: " + task.Level

	button := tgbotapi.NewInlineKeyboardButtonData("Показать ответ", strconv.Itoa(task.ID))

	keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(button))
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ParseMode = "Markdown"
	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}

func RandomMediumTask(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	task, ok := repository.GetMediumTask()
	if !ok {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Нет задач для выбранного уровня сложности.")
		bot.Send(msg)
		return
	}

	text := "Задача:\n```go\n" + task.Question + "\n```\nСложность: " + task.Level

	button := tgbotapi.NewInlineKeyboardButtonData("Показать ответ", strconv.Itoa(task.ID))

	keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(button))
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ParseMode = "Markdown"
	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}

func RandomHardTask(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	task, ok := repository.GetHardTask()
	if !ok {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Нет задач для выбранного уровня сложности.")
		bot.Send(msg)
		return
	}

	text := "Задача:\n```go\n" + task.Question + "\n```\nСложность: " + task.Level

	button := tgbotapi.NewInlineKeyboardButtonData("Показать ответ", strconv.Itoa(task.ID))

	keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(button))
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ParseMode = "Markdown"
	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}

func CallbackQuery(bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery) {
	id := callbackQuery.Data

	taskID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("Error converting task ID: %v", err)
		return
	}

	task, ok := repository.GetAnswerByTaskID(taskID)
	if ok {

		newText := "Задача:\n```go\n" + task.Question + "\n```\nСложность: " + task.Level + "\n\nОтвет: ```go\n" + task.Answer + "\n```"

		editedMessage := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, newText)
		editedMessage.ParseMode = "Markdown"
		bot.Send(editedMessage)
	}

	callback := tgbotapi.NewCallback(callbackQuery.ID, "")
	bot.AnswerCallbackQuery(callback)
}

func HandleTaskAdditionProcess(bot *tgbotapi.BotAPI, update tgbotapi.Update, state models.UserState) {
	switch state.Step {
	case 1:
		state.Task.Question = update.Message.Text
		state.Step = 2

		err := repository.SetUserState(update.Message.Chat.ID, state)
		if err != nil {
			log.Printf("Ошибка при сохранении состояния пользователя: %v, state: %v", err, state)
			return
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отлично! Теперь введите ответ на задачу в формате кода Golang:")

		bot.Send(msg)
	case 2:
		state.Task.Answer = update.Message.Text
		state.Step = 3
		err := repository.SetUserState(update.Message.Chat.ID, state)
		if err != nil {
			log.Printf("Ошибка при сохранении состояния пользователя: %v, state: %v", err, state)
			return
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Теперь выберите уровень сложности задачи:")
		keyboard := tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton(models.EasyLevel),
				tgbotapi.NewKeyboardButton(models.MediumLevel),
				tgbotapi.NewKeyboardButton(models.HardLevel),
			),
		)
		keyboard.OneTimeKeyboard = true
		msg.ReplyMarkup = keyboard
		bot.Send(msg)
	case 3:
		level := strings.ToLower(update.Message.Text)
		if level != models.EasyLevel && level != models.MediumLevel && level != models.HardLevel {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, выберите уровень сложности, используя кнопки ниже.")
			bot.Send(msg)
			return
		}

		state.Task.Level = level
		err := repository.AddTask(state.Task)
		if err != nil {
			log.Printf("Ошибка при добавлении задачи: %v", err)
			return
		}

		err = repository.ClearUserState(update.Message.Chat.ID)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Задача успешно добавлена!")

		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		bot.Send(msg)
	}
}
