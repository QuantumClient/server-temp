package models

import (
	"github.com/google/uuid"
)

type Cape struct {
	Uuid           uuid.UUID   `json:"uuid"`
	Username       string      `json:"username"`
	CapeType       int         `json:"type"`
	Url            *NullString `json:"url,omitempty"`
	Enabled        NullBool    `json:"enabled"`
	Owner_uuid     *uuid.UUID  `json:"owner_uuid,omitempty"`
	Owner_username *string     `json:"owner_username,omitempty"`
}
