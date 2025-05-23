package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func SendSoftSkillsMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	message := `
	Команды:

/teamwork_questions – Как отвечать на вопросы о работе в команде
/conflict_questions – Вопросы про конфликты и способы их решения
/strengths_weaknesses – Как говорить о сильных и слабых сторонах
/stress_questions – Примеры ответов на стресс-вопросы
/motivation_questions – Вопросы про мотивацию и карьерные цели
/leadership_questions – Проверка лидерских качеств
/questions_to_employer – Что спросить у работодателя в конце собеседования
/case_questions – Разбор кейсов и ситуационных вопросов

/self_presentation_tips – Советы по самопрезентации
/interview_etiquette – Как правильно вести себя на собеседовании
/common_mistakes – Топ ошибок на собеседовании
/salary_talk – Как обсуждать зарплатные ожидания
	`

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	bot.Send(msg)
}
