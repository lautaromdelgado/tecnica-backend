package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	dto "github.com/lautaromdelgado/tecnica-backend/delivery/http/http/dto/user"
	"github.com/lautaromdelgado/tecnica-backend/delivery/http/middleware"
	"github.com/lautaromdelgado/tecnica-backend/infrastructure/token"
	model "github.com/lautaromdelgado/tecnica-backend/internal/domain/model/user"
	usecase "github.com/lautaromdelgado/tecnica-backend/usecase/auth"
)

type AuthHandler struct {
	authUC    usecase.AuthUseCase
	jwtSecret string
}

func NewAuthHandler(usecase usecase.AuthUseCase, secret string) *AuthHandler {
	return &AuthHandler{
		authUC:    usecase,
		jwtSecret: secret,
	}
}

// crear un nuevo usuario
func (h *AuthHandler) Register(c echo.Context) error {
	var request model.User
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid input")
	}
	if err := h.authUC.Register(c.Request().Context(), &request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, echo.Map{"message": "user created"})
}

// inicia sesión de usuario
func (h *AuthHandler) Login(c echo.Context) error {
	var req dto.LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid input")
	}
	user, err := h.authUC.Login(c.Request().Context(), req.Username, req.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	token, err := token.GenerateToken(user.ID, user.Role, h.jwtSecret)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{"token": token})
}

// actualizar usuario por id
func (h *AuthHandler) Update(c echo.Context) error {
	var req dto.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid input")
	}
	userID, _, err := middleware.GetUserFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	user := &model.User{
		ID:       userID,
		Username: req.Username,
		Email:    req.Email,
	}

	if err := h.authUC.UpdateByID(c.Request().Context(), user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "user updated"})
}

// eliminar usuario por id
func (h *AuthHandler) Delete(c echo.Context) error {
	userID, _, err := middleware.GetUserFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	}
	if err := h.authUC.DeleteByID(c.Request().Context(), userID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error deleting user")
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "user deleted"})
}
