package dto

type LoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role" validate:"oneof=admin user"`
}
