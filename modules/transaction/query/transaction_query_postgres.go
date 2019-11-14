package query

import (
	"encoding/json"
	"fmt"
	"loket-app/helper"
	"loket-app/modules/transaction/model"

	"github.com/sirupsen/logrus"
)

// InsertTxTicketPurcashing - function for inserting ticket purchasing transaction to transactions table.
func (tq *transactionQueryImpl) InsertTxTicketPurcashing(data *model.PurchaseTicketReq) (uint64, error) {
	logCtx := fmt.Sprintf("%T.InsertTxTicketPurcashing", *tq)

	var id uint64

	bodyBytes, err := json.Marshal(data.Ticket)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_marshal_ticket")
		return 0, err
	}

	sq := `INSERT INTO transactions
	(
		"eventId", "customerId", "ticket"
	) VALUES (
		$1, $2, $3
	) RETURNING id`

	stmt, err := tq.dbWrite.Prepare(sq)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_prepare_database")
		return 0, err
	}

	if err := stmt.QueryRow(
		data.EventID, data.CustomerID, bodyBytes,
	).Scan(&id); err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_exec_database")
		return 0, err
	}

	return id, nil
}

// GetTotalTicketPurchased - function for getting total ticket that have purchased by event id, customer id and ticket id.
func (tq *transactionQueryImpl) GetTotalTicketPurchased(eventID, customerID, ticketID uint64) (uint64, error) {
	logCtx := fmt.Sprintf("%T.GetTotalTicketPurchased", *tq)

	var total uint64

	sq := `select coalesce(sum(cast(ticket->>'quantity' as integer)), 0)
	from transactions 
	where CAST(ticket ->> 'ticketId' as integer) = $1
	and "eventId"=$2 and "customerId"=$3 and "deletedAt" isnull;`

	stmt, err := tq.dbRead.Prepare(sq)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_prepare_database")
		return 0, err
	}

	if err := stmt.QueryRow(ticketID, eventID, customerID).Scan(&total); err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_exec_database")
		return 0, err
	}

	return total, nil
}

// LoadTransactionByID - function for loading transaction data from transactions table by transaction id.
func (tq *transactionQueryImpl) LoadTransactionByID(txID uint64) (*model.Transaction, error) {
	logCtx := fmt.Sprintf("%T.LoadTransactionByID", *tq)

	ticketReqs := make([]model.TicketReq, 0)

	var tx model.Transaction

	sq := `select id, "eventId", "customerId",
			json_array_elements_text(ticket::json)::jsonb->>'qt' as "qt",
			json_array_elements_text(ticket::json)::jsonb->>'ticketId' as "ticketId",
			"createdAt", "updatedAt", "deletedAt"
		   from transactions
		   where id=$1 and "deletedAt" isnull;`

	stmt, err := tq.dbRead.Prepare(sq)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_prepare_database")
		return nil, err
	}

	rows, err := stmt.Query(txID)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_rows")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		var (
			ticketReq model.TicketReq
		)

		if err := stmt.QueryRow(txID).Scan(
			&tx.ID, &tx.EventID, &tx.CustomerID,
			&ticketReq.Quantity, &ticketReq.TicketID, &tx.CreatedAt, &tx.UpdatedAt,
			&tx.DeletedAt,
		); err != nil {
			helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_exec_database")
			return nil, err
		}

		ticketReqs = append(ticketReqs, ticketReq)
	}

	tx.Ticket = ticketReqs

	return &tx, nil
}
