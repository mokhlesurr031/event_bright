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

type SelfEventParticipant struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SelfEventList struct {
	Name             string    `json:"name"`
	Location         string    `json:"location"`
	Date             time.Time `json:"date"`
	Description      string    `json:"description"`
	CreatedBy        string    `json:"created_by"`
	TotalParticipant int       `json:"total_participant"`
	Participant      []SelfEventParticipant
}

type EventRepository interface {
	Event(ctx context.Context, ctr *Event) (*Event, error)
	EventList(ctx context.Context, ctr *EventCriteria) ([]*Event, error)
	MyEventList(ctx context.Context, ctr *EventCriteria) ([]*SelfEventList, error)
	EventDetails(ctx context.Context, ctr *EventCriteria) (*Event, error)
	Participate(ctx context.Context, ctr *Participant) (*Participant, error)
}

type EventUseCase interface {
	Event(ctx context.Context, ctr *Event) (*Event, error)
	EventList(ctx context.Context, ctr *EventCriteria) ([]*Event, error)
	MyEventList(ctx context.Context, ctr *EventCriteria) ([]*SelfEventList, error)
	EventDetails(ctx context.Context, ctr *EventCriteria) (*Event, error)
	Participate(ctx context.Context, ctr *Participant) (*Participant, error)
}
