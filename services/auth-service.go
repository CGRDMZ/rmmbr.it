package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/CGRDMZ/rmmbrit-api/auth"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthService struct {
	Db *pgxpool.Pool
	UserServ *UserService
	JwtService *auth.JwtService
}

var (
	ErrLoginFailed = errors.New("login failed")
)

func (as *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	var err error
	user, err := as.UserServ.FindByUsername(ctx, username)
	if err != nil {
		return "", fmt.Errorf("something wrong happened while finding user: %w", err)
	}
	
	if !auth.CompareHash(password, user.Password) {
		return "", ErrLoginFailed
	}

	token, err := as.JwtService.GenerateToken(user.Id)
	if err != nil {
		return "", fmt.Errorf("something wrong happened while generating token: %w", err)
	}

	return token, nil
}

