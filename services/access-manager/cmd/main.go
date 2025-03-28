package main

import (
	"log"

	"github.com/FelipeStillner/ProjectPilot/services/access-manager/internal/adapter"
	core "github.com/FelipeStillner/ProjectPilot/services/access-manager/internal/core/service"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func main() {
	teamRepo := adapter.NewTeamRepository()
	userRepo := adapter.NewUserRepository()
	accessService := core.NewAccessService(userRepo, teamRepo)
	grpcController := adapter.NewGrpcController(*accessService)

	go grpcController.Run()

	select {}
}
