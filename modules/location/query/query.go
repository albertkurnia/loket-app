package query

import (
	"database/sql"
)

type locationQueryImpl struct {
	dbWrite *sql.DB
	dbRead  *sql.DB
}

func NewLocationQuery(dbWrite, dbRead *sql.DB) LocationQuery {
	return &locationQueryImpl{
		dbWrite: dbWrite,
		dbRead:  dbRead,
	}
}

type LocationQuery interface {
}