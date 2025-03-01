package core

import (
	c "github.com/FelipeStillner/ProjectPilot/services/task-manager/internal/core/entities"
	"github.com/FelipeStillner/ProjectPilot/services/task-manager/internal/port"
)

type TaskService struct {
	taskRepo port.TaskRepositoryInterface
}

func NewTaskService(taksRepo port.TaskRepositoryInterface) *TaskService {
	return &TaskService{taksRepo}
}

func (t *TaskService) CreateTask(name string) error {
	task := c.NewTask(name)

	return t.taskRepo.SaveTask(task)
}
