package sanitizer

import (
	"errors"
	"loket-app/modules/event/model"
	"strings"
)

// ValidateCreateTicket - function for validating create ticket request.
func ValidateCreateTicket(payload *model.CreateTicketReq) error {

	if payload == nil {
		err := errors.New("invalid payload")
		return err
	}

	if strings.TrimSpace(payload.Type) == "" {
		err := errors.New("invalid type")
		return err
	}

	if payload.Quantity < 0 {
		err := errors.New("invalid quantity")
		return err
	}

	if payload.Price < 0 {
		err := errors.New("invalid price")
		return err
	}

	return nil
}

// ValidateCreateEvent - function for validating create event request.
func ValidateCreateEvent(payload *model.CreateEventReq) error {
	if payload == nil {
		err := errors.New("invalid payload")
		return err
	}

	if strings.TrimSpace(payload.Title) == "" {
		err := errors.New("invalid title")
		return err
	}

	if payload.LocationID == 0 {
		err := errors.New("invalid location id")
		return err
	}

	if payload.StartDate.UTC().IsZero() {
		err := errors.New("invalid time. time format is YYYY:MM:dd HH:mm:ss")
		return err
	}

	if payload.EndDate.UTC().IsZero() {
		err := errors.New("invalid time. time format is YYYY:MM:dd HH:mm:ss")
		return err
	}

	if len(payload.TicketID) == 0 {
		err := errors.New("invalid ticket id")
		return err
	}

	return nil
}
