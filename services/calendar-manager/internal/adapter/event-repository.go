package adapter

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	c "github.com/FelipeStillner/ProjectPilot/services/calendar-manager/internal/core/struct"
	_ "github.com/lib/pq"
)

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository() *EventRepository {
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

	return &EventRepository{db: db}
}

func (t *EventRepository) Create(event c.Event) (*c.Event, error) {
	stmt, err := t.db.Prepare("INSERT INTO event (id, name, description, time, duration, attendees, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $7)")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(event.Id, event.Name, event.Description, event.Time, event.Duration, event.Attendees, time.Now())
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (t *EventRepository) Read(id uint32) (*c.Event, error) {
	var event c.Event

	stmt, err := t.db.Prepare("SELECT id, name, description, time, duration, attendees FROM event WHERE id = $1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&event.Id, &event.Name, &event.Description, &event.Time, &event.Duration, &event.Attendees)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (t *EventRepository) Update(id uint32, event c.Event) (*c.Event, error) {
	stmt, err := t.db.Prepare("UPDATE event SET name = $1, description = $2, time = $3, duration = $4, attendees = $5, updated_at = $6 WHERE id = $7")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(event.Name, event.Description, event.Time, event.Duration, event.Attendees, time.Now(), id)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (t *EventRepository) Delete(id uint32) error {
	stmt, err := t.db.Prepare("UPDATE event SET deleted_at = $1 WHERE id = $2")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}
