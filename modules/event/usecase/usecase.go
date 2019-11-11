package usecase

import (
	"loket-app/modules/event/query"
)

type eventUseCaseImpl struct {
	EventQuery query.EventQuery
}

func NewEventUseCase(eventQuery query.EventQuery) EventUseCase {
	return &eventUseCaseImpl{
		EventQuery: eventQuery,
	}
}

type EventUseCase interface {
}
