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

// Create inserta un nuevo evento
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

// Update actualiza un evento existente
func (r *eventRepo) Update(ctx context.Context, event *model.Event) error {
	query := `
		UPDATE events 
		SET 
			organizer = ?, 
			title = ?, 
			long_description = ?, 
			short_description = ?, 
			date = ?, 
			location = ?, 
			is_published = ?, 
			updated_at = NOW()
		WHERE id = ? AND deleted_at IS NULL
	`
	_, err := r.db.ExecContext(ctx, query, event.Organizer,
		event.Title,
		event.LongDescription,
		event.ShortDescription,
		event.Date,
		event.Location,
		event.IsPublished,
		event.ID)
	return err
}

// Delete marca un evento como eliminado (soft delete)
func (r *eventRepo) Delete(ctx context.Context, id uint) error {
	query := `UPDATE events SET deleted_at = NOW(), updated_at = NOW() WHERE id = ? AND deleted_at IS NULL`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// UpdatePublishStatus actualiza el estado de publicación de un evento por ID
func (r *eventRepo) UpdatePublishStatus(ctx context.Context, id uint, publish bool) error {
	query := `UPDATE events SET is_published = ?, updated_at = NOW() WHERE id = ? AND deleted_at IS NULL`
	_, err := r.db.ExecContext(ctx, query, publish, id)
	return err
}

// RestoreByID restaura un evento por ID (soft delete)
func (r *eventRepo) RestoreByID(ctx context.Context, id uint) error {
	query := `UPDATE events SET deleted_at = NULL, updated_at = NOW() WHERE id = ? AND deleted_at IS NOT NULL`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
