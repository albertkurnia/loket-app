package usecase

import (
	"errors"
	"fmt"
	"loket-app/helper"
	"loket-app/modules/event/model"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// CreateTicket - usecase for creating new ticket.
func (impl *eventUseCaseImpl) CreateTicket(data *model.CreateTicketReq) (*model.Ticket, error) {
	logCtx := fmt.Sprintf("%T.CreateTicket", &impl)

	if data == nil {
		err := errors.New("invalid request")
		return nil, err
	}

	t, err := impl.EventQuery.LoadTicketByType(strings.ToUpper(strings.TrimSpace(data.Type)))
	if err != nil {
		if err.Error() != "ticket not found" {
			helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_load_ticket_by_type")
			return nil, err
		}
	}

	if t != nil {
		err := errors.New("ticket with that type already exist")
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_ticket_already_exist")
		return nil, err
	}

	data.Type = strings.ToUpper(strings.TrimSpace(data.Type))

	ticket, err := impl.EventQuery.InsertTicket(data)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_insert_ticket")
		return nil, err
	}

	return ticket, nil
}

// CreateEvent - usecase for creating new event.
func (impl *eventUseCaseImpl) CreateEvent(data *model.CreateEventReq) (*model.Event, error) {
	logCtx := fmt.Sprintf("%T.CreateEvent", &impl)

	if data == nil {
		err := errors.New("invalid request")
		return nil, err
	}

	sd, err := time.Parse(time.RFC3339, data.StartDate.Format(time.RFC3339))
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_time_parse")
		return nil, err
	}

	data.StartDate = sd

	ed, err := time.Parse(time.RFC3339, data.EndDate.Format(time.RFC3339))
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_time_parse")
		return nil, err
	}

	data.EndDate = ed

	event, err := impl.EventQuery.InsertEvent(data)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_insert_event")
		return nil, err
	}

	return event, nil
}

// GetEventInformation - usecase for getting event information.
func (impl *eventUseCaseImpl) GetEventInformation(eventID uint64) (*model.EventInformation, error) {
	logCtx := fmt.Sprintf("%T.GetEventInformation", *impl)

	var resp model.EventInformation

	event, err := impl.EventQuery.LoadEventByID(eventID)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_load_event_by_id")
		return nil, err
	}

	// construct response data
	resp.ID = event.ID
	resp.Title = event.Title
	resp.Description = event.Description
	resp.StartDate = event.StartDate
	resp.EndDate = event.EndDate
	resp.BaseTime = event.BaseTime

	tickets, err := impl.EventQuery.LoadTicketByIDs(event.TicketID)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_load_ticket_by_ids")
		return nil, err
	}

	resp.Ticket = tickets

	loc, err := impl.LocationUseCase.LoadLocationByID(event.LocationID)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_load_location_by_id")
		return nil, err
	}

	resp.Location = loc

	return &resp, nil
}

// GetTicket - usecase for getting ticket
func (impl *eventUseCaseImpl) GetTicket(ticketID uint64) (*model.Ticket, error) {
	return impl.EventQuery.GetTicket(ticketID)
}
