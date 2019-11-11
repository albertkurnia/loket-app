package usecase

import (
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
}
