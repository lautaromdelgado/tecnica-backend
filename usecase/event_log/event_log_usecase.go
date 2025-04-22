package usecase

import (
	"context"

	model "github.com/lautaromdelgado/tecnica-backend/internal/domain/model/event_log"
	repository "github.com/lautaromdelgado/tecnica-backend/internal/domain/repository/event_log"
)

type EventLogUseCase interface {
	GetAllLogs(ctx context.Context) ([]*model.EventLog, error)
	GetLogsByTittle(ctx context.Context, title string) ([]*model.EventLog, error)
}

type eventLogUseCase struct {
	eventLogRepo repository.EventLogRepository
}

func NewEventLogRepository(el repository.EventLogRepository) *eventLogUseCase {
	return &eventLogUseCase{
		eventLogRepo: el,
	}
}

// GetAllLogs obtiene todos los logs de eventos
func (uc *eventLogUseCase) GetAllLogs(ctx context.Context) ([]*model.EventLog, error) {
	return uc.eventLogRepo.GetAllLogs(ctx)
}

// GetLogsByTitle obtiene logs de eventos por título
func (uc *eventLogUseCase) GetLogsByTittle(ctx context.Context, title string) ([]*model.EventLog, error) {
	return uc.eventLogRepo.GetLogsByTittle(ctx, title)
}
