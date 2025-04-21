package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	dto "github.com/lautaromdelgado/tecnica-backend/delivery/http/dto/event"
	"github.com/lautaromdelgado/tecnica-backend/delivery/http/middleware"
	model "github.com/lautaromdelgado/tecnica-backend/internal/domain/model/event"
	usecase "github.com/lautaromdelgado/tecnica-backend/usecase/event"
)

type EventHandler struct {
	eventUC usecase.EventUseCase
}

func NewEventHandler(uc usecase.EventUseCase) *EventHandler {
	return &EventHandler{
		eventUC: uc,
	}
}

// Create crea un nuevo evento
func (h *EventHandler) Create(c echo.Context) error {
	// Obtener el rol del usuario desde el contexto
	_, role, err := middleware.GetUserFromContext(c)
	if err != nil || role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, "admin only")
	}

	var req dto.CreateEventRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	event := &model.Event{
		Organizer:        req.Organizer,
		Title:            req.Title,
		LongDescription:  req.LongDescription,
		ShortDescription: req.ShortDescription,
		Date:             req.Date,
		Location:         req.Location,
		IsPublished:      false, // por defecto, opcional
	}

	// Solo si el campo fue enviado, lo pasás como valor
	if req.IsPublished != nil {
		event.IsPublished = *req.IsPublished
	}

	if err := h.eventUC.CreateEvent(c.Request().Context(), event); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create event")
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "event created successfully",
	})
}
