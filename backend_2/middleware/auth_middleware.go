package middleware

import (
	"github.com/google/uuid"
	"github.com/gookit/rux"
	"net/http"
	"quantumclient.org/backend/v2/models"
	"quantumclient.org/backend/v2/services"
	"strings"
)

type AuthMiddleware struct {
	jwtSvr services.JwtServiceInterface
}

// NewAuthMiddleware creates a new AuthMiddleware instance
func NewAuthMiddleware(jwtSvr services.JwtServiceInterface) *AuthMiddleware {
	return &AuthMiddleware{
		jwtSvr: jwtSvr,
	}
}

func (m AuthMiddleware) ValidateJWT(c *rux.Context)  {
	token := c.Header("Authorization")
	token = strings.Replace(token, "Bearer ", "", 1)
	if token == "" || (len(strings.Split(token, ".")) != 3) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	pass, parsedToken, err := m.jwtSvr.ValidateAccessToken(token)
	if err != nil || !pass {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	tokenClaims := parsedToken.Claims.(*models.AccessTokenClaims)
	user := models.User{
		Uuid:     uuid.MustParse(tokenClaims.Uuid),
		Username: tokenClaims.Username,
		Admin: tokenClaims.Admin,
		Access: tokenClaims.Access,
	}
	c.Set("user", user)
	c.Next()
}

