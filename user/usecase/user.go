package usecase

import (
	"context"

	"github.com/event_bright/domain"
	"github.com/event_bright/domain/dto"
)

// New return new usecase for user
func New(repo domain.AuthRepository) domain.AuthUseCase {
	return &AuthUseCase{
		repo: repo,
	}
}

type AuthUseCase struct {
	repo domain.AuthRepository
}

func (a *AuthUseCase) User(ctx context.Context, ctr *domain.User) (*domain.User, error) {
	return a.repo.User(ctx, ctr)
}

func (a *AuthUseCase) SignIn(ctx context.Context, ctr *dto.SignIn) (*domain.JWTToken, error) {
	return a.repo.SignIn(ctx, ctr)
}
