package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/CGRDMZ/rmmbrit-api/commands"
	"github.com/CGRDMZ/rmmbrit-api/config"
)

func main() {

	env := flag.String("env", "Dev", "This is the environment value.")

	flag.Parse()

	var confFile string

	switch strings.ToUpper(*env) {
	case "DEV":
		confFile = "config.dev"
	case "PROD":
		confFile = "config.prod"
	}

	config.LoadConfig(confFile)

	err := commands.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while starting the application. Err: %s", err)
		os.Exit(-1)
	}
}