package repository

import (
	"context"
	"fmt"
	"github.com/event_bright/domain"
	"github.com/event_bright/domain/dto"
	"github.com/event_bright/internal/config"
	"github.com/event_bright/internal/utils"
	"gorm.io/gorm"
	"log"
	"strconv"
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

func (a *AuthSqlStorage) SignIn(ctx context.Context, ctr *dto.SignIn) (*dto.JWTToken, error) {
	qry := a.db
	jwt := config.JWT()
	user := domain.User{}

	if ctr.Email != "" && ctr.Password != "" {
		qry := qry.Find(&user, "email=?", ctr.Email)
		cred := &domain.User{}
		if err := qry.WithContext(ctx).Take(cred).Error; err != nil {
			log.Println(err)
		}
		if cred.Email == "" {
			reqJwt := &dto.JWTToken{Message: "invalid data"}
			return reqJwt, nil
		}

		if err := utils.VerifyPassword(cred.Password, ctr.Password); err != nil {
			reqJwt := &dto.JWTToken{Message: "invalid password"}
			return reqJwt, err
		}

		token, err := utils.GenerateToken(jwt.ExpiredIn, strconv.Itoa(int(cred.ID)), jwt.Secret)
		if err != nil {
			log.Println(err)
		}

		loggedInData := &dto.LoggerInUserData{}
		loggedInData.Name = user.Name
		loggedInData.Email = user.Email
		loggedInData.ID = user.ID

		reqJwt := &dto.JWTToken{User: loggedInData, Secret: token, MaxAge: jwt.MaxAge, ExpiredIn: jwt.ExpiredIn, Message: "success"}
		return reqJwt, nil
	}

	reqJwt := &dto.JWTToken{Message: "invalid data"}
	return reqJwt, nil

}
