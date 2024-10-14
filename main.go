package main

import (
	"embed"
	"os"

	"github.com/softwarespot/public-holidays/cmd"
	"github.com/softwarespot/public-holidays/internal/env"
	"github.com/softwarespot/public-holidays/internal/logging"
)

//go:embed ".env"
var envFile embed.FS

func main() {
	logger := logging.NewStdoutLogger()
	if err := env.Load(envFile, ".env"); err != nil {
		logger.Fatal(err, 1)
	}

	// Remove the executable name
	if err := cmd.Execute(os.Args[1:], logger); err != nil {
		logger.Fatal(err, 1)
	}
}
