package domain

import "time"

type Participant struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	EventID   uint      `json:"event_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ParticipantCriteria struct {
	Id        uint       `json:"id"`
	Name      *string    `json:"name"`
	Email     *string    `json:"email"`
	EventID   uint       `json:"event_id"`
	CreatedAt *time.Time `json:"created_at"`
}
