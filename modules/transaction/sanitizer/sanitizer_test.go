package sanitizer

import (
	"loket-app/modules/transaction/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePurchaseTicketRequest(t *testing.T) {
	var payload model.PurchaseTicketReq
	t.Run("should return error: invalid event id", func(t *testing.T) {
		err := ValidatePurchaseTicketRequest(&payload)
		assert.Error(t, err)
	})

	t.Run("should return error: invalid customer id", func(t *testing.T) {
		payload.EventID = 1
		err := ValidatePurchaseTicketRequest(&payload)
		assert.Error(t, err)
	})

	t.Run("should return error: invalid ticket request", func(t *testing.T) {
		payload.CustomerID = 1
		err := ValidatePurchaseTicketRequest(&payload)
		assert.Error(t, err)
	})

	t.Run("should return error: invalid request", func(t *testing.T) {
		var payload2 *model.PurchaseTicketReq
		err := ValidatePurchaseTicketRequest(payload2)
		assert.Error(t, err)
	})

	t.Run("Should return nil", func(t *testing.T) {
		payload.Ticket = []model.TicketReq{
			model.TicketReq{
				TicketID: 1,
			},
		}
		err := ValidatePurchaseTicketRequest(&payload)
		assert.NoError(t, err)
	})
}
