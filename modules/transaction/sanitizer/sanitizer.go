package sanitizer

import (
	"errors"
	"loket-app/modules/transaction/model"
)

// ValidatePurchaseTicketRequest - function for validating purchase ticket request.
func ValidatePurchaseTicketRequest(payload *model.PurchaseTicketReq) error {

	if payload == nil {
		err := errors.New("invalid request")
		return err
	}

	if payload.EventID <= 0 {
		err := errors.New("invalid event id")
		return err
	}

	if payload.CustomerID <= 0 {
		err := errors.New("invalid customer id")
		return err
	}

	if len(payload.Ticket) == 0 {
		err := errors.New("invalid ticket request")
		return err
	}

	return nil
}
