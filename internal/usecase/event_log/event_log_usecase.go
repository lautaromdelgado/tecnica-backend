package usecase

import (
	"context"

	model "github.com/lautaromdelgado/tecnica-backend/internal/domain/model/event_log"
	repository "github.com/lautaromdelgado/tecnica-backend/internal/domain/repository/event_log"
)

type EventLogUseCase interface {
	GetAllLogs(ctx context.Context) ([]*model.EventLog, error)
	GetLogsByFilters(ctx context.Context, title, action, organizer string) ([]*model.EventLog, error)
}

type eventLogUseCase struct {
	eventLogRepo repository.EventLogRepository
}

func NewEventLogUseCase(el repository.EventLogRepository) *eventLogUseCase {
	return &eventLogUseCase{
		eventLogRepo: el,
	}
}

// GetAllLogs obtiene todos los logs de eventos
func (uc *eventLogUseCase) GetAllLogs(ctx context.Context) ([]*model.EventLog, error) {
	return uc.eventLogRepo.GetAllLogs(ctx)
}

// GetLogsByTitle obtiene logs de eventos por título
func (uc *eventLogUseCase) GetLogsByFilters(ctx context.Context, title, action, organizer string) ([]*model.EventLog, error) {
	return uc.eventLogRepo.GetLogsByFilters(ctx, title, action, organizer)
}
