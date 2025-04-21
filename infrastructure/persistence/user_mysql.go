package persistence

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	model "github.com/lautaromdelgado/tecnica-backend/internal/domain/model/user"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create inserta un nuevo usuario en la base de datos
func (r *UserRepository) Create(ctx context.Context, u *model.User) error {
	query := `INSERT INTO users (username, email, role) VALUES (?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, u.Username, u.Email, u.Role)
	return err
}

// GetByEmail busca un usuario en la base de datos por su nombre de usuario y correo electrónico
func (r *UserRepository) GetByEmail(ctx context.Context, username, email string) (*model.User, error) {
	var user model.User
	query := `SELECT * FROM users WHERE username = ? AND email = ?`
	/* if err := r.db.GetContext(ctx, &user, query, username, email); err != nil {
		return nil, err
	} */
	err := r.db.GetContext(ctx, &user, query, username, email)
	if err != nil {
		log.Printf("❌ Error en GetByEmail: %v", err)
		return nil, err
	}

	return &user, nil
}

// UpdateByID actualiza un usuario en la base de datos por su ID
func (r *UserRepository) UpdateByID(ctx context.Context, u *model.User) error {
	query := `UPDATE users SET username = ?, email = ?, updated_at = NOW() WHERE id = ? AND deleted_at IS NULL`
	_, err := r.db.ExecContext(ctx, query, u.Username, u.Email, u.ID)
	return err
}
