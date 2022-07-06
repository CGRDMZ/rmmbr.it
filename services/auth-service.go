package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/CGRDMZ/rmmbrit-api/auth"
	"github.com/CGRDMZ/rmmbrit-api/config"
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

const (
	GoogleAuthEndpoint  string = "https://accounts.google.com/o/oauth2/v2/auth"
	GoogleTokenEndpoint string = "https://oauth2.googleapis.com/token"
)

// kinda deprecated and one shall not use.
func (as *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	var err error
	user, err := as.UserServ.FindByUsername(ctx, username)
	if err != nil {
		return "", fmt.Errorf("something wrong happened while finding user: %w", err)
	}
	
	if !auth.CompareHash(password, user.Password) {
		return "", ErrLoginFailed
	}

	token, err := as.JwtService.GenerateIdToken(user.Id)
	if err != nil {
		return "", fmt.Errorf("something wrong happened while generating token: %w", err)
	}

	return token, nil
}

func (as *AuthService) LoginWithExchangeCode(code string) (string, error) {
	// exchange code with token

	reqBody := url.Values{}
	reqBody.Set("code", code)
	reqBody.Set("client_id", config.Conf.OAuth["google"].ClientId)
	reqBody.Set("client_secret", config.Conf.OAuth["google"].ClientSecret)
	reqBody.Set("grant_type", "authorization_code")
	reqBody.Set("redirect_uri", "http://localhost:3000/callback")

	r, err := http.PostForm(GoogleTokenEndpoint, reqBody)
	if err != nil {
		return "", err
	}
	if r.StatusCode != 200 {
		return "", fmt.Errorf("error while acquiring the token from google token endpoint")
	}

	type ExchangeResponse struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		Scope       string `json:"scope"`
		TokenType   string `json:"token_type"`
		IdToken     string `json:"id_token"`
	}
	var res ExchangeResponse
	json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	claims, err := as.JwtService.ParseIdToken(res.IdToken)
	if err != nil {
		return "", fmt.Errorf("Could not parse Id token: %w", err)
	}

	claims.Subject

}

