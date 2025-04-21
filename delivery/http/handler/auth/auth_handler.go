package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	dto "github.com/lautaromdelgado/tecnica-backend/delivery/http/http/dto/user"
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

func (h *AuthHandler) Login(c echo.Context) error {
	// TODO: Implementar un DTO para el login => una estructura usuario solo meial y username
	var req dto.RegisterRequest
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
