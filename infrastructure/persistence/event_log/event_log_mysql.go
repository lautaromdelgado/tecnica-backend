package persistence

import (
	"context"

	"github.com/jmoiron/sqlx"
	model "github.com/lautaromdelgado/tecnica-backend/internal/domain/model/event_log"
)

type eventLogRepo struct {
	db *sqlx.DB
}

func NewEventLogRepository(db *sqlx.DB) *eventLogRepo {
	return &eventLogRepo{db: db}
}

// LogAction registra una acción de evento en la base de datos
func (r *eventLogRepo) LogAction(ctx context.Context, log *model.EventLog) error {
	query := `INSERT INTO event_logs (title, organizer, action) VALUES (?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, log.Title, log.Organizer, log.Action)
	return err
}
