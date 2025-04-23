package persistence

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	dto "github.com/lautaromdelgado/tecnica-backend/internal/delivery/http/dto/user_event"
	model "github.com/lautaromdelgado/tecnica-backend/internal/domain/model/event"
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

// GetEventsByUser obtiene todos los eventos a los que un usuario está suscrito
// y que no han sido eliminados
func (r *UserEventRepository) GetEventsByUser(ctx context.Context, userID uint) ([]*model.Event, error) {
	query := `
		SELECT e.* FROM events e
		INNER JOIN user_event ue ON e.id = ue.event_id
		WHERE ue.user_id = ? AND e.deleted_at IS NULL
		ORDER BY e.date ASC
	`
	var events []*model.Event
	err := r.db.SelectContext(ctx, &events, query, userID)
	return events, err
}

// Delete elimina la suscripción de un usuario a un evento
func (r *UserEventRepository) Delete(ctx context.Context, userID, eventID uint) error {
	query := `DELETE FROM user_event WHERE user_id = ? AND event_id = ?`
	result, err := r.db.ExecContext(ctx, query, userID, eventID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("not subscribed to this event")
	}
	return nil
}
