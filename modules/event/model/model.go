package model

import (
	locModel "loket-app/modules/location/model"
	"time"
)

type (
	Ticket struct {
		ID       uint64 `json:"id"`
		Type     string `json:"type"`
		Quantity uint64 `json:"quantity"`
		Price    uint64 `json:"price"`
		BaseTime
	}

	CreateTicketReq struct {
		Type     string `json:"type"`
		Quantity uint64 `json:"quantity"`
		Price    uint64 `json:"price"`
	}

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

	CreateEventReq struct {
		Title       string    `json:"title"`
		LocationID  uint64    `json:"locationId"`
		Description string    `json:"description"`
		StartDate   time.Time `json:"startDate"`
		EndDate     time.Time `json:"endDate"`
		TicketID    []uint64  `json:"ticketId"`
	}

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

	BaseTime struct {
		CreatedAt time.Time  `json:"createdAt"`
		UpdatedAt *time.Time `json:"updatedAt"`
		DeletedAt *time.Time `json:"deletedAt"`
	}
)
