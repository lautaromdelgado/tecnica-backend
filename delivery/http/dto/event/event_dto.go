package dto

type CreateEventRequest struct {
	Organizer        string `json:"organizer" validate:"required"`
	Title            string `json:"title" validate:"required"`
	LongDescription  string `json:"long_description"`
	ShortDescription string `json:"short_description"`
	Date             int64  `json:"date" validate:"required"`
	Location         string `json:"location" validate:"required"`
	IsPublished      *bool  `json:"is_published,omitempty"`
}

type UpdateEventRequest struct {
	Organizer        string `json:"organizer" validate:"required"`
	Title            string `json:"title" validate:"required"`
	LongDescription  string `json:"long_description"`
	ShortDescription string `json:"short_description"`
	Date             int64  `json:"date" validate:"required"`
	Location         string `json:"location" validate:"required"`
	IsPublished      *bool  `json:"is_published,omitempty"`
}

type EventResponse struct {
	ID               uint   `json:"id"`
	Organizer        string `json:"organizer"`
	Title            string `json:"title"`
	LongDescription  string `json:"long_description"`
	ShortDescription string `json:"short_description"`
	Date             int64  `json:"date"`
	Location         string `json:"location"`
	IsPublished      bool   `json:"is_published"`
}
