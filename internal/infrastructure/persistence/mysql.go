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

	var db *sqlx.DB
	var err error
	for i := 0; i < 10; i++ {
		db, err = sqlx.Connect("mysql", dns)
		if err == nil {
			break
		}
		log.Printf("Intento %d de conexión fallido: %v", i+1, err)
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	// Probar conexión
	if err := db.Ping(); err != nil {
		log.Fatalf("Error haciendo ping a la base de datos: %v", err)
	}

	// Configurar pool de conexiones
	db.SetMaxOpenConns(25)                 // conexiones máximas abiertas
	db.SetMaxIdleConns(25)                 // conexiones en espera
	db.SetConnMaxLifetime(5 * time.Minute) // duración máxima de una conexión
	log.Printf("Conectado a la base de datos %s en %s:%s", cfg.DBName, cfg.DBHost, cfg.DBPort)
	return db, nil
}
