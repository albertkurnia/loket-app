// Code generated by MockGen. DO NOT EDIT.
// Source: modules/location/query/query.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	model "loket-app/modules/location/model"
	reflect "reflect"
)

// MockLocationQuery is a mock of LocationQuery interface
type MockLocationQuery struct {
	ctrl     *gomock.Controller
	recorder *MockLocationQueryMockRecorder
}

// MockLocationQueryMockRecorder is the mock recorder for MockLocationQuery
type MockLocationQueryMockRecorder struct {
	mock *MockLocationQuery
}

// NewMockLocationQuery creates a new mock instance
func NewMockLocationQuery(ctrl *gomock.Controller) *MockLocationQuery {
	mock := &MockLocationQuery{ctrl: ctrl}
	mock.recorder = &MockLocationQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLocationQuery) EXPECT() *MockLocationQueryMockRecorder {
	return m.recorder
}

// InsertLocation mocks base method
func (m *MockLocationQuery) InsertLocation(data *model.CreateLocationReq) (*model.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertLocation", data)
	ret0, _ := ret[0].(*model.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertLocation indicates an expected call of InsertLocation
func (mr *MockLocationQueryMockRecorder) InsertLocation(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertLocation", reflect.TypeOf((*MockLocationQuery)(nil).InsertLocation), data)
}

// LoadLocationByID mocks base method
func (m *MockLocationQuery) LoadLocationByID(id uint64) (*model.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadLocationByID", id)
	ret0, _ := ret[0].(*model.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadLocationByID indicates an expected call of LoadLocationByID
func (mr *MockLocationQueryMockRecorder) LoadLocationByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadLocationByID", reflect.TypeOf((*MockLocationQuery)(nil).LoadLocationByID), id)
}
