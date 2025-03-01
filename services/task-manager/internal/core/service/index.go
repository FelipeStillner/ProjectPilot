package core

import "github.com/FelipeStillner/ProjectPilot/services/task-manager/internal/port"

type TaskService struct {
	taskRepo port.TaskRepositoryInterface
}

func NewTaskService(taksRepo port.TaskRepositoryInterface) *TaskService {
	return &TaskService{taksRepo}
}
