package repository

import (
	"errors"
	"github.com/google/uuid"
	"quantumclient.org/backend/v2/models"
)

type OnlineRepository interface {
	Add(online *models.Online) error
	Get(id uuid.UUID) (*models.Online, error)
	GetAll() ([]*models.Online, error)
	Delete(id uuid.UUID) error
}

type OnlineRepo struct {
	Online []*models.Online
}

// online repo constructor
func NewOnlineRepo() *OnlineRepo {
	// empty online list
	return &OnlineRepo{
		Online: []*models.Online{},
	}
}

func (r *OnlineRepo) Add(online *models.Online) error {
	index := r.findIndex(online.Uuid)
	if index == -1 {
		r.Online = append(r.Online, online)
	} else {
		r.Online[index].Expiration = online.Expiration
	}
	return nil
}

func (r *OnlineRepo) Get(id uuid.UUID) (*models.Online, error) {
	for i, online := range r.Online {
		if online.IsExpired() {
			r.delete(i)
			continue
		}
		if online.Uuid == id {
			return online, nil
		}
	}
	return nil, errors.New("online not found")
}


func (r *OnlineRepo) GetAll() ([]*models.Online, error) {
	//remove all expired online
	for i, online := range r.Online {
		if online.IsExpired() {
			r.delete(i)
		}

	}

	return r.Online, nil
}

func (r *OnlineRepo) Delete(id uuid.UUID) error {
	index := r.findIndex(id)
	if index == -1 {
		return nil
	}
	return r.delete(index)
}

func (r *OnlineRepo) delete(i int) error {
	if i != -1 {
		r.Online = append(r.Online[:i], r.Online[i+1:]...)
	}
	return nil
}

func (r *OnlineRepo) findIndex(id uuid.UUID) int {
	for i, n := range r.Online {
		if n.Uuid.ID() == id.ID() {
			return i
		}
	}

	return -1
}
