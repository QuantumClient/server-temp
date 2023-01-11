package models

import "github.com/google/uuid"

type Cape struct {
	Uuid     uuid.UUID      `json:"uuid"`
	Username string      `json:"username"`
	Type     int         `json:"type"`
	Url     string      `json:"url"`
	Enabled  bool 	`json:"enabled"`
	OwnerUuid     uuid.UUID  `json:"owner_uuid,omitempty"`
}
