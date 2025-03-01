package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Database connection function
func connectDB() (*sql.DB, error) {
	dbHost := os.Getenv("HOST_POSTGRES")
	dbPort := os.Getenv("PORT_POSTGRES")
	dbUser := os.Getenv("USER_POSTGRES")
	dbPassword := os.Getenv("PASSWORD_POSTGRES")
	dbName := os.Getenv("DATABASE_POSTGRES")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	return sql.Open("postgres", dsn)
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO tasks (name) VALUES ('New Task')")
	if err != nil {
		http.Error(w, "Failed to insert task", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Task Created!")
}

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("Loaded .env file")
}

func main() {
	http.HandleFunc("/create-task", createTaskHandler)

	port := os.Getenv("PORT_TASK_MANAGER")
	fmt.Println("Task Creator Service running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
