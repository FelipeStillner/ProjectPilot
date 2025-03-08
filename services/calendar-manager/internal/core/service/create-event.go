package core

import (
	"errors"
	"time"

	core "github.com/FelipeStillner/ProjectPilot/services/calendar-manager/internal/core/struct"
	"github.com/google/uuid"
)

type CreateEventInput struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Time        string   `json:"time"`
	Duration    uint32   `json:"duration"`
	Attendees   []uint32 `json:"attendees"`
}

func (t *CalendarService) CreateEvent(input CreateEventInput) (*core.Event, error) {
	err := input.verify()
	if err != nil {
		return nil, err
	}

	Time, err := time.Parse(time.RFC3339, input.Time)
	if err != nil {
		return nil, err
	}

	event := core.Event{
		Id:          uuid.New().ID(),
		Name:        input.Name,
		Description: input.Description,
		Time:        Time,
		Duration:    input.Duration,
		Attendees:   input.Attendees,
	}

	return t.eventRepo.Create(event)
}

func (input CreateEventInput) verify() error {
	if input.Name == "" {
		return errors.New("event name is required")
	}

	return nil
}
