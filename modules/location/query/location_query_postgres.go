package query

import (
	"fmt"
	"loket-app/helper"
	"loket-app/modules/location/model"

	log "github.com/sirupsen/logrus"
)

func (lq *locationQueryImpl) InsertLocation(data *model.CreateLocationReq) (*model.Location, error) {
	logCtx := fmt.Sprintf("%T.InsertLocation", *lq)

	var resp model.Location

	sq := `INSERT INTO "locations"
	(
		"name", 
		"address",
		"province"
	) VALUES (
		$1, $2, $3
	) RETURNING id, "createdAt"`

	stmt, err := lq.dbWrite.Prepare(sq)
	if err != nil {
		helper.Log(log.ErrorLevel, err.Error(), logCtx, "error_prepare_database")
		return nil, err
	}

	if err = stmt.QueryRow(
		data.Name, data.Address, data.Province,
	).Scan(&resp.ID, &resp.CreatedAt); err != nil {
		helper.Log(log.ErrorLevel, err.Error(), logCtx, "error_exec_database")
		return nil, err
	}

	return &resp, nil
}

func (lq *locationQueryImpl) LoadLocationByID(id uint64) (*model.Location, error) {
	logCtx := fmt.Sprintf("%T.LoadLocationByID", *lq)

	var resp model.Location
	sq := `SELECT id, "name", "address", 
					"province", "createdAt", "updatedAt",
					"deletedAt"
					FROM locations
					WHERE id=$1 AND "deletedAt" ISNULL`

	stmt, err := lq.dbWrite.Prepare(sq)
	if err != nil {
		helper.Log(log.ErrorLevel, err.Error(), logCtx, "error_prepare_database")
		return nil, err
	}

	if err = stmt.QueryRow(id).Scan(
		&resp.ID, &resp.Name, &resp.Address,
		&resp.Province, &resp.CreatedAt, &resp.UpdatedAt,
		&resp.DeletedAt); err != nil {
		helper.Log(log.ErrorLevel, err.Error(), logCtx, "error_exec_database")
		return nil, err
	}

	return &resp, nil
}
