// Code generated by MockGen. DO NOT EDIT.
// Source: group.go

// Package mock is a generated GoMock package.
package mock

import (
	pap "iam/pkg/abac/pap"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGroupController is a mock of GroupController interface.
type MockGroupController struct {
	ctrl     *gomock.Controller
	recorder *MockGroupControllerMockRecorder
}

// MockGroupControllerMockRecorder is the mock recorder for MockGroupController.
type MockGroupControllerMockRecorder struct {
	mock *MockGroupController
}

// NewMockGroupController creates a new mock instance.
func NewMockGroupController(ctrl *gomock.Controller) *MockGroupController {
	mock := &MockGroupController{ctrl: ctrl}
	mock.recorder = &MockGroupControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupController) EXPECT() *MockGroupControllerMockRecorder {
	return m.recorder
}

// CreateOrUpdateSubjectMembers mocks base method.
func (m *MockGroupController) CreateOrUpdateSubjectMembers(_type, id string, members []pap.GroupMember) (map[string]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateSubjectMembers", _type, id, members)
	ret0, _ := ret[0].(map[string]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdateSubjectMembers indicates an expected call of CreateOrUpdateSubjectMembers.
func (mr *MockGroupControllerMockRecorder) CreateOrUpdateSubjectMembers(_type, id, members interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateSubjectMembers", reflect.TypeOf((*MockGroupController)(nil).CreateOrUpdateSubjectMembers), _type, id, members)
}

// DeleteSubjectMembers mocks base method.
func (m *MockGroupController) DeleteSubjectMembers(_type, id string, members []pap.Subject) (map[string]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSubjectMembers", _type, id, members)
	ret0, _ := ret[0].(map[string]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSubjectMembers indicates an expected call of DeleteSubjectMembers.
func (mr *MockGroupControllerMockRecorder) DeleteSubjectMembers(_type, id, members interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubjectMembers", reflect.TypeOf((*MockGroupController)(nil).DeleteSubjectMembers), _type, id, members)
}

// GetMemberCount mocks base method.
func (m *MockGroupController) GetMemberCount(_type, id string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMemberCount", _type, id)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMemberCount indicates an expected call of GetMemberCount.
func (mr *MockGroupControllerMockRecorder) GetMemberCount(_type, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMemberCount", reflect.TypeOf((*MockGroupController)(nil).GetMemberCount), _type, id)
}

// GetMemberCountBeforeExpiredAt mocks base method.
func (m *MockGroupController) GetMemberCountBeforeExpiredAt(_type, id string, expiredAt int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMemberCountBeforeExpiredAt", _type, id, expiredAt)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMemberCountBeforeExpiredAt indicates an expected call of GetMemberCountBeforeExpiredAt.
func (mr *MockGroupControllerMockRecorder) GetMemberCountBeforeExpiredAt(_type, id, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMemberCountBeforeExpiredAt", reflect.TypeOf((*MockGroupController)(nil).GetMemberCountBeforeExpiredAt), _type, id, expiredAt)
}

// ListExistSubjectsBeforeExpiredAt mocks base method.
func (m *MockGroupController) ListExistSubjectsBeforeExpiredAt(subjects []pap.Subject, expiredAt int64) ([]pap.Subject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListExistSubjectsBeforeExpiredAt", subjects, expiredAt)
	ret0, _ := ret[0].([]pap.Subject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListExistSubjectsBeforeExpiredAt indicates an expected call of ListExistSubjectsBeforeExpiredAt.
func (mr *MockGroupControllerMockRecorder) ListExistSubjectsBeforeExpiredAt(subjects, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListExistSubjectsBeforeExpiredAt", reflect.TypeOf((*MockGroupController)(nil).ListExistSubjectsBeforeExpiredAt), subjects, expiredAt)
}

// ListPagingMember mocks base method.
func (m *MockGroupController) ListPagingMember(_type, id string, limit, offset int64) ([]pap.GroupMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingMember", _type, id, limit, offset)
	ret0, _ := ret[0].([]pap.GroupMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingMember indicates an expected call of ListPagingMember.
func (mr *MockGroupControllerMockRecorder) ListPagingMember(_type, id, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingMember", reflect.TypeOf((*MockGroupController)(nil).ListPagingMember), _type, id, limit, offset)
}

// ListPagingMemberBeforeExpiredAt mocks base method.
func (m *MockGroupController) ListPagingMemberBeforeExpiredAt(_type, id string, expiredAt, limit, offset int64) ([]pap.GroupMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingMemberBeforeExpiredAt", _type, id, expiredAt, limit, offset)
	ret0, _ := ret[0].([]pap.GroupMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingMemberBeforeExpiredAt indicates an expected call of ListPagingMemberBeforeExpiredAt.
func (mr *MockGroupControllerMockRecorder) ListPagingMemberBeforeExpiredAt(_type, id, expiredAt, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingMemberBeforeExpiredAt", reflect.TypeOf((*MockGroupController)(nil).ListPagingMemberBeforeExpiredAt), _type, id, expiredAt, limit, offset)
}

// ListSubjectGroups mocks base method.
func (m *MockGroupController) ListSubjectGroups(_type, id string, beforeExpiredAt int64) ([]pap.SubjectGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSubjectGroups", _type, id, beforeExpiredAt)
	ret0, _ := ret[0].([]pap.SubjectGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubjectGroups indicates an expected call of ListSubjectGroups.
func (mr *MockGroupControllerMockRecorder) ListSubjectGroups(_type, id, beforeExpiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubjectGroups", reflect.TypeOf((*MockGroupController)(nil).ListSubjectGroups), _type, id, beforeExpiredAt)
}

// UpdateSubjectMembersExpiredAt mocks base method.
func (m *MockGroupController) UpdateSubjectMembersExpiredAt(_type, id string, members []pap.GroupMember) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSubjectMembersExpiredAt", _type, id, members)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSubjectMembersExpiredAt indicates an expected call of UpdateSubjectMembersExpiredAt.
func (mr *MockGroupControllerMockRecorder) UpdateSubjectMembersExpiredAt(_type, id, members interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSubjectMembersExpiredAt", reflect.TypeOf((*MockGroupController)(nil).UpdateSubjectMembersExpiredAt), _type, id, members)
}
