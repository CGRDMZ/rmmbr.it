package main

import (
	"fmt"
	"os"
	"github.com/CGRDMZ/rmmbrit-api/commands"
	"github.com/CGRDMZ/rmmbrit-api/config"
)

func main() {

	config.LoadConfig()

	err := commands.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while starting the application. Err: %s", err)
		os.Exit(-1)
	}
}