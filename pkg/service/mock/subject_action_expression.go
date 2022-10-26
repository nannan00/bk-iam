// Code generated by MockGen. DO NOT EDIT.
// Source: subject_action_expression.go

// Package mock is a generated GoMock package.
package mock

import (
	types "iam/pkg/service/types"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockSubjectActionExpressionService is a mock of SubjectActionExpressionService interface.
type MockSubjectActionExpressionService struct {
	ctrl     *gomock.Controller
	recorder *MockSubjectActionExpressionServiceMockRecorder
}

// MockSubjectActionExpressionServiceMockRecorder is the mock recorder for MockSubjectActionExpressionService.
type MockSubjectActionExpressionServiceMockRecorder struct {
	mock *MockSubjectActionExpressionService
}

// NewMockSubjectActionExpressionService creates a new mock instance.
func NewMockSubjectActionExpressionService(ctrl *gomock.Controller) *MockSubjectActionExpressionService {
	mock := &MockSubjectActionExpressionService{ctrl: ctrl}
	mock.recorder = &MockSubjectActionExpressionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubjectActionExpressionService) EXPECT() *MockSubjectActionExpressionServiceMockRecorder {
	return m.recorder
}

// BulkDeleteBySubjectPKsWithTx mocks base method.
func (m *MockSubjectActionExpressionService) BulkDeleteBySubjectPKsWithTx(tx *sqlx.Tx, subjectPKs []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteBySubjectPKsWithTx", tx, subjectPKs)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDeleteBySubjectPKsWithTx indicates an expected call of BulkDeleteBySubjectPKsWithTx.
func (mr *MockSubjectActionExpressionServiceMockRecorder) BulkDeleteBySubjectPKsWithTx(tx, subjectPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteBySubjectPKsWithTx", reflect.TypeOf((*MockSubjectActionExpressionService)(nil).BulkDeleteBySubjectPKsWithTx), tx, subjectPKs)
}

// CreateOrUpdateWithTx mocks base method.
func (m *MockSubjectActionExpressionService) CreateOrUpdateWithTx(tx *sqlx.Tx, expression types.SubjectActionExpression) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateWithTx", tx, expression)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateWithTx indicates an expected call of CreateOrUpdateWithTx.
func (mr *MockSubjectActionExpressionServiceMockRecorder) CreateOrUpdateWithTx(tx, expression interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateWithTx", reflect.TypeOf((*MockSubjectActionExpressionService)(nil).CreateOrUpdateWithTx), tx, expression)
}

// DeleteByActionPKWithTx mocks base method.
func (m *MockSubjectActionExpressionService) DeleteByActionPKWithTx(tx *sqlx.Tx, actionPK int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByActionPKWithTx", tx, actionPK)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByActionPKWithTx indicates an expected call of DeleteByActionPKWithTx.
func (mr *MockSubjectActionExpressionServiceMockRecorder) DeleteByActionPKWithTx(tx, actionPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByActionPKWithTx", reflect.TypeOf((*MockSubjectActionExpressionService)(nil).DeleteByActionPKWithTx), tx, actionPK)
}

// DeleteBySubjectActionWithTx mocks base method.
func (m *MockSubjectActionExpressionService) DeleteBySubjectActionWithTx(tx *sqlx.Tx, subjectPK, actionPK int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBySubjectActionWithTx", tx, subjectPK, actionPK)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBySubjectActionWithTx indicates an expected call of DeleteBySubjectActionWithTx.
func (mr *MockSubjectActionExpressionServiceMockRecorder) DeleteBySubjectActionWithTx(tx, subjectPK, actionPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBySubjectActionWithTx", reflect.TypeOf((*MockSubjectActionExpressionService)(nil).DeleteBySubjectActionWithTx), tx, subjectPK, actionPK)
}

// ListBySubjectAction mocks base method.
func (m *MockSubjectActionExpressionService) ListBySubjectAction(subjectPKs []int64, actionPK int64) ([]types.SubjectActionExpression, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySubjectAction", subjectPKs, actionPK)
	ret0, _ := ret[0].([]types.SubjectActionExpression)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySubjectAction indicates an expected call of ListBySubjectAction.
func (mr *MockSubjectActionExpressionServiceMockRecorder) ListBySubjectAction(subjectPKs, actionPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySubjectAction", reflect.TypeOf((*MockSubjectActionExpressionService)(nil).ListBySubjectAction), subjectPKs, actionPK)
}
