package authentication

import (
	"authenncrm/config"
	"authenncrm/internal/entities"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenCustomClaims struct {
	UserId               string `json:"userId,omitempty"`
	Username             string `json:"username,omitempty"`
	FirstName            string `json:"firstName,omitempty"`
	LastName             string `json:"lastName,omitempty"`
	RoleId               string `json:"roleId,omitempty"`
	RoleName             string `json:"roleName,omitempty"`
	jwt.RegisteredClaims `json:"registeredClaims,omitempty"`
}

func GenerateToken(user *entities.User, config *config.Config) (string, error) {
	claims := TokenCustomClaims{
		UserId:    user.Id.String(),
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		RoleId:    user.RoleId.String(),
		RoleName:  user.Role.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "authenncrm",
			Subject:   user.Id.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.Secret))
}

func GenerateRefreshToken(user *entities.User, config *config.Config) (string, error) {
	claims := TokenCustomClaims{
		UserId: user.Id.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.Id.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.Secret))
}
