package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/lelebus/inditext-go-backend/internal/api"
	"github.com/lelebus/inditext-go-backend/internal/database"
	"github.com/lelebus/inditext-go-backend/test/mocks"
)

func main() {
	// Open connection with the in-memory database
	if err := database.InitDatabaseLayer(); err != nil {
		panic(err)
	}

	// Load mocks, if env variable is set
	if os.Getenv("MOCKS") == "true" {
		mocks.ResetDB()
	}

	// Launch Server
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)
	api.StartServer("8080", done)
}
