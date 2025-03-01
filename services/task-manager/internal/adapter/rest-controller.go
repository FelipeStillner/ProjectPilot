package adapter

import (
	"net/http"
	"strconv"

	core "github.com/FelipeStillner/ProjectPilot/services/task-manager/internal/core/service"
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
	r.POST("/task", func(ctx *gin.Context) {
		var requestBody core.CreateTaskInput
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		task, err := c.taskService.CreateTask(requestBody)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, task)
	})
	r.GET("/task/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		idUint, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		input := core.ReadTaskInput{Id: uint32(idUint)}
		task, err := c.taskService.ReadTask(input)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, task)
	})
	r.PATCH("/task", func(ctx *gin.Context) {
		var requestBody core.UpdateTaskInput
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		task, err := c.taskService.ReadTask(core.ReadTaskInput{Id: requestBody.Id})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, task)
	})
	r.DELETE("/task/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		idUint, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		input := core.DeleteTaskInput{Id: uint32(idUint)}
		err = c.taskService.DeleteTask(input)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "task deleted"})
	})
	return r
}
