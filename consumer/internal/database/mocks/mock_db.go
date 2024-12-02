// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "consumer/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockInterfacePadtgresDB is a mock of InterfacePadtgresDB interface.
type MockInterfacePadtgresDB struct {
	ctrl     *gomock.Controller
	recorder *MockInterfacePadtgresDBMockRecorder
}

// MockInterfacePadtgresDBMockRecorder is the mock recorder for MockInterfacePadtgresDB.
type MockInterfacePadtgresDBMockRecorder struct {
	mock *MockInterfacePadtgresDB
}

// NewMockInterfacePadtgresDB creates a new mock instance.
func NewMockInterfacePadtgresDB(ctrl *gomock.Controller) *MockInterfacePadtgresDB {
	mock := &MockInterfacePadtgresDB{ctrl: ctrl}
	mock.recorder = &MockInterfacePadtgresDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterfacePadtgresDB) EXPECT() *MockInterfacePadtgresDBMockRecorder {
	return m.recorder
}

// AddOrder mocks base method.
func (m *MockInterfacePadtgresDB) AddOrder(order models.Order) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrder", order)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOrder indicates an expected call of AddOrder.
func (mr *MockInterfacePadtgresDBMockRecorder) AddOrder(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrder", reflect.TypeOf((*MockInterfacePadtgresDB)(nil).AddOrder), order)
}

// GetOrder mocks base method.
func (m *MockInterfacePadtgresDB) GetOrder(order_uuid string) (models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", order_uuid)
	ret0, _ := ret[0].(models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockInterfacePadtgresDBMockRecorder) GetOrder(order_uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockInterfacePadtgresDB)(nil).GetOrder), order_uuid)
}