package persistence

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/lautaromdelgado/tecnica-backend/pkg/config"
)

func NewDB(cfg *config.Config) (*sqlx.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName)

	db, err := sqlx.Connect("mysql", dns)
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}
	// Probar conexión
	if err := db.Ping(); err != nil {
		log.Fatalf("Error haciendo ping a la base de datos: %v", err)
	}
	// Configurar el pool de conexiones
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	log.Printf("Conectado a la base de datos %s en %s:%s", cfg.DBName, cfg.DBHost, cfg.DBPort)
	return db, nil
}
