package core

import (
	"errors"

	core "github.com/FelipeStillner/ProjectPilot/services/calendar-manager/internal/core/struct"
)

type ReadEventInput struct {
	Id uint32 `json:"id"`
}

func (t *CalendarService) ReadEvent(input ReadEventInput) (*core.Event, error) {
	err := input.verify()
	if err != nil {
		return nil, err
	}

	return t.eventRepo.Read(input.Id)
}

func (input ReadEventInput) verify() error {
	if input.Id == 0 {
		return errors.New("event id is required")
	}

	return nil
}
