package repository

import (
	"log"
	"telegram-golang-tasks-bot/pck/models"
)

func SetUserState(chatID int64, state models.UserState) error {
	query := `INSERT INTO user_states
	(chat_id, step, task_question, task_answer, task_level, message_id) VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (chat_id)
	DO UPDATE SET step = $2, task_question = $3, task_answer = $4, task_level = $5, message_id = $6`

	_, err := DB.Exec(query, chatID, state.Step, state.Task.Question, state.Task.Answer, state.Task.Level, state.MessageID)
	if err != nil {
		log.Printf("Ошибка при сохранении состояния пользователя: %v", err)
		return err
	}

	return nil
}
