// Code generated by MockGen. DO NOT EDIT.
// Source: subject_action_expression.go

// Package mock is a generated GoMock package.
package mock

import (
	dao "iam/pkg/database/dao"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockSubjectActionExpressionManager is a mock of SubjectActionExpressionManager interface.
type MockSubjectActionExpressionManager struct {
	ctrl     *gomock.Controller
	recorder *MockSubjectActionExpressionManagerMockRecorder
}

// MockSubjectActionExpressionManagerMockRecorder is the mock recorder for MockSubjectActionExpressionManager.
type MockSubjectActionExpressionManagerMockRecorder struct {
	mock *MockSubjectActionExpressionManager
}

// NewMockSubjectActionExpressionManager creates a new mock instance.
func NewMockSubjectActionExpressionManager(ctrl *gomock.Controller) *MockSubjectActionExpressionManager {
	mock := &MockSubjectActionExpressionManager{ctrl: ctrl}
	mock.recorder = &MockSubjectActionExpressionManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubjectActionExpressionManager) EXPECT() *MockSubjectActionExpressionManagerMockRecorder {
	return m.recorder
}

// BulkDeleteBySubjectPKsWithTx mocks base method.
func (m *MockSubjectActionExpressionManager) BulkDeleteBySubjectPKsWithTx(tx *sqlx.Tx, pks []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteBySubjectPKsWithTx", tx, pks)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDeleteBySubjectPKsWithTx indicates an expected call of BulkDeleteBySubjectPKsWithTx.
func (mr *MockSubjectActionExpressionManagerMockRecorder) BulkDeleteBySubjectPKsWithTx(tx, pks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteBySubjectPKsWithTx", reflect.TypeOf((*MockSubjectActionExpressionManager)(nil).BulkDeleteBySubjectPKsWithTx), tx, pks)
}

// CreateWithTx mocks base method.
func (m *MockSubjectActionExpressionManager) CreateWithTx(tx *sqlx.Tx, subjectActionExpression dao.SubjectActionExpression) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWithTx", tx, subjectActionExpression)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWithTx indicates an expected call of CreateWithTx.
func (mr *MockSubjectActionExpressionManagerMockRecorder) CreateWithTx(tx, subjectActionExpression interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWithTx", reflect.TypeOf((*MockSubjectActionExpressionManager)(nil).CreateWithTx), tx, subjectActionExpression)
}

// DeleteByActionPKWithTx mocks base method.
func (m *MockSubjectActionExpressionManager) DeleteByActionPKWithTx(tx *sqlx.Tx, actionPK, limit int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByActionPKWithTx", tx, actionPK, limit)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteByActionPKWithTx indicates an expected call of DeleteByActionPKWithTx.
func (mr *MockSubjectActionExpressionManagerMockRecorder) DeleteByActionPKWithTx(tx, actionPK, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByActionPKWithTx", reflect.TypeOf((*MockSubjectActionExpressionManager)(nil).DeleteByActionPKWithTx), tx, actionPK, limit)
}

// GetBySubjectAction mocks base method.
func (m *MockSubjectActionExpressionManager) GetBySubjectAction(subjectPK, actionPK int64) (dao.SubjectActionExpression, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySubjectAction", subjectPK, actionPK)
	ret0, _ := ret[0].(dao.SubjectActionExpression)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySubjectAction indicates an expected call of GetBySubjectAction.
func (mr *MockSubjectActionExpressionManagerMockRecorder) GetBySubjectAction(subjectPK, actionPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySubjectAction", reflect.TypeOf((*MockSubjectActionExpressionManager)(nil).GetBySubjectAction), subjectPK, actionPK)
}

// ListBySubjectAction mocks base method.
func (m *MockSubjectActionExpressionManager) ListBySubjectAction(subjectPKs []int64, actionPK int64) ([]dao.SubjectActionExpression, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySubjectAction", subjectPKs, actionPK)
	ret0, _ := ret[0].([]dao.SubjectActionExpression)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySubjectAction indicates an expected call of ListBySubjectAction.
func (mr *MockSubjectActionExpressionManagerMockRecorder) ListBySubjectAction(subjectPKs, actionPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySubjectAction", reflect.TypeOf((*MockSubjectActionExpressionManager)(nil).ListBySubjectAction), subjectPKs, actionPK)
}

// UpdateExpressionExpiredAtWithTx mocks base method.
func (m *MockSubjectActionExpressionManager) UpdateExpressionExpiredAtWithTx(tx *sqlx.Tx, pk int64, expression, signature string, expiredAt int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExpressionExpiredAtWithTx", tx, pk, expression, signature, expiredAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExpressionExpiredAtWithTx indicates an expected call of UpdateExpressionExpiredAtWithTx.
func (mr *MockSubjectActionExpressionManagerMockRecorder) UpdateExpressionExpiredAtWithTx(tx, pk, expression, signature, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExpressionExpiredAtWithTx", reflect.TypeOf((*MockSubjectActionExpressionManager)(nil).UpdateExpressionExpiredAtWithTx), tx, pk, expression, signature, expiredAt)
}
