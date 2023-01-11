package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Uuid     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Hwid 	NullString    `json:"hwid,omitempty"`
	Admin bool			`json:"admin"`
	Access bool			`json:"access"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
