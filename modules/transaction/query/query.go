package query

import (
	"database/sql"
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
}
