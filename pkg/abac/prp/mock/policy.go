// Code generated by MockGen. DO NOT EDIT.
// Source: policy.go

// Package mock is a generated GoMock package.
package mock

import (
	types "iam/pkg/abac/types"
	debug "iam/pkg/logging/debug"
	types0 "iam/pkg/service/types"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
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

// GetExpressionsFromCache mocks base method.
func (m *MockPolicyManager) GetExpressionsFromCache(actionPK int64, expressionPKs []int64) ([]types0.AuthExpression, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExpressionsFromCache", actionPK, expressionPKs)
	ret0, _ := ret[0].([]types0.AuthExpression)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExpressionsFromCache indicates an expected call of GetExpressionsFromCache.
func (mr *MockPolicyManagerMockRecorder) GetExpressionsFromCache(actionPK, expressionPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExpressionsFromCache", reflect.TypeOf((*MockPolicyManager)(nil).GetExpressionsFromCache), actionPK, expressionPKs)
}

// ListBySubjectAction mocks base method.
func (m *MockPolicyManager) ListBySubjectAction(system string, subject types.Subject, action types.Action, effectGroupPKs []int64, withRbacPolicies, withoutCache bool, entry *debug.Entry) ([]types.AuthPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySubjectAction", system, subject, action, effectGroupPKs, withRbacPolicies, withoutCache, entry)
	ret0, _ := ret[0].([]types.AuthPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySubjectAction indicates an expected call of ListBySubjectAction.
func (mr *MockPolicyManagerMockRecorder) ListBySubjectAction(system, subject, action, effectGroupPKs, withRbacPolicies, withoutCache, entry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySubjectAction", reflect.TypeOf((*MockPolicyManager)(nil).ListBySubjectAction), system, subject, action, effectGroupPKs, withRbacPolicies, withoutCache, entry)
}
