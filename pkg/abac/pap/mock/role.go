// Code generated by MockGen. DO NOT EDIT.
// Source: role.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	pap "iam/pkg/abac/pap"
	reflect "reflect"
)

// MockRoleController is a mock of RoleController interface
type MockRoleController struct {
	ctrl     *gomock.Controller
	recorder *MockRoleControllerMockRecorder
}

// MockRoleControllerMockRecorder is the mock recorder for MockRoleController
type MockRoleControllerMockRecorder struct {
	mock *MockRoleController
}

// NewMockRoleController creates a new mock instance
func NewMockRoleController(ctrl *gomock.Controller) *MockRoleController {
	mock := &MockRoleController{ctrl: ctrl}
	mock.recorder = &MockRoleControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRoleController) EXPECT() *MockRoleControllerMockRecorder {
	return m.recorder
}

// ListSubjectByRole mocks base method
func (m *MockRoleController) ListSubjectByRole(roleType, system string) ([]pap.Subject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSubjectByRole", roleType, system)
	ret0, _ := ret[0].([]pap.Subject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubjectByRole indicates an expected call of ListSubjectByRole
func (mr *MockRoleControllerMockRecorder) ListSubjectByRole(roleType, system interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubjectByRole", reflect.TypeOf((*MockRoleController)(nil).ListSubjectByRole), roleType, system)
}

// BulkAddSubjects mocks base method
func (m *MockRoleController) BulkAddSubjects(roleType, system string, subjects []pap.Subject) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkAddSubjects", roleType, system, subjects)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkAddSubjects indicates an expected call of BulkAddSubjects
func (mr *MockRoleControllerMockRecorder) BulkAddSubjects(roleType, system, subjects interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkAddSubjects", reflect.TypeOf((*MockRoleController)(nil).BulkAddSubjects), roleType, system, subjects)
}

// BulkDeleteSubjects mocks base method
func (m *MockRoleController) BulkDeleteSubjects(roleType, system string, subjects []pap.Subject) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteSubjects", roleType, system, subjects)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDeleteSubjects indicates an expected call of BulkDeleteSubjects
func (mr *MockRoleControllerMockRecorder) BulkDeleteSubjects(roleType, system, subjects interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteSubjects", reflect.TypeOf((*MockRoleController)(nil).BulkDeleteSubjects), roleType, system, subjects)
}
