package application

import (
	"github.com/FelipeStillner/ProjectPilot/services/task-manager/repositories"
	"github.com/gin-gonic/gin"
)

func CreateTaskHandler(c *gin.Context) {
	db, err := repositories.ConnectDB()
	if err != nil {
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO tasks (name) VALUES ('New Task')")
	if err != nil {
		return
	}
}
