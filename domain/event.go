package domain

import (
	"context"
	"time"
)

type Event struct {
	Id               uint      `json:"id"`
	Name             string    `json:"name"`
	Date             time.Time `json:"date"`
	Location         string    `json:"location"`
	Description      string    `json:"description"`
	CreatedBy        uint      `json:"created_by"`
	TotalParticipant int       `json:"total_participant"`
	CreatedAt        time.Time `json:"created_at"`
}

type EventCriteria struct {
	Id               *uint      `json:"id"`
	Name             *string    `json:"name"`
	Date             *time.Time `json:"date"`
	Location         *string    `json:"location"`
	Description      *string    `json:"description"`
	CreatedBy        *uint      `json:"created_by"`
	TotalParticipant int        `json:"total_participant"`
	CreatedAt        *time.Time `json:"created_at"`
}

type EventRepository interface {
	Event(ctx context.Context, ctr *Event) (*Event, error)
	EventList(ctx context.Context, ctr *EventCriteria) ([]*Event, error)
	MyEventList(ctx context.Context, ctr *EventCriteria) ([]*Event, error)
	EventDetails(ctx context.Context, ctr *EventCriteria) (*Event, error)
	Participate(ctx context.Context, ctr *Participant) (*Participant, error)
}

type EventUseCase interface {
	Event(ctx context.Context, ctr *Event) (*Event, error)
	EventList(ctx context.Context, ctr *EventCriteria) ([]*Event, error)
	MyEventList(ctx context.Context, ctr *EventCriteria) ([]*Event, error)
	EventDetails(ctx context.Context, ctr *EventCriteria) (*Event, error)
	Participate(ctx context.Context, ctr *Participant) (*Participant, error)
}
