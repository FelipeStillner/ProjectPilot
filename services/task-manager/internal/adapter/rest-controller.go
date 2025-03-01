package adapter

import (
	"net/http"

	core "github.com/FelipeStillner/ProjectPilot/services/task-manager/internal/core/services"
	"github.com/gin-gonic/gin"
)

type RestController struct {
	taskService core.TaskService
}

func NewRestController(taskService core.TaskService) *RestController {
	return &RestController{taskService: taskService}
}

func (c *RestController) SetRoutes() *gin.Engine {
	r := gin.Default()
	r.POST("/create-task", func(ctx *gin.Context) {
		var requestBody struct {
			Name string `json:"name"`
		}
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := c.taskService.CreateTask(requestBody.Name); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "task created"})
	})
	return r
}
