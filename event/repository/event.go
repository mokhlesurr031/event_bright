package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"strconv"

	"github.com/event_bright/domain"
	"github.com/event_bright/internal/config"
	"github.com/event_bright/internal/utils"
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
	jwt := config.JWT()
	tokenString, _ := ctx.Value("token").(string)
	token, err := utils.ValidateToken(tokenString, jwt.Secret)
	if err != nil {
		return nil, err
	}

	tokenUint, _ := strconv.ParseUint(token, 10, 64)

	ctr.CreatedBy = uint(tokenUint)

	db := e.db
	res := db.Create(ctr)

	if res.Error != nil {
		return nil, res.Error
	}
	return ctr, nil

}

func (e *EventSqlStorage) EventList(ctx context.Context, ctr *domain.EventCriteria) ([]*domain.Event, error) {
	qry := e.db

	eventList := make([]*domain.Event, 0)
	if err := qry.WithContext(ctx).Find(&eventList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return eventList, nil
}

func (e *EventSqlStorage) MyEventList(ctx context.Context, ctr *domain.EventCriteria) ([]*domain.Event, error) {
	qry := e.db
	jwt := config.JWT()
	tokenString, _ := ctx.Value("token").(string)
	token, err := utils.ValidateToken(tokenString, jwt.Secret)
	if err != nil {
		return nil, err
	}

	tokenUint, _ := strconv.ParseUint(token, 10, 64)

	eventList := make([]*domain.Event, 0)
	if err := qry.WithContext(ctx).Where("created_by = ?", tokenUint).Find(&eventList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return eventList, nil
}

func (c *EventSqlStorage) EventDetails(ctx context.Context, ctr *domain.EventCriteria) (*domain.Event, error) {
	qry := c.db
	eventDetails := &domain.Event{}
	if ctr.Id != nil {
		err := qry.First(&eventDetails, "id=?", ctr.Id).Error
		if err != nil {
			return nil, err
		}
		return eventDetails, nil
	}
	return nil, nil
}

func (e *EventSqlStorage) Participate(ctx context.Context, ctr *domain.Participant) (*domain.Participant, error) {
	db := e.db
	var event domain.Event
	if err := db.Create(ctr).Error; err != nil {
		return nil, err
	}

	db.First(&event, "id=?", ctr.EventID)
	fmt.Println(event.TotalParticipant)
	db.Model(&event).Update("TotalParticipant", event.TotalParticipant+1)

	return ctr, nil
}
