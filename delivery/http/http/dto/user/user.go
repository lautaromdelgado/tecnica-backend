package dto

// Iniciar sesión
type LoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Registrar usuario
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role" validate:"oneof=admin user"`
}

// Actualizar usuario
type UpdateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}
