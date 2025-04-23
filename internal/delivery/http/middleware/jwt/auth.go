package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/lautaromdelgado/tecnica-backend/internal/delivery/http/middleware"
)

// OnlyAdmin es un middleware que verifica si el usuario tiene el rol de administrador
// y permite el acceso a la ruta solo si es así. Si no, devuelve un error 403 Forbidden.
func OnlyAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			_, role, err := middleware.GetUserFromContext(c)
			if err != nil || role != "admin" {
				return echo.NewHTTPError(403, "Admin only")
			}
			return next(c)
		}
	}
}
