package model

import "time"

type (
	// Location - data structure for location.
	Location struct {
		ID       uint64 `json:"id"`
		Name     string `json:"name"`
		Address  string `json:"address"`
		Province string `json:"province"`
		BaseTime
	}

	// CreateLocationReq - data structure for create location request.
	CreateLocationReq struct {
		Name     string `json:"name"`
		Address  string `json:"address"`
		Province string `json:"province"`
	}

	// BaseTime - default data structure for base time.
	BaseTime struct {
		CreatedAt time.Time  `json:"createdAt"`
		UpdatedAt *time.Time `json:"updatedAt"`
		DeletedAt *time.Time `json:"deletedAt"`
	}
)
