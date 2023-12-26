// Code generated by MockGen. DO NOT EDIT.
// Source: group.go

// Package mock is a generated GoMock package.
package mock

import (
	pap "iam/pkg/abac/pap"
	types "iam/pkg/abac/types"
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

// BulkCreateSubjectTemplateGroup mocks base method.
func (m *MockGroupController) BulkCreateSubjectTemplateGroup(subjectTemplateGroups []pap.SubjectTemplateGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreateSubjectTemplateGroup", subjectTemplateGroups)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkCreateSubjectTemplateGroup indicates an expected call of BulkCreateSubjectTemplateGroup.
func (mr *MockGroupControllerMockRecorder) BulkCreateSubjectTemplateGroup(subjectTemplateGroups interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreateSubjectTemplateGroup", reflect.TypeOf((*MockGroupController)(nil).BulkCreateSubjectTemplateGroup), subjectTemplateGroups)
}

// BulkDeleteSubjectTemplateGroup mocks base method.
func (m *MockGroupController) BulkDeleteSubjectTemplateGroup(subjectTemplateGroups []pap.SubjectTemplateGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteSubjectTemplateGroup", subjectTemplateGroups)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDeleteSubjectTemplateGroup indicates an expected call of BulkDeleteSubjectTemplateGroup.
func (mr *MockGroupControllerMockRecorder) BulkDeleteSubjectTemplateGroup(subjectTemplateGroups interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteSubjectTemplateGroup", reflect.TypeOf((*MockGroupController)(nil).BulkDeleteSubjectTemplateGroup), subjectTemplateGroups)
}

// CheckSubjectEffectGroups mocks base method.
func (m *MockGroupController) CheckSubjectEffectGroups(_type, id string, groupIDs []string) (map[string]map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSubjectEffectGroups", _type, id, groupIDs)
	ret0, _ := ret[0].(map[string]map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckSubjectEffectGroups indicates an expected call of CheckSubjectEffectGroups.
func (mr *MockGroupControllerMockRecorder) CheckSubjectEffectGroups(_type, id, groupIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSubjectEffectGroups", reflect.TypeOf((*MockGroupController)(nil).CheckSubjectEffectGroups), _type, id, groupIDs)
}

// CreateOrUpdateGroupMembers mocks base method.
func (m *MockGroupController) CreateOrUpdateGroupMembers(_type, id string, members []pap.GroupMember) (map[string]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateGroupMembers", _type, id, members)
	ret0, _ := ret[0].(map[string]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdateGroupMembers indicates an expected call of CreateOrUpdateGroupMembers.
func (mr *MockGroupControllerMockRecorder) CreateOrUpdateGroupMembers(_type, id, members interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateGroupMembers", reflect.TypeOf((*MockGroupController)(nil).CreateOrUpdateGroupMembers), _type, id, members)
}

// DeleteGroupMembers mocks base method.
func (m *MockGroupController) DeleteGroupMembers(_type, id string, members []pap.Subject) (map[string]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGroupMembers", _type, id, members)
	ret0, _ := ret[0].(map[string]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteGroupMembers indicates an expected call of DeleteGroupMembers.
func (mr *MockGroupControllerMockRecorder) DeleteGroupMembers(_type, id, members interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroupMembers", reflect.TypeOf((*MockGroupController)(nil).DeleteGroupMembers), _type, id, members)
}

// FilterGroupsHasMemberBeforeExpiredAt mocks base method.
func (m *MockGroupController) FilterGroupsHasMemberBeforeExpiredAt(subjects []pap.Subject, expiredAt int64) ([]pap.Subject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterGroupsHasMemberBeforeExpiredAt", subjects, expiredAt)
	ret0, _ := ret[0].([]pap.Subject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterGroupsHasMemberBeforeExpiredAt indicates an expected call of FilterGroupsHasMemberBeforeExpiredAt.
func (mr *MockGroupControllerMockRecorder) FilterGroupsHasMemberBeforeExpiredAt(subjects, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterGroupsHasMemberBeforeExpiredAt", reflect.TypeOf((*MockGroupController)(nil).FilterGroupsHasMemberBeforeExpiredAt), subjects, expiredAt)
}

// GetGroupMemberCount mocks base method.
func (m *MockGroupController) GetGroupMemberCount(_type, id string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupMemberCount", _type, id)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupMemberCount indicates an expected call of GetGroupMemberCount.
func (mr *MockGroupControllerMockRecorder) GetGroupMemberCount(_type, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupMemberCount", reflect.TypeOf((*MockGroupController)(nil).GetGroupMemberCount), _type, id)
}

// GetGroupMemberCountBeforeExpiredAt mocks base method.
func (m *MockGroupController) GetGroupMemberCountBeforeExpiredAt(_type, id string, expiredAt int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupMemberCountBeforeExpiredAt", _type, id, expiredAt)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupMemberCountBeforeExpiredAt indicates an expected call of GetGroupMemberCountBeforeExpiredAt.
func (mr *MockGroupControllerMockRecorder) GetGroupMemberCountBeforeExpiredAt(_type, id, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupMemberCountBeforeExpiredAt", reflect.TypeOf((*MockGroupController)(nil).GetGroupMemberCountBeforeExpiredAt), _type, id, expiredAt)
}

// GetGroupSubjectCountBeforeExpiredAt mocks base method.
func (m *MockGroupController) GetGroupSubjectCountBeforeExpiredAt(expiredAt int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupSubjectCountBeforeExpiredAt", expiredAt)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupSubjectCountBeforeExpiredAt indicates an expected call of GetGroupSubjectCountBeforeExpiredAt.
func (mr *MockGroupControllerMockRecorder) GetGroupSubjectCountBeforeExpiredAt(expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupSubjectCountBeforeExpiredAt", reflect.TypeOf((*MockGroupController)(nil).GetGroupSubjectCountBeforeExpiredAt), expiredAt)
}

// GetSubjectGroupCountBeforeExpiredAt mocks base method.
func (m *MockGroupController) GetSubjectGroupCountBeforeExpiredAt(_type, id string, beforeExpiredAt int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubjectGroupCountBeforeExpiredAt", _type, id, beforeExpiredAt)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubjectGroupCountBeforeExpiredAt indicates an expected call of GetSubjectGroupCountBeforeExpiredAt.
func (mr *MockGroupControllerMockRecorder) GetSubjectGroupCountBeforeExpiredAt(_type, id, beforeExpiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubjectGroupCountBeforeExpiredAt", reflect.TypeOf((*MockGroupController)(nil).GetSubjectGroupCountBeforeExpiredAt), _type, id, beforeExpiredAt)
}

// GetSubjectSystemGroupCountBeforeExpiredAt mocks base method.
func (m *MockGroupController) GetSubjectSystemGroupCountBeforeExpiredAt(_type, id, systemID string, expiredAt int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubjectSystemGroupCountBeforeExpiredAt", _type, id, systemID, expiredAt)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubjectSystemGroupCountBeforeExpiredAt indicates an expected call of GetSubjectSystemGroupCountBeforeExpiredAt.
func (mr *MockGroupControllerMockRecorder) GetSubjectSystemGroupCountBeforeExpiredAt(_type, id, systemID, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubjectSystemGroupCountBeforeExpiredAt", reflect.TypeOf((*MockGroupController)(nil).GetSubjectSystemGroupCountBeforeExpiredAt), _type, id, systemID, expiredAt)
}

// GetTemplateGroupMemberCount mocks base method.
func (m *MockGroupController) GetTemplateGroupMemberCount(_type, id string, templateID int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTemplateGroupMemberCount", _type, id, templateID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTemplateGroupMemberCount indicates an expected call of GetTemplateGroupMemberCount.
func (mr *MockGroupControllerMockRecorder) GetTemplateGroupMemberCount(_type, id, templateID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTemplateGroupMemberCount", reflect.TypeOf((*MockGroupController)(nil).GetTemplateGroupMemberCount), _type, id, templateID)
}

// ListPagingGroupMember mocks base method.
func (m *MockGroupController) ListPagingGroupMember(_type, id string, limit, offset int64) ([]pap.GroupMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingGroupMember", _type, id, limit, offset)
	ret0, _ := ret[0].([]pap.GroupMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingGroupMember indicates an expected call of ListPagingGroupMember.
func (mr *MockGroupControllerMockRecorder) ListPagingGroupMember(_type, id, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingGroupMember", reflect.TypeOf((*MockGroupController)(nil).ListPagingGroupMember), _type, id, limit, offset)
}

// ListPagingGroupMemberBeforeExpiredAt mocks base method.
func (m *MockGroupController) ListPagingGroupMemberBeforeExpiredAt(_type, id string, expiredAt, limit, offset int64) ([]pap.GroupMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingGroupMemberBeforeExpiredAt", _type, id, expiredAt, limit, offset)
	ret0, _ := ret[0].([]pap.GroupMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingGroupMemberBeforeExpiredAt indicates an expected call of ListPagingGroupMemberBeforeExpiredAt.
func (mr *MockGroupControllerMockRecorder) ListPagingGroupMemberBeforeExpiredAt(_type, id, expiredAt, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingGroupMemberBeforeExpiredAt", reflect.TypeOf((*MockGroupController)(nil).ListPagingGroupMemberBeforeExpiredAt), _type, id, expiredAt, limit, offset)
}

// ListPagingGroupSubjectBeforeExpiredAt mocks base method.
func (m *MockGroupController) ListPagingGroupSubjectBeforeExpiredAt(expiredAt, limit, offset int64) ([]pap.GroupSubject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingGroupSubjectBeforeExpiredAt", expiredAt, limit, offset)
	ret0, _ := ret[0].([]pap.GroupSubject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingGroupSubjectBeforeExpiredAt indicates an expected call of ListPagingGroupSubjectBeforeExpiredAt.
func (mr *MockGroupControllerMockRecorder) ListPagingGroupSubjectBeforeExpiredAt(expiredAt, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingGroupSubjectBeforeExpiredAt", reflect.TypeOf((*MockGroupController)(nil).ListPagingGroupSubjectBeforeExpiredAt), expiredAt, limit, offset)
}

// ListPagingSubjectGroups mocks base method.
func (m *MockGroupController) ListPagingSubjectGroups(_type, id string, beforeExpiredAt, limit, offset int64) ([]pap.SubjectGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingSubjectGroups", _type, id, beforeExpiredAt, limit, offset)
	ret0, _ := ret[0].([]pap.SubjectGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingSubjectGroups indicates an expected call of ListPagingSubjectGroups.
func (mr *MockGroupControllerMockRecorder) ListPagingSubjectGroups(_type, id, beforeExpiredAt, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingSubjectGroups", reflect.TypeOf((*MockGroupController)(nil).ListPagingSubjectGroups), _type, id, beforeExpiredAt, limit, offset)
}

// ListPagingSubjectSystemGroups mocks base method.
func (m *MockGroupController) ListPagingSubjectSystemGroups(_type, id, systemID string, beforeExpiredAt, limit, offset int64) ([]pap.SubjectGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingSubjectSystemGroups", _type, id, systemID, beforeExpiredAt, limit, offset)
	ret0, _ := ret[0].([]pap.SubjectGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingSubjectSystemGroups indicates an expected call of ListPagingSubjectSystemGroups.
func (mr *MockGroupControllerMockRecorder) ListPagingSubjectSystemGroups(_type, id, systemID, beforeExpiredAt, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingSubjectSystemGroups", reflect.TypeOf((*MockGroupController)(nil).ListPagingSubjectSystemGroups), _type, id, systemID, beforeExpiredAt, limit, offset)
}

// ListPagingTemplateGroupMember mocks base method.
func (m *MockGroupController) ListPagingTemplateGroupMember(_type, id string, templateID, limit, offset int64) ([]pap.GroupMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingTemplateGroupMember", _type, id, templateID, limit, offset)
	ret0, _ := ret[0].([]pap.GroupMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingTemplateGroupMember indicates an expected call of ListPagingTemplateGroupMember.
func (mr *MockGroupControllerMockRecorder) ListPagingTemplateGroupMember(_type, id, templateID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingTemplateGroupMember", reflect.TypeOf((*MockGroupController)(nil).ListPagingTemplateGroupMember), _type, id, templateID, limit, offset)
}

// ListRbacGroupByActionResource mocks base method.
func (m *MockGroupController) ListRbacGroupByActionResource(systemID, actionID string, resource types.Resource) ([]pap.Subject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRbacGroupByActionResource", systemID, actionID, resource)
	ret0, _ := ret[0].([]pap.Subject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRbacGroupByActionResource indicates an expected call of ListRbacGroupByActionResource.
func (mr *MockGroupControllerMockRecorder) ListRbacGroupByActionResource(systemID, actionID, resource interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRbacGroupByActionResource", reflect.TypeOf((*MockGroupController)(nil).ListRbacGroupByActionResource), systemID, actionID, resource)
}

// ListRbacGroupByResource mocks base method.
func (m *MockGroupController) ListRbacGroupByResource(systemID string, resource types.Resource) ([]pap.Subject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRbacGroupByResource", systemID, resource)
	ret0, _ := ret[0].([]pap.Subject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRbacGroupByResource indicates an expected call of ListRbacGroupByResource.
func (mr *MockGroupControllerMockRecorder) ListRbacGroupByResource(systemID, resource interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRbacGroupByResource", reflect.TypeOf((*MockGroupController)(nil).ListRbacGroupByResource), systemID, resource)
}

// UpdateGroupMembersExpiredAt mocks base method.
func (m *MockGroupController) UpdateGroupMembersExpiredAt(_type, id string, members []pap.GroupMember) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGroupMembersExpiredAt", _type, id, members)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGroupMembersExpiredAt indicates an expected call of UpdateGroupMembersExpiredAt.
func (mr *MockGroupControllerMockRecorder) UpdateGroupMembersExpiredAt(_type, id, members interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGroupMembersExpiredAt", reflect.TypeOf((*MockGroupController)(nil).UpdateGroupMembersExpiredAt), _type, id, members)
}

// UpdateSubjectTemplateGroupExpiredAt mocks base method.
func (m *MockGroupController) UpdateSubjectTemplateGroupExpiredAt(subjectTemplateGroups []pap.SubjectTemplateGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSubjectTemplateGroupExpiredAt", subjectTemplateGroups)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSubjectTemplateGroupExpiredAt indicates an expected call of UpdateSubjectTemplateGroupExpiredAt.
func (mr *MockGroupControllerMockRecorder) UpdateSubjectTemplateGroupExpiredAt(subjectTemplateGroups interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSubjectTemplateGroupExpiredAt", reflect.TypeOf((*MockGroupController)(nil).UpdateSubjectTemplateGroupExpiredAt), subjectTemplateGroups)
}
