package main

import (
	"fmt"
	"strings"
	"telegram-golang-tasks-bot/pck/models"
	"telegram-golang-tasks-bot/pck/storage"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Константы для команд бота
const (
	CommandStart      = "/start"
	CommandAddTask    = "/add"
	CommandEasy       = "/easy"
	CommandMedium     = "/medium"
	CommandHard       = "/hard"
	CommandShowAnswer = "show_answer"
	CommandCancelAdd  = "/cancel"
)

func handleUpdates(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel, storage *storage.Storage) {
	for update := range updates {
		if update.CallbackQuery != nil {
			handleCallbackQuery(bot, update.CallbackQuery, storage)
			continue
		}

		if update.Message == nil {
			continue
		}

		handleMessage(bot, update.Message, storage)
	}
}

func handleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message, storage *storage.Storage) {
	userState, exists := storage.GetUserState(message.Chat.ID)

	if exists && userState.Step > 0 {
		handleTaskAdditionProcess(bot, message, storage, userState)
		return
	}

	switch message.Text {
	case CommandStart:
		sendStartMessage(bot, message.Chat.ID)
	case CommandAddTask:
		startTaskAddition(bot, message.Chat.ID, storage)
	case CommandEasy:
		sendRandomTask(bot, message.Chat.ID, models.EasyLevel, storage)
	case CommandMedium:
		sendRandomTask(bot, message.Chat.ID, models.MediumLevel, storage)
	case CommandHard:
		sendRandomTask(bot, message.Chat.ID, models.HardLevel, storage)
	case CommandCancelAdd:
		cancelTaskAddition(bot, message.Chat.ID, storage)
	default:
		sendHelpMessage(bot, message.Chat.ID)
	}
}

func handleCallbackQuery(bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery, st *storage.Storage) {
	data := callbackQuery.Data

	if strings.HasPrefix(data, CommandShowAnswer) {
		parts := strings.Split(data, ":")
		if len(parts) != 2 {
			return
		}

		taskID := parts[1]
		messageText := callbackQuery.Message.Text
		lines := strings.Split(messageText, "\n")
		if len(lines) < 3 {
			return
		}

		answerLine := lines[len(lines)-1]
		if strings.HasPrefix(answerLine, "Ответ: ```go") {
			return
		}

		for _, task := range st.tasks {
			if fmt.Sprintf("%d", task.ID) == taskID {
				newText := messageText + "\n\nОтвет: ```go\n" + task.Answer + "\n```"

				edit := tgbotapi.NewEditMessageText(
					callbackQuery.Message.Chat.ID,
					callbackQuery.Message.MessageID,
					newText,
				)
				edit.ParseMode = "Markdown"
				edit.ReplyMarkup = &tgbotapi.InlineKeyboardMarkup{}

				bot.Send(edit)
				break
			}
		}
	}

	callback := tgbotapi.NewCallback(callbackQuery.ID, "")
	bot.AnswerCallbackQuery(callback)
}

func handleTaskAdditionProcess(bot *tgbotapi.BotAPI, message *tgbotapi.Message, storage *storage.Storage, state models.UserState) {
	switch state.Step {
	case 1:
		state.Task.Question = message.Text
		state.Step = 2
		storage.SetUserState(message.Chat.ID, state)
		msg := tgbotapi.NewMessage(message.Chat.ID, "Отлично! Теперь введите ответ на задачу в формате кода Golang:")
		bot.Send(msg)
	case 2:
		state.Task.Answer = message.Text
		state.Step = 3
		storage.SetUserState(message.Chat.ID, state)
		msg := tgbotapi.NewMessage(message.Chat.ID, "Теперь выберите уровень сложности задачи:")
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
		level := strings.ToLower(message.Text)
		if level != models.EasyLevel && level != models.MediumLevel && level != models.HardLevel {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Пожалуйста, выберите уровень сложности, используя кнопки ниже.")
			bot.Send(msg)
			return
		}
		state.Task.Level = level
		storage.AddTask(state.Task)
		storage.ClearUserState(message.Chat.ID)
		msg := tgbotapi.NewMessage(message.Chat.ID, "Задача успешно добавлена!")
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		bot.Send(msg)
	}
}

func sendStartMessage(bot *tgbotapi.BotAPI, chatID int64) {
	message := `Привет! Я бот для задач по Golang.
Команды:
/add - Добавить новую задачу
/easy - Получить случайную легкую задачу
/medium - Получить случайную задачу средней сложности
/hard - Получить случайную сложную задачу
/cancel - Отменить добавление задачи`
	msg := tgbotapi.NewMessage(chatID, message)
	bot.Send(msg)
}

func sendHelpMessage(bot *tgbotapi.BotAPI, chatID int64) {
	message := `Доступные команды:
/add - Добавить новую задачу
/easy - Получить случайную легкую задачу
/medium - Получить случайную задачу средней сложности
/hard - Получить случайную сложную задачу
/cancel - Отменить добавление задачи`
	msg := tgbotapi.NewMessage(chatID, message)
	bot.Send(msg)
}

func startTaskAddition(bot *tgbotapi.BotAPI, chatID int64, storage *storage.Storage) {
	state := models.UserState{Step: 1, Task: models.Task{}}
	storage.SetUserState(chatID, state)
	msg := tgbotapi.NewMessage(chatID, "Введите текст задачи:")
	bot.Send(msg)
}

func cancelTaskAddition(bot *tgbotapi.BotAPI, chatID int64, storage *storage.Storage) {
	msg := tgbotapi.NewMessage(chatID, "Добавление задачи отменено.")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.Send(msg)
}

func sendRandomTask(bot *tgbotapi.BotAPI, chatID int64, level string, storage *storage.Storage) {
	task, ok := storage.GetRandomTaskByLevel(level)
	if !ok {
		msg := tgbotapi.NewMessage(chatID, "Нет задач для выбранного уровня сложности.")
		bot.Send(msg)
		return
	}
	text := "Задача:\n```go\n" + task.Question + "\n```\nСложность: " + task.Level
	button := tgbotapi.NewInlineKeyboardButtonData("Показать ответ", CommandShowAnswer+":"+fmt.Sprintf("%d", task.ID))
	keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(button))
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "Markdown"
	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}
