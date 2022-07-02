package auth

import (
	"fmt"
	"log"
	"time"
	"github.com/CGRDMZ/rmmbrit-api/config"
	"github.com/golang-jwt/jwt/v4"
)

type JwtPayload struct {
	UserId uint `json:"uid,omitempty"`
	jwt.RegisteredClaims
}

type JwtService struct{}

func (*JwtService) GenerateToken(id uint) (string, error) {

	now := time.Now()

	expAt := now.Add(time.Second * time.Duration(config.Conf.JwtExpiresIn))

	payload := JwtPayload{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expAt),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	log.Println(config.Conf.JwtSecret)
	signedString, err := token.SignedString([]byte(config.Conf.JwtSecret))
	if err != nil {
		return "", fmt.Errorf("something happened while signing the token: %w", err)
	}

	return signedString, nil
}
