package main

import (
	"github.com/labstack/echo/v4"
	handler_user "github.com/lautaromdelgado/tecnica-backend/internal/delivery/http/handler/auth"
	handler_event "github.com/lautaromdelgado/tecnica-backend/internal/delivery/http/handler/event"
	handler_user_event "github.com/lautaromdelgado/tecnica-backend/internal/delivery/http/handler/user_event"
	"github.com/lautaromdelgado/tecnica-backend/internal/delivery/http/router"
	"github.com/lautaromdelgado/tecnica-backend/internal/infrastructure/persistence"
	persistence_event "github.com/lautaromdelgado/tecnica-backend/internal/infrastructure/persistence/event"
	persistence_event_log "github.com/lautaromdelgado/tecnica-backend/internal/infrastructure/persistence/event_log"
	persistence_user "github.com/lautaromdelgado/tecnica-backend/internal/infrastructure/persistence/user"
	persistence_user_event "github.com/lautaromdelgado/tecnica-backend/internal/infrastructure/persistence/user_event"
	usecase_auth "github.com/lautaromdelgado/tecnica-backend/internal/usecase/auth"
	usecase_event "github.com/lautaromdelgado/tecnica-backend/internal/usecase/event"
	usecase_user_event "github.com/lautaromdelgado/tecnica-backend/internal/usecase/user_event"
	"github.com/lautaromdelgado/tecnica-backend/pkg/config"
)

func main() {
	cfg := config.LoadConfig() // Cargar configuración desde el archivo .env
	secret := cfg.JWTSecret    // Clave secreta para JWT
	e := echo.New()

	// Conexión a base de datos
	db, err := persistence.NewDB(cfg)
	if err != nil {
		e.Logger.Fatal("Error al conectar a la base de datos:", err)
	}
	defer db.Close()

	// Dependencias para usuarios
	userRepo := persistence_user.NewUserRepository(db)
	authUC := usecase_auth.NewAuthUseCase(userRepo)
	authHandler := handler_user.NewAuthHandler(authUC, cfg.JWTSecret)

	// Dependencia para logs de eventos
	eventLogRepo := persistence_event_log.NewEventLogRepository(db)

	// Dependencias para eventos
	eventRepo := persistence_event.NewEventRepository(db)
	eventUC := usecase_event.NewEventUseCase(eventRepo, eventLogRepo)
	eventHandler := handler_event.NewEventHandler(eventUC, eventLogRepo)

	// Dependencias para usuarios y eventos
	userEventRepo := persistence_user_event.NewUserEventRepository(db)
	userEventUC := usecase_user_event.NewUserEventUseCase(userEventRepo, eventRepo)
	userEventHandler := handler_user_event.NewUserEventHandler(userEventUC)

	// Inicializar rutas
	router.InitRoutes(e, authHandler, eventHandler, userEventHandler, secret)

	// Iniciar servidor
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
