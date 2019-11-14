package usecase

import (
	"loket-app/modules/location/model"
	"loket-app/modules/location/query"
)

// locationUseCaseImpl - usecase implementation struct for location service.
type locationUseCaseImpl struct {
	LocationQuery query.LocationQuery
}

// NewLocationUseCase - function for initiating new location usecase.
func NewLocationUseCase(locationQuery query.LocationQuery) LocationUseCase {
	return &locationUseCaseImpl{
		LocationQuery: locationQuery,
	}
}

// LocationUseCase - location usecase interface(s).
type LocationUseCase interface {
	CreateLocation(data *model.CreateLocationReq) (*model.Location, error)
	LoadLocationByID(id uint64) (*model.Location, error)
}
