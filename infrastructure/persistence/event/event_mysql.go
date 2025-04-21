package persistence

import (
	"context"

	"github.com/jmoiron/sqlx"
	model "github.com/lautaromdelgado/tecnica-backend/internal/domain/model/event"
)

type eventRepo struct {
	db *sqlx.DB
}

func NewEventRepository(db *sqlx.DB) *eventRepo {
	return &eventRepo{db: db}
}

// Create inserta un nuevo evento en la base de datos
func (r *eventRepo) Create(ctx context.Context, event *model.Event) error {
	query := `
		INSERT INTO events (
			organizer, title, long_description, short_description, date, location, is_published
		)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.ExecContext(ctx, query,
		event.Organizer,
		event.Title,
		event.LongDescription,
		event.ShortDescription,
		event.Date,
		event.Location,
		event.IsPublished)
	return err
}
