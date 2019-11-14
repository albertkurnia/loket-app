package query

import (
	"database/sql"
	"errors"
	"fmt"

	"loket-app/helper"
	"loket-app/modules/event/model"

	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// InsertTicket - function for inserting ticket to tickets table.
func (eq *eventQueryImpl) InsertTicket(data *model.CreateTicketReq) (*model.Ticket, error) {
	logCtx := fmt.Sprintf("%T.InsertTicket", *eq)

	var resp model.Ticket

	sq := `INSERT INTO "tickets"
	(
		"type", "quantity", "price"
	) VALUES (
		$1, $2, $3
	) RETURNING	id, "createdAt"`

	stmt, err := eq.dbWrite.Prepare(sq)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_prepare_database")
		return nil, err
	}

	if err := stmt.QueryRow(
		data.Type, data.Quantity, data.Price,
	).Scan(&resp.ID, &resp.CreatedAt); err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_exec_database")
		return nil, err
	}

	resp.Type = data.Type
	resp.Quantity = data.Quantity
	resp.Price = data.Price

	return &resp, nil
}

// LoadTicketByType - function for loading data from tickets table by ticket type.
func (eq *eventQueryImpl) LoadTicketByType(ticketType string) (*model.Ticket, error) {
	logCtx := fmt.Sprintf("%T.LoadTicketByType", *eq)

	var resp model.Ticket

	sq := `SELECT id, "type", "quantity", 
				"price", "createdAt", "updatedAt",
				"deletedAt"
		   FROM "tickets"
		   WHERE "type"=$1 AND "deletedAt" ISNULL`

	stmt, err := eq.dbRead.Prepare(sq)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_prepare_database")
		return nil, err
	}

	if err := stmt.QueryRow(ticketType).Scan(
		&resp.ID, &resp.Type, &resp.Quantity,
		&resp.Price, &resp.CreatedAt, &resp.UpdatedAt, &resp.DeletedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			err := errors.New("ticket not found")
			helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_ticket_not_found")
			return nil, err
		}
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_exec_database")
		return nil, err
	}

	return &resp, nil
}

// InsertEvent - function for inserting event to events table.
func (eq *eventQueryImpl) InsertEvent(data *model.CreateEventReq) (*model.Event, error) {
	logCtx := fmt.Sprintf("%T.InsertEvent", *eq)

	var resp model.Event

	sq := `INSERT INTO "events"
	(
		"title", "locationId", "description",
		"startDate", "endDate",
		"ticketId"
	) VALUES (
		$1, $2, $3,
		$4, $5, $6
	) RETURNING	id, "createdAt"`

	stmt, err := eq.dbWrite.Prepare(sq)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_prepare_database")
		return nil, err
	}

	if err := stmt.QueryRow(
		data.Title, data.LocationID, data.Description,
		data.StartDate, data.EndDate, pq.Array(data.TicketID),
	).Scan(&resp.ID, &resp.CreatedAt); err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_exec_database")
		return nil, err
	}

	resp.Title = data.Title
	resp.LocationID = data.LocationID
	resp.Description = data.Description
	resp.StartDate = data.StartDate
	resp.EndDate = data.EndDate
	resp.TicketID = data.TicketID

	return &resp, nil
}

// LoadEventByID - function for loading data from event tables by event id.
func (eq *eventQueryImpl) LoadEventByID(id uint64) (*model.Event, error) {
	logCtx := fmt.Sprintf("%T.LoadEventByID", *eq)

	var (
		resp      model.Event
		ticketIds pq.Int64Array
	)

	sq := `SELECT id, "title", "locationId", 
				"description", "startDate", 
				"endDate", "ticketId", 
				"createdAt", "updatedAt", "deletedAt"
		   FROM "events"
		   WHERE "id"=$1 AND "deletedAt" ISNULL`

	stmt, err := eq.dbRead.Prepare(sq)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_prepare_database")
		return nil, err
	}

	if err := stmt.QueryRow(id).Scan(
		&resp.ID, &resp.Title, &resp.LocationID,
		&resp.Description, &resp.StartDate, &resp.EndDate, &ticketIds,
		&resp.CreatedAt, &resp.UpdatedAt, &resp.DeletedAt,
	); err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_exec_database")
		return nil, err
	}

	temp := []int64(ticketIds)
	ticketIDs := make([]uint64, 0)
	for _, i := range temp {
		ticketIDs = append(ticketIDs, uint64(i))
	}
	resp.TicketID = ticketIDs

	return &resp, nil
}

// LoadTicketByIDs - function for loading tickets data from tickets table by ticket ids.
func (eq *eventQueryImpl) LoadTicketByIDs(ids []uint64) ([]*model.Ticket, error) {
	logCtx := fmt.Sprintf("%T.LoadTicketByID", *eq)

	var resp model.Ticket
	tickets := make([]*model.Ticket, 0)

	sq := `SELECT id, "type", "quantity", 
				"price", "createdAt", "updatedAt",
				"deletedAt"
		   FROM "tickets"
		   WHERE id=any($1) AND "deletedAt" ISNULL`

	stmt, err := eq.dbRead.Prepare(sq)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_prepare_database")
		return nil, err
	}

	rows, err := stmt.Query(pq.Array(ids))
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_exec_database")
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&resp.ID,
			&resp.Type,
			&resp.Quantity,
			&resp.Price,
			&resp.CreatedAt,
			&resp.UpdatedAt,
			&resp.DeletedAt,
		); err != nil {
			helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_rows_scan")
			return nil, err
		}

		tickets = append(tickets, &resp)
	}

	return tickets, nil
}

// GetTicket - function for getting ticket from tickets table by ticket id.
func (eq *eventQueryImpl) GetTicket(ticketID uint64) (*model.Ticket, error) {
	logCtx := fmt.Sprintf("%T.GetTicket", *eq)

	var ticket model.Ticket

	sq := `SELECT id, "type", "quantity", 
			"price", "createdAt", "updatedAt",
			"deletedAt"
		   FROM "tickets"
		   WHERE id=$1 AND "deletedAt" ISNULL`

	stmt, err := eq.dbRead.Prepare(sq)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_prepare_database")
		return nil, err
	}

	if err := stmt.QueryRow(ticketID).Scan(
		&ticket.ID, &ticket.Type, &ticket.Quantity,
		&ticket.Price, &ticket.CreatedAt, &ticket.UpdatedAt,
		&ticket.DeletedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			err := errors.New("ticket not found")
			helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_ticket_not_found")
			return nil, err
		}
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_exec_database")
		return nil, err
	}

	return &ticket, nil
}
