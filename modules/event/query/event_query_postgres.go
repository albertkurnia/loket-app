package query

import (
	"fmt"

	"loket-app/helper"
	"loket-app/modules/event/model"

	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

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

	return &resp, nil
}

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
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_exec_database")
		return nil, err
	}

	return &resp, nil
}

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
		$4, $5, %6,
		$7
	) RETURNING	id, "createdAt"`

	stmt, err := eq.dbWrite.Prepare(sq)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_prepare_database")
		return nil, err
	}

	if err := stmt.QueryRow(
		data.Title, data.LocationID, data.Description,
		data.StartDate, data.EndDate,
		data.TicketID,
	).Scan(&resp.ID, &resp.CreatedAt); err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_exec_database")
		return nil, err
	}

	return &resp, nil
}

func (eq *eventQueryImpl) LoadEventByID(id uint64) (*model.Event, error) {
	logCtx := fmt.Sprintf("%T.LoadEventByID", *eq)

	var resp model.Event

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
		&resp.Description, &resp.StartDate, &resp.EndDate, &resp.TicketID,
		&resp.DeletedAt, &resp.UpdatedAt, &resp.DeletedAt,
	); err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_exec_database")
		return nil, err
	}

	return &resp, nil
}

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

func (eq *eventQueryImpl) GetTicket(ticketID uint64) (*model.Ticket, error) {
	logCtx := fmt.Sprintf("%T.GetTicketQuota", *eq)

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
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_exec_database")
		return nil, err
	}

	return &ticket, nil
}
