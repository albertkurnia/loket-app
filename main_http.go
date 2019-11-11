package main

import (
	"fmt"
	"os"
	"strconv"

	eventPresenter "loket-app/modules/event/presenter"
	locationPresenter "loket-app/modules/location/presenter"
	txPresenter "loket-app/modules/transaction/presenter"

	"github.com/labstack/echo"
)

const (
	// defaultPort - default port for HTTP
	defaultPort = 7020
)

// HTTPServeMain - function for serving HTTP
func (s *Service) HTTPServeMain() {
	e := echo.New()

	if os.Getenv("ENV") == "DEV" {
		e.Debug = true
	}

	eventHandler := eventPresenter.NewEventServiceHandler(s.EventUseCase)
	eventGroup := e.Group("/api/event")
	eventHandler.Mount(eventGroup)

	locationHandler := locationPresenter.NewLocationServiceHandler(s.LocationUseCase)
	locationGroup := e.Group("/api/location")
	locationHandler.Mount(locationGroup)

	txHandler := txPresenter.NewTransactionServiceHandler(s.TransactionUseCase)
	txGroup := e.Group("/api/transaction")
	txHandler.Mount(txGroup)

	// set REST port
	var port uint16
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		portInt, err := strconv.Atoi(portEnv)
		if err != nil {
			port = defaultPort
		} else {
			port = uint16(portInt)
		}
	} else {
		port = defaultPort
	}

	listenerPort := fmt.Sprintf(":%d", port)
	e.Logger.Fatal(e.Start(listenerPort))
}
