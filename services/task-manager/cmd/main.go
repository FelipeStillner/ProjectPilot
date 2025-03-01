package main

import (
	"log"
	"os"

	a "github.com/FelipeStillner/ProjectPilot/services/task-manager/internal/adapter"
	c "github.com/FelipeStillner/ProjectPilot/services/task-manager/internal/core/services"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func main() {
	taskRepo := a.NewTaskRepository()
	taskService := c.NewTaskService(taskRepo)
	restController := a.NewRestController(*taskService)
	r := restController.SetRoutes()
	port := os.Getenv("PORT_TASK_MANAGER")
	r.Run(":" + port)
}
