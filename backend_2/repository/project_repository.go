package repository

import "quantumclient.org/backend/v2/models"

type ProjectRepositoryInterface interface {
	GetProjects() ([]*models.Project, error)
	GetProject(string) (*models.Project, error)
	CreateProject(*models.Project) error
	UpdateProject(*models.Project) error
	DeleteProject(*models.Project) error
}

