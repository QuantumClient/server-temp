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

type jwtRefreshClaims struct {
	Uuid string `json:"uuid"`
	jwt.RegisteredClaims
}

func GetJWT(user *models.Permission) string {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaims{
		Uuid:     user.ID.String(),
		Username: user.Username,
		Admin:    user.Admin,
		Access:   user.Access,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwtRefreshClaims{
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
	token, err := jwt.ParseWithClaims(tokenString, &jwtRefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		uuid, _ := uuid.Parse(token.Claims.(*jwtRefreshClaims).Uuid)
		pass := getHashedPassword(uuid)
		return []byte(os.Getenv("REFRESH_SECRET") + pass), nil
	})
	if err != nil {
		log.Println(err)
		return nil, util.ErrToken
	}
	claims, ok := token.Claims.(*jwtRefreshClaims)
	if !ok && !token.Valid {
		return nil, util.ErrToken
	}
	uuid, _ := uuid.Parse(claims.Uuid)
	return getAccountfromUUID(uuid), nil
}
