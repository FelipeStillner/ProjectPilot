package adapter

import (
	c "github.com/FelipeStillner/ProjectPilot/services/calendar-manager/internal/core/struct"
)

type GoogleIntegration struct{}

func NewGoogleIntegration() *GoogleIntegration {
	return &GoogleIntegration{}
}

func (t *GoogleIntegration) Create(event c.Event) error {
	return nil
}

func (t *GoogleIntegration) Update(event c.Event) error {
	return nil
}

func (t *GoogleIntegration) Delete(event c.Event) error {
	return nil
}
