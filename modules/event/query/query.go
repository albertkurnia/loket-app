package query

import (
	"database/sql"
)

type eventQueryImpl struct {
	dbWrite *sql.DB
	dbRead  *sql.DB
}

func NewEventQuery(dbWrite, dbRead *sql.DB) EventQuery {
	return &eventQueryImpl{
		dbWrite: dbWrite,
		dbRead:  dbRead,
	}
}

type EventQuery interface {
}
