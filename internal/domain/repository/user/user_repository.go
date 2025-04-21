package repository

import (
	"context"

	model "github.com/lautaromdelgado/tecnica-backend/internal/domain/model/user"
)

type UserRepository interface {
	Create(ctx context.Context, u *model.User) error                             // Crear un nuevo usuario
	GetByEmail(ctx context.Context, username, email string) (*model.User, error) // Obtener un usuario por su email y username
	UpdateByID(ctx context.Context, user *model.User) error                      // Actualizar un usuario por su ID
	DeleteByID(ctx context.Context, id uint) error                               // Eliminar un usuario por su ID
	FindAllActive(ctx context.Context) ([]*model.User, error)                    // Obtener todos los usuarios activos
	FindAllInactive(ctx context.Context) ([]*model.User, error)                  // Obtener todos los usuarios inactivos
	RestoreByID(ctx context.Context, id uint) error                              // Restaurar un usuario por su ID
}
