package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/lautaromdelgado/tecnica-backend/infrastructure/token"
)

func GetUserFromContext(c echo.Context) (uint, string, error) {
	claims, ok := c.Get("user").(*token.CustomClaims)
	if !ok {
		return 0, "", echo.ErrUnauthorized
	}
	return claims.UserID, claims.Role, nil
}
