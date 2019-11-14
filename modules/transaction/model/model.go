package model

import "time"

type (
	// Transaction - data structure for transaction.
	Transaction struct {
		ID         uint64      `json:"id"`
		EventID    uint64      `json:"eventId"`
		CustomerID uint64      `json:"customerId"`
		Ticket     []TicketReq `json:"ticket"`
		BaseTime
	}

	// PurchaseTicketReq - data structure for purchase ticket request.
	PurchaseTicketReq struct {
		EventID    uint64      `json:"eventId"`
		CustomerID uint64      `json:"customerId"`
		Ticket     []TicketReq `json:"ticket"`
	}

	// TicketReq - data structure for ticket request.
	TicketReq struct {
		TicketID uint64 `json:"ticketId"`
		Quantity uint64 `json:"qt"`
	}

	// BaseTime - default data structure for base time.
	BaseTime struct {
		CreatedAt time.Time  `json:"createdAt"`
		UpdatedAt *time.Time `json:"updatedAt"`
		DeletedAt *time.Time `json:"deletedAt"`
	}
)
