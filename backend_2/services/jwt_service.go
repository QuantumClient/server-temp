package services

import (
	"github.com/golang-jwt/jwt/v4"
	"quantumclient.org/backend/v2/models"
	"quantumclient.org/backend/v2/repository"
	"time"
)

type JwtServiceInterface interface {
	GenerateAccessToken(models.User) (string, error)
	GenerateAccessTokenWithCustomTime(models.User, time.Duration) (string, error)
	ValidateAccessToken(string) (bool, *jwt.Token, error)
	GenerateRefreshToken(models.User) (string, error)
	ValidateRefreshToken(string) (bool, *jwt.Token, error)
}

// JwtService is a struct that implements JwtServiceInterface
type JwtService struct {
	config *models.Config
	authRepo repository.AuthRepositoryInterface

}

// NewJwtService returns a new JwtService
func NewJwtService(config *models.Config, authRepo repository.AuthRepositoryInterface) *JwtService {
	return &JwtService{config: config, authRepo: authRepo}
}

// GenerateAccessToken generates an access token
func (j *JwtService) GenerateAccessToken(user models.User) (string, error) {
	return j.GenerateAccessTokenWithCustomTime(user, time.Hour*2)
}

// GenerateAccessTokenWithCustomTime generates an access token with a custom time
func (j *JwtService) GenerateAccessTokenWithCustomTime(user models.User, dur time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &models.AccessTokenClaims{
		Username: user.Username,
		Uuid:     user.Uuid.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: user.Uuid.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(dur)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})
	return token.SignedString([]byte(j.config.Auth.AccessTokenSecret))
}

// ValidateAccessToken validates an access token
func (j *JwtService) ValidateAccessToken(tokenString string) (bool, *jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.config.Auth.AccessTokenSecret), nil
	})
	if err != nil {
		return false, nil, err
	}
	return true, token, nil
}

func (j *JwtService) GenerateRefreshToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &models.RefreshTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: user.Uuid.String(),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})
	return token.SignedString([]byte(j.config.Auth.RefreshTokenSecret + user.Password)) // password is used as a salt for the refresh token, must be hashed password
}

func (j *JwtService) ValidateRefreshToken(tokenString string) (bool, *jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		pw, err := j.authRepo.GetUserPasswordByUUID(token.Claims.(*models.RefreshTokenClaims).Subject)
		if err != nil {
			return nil, err
		}
		return []byte(j.config.Auth.RefreshTokenSecret+pw), nil
	})
	if err != nil {
		return false, nil, err
	}
	return token.Valid, token, nil
}
