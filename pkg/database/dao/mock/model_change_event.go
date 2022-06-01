// Code generated by MockGen. DO NOT EDIT.
// Source: model_change_event.go

// Package mock is a generated GoMock package.
package mock

import (
	dao "iam/pkg/database/dao"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockModelChangeEventManager is a mock of ModelChangeEventManager interface.
type MockModelChangeEventManager struct {
	ctrl     *gomock.Controller
	recorder *MockModelChangeEventManagerMockRecorder
}

// MockModelChangeEventManagerMockRecorder is the mock recorder for MockModelChangeEventManager.
type MockModelChangeEventManagerMockRecorder struct {
	mock *MockModelChangeEventManager
}

// NewMockModelChangeEventManager creates a new mock instance.
func NewMockModelChangeEventManager(ctrl *gomock.Controller) *MockModelChangeEventManager {
	mock := &MockModelChangeEventManager{ctrl: ctrl}
	mock.recorder = &MockModelChangeEventManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelChangeEventManager) EXPECT() *MockModelChangeEventManagerMockRecorder {
	return m.recorder
}

// BulkCreate mocks base method.
func (m *MockModelChangeEventManager) BulkCreate(modelChangeEvents []dao.ModelChangeEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreate", modelChangeEvents)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkCreate indicates an expected call of BulkCreate.
func (mr *MockModelChangeEventManagerMockRecorder) BulkCreate(modelChangeEvents interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreate", reflect.TypeOf((*MockModelChangeEventManager)(nil).BulkCreate), modelChangeEvents)
}

// DeleteByStatusWithTx mocks base method.
func (m *MockModelChangeEventManager) DeleteByStatusWithTx(tx *sqlx.Tx, status string, beforeUpdatedAt, limit int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByStatusWithTx", tx, status, beforeUpdatedAt, limit)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteByStatusWithTx indicates an expected call of DeleteByStatusWithTx.
func (mr *MockModelChangeEventManagerMockRecorder) DeleteByStatusWithTx(tx, status, beforeUpdatedAt, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByStatusWithTx", reflect.TypeOf((*MockModelChangeEventManager)(nil).DeleteByStatusWithTx), tx, status, beforeUpdatedAt, limit)
}

// GetByTypeModel mocks base method.
func (m *MockModelChangeEventManager) GetByTypeModel(eventType, status, modelType string, modelPK int64) (dao.ModelChangeEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByTypeModel", eventType, status, modelType, modelPK)
	ret0, _ := ret[0].(dao.ModelChangeEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByTypeModel indicates an expected call of GetByTypeModel.
func (mr *MockModelChangeEventManagerMockRecorder) GetByTypeModel(eventType, status, modelType, modelPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByTypeModel", reflect.TypeOf((*MockModelChangeEventManager)(nil).GetByTypeModel), eventType, status, modelType, modelPK)
}

// ListByStatus mocks base method.
func (m *MockModelChangeEventManager) ListByStatus(status string, limit int64) ([]dao.ModelChangeEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByStatus", status, limit)
	ret0, _ := ret[0].([]dao.ModelChangeEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByStatus indicates an expected call of ListByStatus.
func (mr *MockModelChangeEventManagerMockRecorder) ListByStatus(status, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByStatus", reflect.TypeOf((*MockModelChangeEventManager)(nil).ListByStatus), status, limit)
}

// UpdateStatusByModel mocks base method.
func (m *MockModelChangeEventManager) UpdateStatusByModel(eventType, modelType string, modelPK int64, status string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatusByModel", eventType, modelType, modelPK, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatusByModel indicates an expected call of UpdateStatusByModel.
func (mr *MockModelChangeEventManagerMockRecorder) UpdateStatusByModel(eventType, modelType, modelPK, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatusByModel", reflect.TypeOf((*MockModelChangeEventManager)(nil).UpdateStatusByModel), eventType, modelType, modelPK, status)
}

// UpdateStatusByPK mocks base method.
func (m *MockModelChangeEventManager) UpdateStatusByPK(pk int64, status string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatusByPK", pk, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatusByPK indicates an expected call of UpdateStatusByPK.
func (mr *MockModelChangeEventManagerMockRecorder) UpdateStatusByPK(pk, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatusByPK", reflect.TypeOf((*MockModelChangeEventManager)(nil).UpdateStatusByPK), pk, status)
}
