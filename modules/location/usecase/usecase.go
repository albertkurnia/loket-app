package usecase

import (
	"loket-app/modules/location/model"
	"loket-app/modules/location/query"
)

type locationUseCaseImpl struct {
	LocationQuery query.LocationQuery
}

func NewLocationUseCase(locationQuery query.LocationQuery) LocationUseCase {
	return &locationUseCaseImpl{
		LocationQuery: locationQuery,
	}
}

type LocationUseCase interface {
	CreateLocation(data *model.CreateLocationReq) (*model.Location, error)
	LoadLocationByID(id uint64) (*model.Location, error)
}
