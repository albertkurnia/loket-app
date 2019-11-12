package usecase

import (
	"errors"
	"testing"

	"loket-app/modules/location/model"
	queryMck "loket-app/modules/location/query/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	errMck = errors.New("error mock")
)

func TestCreateLocation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	locQuery := queryMck.NewMockLocationQuery(ctrl)

	impl := locationUseCaseImpl{
		LocationQuery: locQuery,
	}

	t.Run("should return error: invalid data", func(t *testing.T) {
		var data *model.CreateLocationReq
		data = nil
		loc, err := impl.CreateLocation(data)
		assert.Error(t, err)
		assert.Equal(t, "invalid data", err.Error())
		assert.Nil(t, loc)
	})

	var data model.CreateLocationReq

	t.Run("error_insert_location", func(t *testing.T) {

		locQuery.EXPECT().InsertLocation(gomock.Any()).Return(nil, errMck)

		loc, err := impl.CreateLocation(&data)
		assert.Error(t, err)
		assert.Nil(t, loc)
		assert.Equal(t, "error mock", err.Error())
	})

	t.Run("should success to return location and error is nil", func(t *testing.T) {
		var location model.Location
		locQuery.EXPECT().InsertLocation(gomock.Any()).Return(&location, nil)

		loc, err := impl.CreateLocation(&data)
		assert.NoError(t, err)
		assert.NotNil(t, loc)
	})
}

func TestLoadLocationByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	locQuery := queryMck.NewMockLocationQuery(ctrl)

	impl := NewLocationUseCase(locQuery)

	var location model.Location
	locQuery.EXPECT().LoadLocationByID(gomock.Any()).Return(&location, nil)

	var id uint64 = 1
	resp, err := impl.LoadLocationByID(id)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
