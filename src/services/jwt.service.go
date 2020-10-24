package services

import (
	"fmt"
	"gorestapi/src/entities"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtCustomClaims struct {
	ID    uint   `json:"user_id"`
	Name  string `json:"username"`
	Admin bool   `json:"user_is_admin"`
	jwt.StandardClaims
}

var appConfig entities.Configuration

// GenerateToken ... JWTService
func GenerateToken(id uint, username string, admin bool) string {
	// services.ConfigInit
	ConfigInit(&appConfig)
	claims := &jwtCustomClaims{
		id,
		username,
		admin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "zealot13.com",
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(appConfig.Jwt.Secret))

	if err != nil {
		panic(err)
	}

	return t
}

// ValidateToken .. JWTService
func ValidateToken(tokenString string) (*jwt.Token, error) {
	// services.ConfigInit
	ConfigInit(&appConfig)

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(appConfig.Jwt.Secret), nil
	})
}
