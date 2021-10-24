package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Uuid      uuid.UUID `json:"uuid"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	Uuid         uuid.UUID `json:"uuid"`
	Username     string    `json:"username"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}

type LegUserCheck struct {
	Uuid     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Hwid     string    `json:"hwid"`
}

type AuthUserReq struct {
	Uuid         uuid.UUID `json:"uuid"`
	RefreshToken string    `json:"refresh_token"`
	CheckSum     string    `json:"sum"`
	Hwid         string    `json:"hwid"`
}

type AuthResponse struct {
	Status   int       `json:"status"`
	Uuid     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	Token    string    `json:"token"`
}

type UserKey struct {
	Uuid         uuid.UUID `json:"uuid"`
	RefreshToken string    `json:"refresh_token"`
	CheckSum     string    `json:"sum"`
}

func NewUser(username, password string) (*User, error) {
	id := uuid.New()
	pw, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Uuid:      id,
		Username:  username,
		Password:  pw,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return user, nil
}

func hashPassword(password string) (string, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(pw), nil
}

func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err != nil
}

func (r LegUserCheck) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(r.Password))
	return err != nil
}
