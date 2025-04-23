package usecase

import (
	"context"
	"errors"

	model "github.com/lautaromdelgado/tecnica-backend/internal/domain/model/user"
	repository "github.com/lautaromdelgado/tecnica-backend/internal/domain/repository/user"
)

type AuthUseCase interface {
	Register(ctx context.Context, user *model.User) error                   // Crear un nuevo usuario
	Login(ctx context.Context, email, password string) (*model.User, error) // Obtener un usuario por su email y password
	UpdateByID(ctx context.Context, user *model.User) error                 // Actualizar un usuario por su ID
	DeleteByID(ctx context.Context, id uint) error                          // Eliminar un usuario por su ID
	ListActiveUsers(ctx context.Context) ([]*model.User, error)             // Obtener todos los usuarios activos
	ListInactiveUsers(ctx context.Context) ([]*model.User, error)           // Obtener todos los usuarios inactivos
	RestoreUser(ctx context.Context, id uint) error                         // Restaurar un usuario por su ID
}

type authUseCase struct {
	userRepo repository.UserRepository
}

func NewAuthUseCase(userRepo repository.UserRepository) *authUseCase {
	return &authUseCase{
		userRepo: userRepo,
	}
}

// Register crea un nuevo usuario en la base de datos
func (uc *authUseCase) Register(ctx context.Context, u *model.User) error {
	if u.Role == "" {
		u.Role = "user"
	}
	return uc.userRepo.Create(ctx, u)
}

// Login busca un usuario en la base de datos por su nombre de usuario y correo electrónico
func (uc *authUseCase) Login(ctx context.Context, username, email string) (*model.User, error) {
	user, err := uc.userRepo.GetByEmail(ctx, username, email)
	if err != nil || user == nil {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

// UpdateByID actualiza un usuario en la base de datos por su ID
func (uc *authUseCase) UpdateByID(ctx context.Context, u *model.User) error {
	return uc.userRepo.UpdateByID(ctx, u)
}

// DeleteByID marca un usuario como eliminado en la base de datos por su ID
func (uc *authUseCase) DeleteByID(ctx context.Context, id uint) error {
	return uc.userRepo.DeleteByID(ctx, id)
}

// ListActiveUsers obtiene todos los usuarios activos de la base de datos
func (uc *authUseCase) ListActiveUsers(ctx context.Context) ([]*model.User, error) {
	return uc.userRepo.FindAllActive(ctx)
}

// ListInactiveUsers obtiene todos los usuarios inactivos de la base de datos
func (uc *authUseCase) ListInactiveUsers(ctx context.Context) ([]*model.User, error) {
	return uc.userRepo.FindAllInactive(ctx)
}

// RestoreUser restaura un usuario en la base de datos por su ID
func (uc *authUseCase) RestoreUser(ctx context.Context, id uint) error {
	return uc.userRepo.RestoreByID(ctx, id)
}
