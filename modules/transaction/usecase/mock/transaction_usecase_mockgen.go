// Code generated by MockGen. DO NOT EDIT.
// Source: modules/transaction/usecase/usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	model "loket-app/modules/transaction/model"
	reflect "reflect"
)

// MockTransactionUseCase is a mock of TransactionUseCase interface
type MockTransactionUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionUseCaseMockRecorder
}

// MockTransactionUseCaseMockRecorder is the mock recorder for MockTransactionUseCase
type MockTransactionUseCaseMockRecorder struct {
	mock *MockTransactionUseCase
}

// NewMockTransactionUseCase creates a new mock instance
func NewMockTransactionUseCase(ctrl *gomock.Controller) *MockTransactionUseCase {
	mock := &MockTransactionUseCase{ctrl: ctrl}
	mock.recorder = &MockTransactionUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTransactionUseCase) EXPECT() *MockTransactionUseCaseMockRecorder {
	return m.recorder
}

// PurchaseTicket mocks base method
func (m *MockTransactionUseCase) PurchaseTicket(data *model.PurchaseTicketReq) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PurchaseTicket", data)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PurchaseTicket indicates an expected call of PurchaseTicket
func (mr *MockTransactionUseCaseMockRecorder) PurchaseTicket(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PurchaseTicket", reflect.TypeOf((*MockTransactionUseCase)(nil).PurchaseTicket), data)
}

// GetTransactionDetail mocks base method
func (m *MockTransactionUseCase) GetTransactionDetail(txId uint64) (*model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionDetail", txId)
	ret0, _ := ret[0].(*model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionDetail indicates an expected call of GetTransactionDetail
func (mr *MockTransactionUseCaseMockRecorder) GetTransactionDetail(txId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionDetail", reflect.TypeOf((*MockTransactionUseCase)(nil).GetTransactionDetail), txId)
}
