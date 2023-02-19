package domain

import (
	"context"
	"time"
)

type Event struct {
	Name        string    `json:"name"`
	Date        time.Time `json:"date"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type EventRepository interface {
	Event(ctx context.Context, ctr *Event) (*Event, error)
}

type EventUseCase interface {
	Event(ctx context.Context, ctr *Event) (*Event, error)
}
