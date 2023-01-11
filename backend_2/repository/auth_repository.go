package repository

import (
	"database/sql"
	"github.com/gookit/slog"
	"quantumclient.org/backend/v2/models"
)

// eh keep it lol
type AuthRepositoryInterface interface {
	GetUserPasswordByUUID(userId string) (string, error)
	GetUserWithPasswordByUUID(userId string) (models.User, error)
	GetUserWithPasswordByUsername(username string) (models.User, error)
	SetUserPassword(userId string, password string) error
	SetHWID(userId string, hwid string) error
	ResetHWID(userId string) error
	UserRepositoryInterface
}

// AuthRepository is the repository for the auth service
type AuthRepository struct {
	db *sql.DB
	UserRepositoryInterface
}

// NewAuthRepository returns a new instance of the auth repository
func NewAuthRepository(db *sql.DB, userRepo UserRepositoryInterface) AuthRepositoryInterface {
	return &AuthRepository{
		db: db,
		UserRepositoryInterface: userRepo,
	}
}

func (a *AuthRepository) GetUserPasswordByUUID(userId string) (string, error) {
	var password string
	err := a.db.QueryRow("SELECT password FROM users WHERE uuid = ?", userId).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

func (a *AuthRepository) GetUserWithPasswordByUUID(userId string) (models.User, error) {
	row := a.db.QueryRow("SELECT u.uuid, username, password, hwid, admin, access, u.created_at, u.updated_at FROM permissions p JOIN users u ON p.uuid = u.uuid AND u.uuid = ?", userId)
	return scanRowToUserPW(row)
}

func (a *AuthRepository) GetUserWithPasswordByUsername(username string) (models.User, error) {
	row := a.db.QueryRow("SELECT u.uuid, username, password, hwid, admin, access, u.created_at, u.updated_at FROM permissions p JOIN users u ON p.uuid = u.uuid AND u.username = ?", username)
	return scanRowToUserPW(row)
}

func (a *AuthRepository) SetUserPassword(userId string, password string) error {
	_, err := a.db.Exec("UPDATE users SET password = ? WHERE uuid = ?", password, userId)
	if err != nil {
		slog.Errorf("Error updating user password: %s", err)
		return err
	}
	return nil
}

func (a *AuthRepository) SetHWID(userId string, hwid string) error {
	_, err := a.db.Exec("UPDATE users SET hwid = ? WHERE uuid = ?", hwid, userId)
	if err != nil {
		slog.Errorf("Error updating user hwid: %s", err)
		return err
	}
	return nil
}

func (a *AuthRepository) ResetHWID(userId string) error {
	_, err := a.db.Exec("UPDATE users SET hwid = NULL WHERE uuid = ?", userId)
	if err != nil {
		slog.Errorf("Error resetting user hwid: %s", err)
		return err
	}
	return nil
}

func scanRowToUserPW(row *sql.Row) (models.User, error) {
	var user models.User
	err := row.Scan(user.Uuid, user.Username, user.Password, user.Hwid, user.Admin, user.Access, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		slog.Error(err)
		return user, err
	}
	return user, nil
}
