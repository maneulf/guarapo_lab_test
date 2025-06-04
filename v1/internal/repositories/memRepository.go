package repositories

import (
	"log"

	"github.com/maneulf/guarapo_lab_test/internal/handlers/models/req"
)

type MemTasksRepository struct {
	memRepo map[string][]req.Task
}

func NewMemTasksRepository() *MemTasksRepository {
	return &MemTasksRepository{
		memRepo: make(map[string][]req.Task),
	}
}

func (mt *MemTasksRepository) GetTasks(token string) ([]req.Task, error) {
	tasks := mt.memRepo[token]
	return tasks, nil
}

func (mt *MemTasksRepository) GetTask(id int, token string) (req.Task, error) {
	for _, t := range mt.memRepo[token] {
		if t.ID == id {
			return t, nil
		}
	}
	return req.Task{}, nil
}

func (mt *MemTasksRepository) Save(task req.Task, token string) error {
	mt.memRepo[token] = append(mt.memRepo[token], task)
	log.Println("Task save successfuly")
	return nil
}

func (mt *MemTasksRepository) Update(task req.Task, id int, token string) error {
	for i, t := range mt.memRepo[token] {
		if t.ID == id {
			mt.memRepo[token][i] = task
		}
	}
	return nil
}

func (mt *MemTasksRepository) Delete(id int, token string) error {
	for i, t := range mt.memRepo[token] {
		if t.ID == id {
			mt.memRepo[token] = remove(mt.memRepo[token], i)
		}
	}
	return nil
}

func remove(s []req.Task, i int) []req.Task {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
