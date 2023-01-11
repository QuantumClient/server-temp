package models

import (
	"github.com/google/uuid"
	"quantumclient.org/backend/v2/util"
	"time"
)

type Online struct {
	Uuid uuid.UUID `json:"uuid"`
	Owner User	  `json:"owner,omitempty"`
	Expiration int64     `json:"expiration"`
}

func NewOnline(uuid2 uuid.UUID, owner User) *Online {
	return &Online{
		Uuid: uuid2,
		Owner: owner,
		Expiration: time.Now().Add(util.OnlineLimmit).Unix(),
	}
}

func (o *Online) IsExpired() bool {
	return o.Expiration < time.Now().Unix()
}

func (o *Online) Populate() {
	o.Expiration = time.Now().Add(util.OnlineLimmit).Unix()
}
