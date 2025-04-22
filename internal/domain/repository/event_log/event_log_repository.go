package repository

import (
	"context"

	model "github.com/lautaromdelgado/tecnica-backend/internal/domain/model/event_log"
)

type EventLogRepository interface {
	LogAction(ctx context.Context, log *model.EventLog) error  // Registra una acción de evento
	GetAllLogs(ctx context.Context) ([]*model.EventLog, error) // Obtiene todos los logs de eventos
}
