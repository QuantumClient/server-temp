package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/gookit/slog"
	"quantumclient.org/backend/v2/models"
	"quantumclient.org/backend/v2/repository"
)

type UserServiceInterface interface {
	GetUser(uuid.UUID) (models.User, error)
	GetUsers() ([]models.User, error)
	CreateUser(models.User) (models.User, error)
	UpdateUser(models.User) (models.User, error)
	UpdateUsername(uuid.UUID, string) error
	DeleteUser(uuid.UUID) error
}

// UserService is a struct that implements the UserServiceInterface
type UserService struct {
	UserRepository repository.UserRepositoryInterface
}

// NewUserService returns a new UserService
func NewUserService(userRepository repository.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (u UserService) GetUser(uuid uuid.UUID) (models.User, error) {
	return u.UserRepository.GetUserByID(uuid)
}

func (u UserService) GetUsers() ([]models.User, error) {
	return u.UserRepository.GetUsers()
}

func (u UserService) CreateUser(user models.User) (models.User, error) {
	if user.Uuid == uuid.Nil {
		user.Uuid = uuid.New()
	}
	return user, u.UserRepository.CreateUser(user)
}

func (u UserService) UpdateUser(user models.User) (models.User, error) {
	return user, u.UserRepository.UpdateUser(user)
}

func (u UserService) UpdateUsername(uuid uuid.UUID, s string) error {
	available, err := u.UserRepository.UsernameAvailable(s)
	if err != nil || !available {
		return errors.New("Username already taken")
	}
	user, err := u.GetUser(uuid)
	if err != nil {
		slog.Error(err)
		return err
	}
	user.Username = s
	return u.UserRepository.UpdateUser(user)
}

func (u UserService) DeleteUser(u2 uuid.UUID) error {
	panic("implement me")
}

