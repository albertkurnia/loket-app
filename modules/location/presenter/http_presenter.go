package presenter

import (
	"loket-app/modules/location/usecase"

	"github.com/labstack/echo"
)

type locationServiceHTTPHandler struct {
	locationUseCase usecase.LocationUseCase
}

func NewLocationServiceHandler(locationUseCase usecase.LocationUseCase) *locationServiceHTTPHandler {
	return &locationServiceHTTPHandler{
		locationUseCase: locationUseCase,
	}
}

func (h *locationServiceHTTPHandler) Mount(group *echo.Group) {
	group.POST("/create", h.CreateLocation)
}

func (h *locationServiceHTTPHandler) CreateLocation(c echo.Context) error {
	return nil
}
