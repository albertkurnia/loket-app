package usecase

import (
	"errors"
	eventModel "loket-app/modules/event/model"
	eventUCMock "loket-app/modules/event/usecase/mock"
	"loket-app/modules/transaction/model"
	queryMock "loket-app/modules/transaction/query/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	errMock = errors.New("error mock")
)

func TestUsecasePurchaseTicket(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	eventUsecase := eventUCMock.NewMockEventUseCase(ctrl)
	txQuery := queryMock.NewMockTransactionQuery(ctrl)

	t.Run("should return error invalid request", func(t *testing.T) {
		var data *model.PurchaseTicketReq

		impl := transactionUseCaseImpl{}

		txID, err := impl.PurchaseTicket(data)
		assert.Error(t, err)
		assert.Zero(t, txID)
	})

	t.Run("should return error since error_get_event_information", func(t *testing.T) {
		var data model.PurchaseTicketReq
		data.EventID = 1

		eventUsecase.
			EXPECT().
			GetEventInformation(gomock.Any()).
			Return(nil, errMock)

		impl := transactionUseCaseImpl{
			TransactionQuery: txQuery,
			EventUsecase:     eventUsecase,
		}

		txID, err := impl.PurchaseTicket(&data)
		assert.Error(t, err)
		assert.Zero(t, txID)
	})

	t.Run("should return error event not exist", func(t *testing.T) {
		var data model.PurchaseTicketReq
		data.EventID = 1

		eventUsecase.
			EXPECT().
			GetEventInformation(gomock.Any()).
			Return(&eventModel.EventInformation{}, nil)

		impl := transactionUseCaseImpl{
			TransactionQuery: txQuery,
			EventUsecase:     eventUsecase,
		}

		txID, err := impl.PurchaseTicket(&data)
		assert.Error(t, err)
		assert.Zero(t, txID)
	})

	t.Run("should return error since error_get_ticket_quota", func(t *testing.T) {
		var data model.PurchaseTicketReq
		data.EventID = 1

		var ei eventModel.EventInformation
		ticket1 := model.TicketReq{
			TicketID: 1,
		}
		ei.ID = 1
		data.Ticket = append(data.Ticket, ticket1)

		eventUsecase.
			EXPECT().
			GetEventInformation(gomock.Any()).
			Return(&ei, nil)

		eventUsecase.
			EXPECT().
			GetTicket(gomock.Any()).
			Return(nil, errMock)

		impl := transactionUseCaseImpl{
			TransactionQuery: txQuery,
			EventUsecase:     eventUsecase,
		}

		txID, err := impl.PurchaseTicket(&data)
		assert.Error(t, err)
		assert.Zero(t, txID)
	})

	t.Run("should return error since error_get_ticket_transaction", func(t *testing.T) {
		var data model.PurchaseTicketReq
		data.EventID = 1
		data.CustomerID = 1

		var ei eventModel.EventInformation
		ticket1 := model.TicketReq{
			TicketID: 1,
		}
		ei.ID = 1
		data.Ticket = append(data.Ticket, ticket1)

		var ticket eventModel.Ticket
		ticket.Quantity = 100

		eventUsecase.
			EXPECT().
			GetEventInformation(gomock.Any()).
			Return(&ei, nil)

		eventUsecase.
			EXPECT().
			GetTicket(gomock.Any()).
			Return(&ticket, nil)

		txQuery.
			EXPECT().
			GetTotalTicketPurchased(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(uint64(0), errMock)

		impl := transactionUseCaseImpl{
			TransactionQuery: txQuery,
			EventUsecase:     eventUsecase,
		}

		txID, err := impl.PurchaseTicket(&data)
		assert.Error(t, err)
		assert.Zero(t, txID)
	})

	t.Run("should return error since error_ticket_not_available", func(t *testing.T) {
		var data model.PurchaseTicketReq
		data.EventID = 1
		data.CustomerID = 1

		var ei eventModel.EventInformation
		ticket1 := model.TicketReq{
			TicketID: 1,
			Quantity: 100,
		}
		ei.ID = 1
		data.Ticket = append(data.Ticket, ticket1)

		var ticket eventModel.Ticket
		ticket.Quantity = 100

		eventUsecase.
			EXPECT().
			GetEventInformation(gomock.Any()).
			Return(&ei, nil)

		eventUsecase.
			EXPECT().
			GetTicket(gomock.Any()).
			Return(&ticket, nil)

		txQuery.
			EXPECT().
			GetTotalTicketPurchased(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(uint64(101), nil)

		impl := transactionUseCaseImpl{
			TransactionQuery: txQuery,
			EventUsecase:     eventUsecase,
		}

		txID, err := impl.PurchaseTicket(&data)
		assert.Error(t, err)
		assert.Zero(t, txID)
	})

	t.Run("should return error since error_insert_tx_ticket_purchasing", func(t *testing.T) {
		var data model.PurchaseTicketReq
		data.EventID = 1
		data.CustomerID = 1

		var ei eventModel.EventInformation
		ticket1 := model.TicketReq{
			TicketID: 1,
			Quantity: 100,
		}
		ei.ID = 1
		data.Ticket = append(data.Ticket, ticket1)

		var ticket eventModel.Ticket
		ticket.Quantity = 100

		eventUsecase.
			EXPECT().
			GetEventInformation(gomock.Any()).
			Return(&ei, nil)

		eventUsecase.
			EXPECT().
			GetTicket(gomock.Any()).
			Return(&ticket, nil)

		txQuery.
			EXPECT().
			GetTotalTicketPurchased(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(uint64(90), nil)

		txQuery.
			EXPECT().
			InsertTxTicketPurcashing(gomock.Any()).
			Return(uint64(0), errMock)

		impl := transactionUseCaseImpl{
			TransactionQuery: txQuery,
			EventUsecase:     eventUsecase,
		}

		txID, err := impl.PurchaseTicket(&data)
		assert.Error(t, err)
		assert.Zero(t, txID)
	})

	t.Run("should return transaction id", func(t *testing.T) {
		var data model.PurchaseTicketReq
		data.EventID = 1
		data.CustomerID = 1

		var ei eventModel.EventInformation
		ticket1 := model.TicketReq{
			TicketID: 1,
			Quantity: 100,
		}
		ei.ID = 1
		data.Ticket = append(data.Ticket, ticket1)

		var ticket eventModel.Ticket
		ticket.Quantity = 100

		eventUsecase.
			EXPECT().
			GetEventInformation(gomock.Any()).
			Return(&ei, nil)

		eventUsecase.
			EXPECT().
			GetTicket(gomock.Any()).
			Return(&ticket, nil)

		txQuery.
			EXPECT().
			GetTotalTicketPurchased(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(uint64(90), nil)

		txQuery.
			EXPECT().
			InsertTxTicketPurcashing(gomock.Any()).
			Return(uint64(1), nil)

		impl := NewTransactionUseCase(
			txQuery, eventUsecase,
		)

		txID, err := impl.PurchaseTicket(&data)
		assert.NoError(t, err)
		assert.Equal(t, uint64(1), txID)
	})
}

func TestUsecaseGetTransactionDetail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var txID uint64 = 1

	txQuery := queryMock.NewMockTransactionQuery(ctrl)

	txQuery.
		EXPECT().
		LoadTransactionByID(txID).
		Return(&model.Transaction{}, nil)

	impl := transactionUseCaseImpl{
		TransactionQuery: txQuery,
	}

	resp, err := impl.GetTransactionDetail(txID)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
