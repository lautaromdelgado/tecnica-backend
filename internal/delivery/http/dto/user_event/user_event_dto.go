package dto

type UserEventRequest struct {
	UserID  uint `json:"user_id"`
	EventID uint `json:"event_id"`
}
