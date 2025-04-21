package dto

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
