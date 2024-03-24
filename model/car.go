package model

import "time"

type (
	Car struct {
		ID        int64     `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	CreateCarRequest struct {
		Name string `json:"name"`
	}
)
