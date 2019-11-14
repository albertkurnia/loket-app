package query

import (
	"database/sql"

	"loket-app/modules/location/model"
)

// locationQueryImpl - query implementation struct for location service.
type locationQueryImpl struct {
	dbWrite *sql.DB
	dbRead  *sql.DB
}

// NewLocationQuery - function for initiating new location query.
func NewLocationQuery(dbWrite, dbRead *sql.DB) LocationQuery {
	return &locationQueryImpl{
		dbWrite: dbWrite,
		dbRead:  dbRead,
	}
}

// LocationQuery - location query interface(s).
type LocationQuery interface {
	InsertLocation(data *model.CreateLocationReq) (*model.Location, error)
	LoadLocationByID(id uint64) (*model.Location, error)
}
