package adapter

import (
	"context"
	"os"
	"time"

	c "github.com/FelipeStillner/ProjectPilot/services/calendar-manager/internal/core/struct"
	"github.com/FelipeStillner/ProjectPilot/services/calendar-manager/internal/port"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type GoogleIntegration struct {
	cache  port.CacheInterface
	config oauth2.Config
}

func NewGoogleIntegration(cache port.CacheInterface) *GoogleIntegration {
	config := oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  os.Getenv("GOOGLE_AUTH_URI"),
			TokenURL: os.Getenv("GOOGLE_TOKEN_URI"),
		},
		RedirectURL: os.Getenv("GOOGLE_REDIRECT_URI"),
		Scopes:      []string{calendar.CalendarScope},
	}
	return &GoogleIntegration{cache, config}
}

func (t *GoogleIntegration) Create(event c.Event) error {
	srv, err := t.getService()
	if err != nil {
		return err
	}

	e := &calendar.Event{
		Summary:     "ProjectPilot" + event.Name,
		Description: event.Description,
		Start: &calendar.EventDateTime{
			DateTime: event.Time.Format(time.RFC3339),
			TimeZone: "America/Sao_Paulo",
		},
		End: &calendar.EventDateTime{
			DateTime: event.Time.Add(time.Duration(event.Duration) * time.Minute).Format(time.RFC3339),
			TimeZone: "America/Sao_Paulo",
		},
		Attendees: []*calendar.EventAttendee{
			{Email: "felipeufranio@gmail.com"},
		},
	}

	_, err = srv.Events.Insert("primary", e).Do()
	if err != nil {
		return err
	}

	return nil
}

func (t *GoogleIntegration) Update(event c.Event) error {
	return nil
}

func (t *GoogleIntegration) Delete(event c.Event) error {
	return nil
}

func (t *GoogleIntegration) getService() (*calendar.Service, error) {
	ctx := context.Background()

	tok, err := t.getToken()
	if err != nil {
		return nil, err
	}

	client := t.config.Client(ctx, tok)

	return calendar.NewService(ctx, option.WithHTTPClient(client))
}

func (t *GoogleIntegration) getToken() (*oauth2.Token, error) {
	accessToken, err := t.cache.Get("token")
	if err != nil {
		return nil, err
	}

	if accessToken == nil {
		accessToken, err = t.createToken()
		if err != nil {
			return nil, err
		}

		err := t.cache.Set("token", *accessToken, time.Hour)
		if err != nil {
			return nil, err
		}
	}

	tok := &oauth2.Token{
		AccessToken: *accessToken,
		TokenType:   "Bearer",
	}

	return tok, nil
}

func (t *GoogleIntegration) createToken() (*string, error) {
	accessToken := ""

	return &accessToken, nil
}
