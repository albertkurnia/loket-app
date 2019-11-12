package sanitizer

import (
	"loket-app/modules/location/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateLocation(t *testing.T) {

	t.Run("should return error: invalid payload", func(t *testing.T) {
		var payload *model.CreateLocationReq
		err := ValidateLocation(payload)
		assert.Error(t, err)
		assert.Equal(t, "invalid payload", err.Error())
	})

	var payload model.CreateLocationReq
	t.Run("should return error: name is required", func(t *testing.T) {
		err := ValidateLocation(&payload)
		assert.Error(t, err)
		assert.Equal(t, "name is required", err.Error())
	})

	t.Run("should return error: address is required", func(t *testing.T) {
		payload.Name = "aa"
		err := ValidateLocation(&payload)
		assert.Error(t, err)
		assert.Equal(t, "address is required", err.Error())
	})

	t.Run("should return error: province name is required", func(t *testing.T) {
		payload.Address = "bb"
		err := ValidateLocation(&payload)
		assert.Error(t, err)
		assert.Equal(t, "province name is required", err.Error())
	})

	t.Run("should return nil", func(t *testing.T) {
		payload.Province = "cc"
		err := ValidateLocation(&payload)
		assert.NoError(t, err)
	})
}
