package repository

import (
	"database/sql"
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

func GetUserState(chatID int64) (models.UserState, bool) {
	query := `SELECT chat_id, step, task_question, task_answer, task_level, message_id FROM user_states WHERE chat_id = $1`
	state := DB.QueryRow(query, chatID)

	var s models.UserState
	err := state.Scan(&s.ChatID, &s.Step, &s.Task.Question, &s.Task.Answer, &s.Task.Level, &s.MessageID)
	if err == sql.ErrNoRows {
		return models.UserState{}, false
	}
	if err != nil {
		log.Printf("Ошибка при получении состояния пользователя: %v", err)
		return models.UserState{}, false
	}
	return s, true
}

func AddTask(task models.Task) error {
	query := `INSERT INTO tasks (question, answer, level) VALUES ($1, $2, $3)`

	_, err := DB.Exec(query, task.Question, task.Answer, task.Level)

	return err
}

func ClearUserState(chatID int64) error {
	query := `DELETE FROM user_states WHERE chat_id = $1`
	_, err := DB.Exec(query, chatID)
	return err
}

func GetEasyTask() (models.Task, bool) {
	query := `SELECT id, question, answer, level FROM tasks WHERE level = $1 ORDER BY RANDOM() LIMIT 1`

	var task models.Task
	err := DB.QueryRow(query, models.EasyLevel).Scan(&task.ID, &task.Question, &task.Answer, &task.Level)

	if err == sql.ErrNoRows {
		return models.Task{}, false
	}

	return task, true
}

func GetAnswerByTaskID(taskID int64) (models.Task, bool) {
	query := `SELECT id, question, answer, level FROM tasks WHERE id = $1`

	var task models.Task
	err := DB.QueryRow(query, taskID).Scan(&task.ID, &task.Question, &task.Answer, &task.Level)
	if err == sql.ErrNoRows {
		return models.Task{}, false
	}
	return task, true
}
