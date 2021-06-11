package util

import (
	"backend/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"strings"
)


func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), err
}

func GetJWT(r *http.Request) string {
	tokenString := r.Header.Get("Authorization")
	if len(tokenString) == 0 {
		return "1"
	}
	return strings.Replace(tokenString, "Bearer ", "", 1)
}

func GetAccountsFromToken(c jwt.MapClaims) (*models.User, *models.Permission) {
	uuid, err := uuid.Parse(c["uuid"].(string))
	if err != nil {
		log.Println(err)
	}
	user := &models.User{
		Uuid:     uuid,
		Username: c["username"].(string),
		}

	perms := models.PermsfromUser(user)

	return user, perms
}

func FullCheck(w http.ResponseWriter, r *http.Request) (bool, *models.Permission) {
	token := GetJWT(r)

	claims, err := ValidateJWT(token)

	if err != nil {
		log.Println(err)
		ErrorResponse(w, r, "Invaild Token")
		return false, nil
	}

	if claims.Valid == nil {
		log.Println(err)
		ErrorResponse(w, r, "Invaild Token")
		return false, nil
	}

	_, perms := GetAccountsFromToken(claims)

	if !perms.Admin {
		ErrorResponse(w, r, "Not Admin")
		return false, nil
	}
	return true, perms
}


