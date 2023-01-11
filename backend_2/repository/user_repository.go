package repository

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/gookit/slog"
	"quantumclient.org/backend/v2/models"
)

type UserRepositoryInterface interface {
	GetUserByName(string) (models.User, error)
	GetUserByID(uuid.UUID) (models.User, error)
	GetUsers() ([]models.User, error)
	GetUsersWithAccess() ([]models.User, error)
	GetUsersWithAdmin() ([]models.User, error)
	CreateUser(models.User) error
	UpdateUser(models.User) error
	DeleteUser(uuid.UUID) error

	UsernameAvailable(string) (bool, error)
}

// UserRepository is the repository for the user model
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository returns a new instance of the user repository
func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	return &UserRepository{db}
}

func (u UserRepository) GetUserByName(s string) (models.User, error) {
	row := u.db.QueryRow("SELECT u.uuid, username, admin, access, u.created_at, u.updated_at FROM permissions p JOIN users u ON p.uuid = u.uuid AND u.username = ?", s)
	return scanRowToUser(row)
}

func (u UserRepository) GetUserByID(u2 uuid.UUID) (models.User, error) {
	row := u.db.QueryRow("SELECT u.uuid, username, admin, access, u.created_at, u.updated_at FROM permissions p JOIN users u ON p.uuid = u.uuid AND u.uuid = ?", u2.String())
	return scanRowToUser(row)
}

func (u UserRepository) GetUsers() ([]models.User, error) {
	rows, err := u.db.Query("SELECT u.uuid, username, admin, access, u.created_at, u.updated_at FROM permissions p JOIN users u ON p.uuid = u.uuid")
	if err != nil {
		slog.Error(err)
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		user, er := scanRowsToUser(rows)
		if er != nil {
			slog.Error(er)
			return nil, er
		}
		users = append(users, user)
	}
	return users, nil
}

func (u UserRepository) GetUsersWithAccess() ([]models.User, error) {
	rows, err := u.db.Query("SELECT u.uuid, username, admin, access, u.created_at, u.updated_at FROM permissions p JOIN users u ON p.uuid = u.uuid AND p.access = 1")
	if err != nil {
		slog.Error(err)
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		user, er := scanRowsToUser(rows)
		if er != nil {
			slog.Error(er)
			return nil, er
		}
		users = append(users, user)
	}
	return users, nil
}

func (u UserRepository) GetUsersWithAdmin() ([]models.User, error) {
	rows, err := u.db.Query("SELECT u.uuid, username, admin, access, u.created_at, u.updated_at FROM permissions p JOIN users u ON p.uuid = u.uuid AND p.admin = 1")
	if err != nil {
		slog.Error(err)
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		user, er := scanRowsToUser(rows)
		if er != nil {
			slog.Error(er)
			return nil, er
		}
		users = append(users, user)
	}
	return users, nil
}

func (u UserRepository) CreateUser(user models.User) error {
	_, err := u.db.Exec("INSERT INTO users (uuid, username, password) VALUES (?, ?, ?, ?, ?, ?)", user.Uuid.String(), user.Username, user.Password)
	if err != nil {
		slog.Error(err)
		return  err
	}

	_, err = u.db.Exec("INSERT INTO permissions (uuid, admin, access) VALUES (?, ?, ?)", user.Uuid.String(), user.Admin, user.Access)
	if err != nil {
		slog.Error(err)
		return err
	}

	return nil
}

func (u UserRepository) UpdateUser(user models.User) error {
	_, err := u.db.Exec("UPDATE users SET username = ?, updated_at = current_timestamp WHERE uuid = ?", user.Username, user.Uuid.String())
	if err != nil {
		slog.Error(err)
		return err
	}

	_, err = u.db.Exec("UPDATE permissions SET admin = ?, access = ?, updated_at = current_timestamp WHERE uuid = ?", user.Admin, user.Access, user.Uuid.String())
	if err != nil {
		slog.Error(err)
		return err
	}

	return nil
}

func (u UserRepository) DeleteUser(u2 uuid.UUID) error {

	_, err := u.db.Exec("DELETE FROM permissions WHERE uuid = ?", u2.String())
	if err != nil {
		slog.Error(err)
		return err
	}

	_, err = u.db.Exec("DELETE FROM users WHERE uuid = ?", u2.String())
	if err != nil {
		slog.Error(err)
		return err
	}

	return nil
}

func (u UserRepository) UsernameAvailable(username string) (bool, error) {
	row := u.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", username)
	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		return false, err
	}
	return !exists, nil
}

func scanRowToUser(row *sql.Row) (models.User, error) {
	var user models.User
	err := row.Scan(user.Uuid, user.Username, user.Admin, user.Access, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		slog.Error(err)
		return user, err
	}
	return user, nil
}

func scanRowsToUser(row *sql.Rows) (models.User, error) {
	var user models.User
	err := row.Scan(user.Uuid, user.Username, user.Admin, user.Access, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		slog.Error(err)
		return user, err
	}
	return user, nil
}