package ports

import (
	"github.com/maneulf/guarapo_lab_test/internal/handlers/models/req"
)

type TasksRepository interface {
	GetTasks(token string) ([]req.Task, error)
	GetTask(id int, token string) (req.Task, error)
	Save(task req.Task, token string) error
	Update(task req.Task, id int, token string) error
	Delete(id int, token string) error
}
