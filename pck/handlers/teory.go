package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func SendTheoryMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	message := `
	Команды:
	/menu - главное меню

	/theory_tasks -  теоритические вопросы с ответами по бэкенду
	/theory_skills - теория по популярным технологиям
 `
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	bot.Send(msg)
}

func SendTheoryTasksMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	message := `
После нажатия на кнопку старта, вы будете получать вопросы.
По кнопке "следующий вопрос" вы будете получать следующий вопрос в рандомном порядке.
По кнопке "посмотреть решение" вы будете получать ответ на вопрос.

	Команды:
	/start_solve_teory -  начать
	/menu - главное меню`

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	bot.Send(msg)
}

func SendTheorySkillsMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	message := `
	Команды:
	/menu - главное меню

	/golang - теория по Golang
	/rest - теория по REST API
	/gRPC - теория по gRPC
	/docker - теория по Docker
	/kubernetes - теория по Kubernetes

	/postgres - теория по Postgres
	/mongo - теория по MongoDB
	/mysql - теория по MySQL
	/redis - теория по Redis
	/orm - теория по ORM

	/rabbitmq - теория по RabbitMQ
	/kafka - теория по Kafka

	/prometheus - теория по Prometheus
	/grafana - теория по Grafana

	/jaeger - теория по Jaeger
	/zipkin - теория по Zipkin
	/nginx - теория по Nginx
	/apache - теория по Apache

 `
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	bot.Send(msg)
}
