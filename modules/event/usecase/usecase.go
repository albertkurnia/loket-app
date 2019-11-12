package usecase

import (
	"loket-app/modules/event/model"
	"loket-app/modules/event/query"
	locUsecase "loket-app/modules/location/usecase"
)

type eventUseCaseImpl struct {
	EventQuery      query.EventQuery
	LocationUseCase locUsecase.LocationUseCase
}

func NewEventUseCase(eventQuery query.EventQuery, locUC locUsecase.LocationUseCase) EventUseCase {
	return &eventUseCaseImpl{
		EventQuery:      eventQuery,
		LocationUseCase: locUC,
	}
}

type EventUseCase interface {
	CreateTicket(data *model.CreateTicketReq) (*model.Ticket, error)
	CreateEvent(data *model.CreateEventReq) (*model.Event, error)
	GetEventInformation(eventID uint64) (*model.EventInformation, error)
	GetTicket(ticketID uint64) (*model.Ticket, error)
}
