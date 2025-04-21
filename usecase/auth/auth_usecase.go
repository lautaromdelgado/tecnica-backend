package usecase

import (
	"context"
	"errors"

	model "github.com/lautaromdelgado/tecnica-backend/internal/domain/model/user"
	repository "github.com/lautaromdelgado/tecnica-backend/internal/domain/repository/user"
)

type AuthUseCase interface {
	Register(ctx context.Context, user *model.User) error
	Login(ctx context.Context, email, password string) (*model.User, error)
}

type authUseCase struct {
	userRepo repository.UserRepository
}

func NewAuthUseCase(userRepo repository.UserRepository) *authUseCase {
	return &authUseCase{
		userRepo: userRepo,
	}
}

func (uc *authUseCase) Register(ctx context.Context, u *model.User) error {
	if u.Role == "" {
		u.Role = "user" // Default role
	}
	return uc.userRepo.Create(ctx, u)
}

func (uc *authUseCase) Login(ctx context.Context, username, email string) (*model.User, error) {
	user, err := uc.userRepo.GetByEmail(ctx, username, email)
	if err != nil || user == nil {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}
