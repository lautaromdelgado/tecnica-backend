package router

import (
	"github.com/labstack/echo/v4"
	handler "github.com/lautaromdelgado/tecnica-backend/delivery/http/handler/auth"
	middleware "github.com/lautaromdelgado/tecnica-backend/delivery/http/middleware/jwt"
)

func InitRoutes(e *echo.Echo, authHandler *handler.AuthHandler, secret string) {
	// Rutas públicas
	public := e.Group("/api")
	public.POST("/register", authHandler.Register)
	public.POST("/login", authHandler.Login)

	// Rutas privadas (requieren JWT)
	private := e.Group("/api")
	private.Use(middleware.JWTMiddleware(secret))
	private.PUT("/update/user", authHandler.Update)
}
