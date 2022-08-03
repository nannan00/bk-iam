// Code generated by MockGen. DO NOT EDIT.
// Source: group_alter_event.go

// Package mock is a generated GoMock package.
package mock

import (
	types "iam/pkg/service/types"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGroupAlterEventService is a mock of GroupAlterEventService interface.
type MockGroupAlterEventService struct {
	ctrl     *gomock.Controller
	recorder *MockGroupAlterEventServiceMockRecorder
}

// MockGroupAlterEventServiceMockRecorder is the mock recorder for MockGroupAlterEventService.
type MockGroupAlterEventServiceMockRecorder struct {
	mock *MockGroupAlterEventService
}

// NewMockGroupAlterEventService creates a new mock instance.
func NewMockGroupAlterEventService(ctrl *gomock.Controller) *MockGroupAlterEventService {
	mock := &MockGroupAlterEventService{ctrl: ctrl}
	mock.recorder = &MockGroupAlterEventServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupAlterEventService) EXPECT() *MockGroupAlterEventServiceMockRecorder {
	return m.recorder
}

// CreateByGroupAction mocks base method.
func (m *MockGroupAlterEventService) CreateByGroupAction(groupPK int64, actionPKs []int64) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateByGroupAction", groupPK, actionPKs)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateByGroupAction indicates an expected call of CreateByGroupAction.
func (mr *MockGroupAlterEventServiceMockRecorder) CreateByGroupAction(groupPK, actionPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateByGroupAction", reflect.TypeOf((*MockGroupAlterEventService)(nil).CreateByGroupAction), groupPK, actionPKs)
}

// CreateByGroupSubject mocks base method.
func (m *MockGroupAlterEventService) CreateByGroupSubject(groupPK int64, subjectPKs []int64) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateByGroupSubject", groupPK, subjectPKs)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateByGroupSubject indicates an expected call of CreateByGroupSubject.
func (mr *MockGroupAlterEventServiceMockRecorder) CreateByGroupSubject(groupPK, subjectPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateByGroupSubject", reflect.TypeOf((*MockGroupAlterEventService)(nil).CreateByGroupSubject), groupPK, subjectPKs)
}

// CreateBySubjectActionGroup mocks base method.
func (m *MockGroupAlterEventService) CreateBySubjectActionGroup(subjectPK, actionPK, groupPK int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBySubjectActionGroup", subjectPK, actionPK, groupPK)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBySubjectActionGroup indicates an expected call of CreateBySubjectActionGroup.
func (mr *MockGroupAlterEventServiceMockRecorder) CreateBySubjectActionGroup(subjectPK, actionPK, groupPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBySubjectActionGroup", reflect.TypeOf((*MockGroupAlterEventService)(nil).CreateBySubjectActionGroup), subjectPK, actionPK, groupPK)
}

// Delete mocks base method.
func (m *MockGroupAlterEventService) Delete(pk int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", pk)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockGroupAlterEventServiceMockRecorder) Delete(pk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGroupAlterEventService)(nil).Delete), pk)
}

// Get mocks base method.
func (m *MockGroupAlterEventService) Get(pk int64) (types.GroupAlterEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", pk)
	ret0, _ := ret[0].(types.GroupAlterEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockGroupAlterEventServiceMockRecorder) Get(pk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGroupAlterEventService)(nil).Get), pk)
}

// IncrCheckCount mocks base method.
func (m *MockGroupAlterEventService) IncrCheckCount(pk int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrCheckCount", pk)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrCheckCount indicates an expected call of IncrCheckCount.
func (mr *MockGroupAlterEventServiceMockRecorder) IncrCheckCount(pk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrCheckCount", reflect.TypeOf((*MockGroupAlterEventService)(nil).IncrCheckCount), pk)
}

// ListPKLessThanCheckCountBeforeCreateAt mocks base method.
func (m *MockGroupAlterEventService) ListPKLessThanCheckCountBeforeCreateAt(CheckCount, createdAt int64) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPKLessThanCheckCountBeforeCreateAt", CheckCount, createdAt)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPKLessThanCheckCountBeforeCreateAt indicates an expected call of ListPKLessThanCheckCountBeforeCreateAt.
func (mr *MockGroupAlterEventServiceMockRecorder) ListPKLessThanCheckCountBeforeCreateAt(CheckCount, createdAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPKLessThanCheckCountBeforeCreateAt", reflect.TypeOf((*MockGroupAlterEventService)(nil).ListPKLessThanCheckCountBeforeCreateAt), CheckCount, createdAt)
}
