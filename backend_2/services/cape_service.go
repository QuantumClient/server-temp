package services

import (
	"errors"
	"github.com/google/uuid"
	"quantumclient.org/backend/v2/models"
	"quantumclient.org/backend/v2/repository"
)

type CapeServiceInterface interface {
	GetCape(uuid.UUID) (*models.Cape, error)
	GetCapeByName(string) (*models.Cape, error)
	GetCapes() ([]*models.Cape, error)
	GetEnabledCapes() ([]*models.Cape, error)
	GetCapesByOwner(uuid.UUID) ([]*models.Cape, error)
	SetCapeEnabled(uuid.UUID) error
	UpdateCape(uuid.UUID, models.Cape) error
	DeleteCape(uuid.UUID) error
}

// CapeService is the service for managing cape data
type CapeService struct {
	cp repository.CapeRepositoryInterface
}

// NewCapeService creates a new cape service
func NewCapeService(cp repository.CapeRepositoryInterface) *CapeService {
	return &CapeService{
		cp: cp,
	}
}

func (c CapeService) GetCape(u uuid.UUID) (*models.Cape, error) {
	return c.cp.FindCapeByUuid(u)
}

func (c CapeService) GetCapeByName(s string) (*models.Cape, error) {
	return c.cp.FindCapeByUsername(s)
}

func (c CapeService) GetCapes() ([]*models.Cape, error) {
	return c.cp.FindAllCapes()
}

func (c CapeService) GetEnabledCapes() ([]*models.Cape, error) {
	capes, err := c.GetCapes()
	if err != nil {
		return nil, err
	}
	var enabledCapes []*models.Cape
	for _, cape := range capes {
		if cape.Enabled {
			enabledCapes = append(enabledCapes, cape)
		}
	}
	return enabledCapes, nil
}

func (c CapeService) GetCapesByOwner(u uuid.UUID) ([]*models.Cape, error) {
	return c.cp.FindCapesByOwnerUuid(u)
}

func (c CapeService) SetCapeEnabled(u uuid.UUID) error {
	cape, err := c.GetCape(u)
	if err != nil {
		return err
	}
	cape.Enabled = !cape.Enabled
	return c.cp.UpdateCape(*cape)
}

func (c CapeService) UpdateCape(u uuid.UUID, cape models.Cape) error {
	if u != cape.Uuid {
		return errors.New("uuid mismatch")
	}
	return c.cp.UpdateCape(cape)
}

func (c CapeService) DeleteCape(u uuid.UUID) error {
	return c.cp.DeleteCape(u)
}

