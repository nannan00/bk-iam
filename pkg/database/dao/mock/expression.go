// Code generated by MockGen. DO NOT EDIT.
// Source: expression.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
	dao "iam/pkg/database/dao"
	reflect "reflect"
)

// MockExpressionManager is a mock of ExpressionManager interface
type MockExpressionManager struct {
	ctrl     *gomock.Controller
	recorder *MockExpressionManagerMockRecorder
}

// MockExpressionManagerMockRecorder is the mock recorder for MockExpressionManager
type MockExpressionManagerMockRecorder struct {
	mock *MockExpressionManager
}

// NewMockExpressionManager creates a new mock instance
func NewMockExpressionManager(ctrl *gomock.Controller) *MockExpressionManager {
	mock := &MockExpressionManager{ctrl: ctrl}
	mock.recorder = &MockExpressionManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExpressionManager) EXPECT() *MockExpressionManagerMockRecorder {
	return m.recorder
}

// ListAuthByPKs mocks base method
func (m *MockExpressionManager) ListAuthByPKs(pks []int64) ([]dao.AuthExpression, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAuthByPKs", pks)
	ret0, _ := ret[0].([]dao.AuthExpression)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAuthByPKs indicates an expected call of ListAuthByPKs
func (mr *MockExpressionManagerMockRecorder) ListAuthByPKs(pks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAuthByPKs", reflect.TypeOf((*MockExpressionManager)(nil).ListAuthByPKs), pks)
}

// ListDistinctBySignaturesType mocks base method
func (m *MockExpressionManager) ListDistinctBySignaturesType(signatures []string, _type int64) ([]dao.Expression, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDistinctBySignaturesType", signatures, _type)
	ret0, _ := ret[0].([]dao.Expression)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDistinctBySignaturesType indicates an expected call of ListDistinctBySignaturesType
func (mr *MockExpressionManagerMockRecorder) ListDistinctBySignaturesType(signatures, _type interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDistinctBySignaturesType", reflect.TypeOf((*MockExpressionManager)(nil).ListDistinctBySignaturesType), signatures, _type)
}

// BulkCreateWithTx mocks base method
func (m *MockExpressionManager) BulkCreateWithTx(tx *sqlx.Tx, expressions []dao.Expression) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreateWithTx", tx, expressions)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkCreateWithTx indicates an expected call of BulkCreateWithTx
func (mr *MockExpressionManagerMockRecorder) BulkCreateWithTx(tx, expressions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreateWithTx", reflect.TypeOf((*MockExpressionManager)(nil).BulkCreateWithTx), tx, expressions)
}

// BulkUpdateWithTx mocks base method
func (m *MockExpressionManager) BulkUpdateWithTx(tx *sqlx.Tx, expressions []dao.Expression) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkUpdateWithTx", tx, expressions)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkUpdateWithTx indicates an expected call of BulkUpdateWithTx
func (mr *MockExpressionManagerMockRecorder) BulkUpdateWithTx(tx, expressions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkUpdateWithTx", reflect.TypeOf((*MockExpressionManager)(nil).BulkUpdateWithTx), tx, expressions)
}

// BulkDeleteByPKsWithTx mocks base method
func (m *MockExpressionManager) BulkDeleteByPKsWithTx(tx *sqlx.Tx, pks []int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteByPKsWithTx", tx, pks)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkDeleteByPKsWithTx indicates an expected call of BulkDeleteByPKsWithTx
func (mr *MockExpressionManagerMockRecorder) BulkDeleteByPKsWithTx(tx, pks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteByPKsWithTx", reflect.TypeOf((*MockExpressionManager)(nil).BulkDeleteByPKsWithTx), tx, pks)
}

// UpdateUnquotedType mocks base method
func (m *MockExpressionManager) UpdateUnquotedType(fromType, toType int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUnquotedType", fromType, toType)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUnquotedType indicates an expected call of UpdateUnquotedType
func (mr *MockExpressionManagerMockRecorder) UpdateUnquotedType(fromType, toType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUnquotedType", reflect.TypeOf((*MockExpressionManager)(nil).UpdateUnquotedType), fromType, toType)
}

// UpdateQuotedType mocks base method
func (m *MockExpressionManager) UpdateQuotedType(fromType, toType, updatedAt int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateQuotedType", fromType, toType, updatedAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateQuotedType indicates an expected call of UpdateQuotedType
func (mr *MockExpressionManagerMockRecorder) UpdateQuotedType(fromType, toType, updatedAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQuotedType", reflect.TypeOf((*MockExpressionManager)(nil).UpdateQuotedType), fromType, toType, updatedAt)
}

// DeleteUnquoted mocks base method
func (m *MockExpressionManager) DeleteUnquoted(_type, updatedAt int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUnquoted", _type, updatedAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUnquoted indicates an expected call of DeleteUnquoted
func (mr *MockExpressionManagerMockRecorder) DeleteUnquoted(_type, updatedAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUnquoted", reflect.TypeOf((*MockExpressionManager)(nil).DeleteUnquoted), _type, updatedAt)
}
