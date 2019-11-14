package query

import (
	"database/sql"
	"loket-app/modules/event/model"
)

// eventQueryImpl - query implementation struct for event.
type eventQueryImpl struct {
	dbWrite *sql.DB
	dbRead  *sql.DB
}

// NewEventQuery - function for initiate new event query.
func NewEventQuery(dbWrite, dbRead *sql.DB) EventQuery {
	return &eventQueryImpl{
		dbWrite: dbWrite,
		dbRead:  dbRead,
	}
}

// EventQuery - event query interface(s).
type EventQuery interface {
	InsertTicket(data *model.CreateTicketReq) (*model.Ticket, error)
	LoadTicketByIDs(ids []uint64) ([]*model.Ticket, error)
	LoadTicketByType(ticketType string) (*model.Ticket, error)
	InsertEvent(data *model.CreateEventReq) (*model.Event, error)
	LoadEventByID(id uint64) (*model.Event, error)
	GetTicket(ticketID uint64) (*model.Ticket, error)
}
