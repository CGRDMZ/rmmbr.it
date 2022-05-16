package commands

import (
	"github.com/CGRDMZ/rmmbrit-api/server"
)

func Run() error {

	s := server.NewServer()

	return s.Start()
}
