package adapter

import (
	"database/sql"
	"fmt"
	"os"

	c "github.com/FelipeStillner/ProjectPilot/services/task-manager/internal/core/entities"
	_ "github.com/lib/pq"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository() *TaskRepository {
	dbHost := os.Getenv("HOST_POSTGRES")
	dbPort := os.Getenv("PORT_POSTGRES")
	dbUser := os.Getenv("USER_POSTGRES")
	dbPassword := os.Getenv("PASSWORD_POSTGRES")
	dbName := os.Getenv("DATABASE_POSTGRES")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return &TaskRepository{db: db}
}

func (t *TaskRepository) SaveTask(task c.Task) error {
	stmt, err := t.db.Prepare("INSERT INTO tasks (id, name, created_at) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(task.Id, task.Name, task.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
