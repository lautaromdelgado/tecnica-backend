package model

import "time"

type User struct {
	ID        uint       `db:"id"`
	Username  string     `db:"username"`
	Email     string     `db:"email"`
	Role      string     `db:"role"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
