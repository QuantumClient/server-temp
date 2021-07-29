package models

import "github.com/google/uuid"

type Online struct {
	Uuid     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
}
