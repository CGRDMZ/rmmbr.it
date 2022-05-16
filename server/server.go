package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/CGRDMZ/rmmbrit-api/config"
	"github.com/CGRDMZ/rmmbrit-api/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	driver *http.Server
}

func NewServer() *Server {
	r := gin.Default()

	if config.Conf.Env == "Prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	r.LoadHTMLGlob("web/templates/*")

	r.Static("/assets", "web/assets")

	routes.RegisterRoutes(r)

	

	return &Server{
		router: r,
		driver: &http.Server{
			Addr: fmt.Sprintf("localhost:%s",config.Conf.Port),
		},
	}
}

func (s *Server) Start() error {

	log.Printf("Server is running on port %s", s.driver.Addr)
	s.driver.Handler = s.router

	return s.driver.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.driver.Shutdown(ctx)
}

