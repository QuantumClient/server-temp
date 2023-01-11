package services

import (
	"quantumclient.org/backend/v2/models"
	"quantumclient.org/backend/v2/repository"
)


type OnlineServiceInterface interface {
	Add(online *models.Online) ([]*models.Online, error)
	Delete(online *models.Online) error
	GetOnlineAccounts() []*models.Online
}

type OnlineService struct {
	repo repository.OnlineRepository
}

func NewOnlineService() *OnlineService {
	return &OnlineService{
		repo: repository.NewOnlineRepo(),
	}

}

func (s OnlineService) Init() {

}


func (s OnlineService) GetOnlineAccounts() []*models.Online {
	v, _ := s.repo.GetAll()
	return v
}

func (s OnlineService) Add(online *models.Online) ([]*models.Online, error) {
	err := s.repo.Add(online)
	if err != nil {
		return nil, err
	}
	return s.repo.GetAll()
}

func (s OnlineService) Delete(online *models.Online) error {
	return s.repo.Delete(online.Uuid)
}



