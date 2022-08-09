// Code generated by MockGen. DO NOT EDIT.
// Source: policy.go

// Package mock is a generated GoMock package.
package mock

import (
	dao "iam/pkg/database/dao"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockPolicyManager is a mock of PolicyManager interface.
type MockPolicyManager struct {
	ctrl     *gomock.Controller
	recorder *MockPolicyManagerMockRecorder
}

// MockPolicyManagerMockRecorder is the mock recorder for MockPolicyManager.
type MockPolicyManagerMockRecorder struct {
	mock *MockPolicyManager
}

// NewMockPolicyManager creates a new mock instance.
func NewMockPolicyManager(ctrl *gomock.Controller) *MockPolicyManager {
	mock := &MockPolicyManager{ctrl: ctrl}
	mock.recorder = &MockPolicyManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPolicyManager) EXPECT() *MockPolicyManagerMockRecorder {
	return m.recorder
}

// BulkCreateWithTx mocks base method.
func (m *MockPolicyManager) BulkCreateWithTx(tx *sqlx.Tx, policies []dao.Policy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreateWithTx", tx, policies)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkCreateWithTx indicates an expected call of BulkCreateWithTx.
func (mr *MockPolicyManagerMockRecorder) BulkCreateWithTx(tx, policies interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreateWithTx", reflect.TypeOf((*MockPolicyManager)(nil).BulkCreateWithTx), tx, policies)
}

// BulkDeleteBySubjectPKsWithTx mocks base method.
func (m *MockPolicyManager) BulkDeleteBySubjectPKsWithTx(tx *sqlx.Tx, subjectPKs []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteBySubjectPKsWithTx", tx, subjectPKs)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDeleteBySubjectPKsWithTx indicates an expected call of BulkDeleteBySubjectPKsWithTx.
func (mr *MockPolicyManagerMockRecorder) BulkDeleteBySubjectPKsWithTx(tx, subjectPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteBySubjectPKsWithTx", reflect.TypeOf((*MockPolicyManager)(nil).BulkDeleteBySubjectPKsWithTx), tx, subjectPKs)
}

// BulkDeleteByTemplatePKsWithTx mocks base method.
func (m *MockPolicyManager) BulkDeleteByTemplatePKsWithTx(tx *sqlx.Tx, subjectPK, templateID int64, pks []int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteByTemplatePKsWithTx", tx, subjectPK, templateID, pks)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkDeleteByTemplatePKsWithTx indicates an expected call of BulkDeleteByTemplatePKsWithTx.
func (mr *MockPolicyManagerMockRecorder) BulkDeleteByTemplatePKsWithTx(tx, subjectPK, templateID, pks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteByTemplatePKsWithTx", reflect.TypeOf((*MockPolicyManager)(nil).BulkDeleteByTemplatePKsWithTx), tx, subjectPK, templateID, pks)
}

// BulkUpdateExpiredAtWithTx mocks base method.
func (m *MockPolicyManager) BulkUpdateExpiredAtWithTx(tx *sqlx.Tx, policies []dao.Policy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkUpdateExpiredAtWithTx", tx, policies)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkUpdateExpiredAtWithTx indicates an expected call of BulkUpdateExpiredAtWithTx.
func (mr *MockPolicyManagerMockRecorder) BulkUpdateExpiredAtWithTx(tx, policies interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkUpdateExpiredAtWithTx", reflect.TypeOf((*MockPolicyManager)(nil).BulkUpdateExpiredAtWithTx), tx, policies)
}

// BulkUpdateExpressionPKWithTx mocks base method.
func (m *MockPolicyManager) BulkUpdateExpressionPKWithTx(tx *sqlx.Tx, policies []dao.Policy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkUpdateExpressionPKWithTx", tx, policies)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkUpdateExpressionPKWithTx indicates an expected call of BulkUpdateExpressionPKWithTx.
func (mr *MockPolicyManagerMockRecorder) BulkUpdateExpressionPKWithTx(tx, policies interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkUpdateExpressionPKWithTx", reflect.TypeOf((*MockPolicyManager)(nil).BulkUpdateExpressionPKWithTx), tx, policies)
}

// DeleteByActionPKWithTx mocks base method.
func (m *MockPolicyManager) DeleteByActionPKWithTx(tx *sqlx.Tx, actionPK, limit int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByActionPKWithTx", tx, actionPK, limit)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteByActionPKWithTx indicates an expected call of DeleteByActionPKWithTx.
func (mr *MockPolicyManagerMockRecorder) DeleteByActionPKWithTx(tx, actionPK, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByActionPKWithTx", reflect.TypeOf((*MockPolicyManager)(nil).DeleteByActionPKWithTx), tx, actionPK, limit)
}

// Get mocks base method.
func (m *MockPolicyManager) Get(pk int64) (dao.Policy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", pk)
	ret0, _ := ret[0].(dao.Policy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPolicyManagerMockRecorder) Get(pk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPolicyManager)(nil).Get), pk)
}

// GetCountByActionBeforeExpiredAt mocks base method.
func (m *MockPolicyManager) GetCountByActionBeforeExpiredAt(actionPK, expiredAt int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCountByActionBeforeExpiredAt", actionPK, expiredAt)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCountByActionBeforeExpiredAt indicates an expected call of GetCountByActionBeforeExpiredAt.
func (mr *MockPolicyManagerMockRecorder) GetCountByActionBeforeExpiredAt(actionPK, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCountByActionBeforeExpiredAt", reflect.TypeOf((*MockPolicyManager)(nil).GetCountByActionBeforeExpiredAt), actionPK, expiredAt)
}

// HasAnyByActionPK mocks base method.
func (m *MockPolicyManager) HasAnyByActionPK(actionPK int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasAnyByActionPK", actionPK)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasAnyByActionPK indicates an expected call of HasAnyByActionPK.
func (mr *MockPolicyManagerMockRecorder) HasAnyByActionPK(actionPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasAnyByActionPK", reflect.TypeOf((*MockPolicyManager)(nil).HasAnyByActionPK), actionPK)
}

// ListAuthBySubjectAction mocks base method.
func (m *MockPolicyManager) ListAuthBySubjectAction(subjectPKs []int64, actionPK, expiredAt int64) ([]dao.AuthPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAuthBySubjectAction", subjectPKs, actionPK, expiredAt)
	ret0, _ := ret[0].([]dao.AuthPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAuthBySubjectAction indicates an expected call of ListAuthBySubjectAction.
func (mr *MockPolicyManagerMockRecorder) ListAuthBySubjectAction(subjectPKs, actionPK, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAuthBySubjectAction", reflect.TypeOf((*MockPolicyManager)(nil).ListAuthBySubjectAction), subjectPKs, actionPK, expiredAt)
}

// ListByPKs mocks base method.
func (m *MockPolicyManager) ListByPKs(pks []int64) ([]dao.Policy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByPKs", pks)
	ret0, _ := ret[0].([]dao.Policy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByPKs indicates an expected call of ListByPKs.
func (mr *MockPolicyManagerMockRecorder) ListByPKs(pks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByPKs", reflect.TypeOf((*MockPolicyManager)(nil).ListByPKs), pks)
}

// ListBySubjectActionTemplate mocks base method.
func (m *MockPolicyManager) ListBySubjectActionTemplate(subjectPK int64, actionPKs []int64, templateID int64) ([]dao.Policy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySubjectActionTemplate", subjectPK, actionPKs, templateID)
	ret0, _ := ret[0].([]dao.Policy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySubjectActionTemplate indicates an expected call of ListBySubjectActionTemplate.
func (mr *MockPolicyManagerMockRecorder) ListBySubjectActionTemplate(subjectPK, actionPKs, templateID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySubjectActionTemplate", reflect.TypeOf((*MockPolicyManager)(nil).ListBySubjectActionTemplate), subjectPK, actionPKs, templateID)
}

// ListBySubjectPKAndPKs mocks base method.
func (m *MockPolicyManager) ListBySubjectPKAndPKs(subjectPK int64, pks []int64) ([]dao.Policy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySubjectPKAndPKs", subjectPK, pks)
	ret0, _ := ret[0].([]dao.Policy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySubjectPKAndPKs indicates an expected call of ListBySubjectPKAndPKs.
func (mr *MockPolicyManagerMockRecorder) ListBySubjectPKAndPKs(subjectPK, pks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySubjectPKAndPKs", reflect.TypeOf((*MockPolicyManager)(nil).ListBySubjectPKAndPKs), subjectPK, pks)
}

// ListBySubjectTemplateBeforeExpiredAt mocks base method.
func (m *MockPolicyManager) ListBySubjectTemplateBeforeExpiredAt(subjectPK, templateID, expiredAt int64) ([]dao.Policy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySubjectTemplateBeforeExpiredAt", subjectPK, templateID, expiredAt)
	ret0, _ := ret[0].([]dao.Policy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySubjectTemplateBeforeExpiredAt indicates an expected call of ListBySubjectTemplateBeforeExpiredAt.
func (mr *MockPolicyManagerMockRecorder) ListBySubjectTemplateBeforeExpiredAt(subjectPK, templateID, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySubjectTemplateBeforeExpiredAt", reflect.TypeOf((*MockPolicyManager)(nil).ListBySubjectTemplateBeforeExpiredAt), subjectPK, templateID, expiredAt)
}

// ListExpressionBySubjectsTemplate mocks base method.
func (m *MockPolicyManager) ListExpressionBySubjectsTemplate(subjectPKs []int64, templateID int64) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListExpressionBySubjectsTemplate", subjectPKs, templateID)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListExpressionBySubjectsTemplate indicates an expected call of ListExpressionBySubjectsTemplate.
func (mr *MockPolicyManagerMockRecorder) ListExpressionBySubjectsTemplate(subjectPKs, templateID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListExpressionBySubjectsTemplate", reflect.TypeOf((*MockPolicyManager)(nil).ListExpressionBySubjectsTemplate), subjectPKs, templateID)
}

// ListPagingByActionPKBeforeExpiredAt mocks base method.
func (m *MockPolicyManager) ListPagingByActionPKBeforeExpiredAt(actionPK, expiredAt, offset, limit int64) ([]dao.Policy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingByActionPKBeforeExpiredAt", actionPK, expiredAt, offset, limit)
	ret0, _ := ret[0].([]dao.Policy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingByActionPKBeforeExpiredAt indicates an expected call of ListPagingByActionPKBeforeExpiredAt.
func (mr *MockPolicyManagerMockRecorder) ListPagingByActionPKBeforeExpiredAt(actionPK, expiredAt, offset, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingByActionPKBeforeExpiredAt", reflect.TypeOf((*MockPolicyManager)(nil).ListPagingByActionPKBeforeExpiredAt), actionPK, expiredAt, offset, limit)
}
