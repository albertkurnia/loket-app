package model

import (
	locModel "loket-app/modules/location/model"
	"time"
)

type (
	// Ticket - data structure for ticket.
	Ticket struct {
		ID       uint64 `json:"id"`
		Type     string `json:"type"`
		Quantity uint64 `json:"quantity"`
		Price    uint64 `json:"price"`
		BaseTime
	}

	// CreateTicketReq - data structure for create ticket request.
	CreateTicketReq struct {
		Type     string `json:"type"`
		Quantity uint64 `json:"quantity"`
		Price    uint64 `json:"price"`
	}

	// Event - data structure for event.
	Event struct {
		ID          uint64    `json:"id"`
		Title       string    `json:"title"`
		LocationID  uint64    `json:"locationId"`
		Description string    `json:"description"`
		StartDate   time.Time `json:"startDate"`
		EndDate     time.Time `json:"endDate"`
		TicketID    []uint64  `json:"ticketId"`
		BaseTime
	}

	// CreateEventReq - data structure for create event request.
	CreateEventReq struct {
		Title       string    `json:"title"`
		LocationID  uint64    `json:"locationId"`
		Description string    `json:"description"`
		StartDate   time.Time `json:"startDate"`
		EndDate     time.Time `json:"endDate"`
		TicketID    []uint64  `json:"ticketId"`
	}

	// EventInformation - data structure for getting event information.
	EventInformation struct {
		ID          uint64             `json:"id"`
		Title       string             `json:"title"`
		Description string             `json:"description"`
		StartDate   time.Time          `json:"startDate"`
		EndDate     time.Time          `json:"endDate"`
		Ticket      []*Ticket          `json:"tickets"`
		Location    *locModel.Location `json:"location"`
		BaseTime
	}

	// BaseTime - default data structure for base time.
	BaseTime struct {
		CreatedAt time.Time  `json:"createdAt"`
		UpdatedAt *time.Time `json:"updatedAt"`
		DeletedAt *time.Time `json:"deletedAt"`
	}
)
