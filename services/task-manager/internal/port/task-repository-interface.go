package port

import c "github.com/FelipeStillner/ProjectPilot/services/task-manager/internal/core/entities"

type TaskRepositoryInterface interface {
	SaveTask(task c.Task) error
}
