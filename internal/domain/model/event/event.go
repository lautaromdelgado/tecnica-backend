package model

type Event struct {
	ID               uint    `db:"id"`
	Organizer        string  `db:"organizer"`
	Title            string  `db:"title"`
	LongDescription  string  `db:"long_description"`
	ShortDescription string  `db:"short_description"`
	Date             int64   `db:"date"` // timestamp UNIX
	Location         string  `db:"location"`
	IsPublished      bool    `db:"is_published"`
	CreatedAt        string  `db:"created_at"`
	UpdatedAt        string  `db:"updated_at"`
	DeletedAt        *string `db:"deleted_at"` // NIL POR DEFECTO = Nil dice que el evento no está eliminado
}
