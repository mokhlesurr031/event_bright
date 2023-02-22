package usecase

import (
	"context"

	"github.com/event_bright/domain"
)

// New return new usecase for user
func New(repo domain.EventRepository) domain.EventUseCase {
	return &EventUseCase{
		repo: repo,
	}
}

type EventUseCase struct {
	repo domain.EventRepository
}

func (e *EventUseCase) Event(ctx context.Context, ctr *domain.Event) (*domain.Event, error) {
	return e.repo.Event(ctx, ctr)
}

func (e *EventUseCase) EventList(ctx context.Context, ctr *domain.EventCriteria) ([]*domain.Event, error) {
	return e.repo.EventList(ctx, ctr)
}

func (e *EventUseCase) MyEventList(ctx context.Context, ctr *domain.EventCriteria) ([]*domain.Event, error) {
	return e.repo.MyEventList(ctx, ctr)
}
