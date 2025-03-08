package core

import (
	"errors"
	"time"

	core "github.com/FelipeStillner/ProjectPilot/services/calendar-manager/internal/core/struct"
)

type UpdateEventInput struct {
	Id          uint32 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Time        string `json:"time"`
	Duration    uint32 `json:"duration"`
}

func (t *CalendarService) UpdateEvent(input UpdateEventInput) (*core.Event, error) {
	err := input.verify()
	if err != nil {
		return nil, err
	}

	Time, err := time.Parse(time.RFC3339, input.Time)
	if err != nil {
		return nil, err
	}

	event := core.Event{
		Id:          input.Id,
		Name:        input.Name,
		Description: input.Description,
		Time:        Time,
		Duration:    input.Duration,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return t.eventRepo.Update(event.Id, event)
}

func (input UpdateEventInput) verify() error {
	if input.Id == 0 {
		return errors.New("event id is required")
	}

	if input.Name == "" {
		return errors.New("event name is required")
	}

	return nil
}
