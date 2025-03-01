package port

import c "github.com/FelipeStillner/ProjectPilot/services/task-manager/internal/core/struct"

type TaskRepositoryInterface interface {
	Create(task c.Task) (*c.Task, error)
	Read(id uint32) (*c.Task, error)
	Update(id uint32, task c.Task) (*c.Task, error)
	Delete(id uint32) error
}
