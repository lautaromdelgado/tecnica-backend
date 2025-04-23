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

// FindByID busca un evento por ID
func (r *eventRepo) FindByID(ctx context.Context, id uint) (*model.Event, error) {
	query := `SELECT * FROM events WHERE id = ? LIMIT 1`
	var event model.Event
	if err := r.db.GetContext(ctx, &event, query, id); err != nil {
		return nil, err
	}
	return &event, nil
}

// FindWhitFilters busca eventos por filtros
// Titulo del evento : organizador : ubicacion
// incluideDrafts: si se incluyen eventos no publicados (solo para administradores)
func (r *eventRepo) FindWhitFilters(ctx context.Context, title, organizer, location string, incluideDrafts bool) ([]*model.Event, error) {
	query := `SELECT * FROM events WHERE deleted_at IS NULL`
	var args []interface{}

	// Para mostrar los eventos no publicados
	// Solo pueden ver estos eventos los administradores
	if !incluideDrafts {
		query += ` AND is_published = true`
	}
	if title != "" {
		query += ` AND LOWER(title) LIKE LOWER(?)`
		args = append(args, "%"+title+"%")
	}
	if organizer != "" {
		query += ` AND LOWER(organizer) LIKE LOKER(?)`
		args = append(args, "%"+organizer+"%")
	}
	if location != "" {
		query += ` AND LOWER(location) LIKE LOWER(?)`
		args = append(args, "%"+location+"%")
	}
	query += ` ORDER BY date DESC`

	var events []*model.Event
	err := r.db.SelectContext(ctx, &events, query, args...)
	return events, err
}
