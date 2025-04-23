package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lautaromdelgado/tecnica-backend/delivery/http/middleware"
	usecase "github.com/lautaromdelgado/tecnica-backend/usecase/user_event"
)

type UserEventHandler struct {
	eventUserEvent usecase.UserEventUseCase
}

func NewUserEventHandler(ue usecase.UserEventUseCase) *UserEventHandler {
	return &UserEventHandler{
		eventUserEvent: ue,
	}
}

// Subscribe suscribe a un usuario a un evento
func (h *UserEventHandler) Subscribe(c echo.Context) error {
	userID, _, err := middleware.GetUserFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	}

	idParam := c.Param("id")
	eventID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid event ID")
	}

	err = h.eventUserEvent.SuscribeUserToEvent(c.Request().Context(), userID, uint(eventID))
	if err != nil {
		switch err.Error() {
		case "event not found":
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		case "event is not published", "event has already occurred", "event has been deleted":
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		case "user already subscribed to this event":
			return echo.NewHTTPError(http.StatusConflict, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, "could not subscribe")
		}
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "user subscribed to event successfully"})
}
