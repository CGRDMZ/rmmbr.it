package routes

import (
	"context"
	"fmt"
	"os"

	"github.com/CGRDMZ/rmmbrit-api/config"
	"github.com/CGRDMZ/rmmbrit-api/controllers"
	"github.com/CGRDMZ/rmmbrit-api/middlewares"
	"github.com/CGRDMZ/rmmbrit-api/services"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func RegisterRoutes(router *gin.Engine) {

	p, err := pgxpool.Connect(context.Background(), config.Conf.DbConnectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	sc := &controllers.ShortenerController{
		Ss: &services.ShortenerService{
			Db: p,
		},
	}

	router.Use(middlewares.ErrorHandler)

	// Url shortener endpoints
	router.GET("/:id", sc.RedirectToOriginalUrl)
	router.GET("/info/:id", sc.GetUrlMapInfo)
	router.POST("/add", sc.AddNewUrlMap)
	// -----------------------

}
