package core

import (
	"errors"
)

type DeleteTaskInput struct {
	Id uint32 `json:"id"`
}

func (t *TaskService) DeleteTask(input DeleteTaskInput) error {
	err := input.verify()
	if err != nil {
		return err
	}

	return t.taskRepo.Delete(input.Id)
}

func (input DeleteTaskInput) verify() error {
	if input.Id == 0 {
		return errors.New("task id is required")
	}

	return nil
}
