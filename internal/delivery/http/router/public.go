package router

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	handler_auth "github.com/lautaromdelgado/tecnica-backend/internal/delivery/http/handler/auth"
	persistence "github.com/lautaromdelgado/tecnica-backend/internal/infrastructure/persistence/user"
	usecase "github.com/lautaromdelgado/tecnica-backend/internal/usecase/auth"
	"github.com/lautaromdelgado/tecnica-backend/pkg/config"
)

// Register registra las rutas públicas en el servidor Echo
// y asocia los controladores correspondientes a cada ruta.
func RegisterPublic(e *echo.Echo, db *sqlx.DB, cfg *config.Config) {
	userRepo := persistence.NewUserRepository(db)
	authUC := usecase.NewAuthUseCase(userRepo)
	authHandler := handler_auth.NewAuthHandler(authUC, cfg.JWTSecret)

	public := e.Group("/api")
	public.POST("/register", authHandler.Register)
	public.POST("/login", authHandler.Login)
}
