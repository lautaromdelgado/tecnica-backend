package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	dto "github.com/lautaromdelgado/tecnica-backend/delivery/http/dto/event"
	"github.com/lautaromdelgado/tecnica-backend/delivery/http/middleware"
	model "github.com/lautaromdelgado/tecnica-backend/internal/domain/model/event"
	usecase_event "github.com/lautaromdelgado/tecnica-backend/usecase/event"
	usecase_event_log "github.com/lautaromdelgado/tecnica-backend/usecase/event_log"
)

type EventHandler struct {
	eventUC    usecase_event.EventUseCase
	eventLogUC usecase_event_log.EventLogUseCase // Añadido para el EventLogUseCase
}

func NewEventHandler(uc usecase_event.EventUseCase, eventLogUC usecase_event_log.EventLogUseCase) *EventHandler {
	return &EventHandler{
		eventUC:    uc,
		eventLogUC: eventLogUC, // Inicializa el EventLogUseCase
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

// Update actualiza un evento existente
func (h *EventHandler) Update(c echo.Context) error {
	_, role, err := middleware.GetUserFromContext(c)
	if err != nil || role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, "admin only")
	}

	idParam := c.Param("id")
	eventID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid event id")
	}

	var req dto.UpdateEventRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	updated := &model.Event{
		ID:               uint(eventID),
		Organizer:        req.Organizer,
		Title:            req.Title,
		LongDescription:  req.LongDescription,
		ShortDescription: req.ShortDescription,
		Date:             req.Date,
		Location:         req.Location,
		IsPublished:      false, // por defecto
	}

	if req.IsPublished != nil {
		updated.IsPublished = *req.IsPublished
	}

	if err := h.eventUC.UpdateEvent(c.Request().Context(), updated); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update event")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "event updated successfully"})
}

// Delete elimina un evento por ID (marcando como eliminado)
func (h *EventHandler) Delete(c echo.Context) error {
	_, role, err := middleware.GetUserFromContext(c)
	if err != nil || role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, "admin only")
	}

	idParam := c.Param("id")
	eventID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid event id")
	}

	err = h.eventUC.DeleteEvent(c.Request().Context(), uint(eventID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not delete event")

	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "event deleted successfully"})
}

// UpdatePublishStatus actualiza el estado de publicación de un evento por ID
func (h *EventHandler) Publish(c echo.Context) error {
	return h.togglePublish(c, true)
}

// Unpublish despublica un evento por ID
func (h *EventHandler) Unpublish(c echo.Context) error {
	return h.togglePublish(c, false)
}

// togglePublish actualiza el estado de publicación de un evento por ID
// y verifica si el usuario tiene el rol de administrador
func (h *EventHandler) togglePublish(c echo.Context, publish bool) error {
	_, role, err := middleware.GetUserFromContext(c)
	if err != nil || role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, "admin only")
	}

	idParam := c.Param("id")
	eventID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid event id")
	}

	err = h.eventUC.UpdatePublishStatus(c.Request().Context(), uint(eventID), publish)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not update event publish status")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "event publish status updated successfully"})
}

// Restore restaura un evento por ID (soft delete)
func (h *EventHandler) Restore(c echo.Context) error {
	_, role, err := middleware.GetUserFromContext(c)
	if err != nil || role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, "admin only")
	}

	idParam := c.Param("id")
	eventID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid event id")
	}

	err = h.eventUC.RestoreByID(c.Request().Context(), uint(eventID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not restore event")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "event restored successfully"})
}

// GetLogs obtiene todos los logs de eventos
func (h *EventHandler) GetLogs(c echo.Context) error {
	_, role, err := middleware.GetUserFromContext(c)
	if err != nil || role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, "admin only")
	}
	logs, err := h.eventLogUC.GetAllLogs(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not fetch logs")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"logs": logs,
	})
}

// GetLogsFiltered obtiene logs de eventos por filtros (título, acción, organizador)
func (h *EventHandler) GetLogFiltered(c echo.Context) error {
	_, role, err := middleware.GetUserFromContext(c)
	if err != nil || role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, "admin only")
	}
	title := c.QueryParam("title")
	action := c.QueryParam("action")
	organizer := c.QueryParam("organizer")

	logs, err := h.eventLogUC.GetLogsByFilters(c.Request().Context(), title, action, organizer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not fetch logs")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"logs": logs})
}

func (h *EventHandler) GetEventByID(c echo.Context) error {
	role := "user"
	if user := c.Get("user"); user != nil {
		_, roleFromToken, err := middleware.GetUserFromContext(c)
		if err == nil {
			role = roleFromToken
		}
	}

	idParam := c.Param("id")
	eventID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid event id")
	}

	event, err := h.eventUC.GetByIDWhitPermissions(c.Request().Context(), uint(eventID), role)
	if err != nil {
		if err.Error() == "unauthorized to view this event" {
			return echo.NewHTTPError(http.StatusForbidden, "not authorized to view this event")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "event not found")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"event": event,
	})
}
