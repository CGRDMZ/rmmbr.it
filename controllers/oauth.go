package controllers

import (
	"errors"
	"fmt"
	"net/http"
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

	

	c.JSON(200, nil)
}
