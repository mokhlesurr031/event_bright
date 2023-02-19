package domain

import (
	"context"
	"github.com/event_bright/domain/dto"
)

type User struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type AuthRepository interface {
	User(ctx context.Context, ctr *User) string
	SignIn(ctx context.Context, ctr *dto.SignIn)
}

type AuthUseCase interface {
	User(ctx context.Context, ctr *User) string
}
