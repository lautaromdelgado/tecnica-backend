package repository

import (
	"context"

	dto "github.com/lautaromdelgado/tecnica-backend/delivery/http/dto/user_event"
)

type UserEventRepository interface {
	Insert(ctx context.Context, ue *dto.UserEventRequest) error     // Suscribir un usuario a un evento
	Exists(ctx context.Context, userID, eventID uint) (bool, error) // Verificar si un usuario ya está suscrito a un evento
}
