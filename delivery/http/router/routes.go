package router

import (
	"github.com/labstack/echo/v4"
	handler_auth "github.com/lautaromdelgado/tecnica-backend/delivery/http/handler/auth"
	handler_event "github.com/lautaromdelgado/tecnica-backend/delivery/http/handler/event"
	handler_user_event "github.com/lautaromdelgado/tecnica-backend/delivery/http/handler/user_event"
	middleware "github.com/lautaromdelgado/tecnica-backend/delivery/http/middleware/jwt"
)

func InitRoutes(e *echo.Echo, authHandler *handler_auth.AuthHandler, eventHandler *handler_event.EventHandler, userEventHandler *handler_user_event.UserEventHandler, secret string) {
	// Rutas públicas
	// Ruta: /api
	public := e.Group("/api")
	public.POST("/register", authHandler.Register)
	public.POST("/login", authHandler.Login)

	// Rutas privadas (requieren JWT)
	private := e.Group("/api")
	private.Use(middleware.JWTMiddleware(secret))

	// Rutas privadas para administradores
	admin := private.Group("/admin")
	admin.Use(middleware.OnlyAdmin()) // Middleware para verificar si el usuario es admin

	// USUARIOS /api
	private.PUT("/update/user", authHandler.Update)                   // Actualizar usuario por ID
	private.DELETE("/delete/user", authHandler.Delete)                // Eliminar usuario por ID
	private.GET("/events/:id", eventHandler.GetEventByID)             // Obtener evento por ID
	private.GET("/events/search", eventHandler.SearchEvents)          // Buscar eventos por filtros
	private.POST("/events/:id/subscribe", userEventHandler.Subscribe) // Suscribir usuario a evento por ID

	// ADMINISTRADORES = /api/admin/

	// Relacionado a los usuarios
	admin.DELETE("/delete/user/:id", authHandler.DeleteByID) // Eliminar usuario por ID siendo admin
	admin.GET("/users", authHandler.ListAll)                 // Listar todos los usuarios siendo admin
	admin.GET("/users/inactive", authHandler.ListInactive)   // Listar usuarios inactivos siendo admin
	admin.PUT("/users/:id/restore", authHandler.RestoreUser) // Restaurar usuario por ID siendo admin

	// Relacionado a los eventos
	admin.POST("/events/create", eventHandler.Create)          // Crear evento
	admin.PUT("/events/:id/update", eventHandler.Update)       // Actualizar evento por ID siendo admin
	admin.DELETE("/events/:id/delete", eventHandler.Delete)    // Eliminar evento por ID siendo admin
	admin.PUT("/events/:id/publish", eventHandler.Publish)     // Publicar evento por ID siendo admin
	admin.PUT("/events/:id/unpublish", eventHandler.Unpublish) // Despublicar evento por ID siendo admin
	admin.PUT("/events/:id/restore", eventHandler.Restore)     // Restaurar evento por ID siendo admin

	// Relacionado a los logs de eventos
	admin.GET("/events/logs", eventHandler.GetLogs)               // Obtener todos los logs de eventos siendo admin
	admin.GET("/events/logs/filter", eventHandler.GetLogFiltered) // Obtener logs de eventos por filtros siendo admin
}
