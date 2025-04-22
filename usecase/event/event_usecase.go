package usecase

import (
	"context"
	"errors"

	model "github.com/lautaromdelgado/tecnica-backend/internal/domain/model/event"
	model_event_log "github.com/lautaromdelgado/tecnica-backend/internal/domain/model/event_log"
	repository_event "github.com/lautaromdelgado/tecnica-backend/internal/domain/repository/event"
	repository_event_log "github.com/lautaromdelgado/tecnica-backend/internal/domain/repository/event_log"
)

type EventUseCase interface {
	CreateEvent(ctx context.Context, event *model.Event) error            // Crea un nuevo evento
	UpdateEvent(ctx context.Context, event *model.Event) error            // Actualiza un evento existente
	DeleteEvent(ctx context.Context, id uint) error                       // Elimina un evento por ID (marcando como eliminado)
	UpdatePublishStatus(ctx context.Context, id uint, publish bool) error // Actualiza el estado de publicación de un evento por ID
	RestoreByID(ctx context.Context, id uint) error                       // Restaurar evento por ID (soft delete)
}

type eventUseCase struct {
	eventRepo    repository_event.EventRepository        // Repositorio de eventos
	eventLogRepo repository_event_log.EventLogRepository // Repositorio para registrar acciones de eventos
}

func NewEventUseCase(er repository_event.EventRepository, event_log repository_event_log.EventLogRepository) *eventUseCase {
	return &eventUseCase{
		eventRepo:    er,
		eventLogRepo: event_log,
	}
}

// CreateEvent crea un nuevo evento
func (uc *eventUseCase) CreateEvent(ctx context.Context, event *model.Event) error {
	if event.Title == "" || event.Organizer == "" || event.Date == 0 {
		return errors.New("missing required fields")
	}
	if err := uc.eventRepo.Create(ctx, event); err != nil {
		return err
	}
	log := &model_event_log.EventLog{
		Title:     event.Title,
		Organizer: event.Organizer,
		Action:    "create",
	}
	return uc.eventLogRepo.LogAction(ctx, log)
}

// UpdateEvent actualiza un evento existente
func (uc *eventUseCase) UpdateEvent(ctx context.Context, event *model.Event) error {
	if event.ID == 0 || event.Title == "" || event.Organizer == "" || event.Date <= 0 {
		return errors.New("missing required fields")
	}
	if err := uc.eventRepo.Update(ctx, event); err != nil {
		return err
	}
	log := &model_event_log.EventLog{
		Title:     event.Title,
		Organizer: event.Organizer,
		Action:    "update",
	}
	return uc.eventLogRepo.LogAction(ctx, log)
}

// TODO: Implementar event_log para delete y restore y cambiar estado de evento si publicado o no publicado
// DeleteEvent elimina un evento por ID (marcando como eliminado)
func (uc *eventUseCase) DeleteEvent(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid event id")
	}
	return uc.eventRepo.Delete(ctx, id)
}

// UpdatePublishStatus actualiza el estado de publicación de un evento por ID
func (uc *eventUseCase) UpdatePublishStatus(ctx context.Context, id uint, publish bool) error {
	if id == 0 {
		return errors.New("invalid event id")
	}
	return uc.eventRepo.UpdatePublishStatus(ctx, id, publish)
}

// RestoreByID restaura un evento por ID (soft delete)
func (uc *eventUseCase) RestoreByID(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid event id")
	}
	return uc.eventRepo.RestoreByID(ctx, id)
}
