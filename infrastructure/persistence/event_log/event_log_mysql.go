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

// LogAction registra una acción de evento
func (r *eventLogRepo) LogAction(ctx context.Context, log *model.EventLog) error {
	query := `INSERT INTO event_logs (title, organizer, action) VALUES (?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, log.Title, log.Organizer, log.Action)
	return err
}

// GetAllLogs obtiene todos los logs de eventos
func (r *eventLogRepo) GetAllLogs(ctx context.Context) ([]*model.EventLog, error) {
	query := `SELECT * FROM event_logs ORDER BY timestamp DESC`
	var logs []*model.EventLog
	if err := r.db.SelectContext(ctx, &logs, query); err != nil {
		return nil, err
	}
	return logs, nil
}

// GetLogsByTittle obtiene logs de eventos por título
func (r *eventLogRepo) GetLogsByFilters(ctx context.Context, title, action, organizer string) ([]*model.EventLog, error) {
	query := `SELECT * FROM event_logs WHERE 1=1`
	var args []interface{}

	if title != "" {
		query += ` AND LOWER(title) LIKE LOWER(?)`
		args = append(args, "%"+title+"%")
	}
	if action != "" {
		query += ` AND action = ?`
		args = append(args, action)
	}
	if organizer != "" {
		query += ` AND LOWER(organizer) LIKE LOWER(?)`
		args = append(args, "%"+organizer+"%")
	}

	query += ` ORDER BY timestamp DESC`

	var logs []*model.EventLog
	err := r.db.SelectContext(ctx, &logs, query, args...)
	return logs, err
}
