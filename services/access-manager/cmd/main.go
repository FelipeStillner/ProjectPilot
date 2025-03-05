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
	redis := adapter.NewRedis()
	accessService := core.NewAccessService(userRepo, teamRepo, redis)
	restController := adapter.NewRestController(*accessService)
	grpcController := adapter.NewGrpcController(*accessService)

	go restController.Run()
	go grpcController.Run()

	select {}
}
