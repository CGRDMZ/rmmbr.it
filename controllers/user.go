package controllers

import (
	"log"
	"net/http"
	"github.com/CGRDMZ/rmmbrit-api/errors"
	"github.com/CGRDMZ/rmmbrit-api/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserServ *services.UserService
	AuthServ *services.AuthService
}

type CreateUserRequest struct {
	Username string `json:"username" xml:"username" form:"username" binding:"required"`
	Email    string `json:"email" xml:"email" form:"email" binding:"required"`
	Password string `json:"password" xml:"password" form:"password" binding:"required"`
}

func (uc *UserController) CreateUser(c *gin.Context) {
	log.Println("CreateUser")
	var err error
	var reqDto CreateUserRequest
	err = c.ShouldBind(&reqDto)
	if err != nil {
		c.Error(
			errors.NewCustomError(
				"invalid request body",
				"the request body does not contain all the necessary data to start the sign up process.",
				http.StatusBadRequest, err.Error()),
		)
		return
	}

	user, err := uc.UserServ.CreateUser(
		c.Request.Context(),
		services.CreateUserParams{Username: reqDto.Username, Email: reqDto.Email, Password: reqDto.Password},
	)
	if err != nil {
		c.Error(err)
		return
	}

	user.Password = ""
	c.JSON(http.StatusCreated, user)
}

type LoginRequest struct {
	Username string `json:"username" xml:"username" form:"username" binding:"required"`
	Password string `json:"password" xml:"password" form:"password" binding:"required"`
}

func (uc *UserController) Login(c *gin.Context) {
	log.Println("Login")
	var err error
	var reqDto LoginRequest
	err = c.ShouldBind(&reqDto)
	if err != nil {
		c.Error(
			errors.NewCustomError(
				"invalid request body",
				"the request body does not contain all the necessary data to start the login process.",
				http.StatusBadRequest, err.Error()),
		)
		return
	}

	token, err := uc.AuthServ.Login(
		c.Request.Context(),
		reqDto.Username,
		reqDto.Password,
	)

	if err != nil {
		if err == services.ErrLoginFailed {
			c.Error(services.ErrLoginFailed)
			return
		}
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
