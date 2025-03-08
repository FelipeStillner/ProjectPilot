package core

import (
	"errors"
)

type DeleteEventInput struct {
	Id uint32 `json:"id"`
}

func (t *CalendarService) DeleteEvent(input DeleteEventInput) error {
	err := input.verify()
	if err != nil {
		return err
	}

	return t.eventRepo.Delete(input.Id)
}

func (input DeleteEventInput) verify() error {
	if input.Id == 0 {
		return errors.New("event id is required")
	}

	return nil
}
