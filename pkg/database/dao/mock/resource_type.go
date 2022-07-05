// Code generated by MockGen. DO NOT EDIT.
// Source: resource_type.go

// Package mock is a generated GoMock package.
package mock

import (
	dao "iam/pkg/database/dao"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockResourceTypeManager is a mock of ResourceTypeManager interface.
type MockResourceTypeManager struct {
	ctrl     *gomock.Controller
	recorder *MockResourceTypeManagerMockRecorder
}

// MockResourceTypeManagerMockRecorder is the mock recorder for MockResourceTypeManager.
type MockResourceTypeManagerMockRecorder struct {
	mock *MockResourceTypeManager
}

// NewMockResourceTypeManager creates a new mock instance.
func NewMockResourceTypeManager(ctrl *gomock.Controller) *MockResourceTypeManager {
	mock := &MockResourceTypeManager{ctrl: ctrl}
	mock.recorder = &MockResourceTypeManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceTypeManager) EXPECT() *MockResourceTypeManagerMockRecorder {
	return m.recorder
}

// BulkCreateWithTx mocks base method.
func (m *MockResourceTypeManager) BulkCreateWithTx(tx *sqlx.Tx, resourceTypes []dao.ResourceType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreateWithTx", tx, resourceTypes)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkCreateWithTx indicates an expected call of BulkCreateWithTx.
func (mr *MockResourceTypeManagerMockRecorder) BulkCreateWithTx(tx, resourceTypes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreateWithTx", reflect.TypeOf((*MockResourceTypeManager)(nil).BulkCreateWithTx), tx, resourceTypes)
}

// BulkDeleteWithTx mocks base method.
func (m *MockResourceTypeManager) BulkDeleteWithTx(tx *sqlx.Tx, system string, ids []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteWithTx", tx, system, ids)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDeleteWithTx indicates an expected call of BulkDeleteWithTx.
func (mr *MockResourceTypeManagerMockRecorder) BulkDeleteWithTx(tx, system, ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteWithTx", reflect.TypeOf((*MockResourceTypeManager)(nil).BulkDeleteWithTx), tx, system, ids)
}

// GetPK mocks base method.
func (m *MockResourceTypeManager) GetPK(system, id string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPK", system, id)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPK indicates an expected call of GetPK.
func (mr *MockResourceTypeManagerMockRecorder) GetPK(system, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPK", reflect.TypeOf((*MockResourceTypeManager)(nil).GetPK), system, id)
}
