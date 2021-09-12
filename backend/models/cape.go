package models

import (
	"github.com/google/uuid"
)

type Cape struct {
	Uuid      uuid.UUID `json:"uuid"`
	CapeType  int 		`json:"type"`
}
