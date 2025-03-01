package core

import (
	"errors"
	"time"

	core "github.com/FelipeStillner/ProjectPilot/services/task-manager/internal/core/struct"
	"github.com/google/uuid"
)

type CreateTaskInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	Assignee    string `json:"assignee"`
	Status      string `json:"status"`
}

func (t *TaskService) CreateTask(input CreateTaskInput) error {
	err := input.verify()
	if err != nil {
		return err
	}

	task := core.Task{
		Id:          uuid.New().ID(),
		Name:        input.Name,
		Description: input.Description,
		Priority:    input.Priority,
		Assignee:    input.Assignee,
		Status:      input.Status,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	return t.taskRepo.SaveTask(task)
}

func (input CreateTaskInput) verify() error {
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
