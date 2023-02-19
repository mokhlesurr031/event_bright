package repository

import (
	"context"
	"github.com/event_bright/domain"
	"gorm.io/gorm"
)

func New(db *gorm.DB) domain.EventRepository {
	return &EventSqlStorage{
		db: db,
	}
}

type EventSqlStorage struct {
	db *gorm.DB
}

func (e *EventSqlStorage) Event(ctx context.Context, ctr *domain.Event) (*domain.Event, error) {
	db := e.db
	res := db.Create(ctr)

	if res.Error != nil {
		return nil, res.Error
	}
	return ctr, nil

}
