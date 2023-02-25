package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ahmadaidin/echoscratch/config"
	"github.com/ahmadaidin/echoscratch/controller/http"
	"github.com/ahmadaidin/echoscratch/controller/http/book"
	"github.com/ahmadaidin/echoscratch/domain/repository/mongorepo"
	"github.com/ahmadaidin/echoscratch/infra"
)

func beforeTerminate() {
	fmt.Println("Good bye!")
}

// SetupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by calling
// our clean up procedure and exiting the program.
func setupCloseHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		beforeTerminate()
		os.Exit(0)
	}()
}

// @title EchoScratch Service Documentation
// @version 1.0
// @description This is a server for EchoScratch API.

// @contact.name Ahmad Aidin
// @contact.email ahmadaidin08.aa@gmail.com

// @host localhost:8000
// @BasePath /
func main() {
	setupCloseHandler()

	// load configuration
	config.Read(".env")
	cfg := config.GetConfig()

	// initiate all infrastructures
	mongoDb := infra.NewMongoConnection(cfg.DatabaseURI, cfg.DatabaseConnectionTimeout)

	// initiate all repostiroies
	bookRepo := mongorepo.NewBookRepository(mongoDb)

	// initiate all services

	// initiate all controllers
	bookCtrl := book.NewBookController(
		bookRepo,
	)

	http.NewFiberHttpHandler(bookCtrl).Listen(config.GetConfig().Port)
}
