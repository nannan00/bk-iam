// Code generated by MockGen. DO NOT EDIT.
// Source: policy_v2.go

// Package mock is a generated GoMock package.
package mock

import (
	types "iam/pkg/abac/types"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPolicyControllerV2 is a mock of PolicyControllerV2 interface.
type MockPolicyControllerV2 struct {
	ctrl     *gomock.Controller
	recorder *MockPolicyControllerV2MockRecorder
}

// MockPolicyControllerV2MockRecorder is the mock recorder for MockPolicyControllerV2.
type MockPolicyControllerV2MockRecorder struct {
	mock *MockPolicyControllerV2
}

// NewMockPolicyControllerV2 creates a new mock instance.
func NewMockPolicyControllerV2(ctrl *gomock.Controller) *MockPolicyControllerV2 {
	mock := &MockPolicyControllerV2{ctrl: ctrl}
	mock.recorder = &MockPolicyControllerV2MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPolicyControllerV2) EXPECT() *MockPolicyControllerV2MockRecorder {
	return m.recorder
}

// Alter mocks base method.
func (m *MockPolicyControllerV2) Alter(systemID, subjectType, subjectID string, templateID int64, createPolicies, updatePolicies []types.Policy, deletePolicyIDs []int64, resourceChangedActions []types.ResourceChangedAction, groupAuthType int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Alter", systemID, subjectType, subjectID, templateID, createPolicies, updatePolicies, deletePolicyIDs, resourceChangedActions, groupAuthType)
	ret0, _ := ret[0].(error)
	return ret0
}

// Alter indicates an expected call of Alter.
func (mr *MockPolicyControllerV2MockRecorder) Alter(systemID, subjectType, subjectID, templateID, createPolicies, updatePolicies, deletePolicyIDs, resourceChangedActions, groupAuthType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Alter", reflect.TypeOf((*MockPolicyControllerV2)(nil).Alter), systemID, subjectType, subjectID, templateID, createPolicies, updatePolicies, deletePolicyIDs, resourceChangedActions, groupAuthType)
}
