package main

import (
	"log"

	a "github.com/FelipeStillner/ProjectPilot/services/calendar-manager/internal/adapter"
	c "github.com/FelipeStillner/ProjectPilot/services/calendar-manager/internal/core/service"
	"github.com/FelipeStillner/ProjectPilot/services/calendar-manager/internal/port"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func main() {
	redis := a.NewRedis()
	eventRepo := a.NewEventRepository()
	googleIntegration := a.NewGoogleIntegration(redis)
	integrations := []port.IntegrationInterface{googleIntegration}
	calendarService := c.NewCalendarService(eventRepo, integrations)
	grpcController := a.NewGrpcController(*calendarService)

	go grpcController.Run()

	select {}
}
