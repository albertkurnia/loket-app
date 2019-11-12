package main

import (
	"errors"
	"loket-app/config"
	"loket-app/helper"
	"os"

	eventQ "loket-app/modules/event/query"
	eventUC "loket-app/modules/event/usecase"
	locationQ "loket-app/modules/location/query"
	locationUC "loket-app/modules/location/usecase"
	txQ "loket-app/modules/transaction/query"
	txUC "loket-app/modules/transaction/usecase"

	log "github.com/sirupsen/logrus"
)

// Service - data structure for wrapping all use case layers
type Service struct {
	EventUseCase       eventUC.EventUseCase
	LocationUseCase    locationUC.LocationUseCase
	TransactionUseCase txUC.TransactionUseCase
}

func MakeHandler() *Service {
	// initiate database connection
	sqlInfo := new(config.PSQLInfo)
	dbMaster, err := config.CreateDBConn(sqlInfo, config.DBMasterType)
	if err != nil {
		err := errors.New("failed to created DB master connection")
		helper.Log(log.PanicLevel, err.Error(), "MakeHandler", "initiate_database")
		os.Exit(1)
	}

	dbSlave, err := config.CreateDBConn(sqlInfo, config.DBSlaveType)
	if err != nil {
		err := errors.New("failed to created DB slave connection")
		helper.Log(log.PanicLevel, err.Error(), "MakeHandler", "initiate_database")
		os.Exit(1)
	}

	eventQuery := eventQ.NewEventQuery(dbMaster, dbSlave)
	locationQuery := locationQ.NewLocationQuery(dbMaster, dbSlave)
	txQuery := txQ.NewTransactionQuery(dbMaster, dbSlave)

	locationUsecase := locationUC.NewLocationUseCase(locationQuery)
	eventUsecase := eventUC.NewEventUseCase(eventQuery, locationUsecase)
	txUsecase := txUC.NewTransactionUseCase(txQuery, eventUsecase)

	return &Service{
		EventUseCase:       eventUsecase,
		LocationUseCase:    locationUsecase,
		TransactionUseCase: txUsecase,
	}
}
