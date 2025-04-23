package router

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	authHandler "github.com/lautaromdelgado/tecnica-backend/internal/delivery/http/handler/auth"
	eventHandler "github.com/lautaromdelgado/tecnica-backend/internal/delivery/http/handler/event"
	middleware "github.com/lautaromdelgado/tecnica-backend/internal/delivery/http/middleware/jwt"
	persistence_event "github.com/lautaromdelgado/tecnica-backend/internal/infrastructure/persistence/event"
	persistence_event_log "github.com/lautaromdelgado/tecnica-backend/internal/infrastructure/persistence/event_log"
	persistence "github.com/lautaromdelgado/tecnica-backend/internal/infrastructure/persistence/user"
	usecase_auth "github.com/lautaromdelgado/tecnica-backend/internal/usecase/auth"
	usecase_event "github.com/lautaromdelgado/tecnica-backend/internal/usecase/event"
	usecase_event_log "github.com/lautaromdelgado/tecnica-backend/internal/usecase/event_log"
	"github.com/lautaromdelgado/tecnica-backend/pkg/config"
)

func RegisterAdmin(e *echo.Echo, db *sqlx.DB, cfg *config.Config) {
	private := e.Group("/api")
	private.Use(middleware.JWTMiddleware(cfg.JWTSecret))
	admin := private.Group("/admin")
	admin.Use(middleware.OnlyAdmin())

	// Repos
	userRepo := persistence.NewUserRepository(db)
	eventRepo := persistence_event.NewEventRepository(db)
	eventLogRepo := persistence_event_log.NewEventLogRepository(db)
	// Use Cases
	authUC := usecase_auth.NewAuthUseCase(userRepo)
	eventUC := usecase_event.NewEventUseCase(eventRepo, eventLogRepo)
	eventLogUC := usecase_event_log.NewEventLogUseCase(eventLogRepo)
	// Handlers
	authHandler := authHandler.NewAuthHandler(authUC, cfg.JWTSecret)
	eventHandler := eventHandler.NewEventHandler(eventUC, eventLogUC)

	// Administración de usuarios
	admin.DELETE("/delete/user/:id", authHandler.DeleteByID) // Eliminar usuario
	admin.GET("/users", authHandler.ListAll)                 // Listar usuarios activos
	admin.GET("/users/inactive", authHandler.ListInactive)   // Listar usuarios inactivos
	admin.PUT("/users/:id/restore", authHandler.RestoreUser) // Restaurar usuario eliminado

	// Administración de eventos
	admin.POST("/events/create", eventHandler.Create)          // Crear evento
	admin.PUT("/events/:id/update", eventHandler.Update)       // Actualizar evento
	admin.DELETE("/events/:id/delete", eventHandler.Delete)    // Eliminar evento
	admin.PUT("/events/:id/publish", eventHandler.Publish)     // Publicar evento
	admin.PUT("/events/:id/unpublish", eventHandler.Unpublish) // Despublicar evento
	admin.PUT("/events/:id/restore", eventHandler.Restore)     // Restaurar evento

	// Logs de eventos
	admin.GET("/events/logs", eventHandler.GetLogs)               // Obtener todos los logs
	admin.GET("/events/logs/filter", eventHandler.GetLogFiltered) // Filtrar logs por title, organizer, action

}
