package main

import (
	"log"

	a "github.com/FelipeStillner/ProjectPilot/services/task-manager/internal/adapter"
	c "github.com/FelipeStillner/ProjectPilot/services/task-manager/internal/core/service"
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
	grpcController := a.NewGrpcController(*taskService)

	go restController.Run()
	go grpcController.Run()

	select {}
}
