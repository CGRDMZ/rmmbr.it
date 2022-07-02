package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"github.com/CGRDMZ/rmmbrit-api/config"
	"github.com/CGRDMZ/rmmbrit-api/sherrors"
	"github.com/gin-gonic/gin"
)

const (
	GoogleAuthEndpoint  string = "https://accounts.google.com/o/oauth2/v2/auth"
	GoogleTokenEndpoint string = "https://oauth2.googleapis.com/token"
)

type OAuthController struct{}

func (oc *OAuthController) RedirectToExternalLogin(c *gin.Context) {
	redirectUrl := fmt.Sprintf(`%s?client_id=%s&redirect_uri=%s&response_type=code&scope=https://www.googleapis.com/auth/userinfo.email openid email`, GoogleAuthEndpoint, config.Conf.OAuth["google"].ClientId, "http://localhost:3000/callback")

	c.Redirect(http.StatusFound, redirectUrl)
}

func (oc *OAuthController) HandleCallback(c *gin.Context) {
	if c.Query("error") != "" {
		c.Error(sherrors.ErrLoginFailed)
		return
	}

	// exchange code with token
	code := c.Query("code")
	if code == "" {
		c.Error(errors.New("google did not provide code parameter"))
		return
	}

	reqBody := url.Values{}
	reqBody.Set("code", code)
	reqBody.Set("client_id", config.Conf.OAuth["google"].ClientId)
	reqBody.Set("client_secret", config.Conf.OAuth["google"].ClientSecret)
	reqBody.Set("grant_type", "authorization_code")
	reqBody.Set("redirect_uri", "http://localhost:3000/callback")

	r, err := http.PostForm(GoogleTokenEndpoint, reqBody)
	if err != nil {
		c.Error(err)
		return
	}
	if r.StatusCode != 200 {
		c.Error(fmt.Errorf("error while acquiring the token from google token endpoint"))
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
		c.Error(err)
		return
	}
	defer r.Body.Close()

	fmt.Println(res)

	c.JSON(200, res)
}
