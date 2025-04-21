package router

import (
	"github.com/labstack/echo/v4"
	handler "github.com/lautaromdelgado/tecnica-backend/delivery/http/handler/auth"
	middleware "github.com/lautaromdelgado/tecnica-backend/delivery/http/middleware/jwt"
)

func InitRoutes(e *echo.Echo, authHandler *handler.AuthHandler, secret string) {
	// Rutas públicas
	// Ruta: /api
	public := e.Group("/api")
	public.POST("/register", authHandler.Register)
	public.POST("/login", authHandler.Login)

	// Rutas privadas (requieren JWT)
	// Ruta: /api
	private := e.Group("/api")
	private.Use(middleware.JWTMiddleware(secret))

	// Rutas privadas para administradores
	// Ruta: /api/admin/
	admin := private.Group("/admin")
	admin.Use(middleware.OnlyAdmin()) // Middleware para verificar si el usuario es admin

	// USUARIOS
	private.PUT("/update/user", authHandler.Update)    // Actualizar usuario por ID
	private.DELETE("/delete/user", authHandler.Delete) // Eliminar usuario por ID

	// ADMINISTRADORES
	admin.DELETE("/delete/user/:id", authHandler.DeleteByID) // Eliminar usuario por ID siendo admin
}
