package adapter

import (
	"fmt"
	"net/http"
	"os"
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

func (c *RestController) Run() error {
	gin.SetMode(gin.ReleaseMode)
	r := c.createEngine()
	port := os.Getenv("PORT_HTTP_TASK_MANAGER")
	fmt.Println("Starting HTTP server on port " + port)
	return r.Run(":" + port)
}

func (c *RestController) createEngine() *gin.Engine {
	r := gin.Default()
	setRouteCreateTask(r, c)
	setRouteReadTask(r, c)
	setRouteUpdateTask(r, c)
	setRouteDeleteTask(r, c)
	return r
}

func setRouteCreateTask(r *gin.Engine, c *RestController) {
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
}

func setRouteReadTask(r *gin.Engine, c *RestController) {
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
}

func setRouteUpdateTask(r *gin.Engine, c *RestController) {
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
}

func setRouteDeleteTask(r *gin.Engine, c *RestController) {
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
}
