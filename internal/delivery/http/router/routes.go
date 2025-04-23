package router

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/lautaromdelgado/tecnica-backend/pkg/config"
)

func InitRoutes(e *echo.Echo, db *sqlx.DB, cfg *config.Config) {
	RegisterPublic(e, db, cfg)
	RegisterUser(e, db, cfg)
	RegisterAdmin(e, db, cfg)
}
