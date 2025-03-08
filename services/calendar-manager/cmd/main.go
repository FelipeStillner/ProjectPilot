package main

import (
	"log"

	a "github.com/FelipeStillner/ProjectPilot/services/calendar-manager/internal/adapter"
	c "github.com/FelipeStillner/ProjectPilot/services/calendar-manager/internal/core/service"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func main() {
	eventRepo := a.NewEventRepository()
	calendarService := c.NewCalendarService(eventRepo)
	restController := a.NewRestController(*calendarService)
	grpcController := a.NewGrpcController(*calendarService)

	go restController.Run()
	go grpcController.Run()

	select {}
}
