package presenter

import (
	"loket-app/modules/event/usecase"

	"github.com/labstack/echo"
)

type eventServiceHTTPHandler struct {
	eventUseCase usecase.EventUseCase
}

func NewEventServiceHandler(eventUseCase usecase.EventUseCase) *eventServiceHTTPHandler {
	return &eventServiceHTTPHandler{
		eventUseCase: eventUseCase,
	}
}

func (h *eventServiceHTTPHandler) Mount(group *echo.Group) {
	group.POST("/create", h.CreateEvent)
	group.POST("/ticket/create", h.CreateTicket)
	group.GET("/get_info", h.GetEventInfo)
}

func (h *eventServiceHTTPHandler) CreateEvent(c echo.Context) error {
	return nil
}

func (h *eventServiceHTTPHandler) CreateTicket(c echo.Context) error {
	return nil
}

func (h *eventServiceHTTPHandler) GetEventInfo(c echo.Context) error {
	return nil
}
