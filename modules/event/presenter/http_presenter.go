package presenter

import (
	"errors"
	"fmt"
	"loket-app/helper"
	"loket-app/modules/event/model"
	"loket-app/modules/event/sanitizer"
	"loket-app/modules/event/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
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
	logCtx := fmt.Sprintf("%T.CreateEvent", *h)

	params := model.CreateEventReq{}
	if err := c.Bind(&params); err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_bind_params")
		return helper.NewResponse(http.StatusBadRequest, http.StatusBadRequest, err.Error(), nil).WriteResponse(c)
	}

	if err := sanitizer.ValidateCreateEvent(&params); err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_validate_params")
		return helper.NewResponse(http.StatusBadRequest, http.StatusBadRequest, err.Error(), nil).WriteResponse(c)
	}

	resp, err := h.eventUseCase.CreateEvent(&params)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_create_event")
		return helper.NewResponse(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), nil).WriteResponse(c)
	}

	data := make(map[string]interface{})
	data["event"] = resp
	return helper.NewResponse(http.StatusCreated, http.StatusCreated, "Success", data).WriteResponse(c)
}

func (h *eventServiceHTTPHandler) CreateTicket(c echo.Context) error {
	logCtx := fmt.Sprintf("%T.CreateTicket", *h)

	params := model.CreateTicketReq{}
	if err := c.Bind(&params); err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_bind_params")
		return helper.NewResponse(http.StatusBadRequest, http.StatusBadRequest, err.Error(), nil).WriteResponse(c)
	}

	if err := sanitizer.ValidateCreateTicket(&params); err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_validate_params")
		return helper.NewResponse(http.StatusBadRequest, http.StatusBadRequest, err.Error(), nil).WriteResponse(c)
	}

	resp, err := h.eventUseCase.CreateTicket(&params)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_create_ticket")
		return helper.NewResponse(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), nil).WriteResponse(c)
	}

	data := make(map[string]interface{})
	data["ticket"] = resp
	return helper.NewResponse(http.StatusCreated, http.StatusCreated, "Success", data).WriteResponse(c)
}

func (h *eventServiceHTTPHandler) GetEventInfo(c echo.Context) error {
	logCtx := fmt.Sprintf("%T.GetEventInfo", *h)

	eventID, err := strconv.ParseUint(c.QueryParam("id"), 10, 64)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_parse_uint")
		return helper.NewResponse(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), nil).WriteResponse(c)
	}

	if eventID <= 0 {
		err := errors.New("invalid event id")
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_event_id")
		return helper.NewResponse(http.StatusBadRequest, http.StatusBadRequest, err.Error(), nil).WriteResponse(c)
	}

	resp, err := h.eventUseCase.GetEventInformation(eventID)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_get_event_information")
		return helper.NewResponse(http.StatusBadRequest, http.StatusBadRequest, err.Error(), nil).WriteResponse(c)
	}

	data := make(map[string]interface{})
	data["event"] = resp
	return helper.NewResponse(http.StatusOK, http.StatusOK, "Success", data).WriteResponse(c)
}
