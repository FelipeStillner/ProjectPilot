package adapter

import (
	"database/sql"
	"fmt"
	"os"

	c "github.com/FelipeStillner/ProjectPilot/services/task-manager/internal/core/struct"
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
	stmt, err := t.db.Prepare("INSERT INTO tasks (id, name, description, priority, assignee, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(task.Id, task.Name, task.Description, task.Priority, task.Assignee, task.Status, task.CreatedAt, task.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
