package repository

import (
	"context"
	"fmt"
	"github.com/event_bright/domain"
	"github.com/event_bright/internal/utils"
	"gorm.io/gorm"
	"log"
)

func New(db *gorm.DB) domain.AuthRepository {
	return &AuthSqlStorage{
		db: db,
	}
}

type AuthSqlStorage struct {
	db *gorm.DB
}

func (a *AuthSqlStorage) User(ctx context.Context, ctr *domain.User) string {
	db := a.db
	fmt.Println(ctr.Password, ctr.PasswordConfirm)
	hashedPassword := ""

	if ctr.Password != "" && ctr.PasswordConfirm != "" {
		if ctr.Password != ctr.PasswordConfirm {
			return "password doesn't match"
		} else {
			hash, err := utils.HashPassword(ctr.Password)
			if err != nil {
				log.Println("Password hashing failed")
			}
			hashedPassword = hash
		}
	} else {
		return "please provide both password"
	}

	ctr.Password = hashedPassword
	ctr.PasswordConfirm = hashedPassword

	user := domain.User{}

	if ctr.Email != "" {
		//Check if email already exists
		mail := db.First(&user, "email=?", ctr.Email)
		cred := &domain.User{}
		if err := mail.WithContext(ctx).Take(cred).Error; err != nil {
			log.Println(err)
		}
		if cred.Email != "" {
			return "email already exists"
		}
	}

	db.Create(ctr)

	return "success"
}
