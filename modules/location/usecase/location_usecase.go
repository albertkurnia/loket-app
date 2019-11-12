package usecase

import (
	"errors"
	"fmt"

	"loket-app/helper"
	"loket-app/modules/location/model"

	"github.com/sirupsen/logrus"
)

func (impl *locationUseCaseImpl) CreateLocation(data *model.CreateLocationReq) (*model.Location, error) {
	logCtx := fmt.Sprintf("%T.CreateLocation", *impl)

	if data == nil {
		err := errors.New("invalid data")
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_invalid_data")
		return nil, err
	}

	location, err := impl.LocationQuery.InsertLocation(data)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_insert_location")
		return nil, err
	}

	return location, err
}

func (impl *locationUseCaseImpl) LoadLocationByID(id uint64) (*model.Location, error) {
	return impl.LocationQuery.LoadLocationByID(id)
}
