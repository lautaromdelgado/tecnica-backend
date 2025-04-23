package persistence

import (
	"context"

	"github.com/jmoiron/sqlx"
	dto "github.com/lautaromdelgado/tecnica-backend/delivery/http/dto/user_event"
)

type UserEventRepository struct {
	db *sqlx.DB
}

func NewUserEventRepository(db *sqlx.DB) *UserEventRepository {
	return &UserEventRepository{db: db}
}

// Insert inserta un usuario id y un evento id, para manejar la suscripción de los eventos
// a los que el usuario se suscribe
func (r *UserEventRepository) Insert(ctx context.Context, ue *dto.UserEventRequest) error {
	query := `INSERT INTO user_event (user_id, event_id) VALUES (?, ?)`
	_, err := r.db.ExecContext(ctx, query, ue.UserID, ue.EventID)
	return err
}

// Exists verifica si un usuario ya está suscrito a un evento
// para evitar duplicados en la base de datos
func (r *UserEventRepository) Exists(ctx context.Context, userID, eventID uint) (bool, error) {
	query := `SELECT COUNT(*) FROM user_event WHERE user_id = ? AND event_id = ?`
	var count int
	err := r.db.GetContext(ctx, &count, query, userID, eventID)
	return count > 0, err
}
