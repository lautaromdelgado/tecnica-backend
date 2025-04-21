package usecase

import (
	"context"
	"errors"

	model "github.com/lautaromdelgado/tecnica-backend/internal/domain/model/event"
	repository "github.com/lautaromdelgado/tecnica-backend/internal/domain/repository/event"
)

type EventUseCase interface {
	CreateEvent(ctx context.Context, event *model.Event) error
}

type eventUseCase struct {
	eventRepo repository.EventRepository
}

func NewEventUseCase(er repository.EventRepository) *eventUseCase {
	return &eventUseCase{
		eventRepo: er,
	}
}

func (uc *eventUseCase) CreateEvent(ctx context.Context, event *model.Event) error {
	if event.Title == "" || event.Organizer == "" || event.Date == 0 {
		return errors.New("missing required fields")
	}
	return uc.eventRepo.Create(ctx, event)
}
