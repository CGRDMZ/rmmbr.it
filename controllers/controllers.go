package controllers

import (
	"github.com/CGRDMZ/rmmbrit-api/auth"
	"github.com/CGRDMZ/rmmbrit-api/services"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ControllerFactory struct {
	Db *pgxpool.Pool
}

func (cf *ControllerFactory) CreateShortenerController() *ShortenerController {
	return &ShortenerController{
		Ss: &services.ShortenerService{
			Db: cf.Db,
		},
	}
}

func (cf *ControllerFactory) CreateUserController() *UserController {
	return &UserController{
		UserServ: &services.UserService{
			Db: cf.Db,
		},
		AuthServ: &services.AuthService{
			Db: cf.Db,
			UserServ: &services.UserService{
				Db: cf.Db,
			},
			JwtService: &auth.JwtService{},
		},
	}
}

func (cf *ControllerFactory) CreateOAuthController() *OAuthController {
	return &OAuthController{}
}
