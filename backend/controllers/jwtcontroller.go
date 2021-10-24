package controllers

import (
	"backend/models"
	"backend/util"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type JwtCustomClaims struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
	Access   bool   `json:"access"`
	jwt.RegisteredClaims
}

type JwtRefreshClaims struct {
	Uuid string `json:"uuid"`
	jwt.RegisteredClaims
}

func GetJWT(user *models.Permission) string {
	return GetJWTCustomTime(user, 2)
}

func GetJWTCustomTime(user *models.Permission, hourX int64) string {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaims{
		Uuid:     user.ID.String(),
		Username: user.Username,
		Admin:    user.Admin,
		Access:   user.Access,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(hourX) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		log.Println(err)
	}
	return tokenString
}

func GetToken(r *http.Request) (*jwt.Token, error) {
	tokenString := r.Header.Get("Authorization")
	if len(tokenString) == 0 {
		return nil, util.ErrToken
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	return jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

func GetRefreshToken(uuid uuid.UUID, hashedPassword string) string {

	signingKey := []byte(os.Getenv("REFRESH_SECRET") + hashedPassword)

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &JwtRefreshClaims{
		Uuid: uuid.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		log.Println(err)
	}
	return tokenString
}

func RefreshFromUUID(uuid uuid.UUID) string {
	return GetRefreshToken(uuid, getHashedPassword(uuid))
}

func AccountFromRefresh(tokenString string) (*models.Permission, error) {
	token, err := JWTFromRefresh(tokenString)
	if err != nil {
		log.Println(err)
		return nil, util.ErrToken
	}
	claims, ok := token.Claims.(*JwtRefreshClaims)
	if !ok && !token.Valid {
		return nil, util.ErrToken
	}
	uuid, _ := uuid.Parse(claims.Uuid)
	return getAccountfromUUID(uuid), nil
}

func JWTFromRefresh(refreshToken string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(refreshToken, &JwtRefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		uuid, _ := uuid.Parse(token.Claims.(*JwtRefreshClaims).Uuid)
		pass := getHashedPassword(uuid)
		return []byte(os.Getenv("REFRESH_SECRET") + pass), nil
	})
}
