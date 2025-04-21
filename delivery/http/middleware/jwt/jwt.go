package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/lautaromdelgado/tecnica-backend/infrastructure/token"
)

func JWTMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid Authorization header")
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

			parsedToken, err := jwt.ParseWithClaims(tokenStr, &token.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})
			if err != nil || !parsedToken.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
			}

			claims, ok := parsedToken.Claims.(*token.CustomClaims)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
			}

			// Guardar claims en contexto
			c.Set("user", claims)
			return next(c)
		}
	}
}
