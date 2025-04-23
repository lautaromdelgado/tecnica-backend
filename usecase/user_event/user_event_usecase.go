package usecase

import (
	"context"
	"errors"
	"time"

	dto "github.com/lautaromdelgado/tecnica-backend/delivery/http/dto/user_event"
	repository_event "github.com/lautaromdelgado/tecnica-backend/internal/domain/repository/event"
	repository_user_event "github.com/lautaromdelgado/tecnica-backend/internal/domain/repository/user_event"
)

type UserEventUseCase interface {
	SuscribeUserToEvent(ctx context.Context, userID, eventID uint) error // Suscribir un usuario a un evento
}

type userEventUseCase struct {
	userEventRepo repository_user_event.UserEventRepository
	eventRepo     repository_event.EventRepository
}

func NewUserEventUseCase(ue repository_user_event.UserEventRepository, er repository_event.EventRepository) *userEventUseCase {
	return &userEventUseCase{
		userEventRepo: ue,
		eventRepo:     er,
	}
}

// SubscribeUserToEvent suscribe a un usuario a un evento
func (uc *userEventUseCase) SuscribeUserToEvent(ctx context.Context, userID, eventID uint) error {
	// Buscar evento
	event, err := uc.eventRepo.FindByID(ctx, eventID)
	if err != nil {
		return errors.New("event not found")
	}
	// Validar si está eliminado
	if event.DeletedAt != nil {
		return errors.New("event has been deleted")
	}
	// Validar si esta publicado
	if !event.IsPublished {
		return errors.New("event is not published")
	}
	// Validar fecha futura
	if event.Date <= time.Now().Unix() {
		return errors.New("event has already occurred")
	}
	// Verificar si el usuario ya está suscrito al evento
	exists, err := uc.userEventRepo.Exists(ctx, userID, eventID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user is already subscribed to this event")
	}

	// Hacer la suscripción
	sub := &dto.UserEventRequest{
		UserID:  userID,
		EventID: eventID,
	}
	return uc.userEventRepo.Insert(ctx, sub)
}
