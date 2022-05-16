package auth

import (
	"errors"
	"time"

	"github.com/CGRDMZ/rmmbrit-api/config"
	"github.com/golang-jwt/jwt/v4"
)

type JwtPayload struct {
	UserId string `json:"uid,omitempty"`
	jwt.RegisteredClaims
}

type JwtService struct{}

func (*JwtService) GenerateToken(id string) (string, error) {

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

	signedString, err := token.SignedString(config.Conf.JwtSecret)
	if err != nil {
		return "", errors.New("something happened while signing the token")
	}

	return signedString, nil
}
