package model

type EventLog struct {
	ID        uint   `db:"id"`
	Title     string `db:"title"`
	Organizer string `db:"organizer"`
	Action    string `db:"action"`
	Timestamp string `db:"timestamp"`
}
