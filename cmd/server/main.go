package main

import (
	"github.com/labstack/echo/v4"
	handler "github.com/lautaromdelgado/tecnica-backend/delivery/http/handler/auth"
	"github.com/lautaromdelgado/tecnica-backend/delivery/http/router"
	"github.com/lautaromdelgado/tecnica-backend/infrastructure/persistence"
	"github.com/lautaromdelgado/tecnica-backend/pkg/config"
	usecase "github.com/lautaromdelgado/tecnica-backend/usecase/auth"
)

func main() {
	cfg := config.LoadConfig()
	secret := cfg.JWTSecret
	e := echo.New()

	// Conexión a base de datos
	db, err := persistence.NewDB(cfg)
	if err != nil {
		e.Logger.Fatal("Error al conectar a la base de datos:", err)
	}
	defer db.Close()

	// Dependencias
	userRepo := persistence.NewUserRepository(db)
	authUC := usecase.NewAuthUseCase(userRepo)
	authHandler := handler.NewAuthHandler(authUC, cfg.JWTSecret)

	router.InitRoutes(e, authHandler, secret)

	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
