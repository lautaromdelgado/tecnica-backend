package model

import "time"

type UserEvent struct {
	UserID   uint      `db:"user_id"`
	EventID  uint      `db:"event_id"`
	JoinedAt time.Time `db:"joined_at"`
}
