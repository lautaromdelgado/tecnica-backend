package usecase

import (
	"context"
	"errors"

	dto "github.com/lautaromdelgado/tecnica-backend/delivery/http/dto/event"
	model "github.com/lautaromdelgado/tecnica-backend/internal/domain/model/event"
	model_event_log "github.com/lautaromdelgado/tecnica-backend/internal/domain/model/event_log"
	repository_event "github.com/lautaromdelgado/tecnica-backend/internal/domain/repository/event"
	repository_event_log "github.com/lautaromdelgado/tecnica-backend/internal/domain/repository/event_log"
)

type EventUseCase interface {
	CreateEvent(ctx context.Context, event *model.Event) error                                    // Crea un nuevo evento
	UpdateEvent(ctx context.Context, event *model.Event) error                                    // Actualiza un evento existente
	DeleteEvent(ctx context.Context, id uint) error                                               // Elimina un evento por ID (marcando como eliminado)
	UpdatePublishStatus(ctx context.Context, id uint, publish bool) error                         // Actualiza el estado de publicación de un evento por ID
	RestoreByID(ctx context.Context, id uint) error                                               // Restaurar evento por ID (soft delete)
	GetByIDWhitPermissions(ctx context.Context, id uint, role string) (*dto.EventResponse, error) // Obtiene un evento por ID con permisos
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

// DeleteEvent elimina un evento por ID (marcando como eliminado)
func (uc *eventUseCase) DeleteEvent(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid event id")
	}
	if err := uc.eventRepo.Delete(ctx, id); err != nil {
		return err
	}
	event, err := uc.eventRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	log := &model_event_log.EventLog{
		Title:     event.Title,
		Organizer: event.Organizer,
		Action:    "delete",
	}
	return uc.eventLogRepo.LogAction(ctx, log)
}

// UpdatePublishStatus actualiza el estado de publicación de un evento por ID
func (uc *eventUseCase) UpdatePublishStatus(ctx context.Context, id uint, publish bool) error {
	if id == 0 {
		return errors.New("invalid event id")
	}
	event, err := uc.eventRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if event.DeletedAt != nil {
		return errors.New("event is deleted and cannot be updated")
	}
	if err := uc.eventRepo.UpdatePublishStatus(ctx, id, publish); err != nil {
		return err
	}
	action := "unpublish"
	if publish {
		action = "publish"
	}
	log := &model_event_log.EventLog{
		Title:     event.Title,
		Organizer: event.Organizer,
		Action:    action,
	}
	return uc.eventLogRepo.LogAction(ctx, log)
}

// RestoreByID restaura un evento por ID (soft delete)
func (uc *eventUseCase) RestoreByID(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid event id")
	}
	if err := uc.eventRepo.RestoreByID(ctx, id); err != nil {
		return err
	}
	event, err := uc.eventRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	log := &model_event_log.EventLog{
		Title:     event.Title,
		Organizer: event.Organizer,
		Action:    "restore",
	}
	return uc.eventLogRepo.LogAction(ctx, log)
}

// GetByIDWhitPermissions obtiene un evento por ID con permisos
// Verifica si el evento fue eliminado y si está publicado o no
func (uc *eventUseCase) GetByIDWhitPermissions(ctx context.Context, id uint, role string) (*dto.EventResponse, error) {
	event, err := uc.eventRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	// Si el evento fue eliminado
	if event.DeletedAt != nil {
		return nil, errors.New("event deleted")
	}
	// Si no está publicado y no sos admin
	if !event.IsPublished && role != "admin" {
		return nil, errors.New("unauthorized to view this event")
	}
	event_dto := dto.EventResponse{
		ID:               event.ID,
		Organizer:        event.Organizer,
		Title:            event.Title,
		LongDescription:  event.LongDescription,
		ShortDescription: event.ShortDescription,
		Date:             event.Date,
		Location:         event.Location,
		IsPublished:      event.IsPublished,
	}
	return &event_dto, nil
}
