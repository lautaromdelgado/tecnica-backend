package main

import (
	"github.com/lautaromdelgado/tecnica-backend/internal/delivery/http/server"
	"github.com/lautaromdelgado/tecnica-backend/pkg/config"
)

func main() {
	cfg := config.LoadConfig() // Cargar configuración desde el archivo .env
	app := server.New(cfg)     // Crear una nueva instancia de Echo
	app.Run()                  // Iniciar el servidor
}
