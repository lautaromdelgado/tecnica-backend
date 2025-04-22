package usecase

import (
	"context"
	"errors"

	model "github.com/lautaromdelgado/tecnica-backend/internal/domain/model/event"
	repository "github.com/lautaromdelgado/tecnica-backend/internal/domain/repository/event"
)

type EventUseCase interface {
	CreateEvent(ctx context.Context, event *model.Event) error // Crea un nuevo evento
	UpdateEvent(ctx context.Context, event *model.Event) error // Actualiza un evento existente
	DeleteEvent(ctx context.Context, id uint) error            // Elimina un evento por ID (marcando como eliminado)
}

type eventUseCase struct {
	eventRepo repository.EventRepository
}

func NewEventUseCase(er repository.EventRepository) *eventUseCase {
	return &eventUseCase{
		eventRepo: er,
	}
}

// CreateEvent crea un nuevo evento
func (uc *eventUseCase) CreateEvent(ctx context.Context, event *model.Event) error {
	if event.Title == "" || event.Organizer == "" || event.Date == 0 {
		return errors.New("missing required fields")
	}
	return uc.eventRepo.Create(ctx, event)
}

// UpdateEvent actualiza un evento existente
func (uc *eventUseCase) UpdateEvent(ctx context.Context, event *model.Event) error {
	if event.ID == 0 || event.Title == "" || event.Organizer == "" || event.Date <= 0 {
		return errors.New("missing required fields")
	}
	return uc.eventRepo.Update(ctx, event)
}

// DeleteEvent elimina un evento por ID (marcando como eliminado)
func (uc *eventUseCase) DeleteEvent(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid event id")
	}
	return uc.eventRepo.Delete(ctx, id)
}
