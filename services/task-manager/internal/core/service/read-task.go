package core

import (
	"errors"

	core "github.com/FelipeStillner/ProjectPilot/services/task-manager/internal/core/struct"
)

type ReadTaskInput struct {
	Id uint32 `json:"id"`
}

func (t *TaskService) ReadTask(input ReadTaskInput) (*core.Task, error) {
	err := input.verify()
	if err != nil {
		return nil, err
	}

	return t.taskRepo.Read(input.Id)
}

func (input ReadTaskInput) verify() error {
	if input.Id == 0 {
		return errors.New("task id is required")
	}

	return nil
}
