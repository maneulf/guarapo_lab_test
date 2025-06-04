package services

import (
	"log"

	"github.com/maneulf/guarapo_lab_test/internal/handlers/models/req"
	"github.com/maneulf/guarapo_lab_test/internal/ports"
)

type TasksService struct {
	tasksRepository ports.TasksRepository
}

func NewTasksService(taskRepository ports.TasksRepository) *TasksService {
	return &TasksService{
		tasksRepository: taskRepository,
	}
}

func (ts *TasksService) GetTasks(token string) ([]req.Task, error) {
	tasks, err := ts.tasksRepository.GetTasks(token)

	if err != nil {
		log.Printf("Error trying to get tasks, Err: %s", err)
		return []req.Task{}, err
	}

	return tasks, nil

}

func (ts *TasksService) GetTask(id int, token string) (req.Task, error) {
	task, err := ts.tasksRepository.GetTask(id, token)

	if err != nil {
		log.Printf("Error trying to get task, Error: %s", err.Error())
		return req.Task{}, err

	}
	log.Printf("Getting task successfuly")
	return task, nil
}

func (ts *TasksService) Save(task req.Task, token string) error {
	err := ts.tasksRepository.Save(task, token)

	if err != nil {
		log.Printf("Error trying to save task, Error: %s", err.Error())
		return err
	}

	return nil

}

func (ts *TasksService) Update(task req.Task, id int, token string) error {
	err := ts.tasksRepository.Update(task, id, token)
	if err != nil {
		log.Printf("Error tryinig to update task, Err: %s", err)
		return err
	}
	return nil

}

func (ts *TasksService) Delete(id int, token string) error {

	err := ts.tasksRepository.Delete(id, token)
	if err != nil {
		log.Printf("Error tryinig to delete task, Err: %s", err)
		return err
	}
	return nil
}
