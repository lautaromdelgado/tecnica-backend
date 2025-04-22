package repository

import (
	"context"

	model "github.com/lautaromdelgado/tecnica-backend/internal/domain/model/event"
)

type EventRepository interface {
	Create(ctx context.Context, event *model.Event) error                 // Crea un nuevo evento
	Update(ctx context.Context, event *model.Event) error                 // Actualiza un evento existente
	Delete(ctx context.Context, id uint) error                            // Elimina un evento por ID (marcando como eliminado)
	UpdatePublishStatus(ctx context.Context, id uint, publish bool) error // Actualiza el estado de publicación de un evento por ID
	RestoreByID(ctx context.Context, id uint) error                       // Restaurar evento por ID (soft delete)

	FindByID(ctx context.Context, id uint) (*model.Event, error) // Busca un evento por ID
}
