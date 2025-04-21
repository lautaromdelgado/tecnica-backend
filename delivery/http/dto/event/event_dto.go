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
