package usecase

import (
	"context"

	model "github.com/lautaromdelgado/tecnica-backend/internal/domain/model/event_log"
	repository "github.com/lautaromdelgado/tecnica-backend/internal/domain/repository/event_log"
)

type EventLogUseCase interface {
	GetAllLogs(ctx context.Context) ([]*model.EventLog, error)
}

type eventLogUseCase struct {
	eventLogRepo repository.EventLogRepository
}

func NewEventLogRepository(el repository.EventLogRepository) *eventLogUseCase {
	return &eventLogUseCase{
		eventLogRepo: el,
	}
}

func (uc *eventLogUseCase) GetAllLogs(ctx context.Context) ([]*model.EventLog, error) {
	return uc.eventLogRepo.GetAllLogs(ctx)
}
