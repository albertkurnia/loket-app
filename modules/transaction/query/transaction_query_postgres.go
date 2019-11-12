package query

import (
	"fmt"
	"loket-app/helper"
	"loket-app/modules/transaction/model"

	"github.com/sirupsen/logrus"
)

func (tq *transactionQueryImpl) InsertTxTicketPurcashing(data *model.PurchaseTicketReq) (uint64, error) {
	logCtx := fmt.Sprintf("%T.InsertTxTicketPurcashing", *tq)

	var id uint64
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
		data.EventID, data.CustomerID, data.Ticket,
	).Scan(&id); err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_exec_database")
		return 0, err
	}

	return id, nil
}

func (tq *transactionQueryImpl) GetTotalTicketPurchased(eventID, customerID, ticketID uint64) (uint64, error) {
	logCtx := fmt.Sprintf("%T.GetTotalTicketPurchased", *tq)

	var total uint64

	sq := `select sum(cast(ticket->>'quantity' as integer)) as total 
			from transactions 
			where CAST(ticket ->> 'ticketId' as integer) = $1
			and "eventId"=$2 and "customerId"=$3 and "deletedAt" isnull`

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

func (tq *transactionQueryImpl) LoadTransactionByID(txID uint64) (*model.Transaction, error) {
	logCtx := fmt.Sprintf("%T.LoadTransactionByID", *tq)

	var tx model.Transaction

	sq := `select id, "eventId", "customerId",
			ticket, "createdAt", "updatedAt",
			"deletedAt"
			from transactions
			where id=$1 and "deletedAt" isnull`

	stmt, err := tq.dbRead.Prepare(sq)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_prepare_database")
		return nil, err
	}

	if err := stmt.QueryRow(txID).Scan(
		&tx.ID, &tx.EventID, &tx.CustomerID,
		&tx.Ticket, &tx.CreatedAt, &tx.UpdatedAt,
		&tx.DeletedAt,
	); err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_exec_database")
		return nil, err
	}

	return &tx, nil
}
