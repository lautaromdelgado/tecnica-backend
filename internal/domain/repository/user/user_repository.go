package repository

import (
	"context"

	model "github.com/lautaromdelgado/tecnica-backend/internal/domain/model/user"
)

type UserRepository interface {
	Create(ctx context.Context, u *model.User) error
	GetByEmail(ctx context.Context, username, email string) (*model.User, error)
	UpdateByID(ctx context.Context, user *model.User) error
}
