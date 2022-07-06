package auth

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/CGRDMZ/rmmbrit-api/config"
	"github.com/golang-jwt/jwt/v4"
)

type JwtPayload struct {
	jwt.RegisteredClaims
}


type JwtService struct{}

func (*JwtService) GenerateIdToken(id uint) (string, error) {

	now := time.Now()

	expAt := now.Add(time.Second * time.Duration(config.Conf.JwtExpiresIn))

	payload := JwtPayload{
		jwt.RegisteredClaims{
			Subject: fmt.Sprint(id),
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

func (*JwtService) ParseIdToken(token string) (*JwtPayload, error) {
	if strings.Trim(token, " ") == "" {
		return nil, errors.New("the id token cannot be empty")
	}

	t, err := jwt.NewParser().ParseWithClaims(token, &JwtPayload{}, func(t *jwt.Token) (interface{}, error) { return config.Conf.JwtSecret, nil})
	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, errors.New("The provided token is not a valid token.")
	}

	if claims, ok :=t.Claims.(*JwtPayload); ok {
		return claims, nil
	} else {
		return nil, errors.New("invalid jwt token")
	}
}
