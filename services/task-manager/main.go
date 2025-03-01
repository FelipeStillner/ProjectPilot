package main

import (
	"log"
	"os"

	"github.com/FelipeStillner/ProjectPilot/services/task-manager/api"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func main() {
	r := api.SetupRouter()
	port := os.Getenv("PORT_TASK_MANAGER")
	r.Run(":" + port)
}
