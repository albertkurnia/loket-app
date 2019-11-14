package query

import (
	"database/sql"
	"loket-app/modules/transaction/model"
)

// transactionQueryImpl - query implementation structure for transaction service.
type transactionQueryImpl struct {
	dbWrite *sql.DB
	dbRead  *sql.DB
}

// NewTransactionQuery - function for initiating new transaction query.
func NewTransactionQuery(dbWrite, dbRead *sql.DB) TransactionQuery {
	return &transactionQueryImpl{
		dbWrite: dbWrite,
		dbRead:  dbRead,
	}
}

// TransactionQuery - transaction query interface(s)
type TransactionQuery interface {
	InsertTxTicketPurcashing(data *model.PurchaseTicketReq) (uint64, error)
	GetTotalTicketPurchased(eventID, customerID, ticketID uint64) (uint64, error)
	LoadTransactionByID(txID uint64) (*model.Transaction, error)
}
