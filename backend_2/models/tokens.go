package models

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type AccessTokenClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	Uuid     string `json:"uuid"` // should be same ad jwt.RegisteredClaims.Subject
	Admin    bool   `json:"admin"`
	Access   bool   `json:"access"`
}

type RefreshTokenClaims struct {
	jwt.RegisteredClaims
}

func GetJWT(user *User) *jwt.Token {
	return GetJWTCustomTime(user, 2)
}

func GetJWTCustomTime(user *User, hourX int64) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &AccessTokenClaims{
		Username: user.Username,
		Uuid:     user.Uuid.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: user.Uuid.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(hourX) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})
}

