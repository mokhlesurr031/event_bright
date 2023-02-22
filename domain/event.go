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
	CreatedBy   uint      `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type EventCriteria struct {
	Name        *string    `json:"name"`
	Date        *time.Time `json:"date"`
	Location    *string    `json:"location"`
	Description *string    `json:"description"`
	CreatedBy   *uint      `json:"created_by"`
	CreatedAt   *time.Time `json:"created_at"`
}

type EventRepository interface {
	Event(ctx context.Context, ctr *Event) (*Event, error)
	EventList(ctx context.Context, ctr *EventCriteria) ([]*Event, error)
	MyEventList(ctx context.Context, ctr *EventCriteria) ([]*Event, error)
}

type EventUseCase interface {
	Event(ctx context.Context, ctr *Event) (*Event, error)
	EventList(ctx context.Context, ctr *EventCriteria) ([]*Event, error)
	MyEventList(ctx context.Context, ctr *EventCriteria) ([]*Event, error)
}
