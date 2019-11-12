package model

import "time"

type (
	Location struct {
		ID       uint64 `json:"id"`
		Name     string `json:"name"`
		Address  string `json:"address"`
		Province string `json:"province"`
		BaseTime
	}

	CreateLocationReq struct {
		Name     string `json:"name"`
		Address  string `json:"address"`
		Province string `json:"province"`
	}

	BaseTime struct {
		CreatedAt time.Time  `json:"createdAt"`
		UpdatedAt *time.Time `json:"updatedAt"`
		DeletedAt *time.Time `json:"deletedAt"`
	}
)
