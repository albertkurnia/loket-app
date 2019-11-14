package presenter

import (
	"fmt"
	"loket-app/helper"
	"loket-app/modules/location/model"
	"loket-app/modules/location/sanitizer"
	"loket-app/modules/location/usecase"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo"
)

// locationServiceHTTPHandler - http handler struct for location service.
type locationServiceHTTPHandler struct {
	locationUseCase usecase.LocationUseCase
}

// NewLocationServiceHandler - function for initiating new location service handler.
func NewLocationServiceHandler(locationUseCase usecase.LocationUseCase) *locationServiceHTTPHandler {
	return &locationServiceHTTPHandler{
		locationUseCase: locationUseCase,
	}
}

// Mount - mounting endpoint(s) by echo framework grouping.
func (h *locationServiceHTTPHandler) Mount(group *echo.Group) {
	group.POST("/create", h.CreateLocation)
}

// CreateLocation - http handler function for creating location.
func (h *locationServiceHTTPHandler) CreateLocation(c echo.Context) error {
	logCtx := fmt.Sprintf("%T.CreateLocation", *h)

	params := model.CreateLocationReq{}
	if err := c.Bind(&params); err != nil {
		helper.Log(log.ErrorLevel, err.Error(), logCtx, "error_bind_params")
		return helper.NewResponse(http.StatusBadRequest, http.StatusBadRequest, err.Error(), nil).WriteResponse(c)
	}

	if err := sanitizer.ValidateLocation(&params); err != nil {
		helper.Log(log.ErrorLevel, err.Error(), logCtx, "error_validate_location")
		return helper.NewResponse(http.StatusBadRequest, http.StatusBadRequest, err.Error(), nil).WriteResponse(c)
	}

	resp, err := h.locationUseCase.CreateLocation(&params)
	if err != nil {
		helper.Log(log.ErrorLevel, err.Error(), logCtx, "error_create_location")
		return helper.NewResponse(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), nil).WriteResponse(c)
	}

	data := make(map[string]interface{})
	data["location"] = resp
	return helper.NewResponse(http.StatusCreated, http.StatusCreated, "Success", data).WriteResponse(c)
}
