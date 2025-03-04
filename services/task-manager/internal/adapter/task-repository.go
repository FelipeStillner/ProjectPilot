package adapter

import (
	"database/sql"
	"fmt"
	"os"
	"time"

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

func (t *TaskRepository) Create(task c.Task) (*c.Task, error) {
	stmt, err := t.db.Prepare("INSERT INTO tasks (id, name, description, priority, assignee, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(task.Id, task.Name, task.Description, task.Priority, task.Assignee, task.Status, task.CreatedAt, task.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *TaskRepository) Read(id uint32) (*c.Task, error) {
	var task c.Task

	stmt, err := t.db.Prepare("SELECT id, name, description, priority, assignee, status, created_at, updated_at FROM tasks WHERE id = $1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&task.Id, &task.Name, &task.Description, &task.Priority, &task.Assignee, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *TaskRepository) Update(id uint32, task c.Task) (*c.Task, error) {
	stmt, err := t.db.Prepare("UPDATE tasks SET name = $1, description = $2, priority = $3, assignee = $4, status = $5, updated_at = $6 WHERE id = $7")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(task.Name, task.Description, task.Priority, task.Assignee, task.Status, task.UpdatedAt, id)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *TaskRepository) Delete(id uint32) error {
	stmt, err := t.db.Prepare("UPDATE tasks SET deleted_at = $1 WHERE id = $2")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}
