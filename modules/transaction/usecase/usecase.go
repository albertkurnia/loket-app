package usecase

import (
	"errors"
	"fmt"
	"loket-app/helper"
	eventUC "loket-app/modules/event/usecase"
	"loket-app/modules/transaction/model"
	"loket-app/modules/transaction/query"

	"github.com/sirupsen/logrus"
)

type transactionUseCaseImpl struct {
	TransactionQuery query.TransactionQuery
	EventUsecase     eventUC.EventUseCase
}

func NewTransactionUseCase(txQuery query.TransactionQuery, eventUc eventUC.EventUseCase) TransactionUseCase {
	return &transactionUseCaseImpl{
		TransactionQuery: txQuery,
		EventUsecase:     eventUc,
	}
}

type TransactionUseCase interface {
	PurchaseTicket(data *model.PurchaseTicketReq) (uint64, error)
	GetTransactionDetail(txId uint64) (*model.Transaction, error)
}

func (impl *transactionUseCaseImpl) PurchaseTicket(data *model.PurchaseTicketReq) (uint64, error) {
	logCtx := fmt.Sprintf("%T.PurchaseTicket", *impl)

	if data == nil {
		err := errors.New("invalid request")
		return 0, err
	}

	// validate if event is exist
	event, err := impl.EventUsecase.GetEventInformation(data.EventID)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_get_event_information")
		return 0, err
	}

	if event.ID == 0 {
		err := errors.New("event not exist")
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_event_not_exist")
		return 0, err
	}

	for _, ticket := range data.Ticket {
		// get ticket quota
		ticket, err := impl.EventUsecase.GetTicket(ticket.TicketID)
		if err != nil {
			helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_get_ticket_quota")
			return 0, err
		}

		// get quantity ticket that already purchased
		totalTicket, err := impl.TransactionQuery.GetTotalTicketPurchased(event.ID, data.CustomerID, ticket.ID)
		if err != nil {
			helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_get_ticket_transaction")
			return 0, err
		}

		remainingTicket := int64(ticket.Quantity) - int64(totalTicket)
		if remainingTicket < 1 {
			err := fmt.Errorf("%s ticket is not available", ticket.Type)
			helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_ticket_not_available")
			return 0, err
		}
	}

	txID, err := impl.TransactionQuery.InsertTxTicketPurcashing(data)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_insert_tx_ticket_purchasing")
		return 0, err
	}

	return txID, nil
}

func (impl *transactionUseCaseImpl) GetTransactionDetail(txId uint64) (*model.Transaction, error) {
	return impl.TransactionQuery.LoadTransactionByID(txId)
}
