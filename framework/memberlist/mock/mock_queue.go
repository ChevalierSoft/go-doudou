// Code generated by MockGen. DO NOT EDIT.
// Source: ./queue.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	memberlist "github.com/unionj-cloud/go-doudou/framework/memberlist"
)

// MockBroadcast is a mock of Broadcast interface.
type MockBroadcast struct {
	ctrl     *gomock.Controller
	recorder *MockBroadcastMockRecorder
}

// MockBroadcastMockRecorder is the mock recorder for MockBroadcast.
type MockBroadcastMockRecorder struct {
	mock *MockBroadcast
}

// NewMockBroadcast creates a new mock instance.
func NewMockBroadcast(ctrl *gomock.Controller) *MockBroadcast {
	mock := &MockBroadcast{ctrl: ctrl}
	mock.recorder = &MockBroadcastMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBroadcast) EXPECT() *MockBroadcastMockRecorder {
	return m.recorder
}

// Finished mocks base method.
func (m *MockBroadcast) Finished() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Finished")
}

// Finished indicates an expected call of Finished.
func (mr *MockBroadcastMockRecorder) Finished() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finished", reflect.TypeOf((*MockBroadcast)(nil).Finished))
}

// Invalidates mocks base method.
func (m *MockBroadcast) Invalidates(b memberlist.Broadcast) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Invalidates", b)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Invalidates indicates an expected call of Invalidates.
func (mr *MockBroadcastMockRecorder) Invalidates(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invalidates", reflect.TypeOf((*MockBroadcast)(nil).Invalidates), b)
}

// Message mocks base method.
func (m *MockBroadcast) Message() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Message")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Message indicates an expected call of Message.
func (mr *MockBroadcastMockRecorder) Message() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Message", reflect.TypeOf((*MockBroadcast)(nil).Message))
}

// MockNamedBroadcast is a mock of NamedBroadcast interface.
type MockNamedBroadcast struct {
	ctrl     *gomock.Controller
	recorder *MockNamedBroadcastMockRecorder
}

// MockNamedBroadcastMockRecorder is the mock recorder for MockNamedBroadcast.
type MockNamedBroadcastMockRecorder struct {
	mock *MockNamedBroadcast
}

// NewMockNamedBroadcast creates a new mock instance.
func NewMockNamedBroadcast(ctrl *gomock.Controller) *MockNamedBroadcast {
	mock := &MockNamedBroadcast{ctrl: ctrl}
	mock.recorder = &MockNamedBroadcastMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamedBroadcast) EXPECT() *MockNamedBroadcastMockRecorder {
	return m.recorder
}

// Finished mocks base method.
func (m *MockNamedBroadcast) Finished() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Finished")
}

// Finished indicates an expected call of Finished.
func (mr *MockNamedBroadcastMockRecorder) Finished() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finished", reflect.TypeOf((*MockNamedBroadcast)(nil).Finished))
}

// Invalidates mocks base method.
func (m *MockNamedBroadcast) Invalidates(b memberlist.Broadcast) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Invalidates", b)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Invalidates indicates an expected call of Invalidates.
func (mr *MockNamedBroadcastMockRecorder) Invalidates(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invalidates", reflect.TypeOf((*MockNamedBroadcast)(nil).Invalidates), b)
}

// Message mocks base method.
func (m *MockNamedBroadcast) Message() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Message")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Message indicates an expected call of Message.
func (mr *MockNamedBroadcastMockRecorder) Message() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Message", reflect.TypeOf((*MockNamedBroadcast)(nil).Message))
}

// Name mocks base method.
func (m *MockNamedBroadcast) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockNamedBroadcastMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockNamedBroadcast)(nil).Name))
}

// MockUniqueBroadcast is a mock of UniqueBroadcast interface.
type MockUniqueBroadcast struct {
	ctrl     *gomock.Controller
	recorder *MockUniqueBroadcastMockRecorder
}

// MockUniqueBroadcastMockRecorder is the mock recorder for MockUniqueBroadcast.
type MockUniqueBroadcastMockRecorder struct {
	mock *MockUniqueBroadcast
}

// NewMockUniqueBroadcast creates a new mock instance.
func NewMockUniqueBroadcast(ctrl *gomock.Controller) *MockUniqueBroadcast {
	mock := &MockUniqueBroadcast{ctrl: ctrl}
	mock.recorder = &MockUniqueBroadcastMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUniqueBroadcast) EXPECT() *MockUniqueBroadcastMockRecorder {
	return m.recorder
}

// Finished mocks base method.
func (m *MockUniqueBroadcast) Finished() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Finished")
}

// Finished indicates an expected call of Finished.
func (mr *MockUniqueBroadcastMockRecorder) Finished() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finished", reflect.TypeOf((*MockUniqueBroadcast)(nil).Finished))
}

// Invalidates mocks base method.
func (m *MockUniqueBroadcast) Invalidates(b memberlist.Broadcast) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Invalidates", b)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Invalidates indicates an expected call of Invalidates.
func (mr *MockUniqueBroadcastMockRecorder) Invalidates(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invalidates", reflect.TypeOf((*MockUniqueBroadcast)(nil).Invalidates), b)
}

// Message mocks base method.
func (m *MockUniqueBroadcast) Message() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Message")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Message indicates an expected call of Message.
func (mr *MockUniqueBroadcastMockRecorder) Message() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Message", reflect.TypeOf((*MockUniqueBroadcast)(nil).Message))
}

// UniqueBroadcast mocks base method.
func (m *MockUniqueBroadcast) UniqueBroadcast() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UniqueBroadcast")
}

// UniqueBroadcast indicates an expected call of UniqueBroadcast.
func (mr *MockUniqueBroadcastMockRecorder) UniqueBroadcast() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UniqueBroadcast", reflect.TypeOf((*MockUniqueBroadcast)(nil).UniqueBroadcast))
}
