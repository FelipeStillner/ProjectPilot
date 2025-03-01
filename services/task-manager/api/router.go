package api

import (
	"github.com/FelipeStillner/ProjectPilot/services/task-manager/application"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/create-task", application.CreateTaskHandler)
	return r
}
