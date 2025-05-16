package main

const (
	EasyLevel   = "easy"
	MediumLevel = "medium"
	HardLevel   = "hard"
)

type Task struct {
	ID       int
	Question string
	Answer   string
	Level    string
}

type UserState struct {
	Step      int
	Task      Task
	ChatID    int64
	MessageID int
}
