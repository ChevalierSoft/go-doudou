// Code generated by MockGen. DO NOT EDIT.
// Source: store/freecache/freecache.go

// Package freecache is a generated GoMock package.
package freecache

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockFreecacheClientInterface is a mock of FreecacheClientInterface interface.
type MockFreecacheClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockFreecacheClientInterfaceMockRecorder
}

// MockFreecacheClientInterfaceMockRecorder is the mock recorder for MockFreecacheClientInterface.
type MockFreecacheClientInterfaceMockRecorder struct {
	mock *MockFreecacheClientInterface
}

// NewMockFreecacheClientInterface creates a new mock instance.
func NewMockFreecacheClientInterface(ctrl *gomock.Controller) *MockFreecacheClientInterface {
	mock := &MockFreecacheClientInterface{ctrl: ctrl}
	mock.recorder = &MockFreecacheClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFreecacheClientInterface) EXPECT() *MockFreecacheClientInterfaceMockRecorder {
	return m.recorder
}

// Clear mocks base method.
func (m *MockFreecacheClientInterface) Clear() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Clear")
}

// Clear indicates an expected call of Clear.
func (mr *MockFreecacheClientInterfaceMockRecorder) Clear() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockFreecacheClientInterface)(nil).Clear))
}

// Del mocks base method.
func (m *MockFreecacheClientInterface) Del(key []byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", key)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockFreecacheClientInterfaceMockRecorder) Del(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockFreecacheClientInterface)(nil).Del), key)
}

// DelInt mocks base method.
func (m *MockFreecacheClientInterface) DelInt(key int64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelInt", key)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DelInt indicates an expected call of DelInt.
func (mr *MockFreecacheClientInterfaceMockRecorder) DelInt(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelInt", reflect.TypeOf((*MockFreecacheClientInterface)(nil).DelInt), key)
}

// Get mocks base method.
func (m *MockFreecacheClientInterface) Get(key []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockFreecacheClientInterfaceMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockFreecacheClientInterface)(nil).Get), key)
}

// GetInt mocks base method.
func (m *MockFreecacheClientInterface) GetInt(key int64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInt", key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInt indicates an expected call of GetInt.
func (mr *MockFreecacheClientInterfaceMockRecorder) GetInt(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInt", reflect.TypeOf((*MockFreecacheClientInterface)(nil).GetInt), key)
}

// Set mocks base method.
func (m *MockFreecacheClientInterface) Set(key, value []byte, expireSeconds int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, value, expireSeconds)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockFreecacheClientInterfaceMockRecorder) Set(key, value, expireSeconds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockFreecacheClientInterface)(nil).Set), key, value, expireSeconds)
}

// SetInt mocks base method.
func (m *MockFreecacheClientInterface) SetInt(key int64, value []byte, expireSeconds int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetInt", key, value, expireSeconds)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetInt indicates an expected call of SetInt.
func (mr *MockFreecacheClientInterfaceMockRecorder) SetInt(key, value, expireSeconds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInt", reflect.TypeOf((*MockFreecacheClientInterface)(nil).SetInt), key, value, expireSeconds)
}

// TTL mocks base method.
func (m *MockFreecacheClientInterface) TTL(key []byte) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TTL", key)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TTL indicates an expected call of TTL.
func (mr *MockFreecacheClientInterfaceMockRecorder) TTL(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TTL", reflect.TypeOf((*MockFreecacheClientInterface)(nil).TTL), key)
}
