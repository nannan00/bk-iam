// Code generated by MockGen. DO NOT EDIT.
// Source: system.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	types "iam/pkg/service/types"
	reflect "reflect"
)

// MockSystemService is a mock of SystemService interface
type MockSystemService struct {
	ctrl     *gomock.Controller
	recorder *MockSystemServiceMockRecorder
}

// MockSystemServiceMockRecorder is the mock recorder for MockSystemService
type MockSystemServiceMockRecorder struct {
	mock *MockSystemService
}

// NewMockSystemService creates a new mock instance
func NewMockSystemService(ctrl *gomock.Controller) *MockSystemService {
	mock := &MockSystemService{ctrl: ctrl}
	mock.recorder = &MockSystemServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSystemService) EXPECT() *MockSystemServiceMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockSystemService) Get(id string) (types.System, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(types.System)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockSystemServiceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSystemService)(nil).Get), id)
}

// Exists mocks base method
func (m *MockSystemService) Exists(id string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exists indicates an expected call of Exists
func (mr *MockSystemServiceMockRecorder) Exists(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockSystemService)(nil).Exists), id)
}

// ListAll mocks base method
func (m *MockSystemService) ListAll() ([]types.System, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll")
	ret0, _ := ret[0].([]types.System)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll
func (mr *MockSystemServiceMockRecorder) ListAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockSystemService)(nil).ListAll))
}

// Create mocks base method
func (m *MockSystemService) Create(system types.System) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", system)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockSystemServiceMockRecorder) Create(system interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSystemService)(nil).Create), system)
}

// Update mocks base method
func (m *MockSystemService) Update(id string, system types.System) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, system)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockSystemServiceMockRecorder) Update(id, system interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSystemService)(nil).Update), id, system)
}
