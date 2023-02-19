package usecase

import (
	"context"
	"github.com/event_bright/domain"
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

func (a *AuthUseCase) User(ctx context.Context, ctr *domain.User) string {
	return a.repo.User(ctx, ctr)
}
