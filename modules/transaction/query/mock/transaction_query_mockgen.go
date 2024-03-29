// Code generated by MockGen. DO NOT EDIT.
// Source: modules/transaction/query/query.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	model "loket-app/modules/transaction/model"
	reflect "reflect"
)

// MockTransactionQuery is a mock of TransactionQuery interface
type MockTransactionQuery struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionQueryMockRecorder
}

// MockTransactionQueryMockRecorder is the mock recorder for MockTransactionQuery
type MockTransactionQueryMockRecorder struct {
	mock *MockTransactionQuery
}

// NewMockTransactionQuery creates a new mock instance
func NewMockTransactionQuery(ctrl *gomock.Controller) *MockTransactionQuery {
	mock := &MockTransactionQuery{ctrl: ctrl}
	mock.recorder = &MockTransactionQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTransactionQuery) EXPECT() *MockTransactionQueryMockRecorder {
	return m.recorder
}

// InsertTxTicketPurcashing mocks base method
func (m *MockTransactionQuery) InsertTxTicketPurcashing(data *model.PurchaseTicketReq) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertTxTicketPurcashing", data)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertTxTicketPurcashing indicates an expected call of InsertTxTicketPurcashing
func (mr *MockTransactionQueryMockRecorder) InsertTxTicketPurcashing(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertTxTicketPurcashing", reflect.TypeOf((*MockTransactionQuery)(nil).InsertTxTicketPurcashing), data)
}

// GetTotalTicketPurchased mocks base method
func (m *MockTransactionQuery) GetTotalTicketPurchased(eventID, customerID, ticketID uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalTicketPurchased", eventID, customerID, ticketID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalTicketPurchased indicates an expected call of GetTotalTicketPurchased
func (mr *MockTransactionQueryMockRecorder) GetTotalTicketPurchased(eventID, customerID, ticketID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalTicketPurchased", reflect.TypeOf((*MockTransactionQuery)(nil).GetTotalTicketPurchased), eventID, customerID, ticketID)
}

// LoadTransactionByID mocks base method
func (m *MockTransactionQuery) LoadTransactionByID(txID uint64) (*model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadTransactionByID", txID)
	ret0, _ := ret[0].(*model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadTransactionByID indicates an expected call of LoadTransactionByID
func (mr *MockTransactionQueryMockRecorder) LoadTransactionByID(txID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadTransactionByID", reflect.TypeOf((*MockTransactionQuery)(nil).LoadTransactionByID), txID)
}
