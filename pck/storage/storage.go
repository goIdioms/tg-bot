package storage

import (
	"math/rand"
	"sync"
	"telegram-golang-tasks-bot/pck/models"
)

type Storage struct {
	tasks      []models.Task              // List of all tasks
	userStates map[int64]models.UserState // User states in the task adding process
	nextID     int                        // Next available ID for a new task
	mutex      sync.RWMutex               // Mutex for safe access to data
}

func NewStorage() *Storage {
	return &Storage{
		tasks:      make([]models.Task, 0),
		userStates: make(map[int64]models.UserState),
		nextID:     1,
	}
}

func (s *Storage) AddTask(task models.Task) int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	task.ID = s.nextID
	s.nextID++
	s.tasks = append(s.tasks, task)

	return task.ID
}

func (s *Storage) GetRandomTaskByLevel(level string) (models.Task, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var levelTasks []models.Task
	for _, task := range s.tasks {
		if task.Level == level {
			levelTasks = append(levelTasks, task)
		}
	}

	if len(levelTasks) == 0 {
		return models.Task{}, false
	}

	randomIndex := rand.Intn(len(levelTasks))
	return levelTasks[randomIndex], true
}

func (s *Storage) SetUserState(chatID int64, state models.UserState) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.userStates[chatID] = state
}

func (s *Storage) GetUserState(chatID int64) (models.UserState, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	state, exists := s.userStates[chatID]
	return state, exists
}

func (s *Storage) ClearUserState(chatID int64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.userStates, chatID)
}
