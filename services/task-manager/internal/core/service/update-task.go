package core

import (
	"errors"
	"time"

	core "github.com/FelipeStillner/ProjectPilot/services/task-manager/internal/core/struct"
)

type UpdateTaskInput struct {
	Id          uint32 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	Assignee    uint32 `json:"assignee"`
	Status      string `json:"status"`
}

func (t *TaskService) UpdateTask(input UpdateTaskInput) (*core.Task, error) {
	err := input.verify()
	if err != nil {
		return nil, err
	}

	task := core.Task{
		Id:          input.Id,
		Name:        input.Name,
		Description: input.Description,
		Priority:    input.Priority,
		Assignee:    input.Assignee,
		Status:      input.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return t.taskRepo.Update(task.Id, task)
}

func (input UpdateTaskInput) verify() error {
	if input.Id == 0 {
		return errors.New("task id is required")
	}

	if input.Name == "" {
		return errors.New("task name is required")
	}

	if input.Priority == "" {
		return errors.New("task priority is required")
	}

	if input.Status == "" {
		return errors.New("task status is required")
	}

	return nil
}
