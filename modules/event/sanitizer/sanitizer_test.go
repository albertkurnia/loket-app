package sanitizer

import (
	"loket-app/modules/event/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCreateTicket(t *testing.T) {
	var payload model.CreateTicketReq

	t.Run("should return error: invalid type", func(t *testing.T) {
		err := ValidateCreateTicket(&payload)
		assert.Error(t, err)
		assert.Equal(t, "invalid type", err.Error())
	})

	t.Run("should return error nil", func(t *testing.T) {
		payload.Type = "test"
		err := ValidateCreateTicket(&payload)
		assert.NoError(t, err)
	})

	t.Run("should return error: invalid payload", func(t *testing.T) {
		var payload *model.CreateTicketReq
		err := ValidateCreateTicket(payload)
		assert.Error(t, err)
		assert.Equal(t, "invalid payload", err.Error())
	})
}
