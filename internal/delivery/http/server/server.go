package server

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/lautaromdelgado/tecnica-backend/internal/delivery/http/router"
	"github.com/lautaromdelgado/tecnica-backend/internal/infrastructure/persistence"
	"github.com/lautaromdelgado/tecnica-backend/pkg/config"
)

// Server representa la estructura del servidor HTTP
type Server struct {
	app *echo.Echo
	cfg *config.Config
	db  *sqlx.DB
}

// New crea una nueva instancia de Server
func New(cfg *config.Config) *Server {
	app := echo.New()

	db, err := persistence.NewDB(cfg)
	if err != nil {
		app.Logger.Fatal("Error connecting to database:", err)
	}

	// Iniciar rutas
	router.InitRoutes(app, db, cfg)
	return &Server{
		app: app,
		cfg: cfg,
		db:  db,
	}
}

func (s *Server) Run() {
	s.app.Logger.Fatal(s.app.Start(":" + s.cfg.Port))
}
