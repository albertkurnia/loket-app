package query

import (
	"database/sql"

	"loket-app/modules/location/model"
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
	InsertLocation(data *model.CreateLocationReq) (*model.Location, error)
	LoadLocationByID(id uint64) (*model.Location, error)
}
