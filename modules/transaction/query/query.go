package query

import (
	"database/sql"
	"loket-app/modules/transaction/model"
)

type transactionQueryImpl struct {
	dbWrite *sql.DB
	dbRead  *sql.DB
}

func NewTransactionQuery(dbWrite, dbRead *sql.DB) TransactionQuery {
	return &transactionQueryImpl{
		dbWrite: dbWrite,
		dbRead:  dbRead,
	}
}

type TransactionQuery interface {
	InsertTxTicketPurcashing(data *model.PurchaseTicketReq) (uint64, error)
	GetTotalTicketPurchased(eventID, customerID, ticketID uint64) (uint64, error)
	LoadTransactionByID(txID uint64) (*model.Transaction, error)
}
