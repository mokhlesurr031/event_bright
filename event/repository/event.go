package repository

import (
	"context"
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
