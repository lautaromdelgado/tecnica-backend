package router

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	middleware "github.com/lautaromdelgado/tecnica-backend/internal/delivery/http/middleware/jwt"
	persistence_event "github.com/lautaromdelgado/tecnica-backend/internal/infrastructure/persistence/event"
	persistence_event_log "github.com/lautaromdelgado/tecnica-backend/internal/infrastructure/persistence/event_log"
	persistence "github.com/lautaromdelgado/tecnica-backend/internal/infrastructure/persistence/user"
	persistence_user_event "github.com/lautaromdelgado/tecnica-backend/internal/infrastructure/persistence/user_event"
	usecase_auth "github.com/lautaromdelgado/tecnica-backend/internal/usecase/auth"
	usecase_event "github.com/lautaromdelgado/tecnica-backend/internal/usecase/event"
	usecase_event_log "github.com/lautaromdelgado/tecnica-backend/internal/usecase/event_log"
	usecase_user_event "github.com/lautaromdelgado/tecnica-backend/internal/usecase/user_event"
	"github.com/lautaromdelgado/tecnica-backend/pkg/config"

	authHandler "github.com/lautaromdelgado/tecnica-backend/internal/delivery/http/handler/auth"
	eventHandler "github.com/lautaromdelgado/tecnica-backend/internal/delivery/http/handler/event"
	userEventHandler "github.com/lautaromdelgado/tecnica-backend/internal/delivery/http/handler/user_event"
)

func RegisterUser(e *echo.Echo, db *sqlx.DB, cfg *config.Config) {
	// Repos
	userRepo := persistence.NewUserRepository(db)
	userEventRepo := persistence_user_event.NewUserEventRepository(db)
	eventRepo := persistence_event.NewEventRepository(db)
	eventLogRepo := persistence_event_log.NewEventLogRepository(db)
	// Use Cases
	authUC := usecase_auth.NewAuthUseCase(userRepo)
	userEventUC := usecase_user_event.NewUserEventUseCase(userEventRepo, eventRepo)
	eventUC := usecase_event.NewEventUseCase(eventRepo, eventLogRepo)
	eventLogUC := usecase_event_log.NewEventLogUseCase(eventLogRepo)
	// Handlers
	authHandler := authHandler.NewAuthHandler(authUC, cfg.JWTSecret)
	eventHandler := eventHandler.NewEventHandler(eventUC, eventLogUC)
	userEventHandler := userEventHandler.NewUserEventHandler(userEventUC)

	private := e.Group("/api")
	private.Use(middleware.JWTMiddleware(cfg.JWTSecret))

	// Usuario autenticado
	private.PUT("/update/user", authHandler.Update)    // Actualizar mi usuario
	private.DELETE("/delete/user", authHandler.Delete) // Eliminar mi cuenta

	// Eventos accesibles para usuarios
	private.GET("/events/:id", eventHandler.GetEventByID)                   // Obtener evento por ID (si está publicado)
	private.GET("/events/search", eventHandler.SearchEvents)                // Buscar eventos (solo los publicados para usuarios)
	private.POST("/events/:id/subscribe", userEventHandler.Subscribe)       // Inscribirse a evento (si es futuro y publicado)
	private.GET("/events/myevents", userEventHandler.MyEvents)              // Ver eventos en los que estoy inscrito
	private.DELETE("/events/:id/unsubscribe", userEventHandler.Unsubscribe) // Cancelar inscripción

}
