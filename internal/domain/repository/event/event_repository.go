package repository

import (
	"context"

	model "github.com/lautaromdelgado/tecnica-backend/internal/domain/model/event"
)

type EventRepository interface {
	Create(ctx context.Context, event *model.Event) error // Crea un nuevo evento
	// Update(ctx context.Context, event *model.Event) error                                         // Actualiza un evento existente
	// Delete(ctx context.Context, id uint) error                                                    // Elimina un evento por ID (marcando como eliminado)
	// FindByID(ctx context.Context, id uint) (*model.Event, error)                                  // Busca un evento por ID
	// FindAllPublished(ctx context.Context, filters map[string]interface{}) ([]*model.Event, error) // Busca todos los eventos publicados
	// FindAllForAdmin(ctx context.Context) ([]*model.Event, error)                                  // Busca todos los eventos (incluidos los no publicados) para el administrador
}
