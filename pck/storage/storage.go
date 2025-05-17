package storage

import (
	"math/rand"
	"sync"
	"telegram-golang-tasks-bot/pck/models"
)

type Storage struct {
	Tasks      []models.Task              // List of all tasks
	UserStates map[int64]models.UserState // User states in the task adding process
	NextID     int                        // Next available ID for a new task
	Mutex      sync.RWMutex               // Mutex for safe access to data
}

func NewStorage() *Storage {
	return &Storage{
		Tasks:      make([]models.Task, 0),
		UserStates: make(map[int64]models.UserState),
		NextID:     1,
	}
}

func (s *Storage) AddTask(task models.Task) int {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	task.ID = s.NextID
	s.NextID++
	s.Tasks = append(s.Tasks, task)

	return task.ID
}

func (s *Storage) GetRandomTaskByLevel(level string) (models.Task, bool) {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()

	var levelTasks []models.Task
	for _, task := range s.Tasks {
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
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.UserStates[chatID] = state
}

func (s *Storage) GetUserState(chatID int64) (models.UserState, bool) {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()
	state, exists := s.UserStates[chatID]
	return state, exists
}

func (s *Storage) ClearUserState(chatID int64) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	delete(s.UserStates, chatID)
}
