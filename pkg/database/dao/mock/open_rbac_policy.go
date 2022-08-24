// Code generated by MockGen. DO NOT EDIT.
// Source: open_rbac_policy.go

// Package mock is a generated GoMock package.
package mock

import (
	dao "iam/pkg/database/dao"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOpenRbacPolicyManager is a mock of OpenRbacPolicyManager interface.
type MockOpenRbacPolicyManager struct {
	ctrl     *gomock.Controller
	recorder *MockOpenRbacPolicyManagerMockRecorder
}

// MockOpenRbacPolicyManagerMockRecorder is the mock recorder for MockOpenRbacPolicyManager.
type MockOpenRbacPolicyManagerMockRecorder struct {
	mock *MockOpenRbacPolicyManager
}

// NewMockOpenRbacPolicyManager creates a new mock instance.
func NewMockOpenRbacPolicyManager(ctrl *gomock.Controller) *MockOpenRbacPolicyManager {
	mock := &MockOpenRbacPolicyManager{ctrl: ctrl}
	mock.recorder = &MockOpenRbacPolicyManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenRbacPolicyManager) EXPECT() *MockOpenRbacPolicyManagerMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockOpenRbacPolicyManager) Get(pk int64) (dao.OpenRbacPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", pk)
	ret0, _ := ret[0].(dao.OpenRbacPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockOpenRbacPolicyManagerMockRecorder) Get(pk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOpenRbacPolicyManager)(nil).Get), pk)
}

// GetCountByActionBeforeExpiredAt mocks base method.
func (m *MockOpenRbacPolicyManager) GetCountByActionBeforeExpiredAt(actionPK, expiredAt int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCountByActionBeforeExpiredAt", actionPK, expiredAt)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCountByActionBeforeExpiredAt indicates an expected call of GetCountByActionBeforeExpiredAt.
func (mr *MockOpenRbacPolicyManagerMockRecorder) GetCountByActionBeforeExpiredAt(actionPK, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCountByActionBeforeExpiredAt", reflect.TypeOf((*MockOpenRbacPolicyManager)(nil).GetCountByActionBeforeExpiredAt), actionPK, expiredAt)
}

// ListByPKs mocks base method.
func (m *MockOpenRbacPolicyManager) ListByPKs(pks []int64) ([]dao.OpenRbacPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByPKs", pks)
	ret0, _ := ret[0].([]dao.OpenRbacPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByPKs indicates an expected call of ListByPKs.
func (mr *MockOpenRbacPolicyManagerMockRecorder) ListByPKs(pks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByPKs", reflect.TypeOf((*MockOpenRbacPolicyManager)(nil).ListByPKs), pks)
}

// ListPagingByActionPKBeforeExpiredAt mocks base method.
func (m *MockOpenRbacPolicyManager) ListPagingByActionPKBeforeExpiredAt(actionPK, expiredAt, offset, limit int64) ([]dao.OpenRbacPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingByActionPKBeforeExpiredAt", actionPK, expiredAt, offset, limit)
	ret0, _ := ret[0].([]dao.OpenRbacPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingByActionPKBeforeExpiredAt indicates an expected call of ListPagingByActionPKBeforeExpiredAt.
func (mr *MockOpenRbacPolicyManagerMockRecorder) ListPagingByActionPKBeforeExpiredAt(actionPK, expiredAt, offset, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingByActionPKBeforeExpiredAt", reflect.TypeOf((*MockOpenRbacPolicyManager)(nil).ListPagingByActionPKBeforeExpiredAt), actionPK, expiredAt, offset, limit)
}
