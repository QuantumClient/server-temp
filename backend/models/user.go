package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"os"
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
	Uuid     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	Token    string    `json:"token"`
}

type ReUser struct {
	Uuid     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Hwid     string    `json:"hwid"`
}

type AuthResponse struct {
	Status   int       `json:"status"`
	Uuid     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	Token    string    `json:"token"`
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

func (u User) GenerateJWT() (string, error) {
	perms := PermsfromUser(&u)
	return GenerateJWT(u.Uuid.String(), u.Username, perms.Admin, perms.Access)
}

func (r ReUser) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(r.Password))
	return err != nil
}

func GenerateJWT(uuid, username string, admin, access bool) (string, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
		"uuid":     uuid,
		"username": username,
		"admin":    admin,
		"access":   access,
	})
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}
