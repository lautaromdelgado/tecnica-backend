package main

import (
	"github.com/labstack/echo/v4"
	handler_user "github.com/lautaromdelgado/tecnica-backend/delivery/http/handler/auth"
	handler_event "github.com/lautaromdelgado/tecnica-backend/delivery/http/handler/event"
	"github.com/lautaromdelgado/tecnica-backend/delivery/http/router"
	"github.com/lautaromdelgado/tecnica-backend/infrastructure/persistence"
	persistence_event "github.com/lautaromdelgado/tecnica-backend/infrastructure/persistence/event"
	persistence_user "github.com/lautaromdelgado/tecnica-backend/infrastructure/persistence/user"
	"github.com/lautaromdelgado/tecnica-backend/pkg/config"
	usecase_auth "github.com/lautaromdelgado/tecnica-backend/usecase/auth"
	usecase_event "github.com/lautaromdelgado/tecnica-backend/usecase/event"
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

	// Dependencias para eventos
	eventRepo := persistence_event.NewEventRepository(db)
	eventUC := usecase_event.NewEventUseCase(eventRepo)
	eventHandler := handler_event.NewEventHandler(eventUC)

	// Inicializar rutas
	router.InitRoutes(e, authHandler, eventHandler, secret)

	// Iniciar servidor
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
