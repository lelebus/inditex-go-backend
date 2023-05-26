package user

import (
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/lelebus/inditext-go-backend/internal/api"
	"github.com/lelebus/inditext-go-backend/internal/database"
	"github.com/lelebus/inditext-go-backend/test/mocks"
)

const port = "8088"

var done chan os.Signal

func setup() {
	// start database
	err := database.InitDatabaseLayer()
	if err != nil {
		panic(err)
	}

	// prepare database
	mocks.ResetDB()

	// start server
	done = make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)
	go api.StartServer(port, done)
}

func shutdown() {
	done <- os.Interrupt
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
