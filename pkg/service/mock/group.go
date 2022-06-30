// Code generated by MockGen. DO NOT EDIT.
// Source: group.go

// Package mock is a generated GoMock package.
package mock

import (
	types "iam/pkg/service/types"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockGroupService is a mock of GroupService interface.
type MockGroupService struct {
	ctrl     *gomock.Controller
	recorder *MockGroupServiceMockRecorder
}

// MockGroupServiceMockRecorder is the mock recorder for MockGroupService.
type MockGroupServiceMockRecorder struct {
	mock *MockGroupService
}

// NewMockGroupService creates a new mock instance.
func NewMockGroupService(ctrl *gomock.Controller) *MockGroupService {
	mock := &MockGroupService{ctrl: ctrl}
	mock.recorder = &MockGroupServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupService) EXPECT() *MockGroupServiceMockRecorder {
	return m.recorder
}

// AlterGroupAuthType mocks base method.
func (m *MockGroupService) AlterGroupAuthType(tx *sqlx.Tx, systemID string, groupPK, authType int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlterGroupAuthType", tx, systemID, groupPK, authType)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AlterGroupAuthType indicates an expected call of AlterGroupAuthType.
func (mr *MockGroupServiceMockRecorder) AlterGroupAuthType(tx, systemID, groupPK, authType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlterGroupAuthType", reflect.TypeOf((*MockGroupService)(nil).AlterGroupAuthType), tx, systemID, groupPK, authType)
}

// BulkCreateGroupMembersWithTx mocks base method.
func (m *MockGroupService) BulkCreateGroupMembersWithTx(tx *sqlx.Tx, groupPK int64, relations []types.SubjectRelationForCreate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreateGroupMembersWithTx", tx, groupPK, relations)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkCreateGroupMembersWithTx indicates an expected call of BulkCreateGroupMembersWithTx.
func (mr *MockGroupServiceMockRecorder) BulkCreateGroupMembersWithTx(tx, groupPK, relations interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreateGroupMembersWithTx", reflect.TypeOf((*MockGroupService)(nil).BulkCreateGroupMembersWithTx), tx, groupPK, relations)
}

// BulkDeleteBySubjectPKsWithTx mocks base method.
func (m *MockGroupService) BulkDeleteBySubjectPKsWithTx(tx *sqlx.Tx, pks []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteBySubjectPKsWithTx", tx, pks)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDeleteBySubjectPKsWithTx indicates an expected call of BulkDeleteBySubjectPKsWithTx.
func (mr *MockGroupServiceMockRecorder) BulkDeleteBySubjectPKsWithTx(tx, pks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteBySubjectPKsWithTx", reflect.TypeOf((*MockGroupService)(nil).BulkDeleteBySubjectPKsWithTx), tx, pks)
}

// BulkDeleteGroupMembers mocks base method.
func (m *MockGroupService) BulkDeleteGroupMembers(groupPK int64, userPKs, departmentPKs []int64) (map[string]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteGroupMembers", groupPK, userPKs, departmentPKs)
	ret0, _ := ret[0].(map[string]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkDeleteGroupMembers indicates an expected call of BulkDeleteGroupMembers.
func (mr *MockGroupServiceMockRecorder) BulkDeleteGroupMembers(groupPK, userPKs, departmentPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteGroupMembers", reflect.TypeOf((*MockGroupService)(nil).BulkDeleteGroupMembers), groupPK, userPKs, departmentPKs)
}

// FilterExistEffectSubjectGroupPKs mocks base method.
func (m *MockGroupService) FilterExistEffectSubjectGroupPKs(subjectPKs, groupPKs []int64) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterExistEffectSubjectGroupPKs", subjectPKs, groupPKs)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterExistEffectSubjectGroupPKs indicates an expected call of FilterExistEffectSubjectGroupPKs.
func (mr *MockGroupServiceMockRecorder) FilterExistEffectSubjectGroupPKs(subjectPKs, groupPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterExistEffectSubjectGroupPKs", reflect.TypeOf((*MockGroupService)(nil).FilterExistEffectSubjectGroupPKs), subjectPKs, groupPKs)
}

// FilterGroupPKsHasMemberBeforeExpiredAt mocks base method.
func (m *MockGroupService) FilterGroupPKsHasMemberBeforeExpiredAt(groupPKs []int64, expiredAt int64) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterGroupPKsHasMemberBeforeExpiredAt", groupPKs, expiredAt)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterGroupPKsHasMemberBeforeExpiredAt indicates an expected call of FilterGroupPKsHasMemberBeforeExpiredAt.
func (mr *MockGroupServiceMockRecorder) FilterGroupPKsHasMemberBeforeExpiredAt(groupPKs, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterGroupPKsHasMemberBeforeExpiredAt", reflect.TypeOf((*MockGroupService)(nil).FilterGroupPKsHasMemberBeforeExpiredAt), groupPKs, expiredAt)
}

// GetMemberCount mocks base method.
func (m *MockGroupService) GetMemberCount(groupPK int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMemberCount", groupPK)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMemberCount indicates an expected call of GetMemberCount.
func (mr *MockGroupServiceMockRecorder) GetMemberCount(groupPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMemberCount", reflect.TypeOf((*MockGroupService)(nil).GetMemberCount), groupPK)
}

// GetMemberCountBeforeExpiredAt mocks base method.
func (m *MockGroupService) GetMemberCountBeforeExpiredAt(groupPK, expiredAt int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMemberCountBeforeExpiredAt", groupPK, expiredAt)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMemberCountBeforeExpiredAt indicates an expected call of GetMemberCountBeforeExpiredAt.
func (mr *MockGroupServiceMockRecorder) GetMemberCountBeforeExpiredAt(groupPK, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMemberCountBeforeExpiredAt", reflect.TypeOf((*MockGroupService)(nil).GetMemberCountBeforeExpiredAt), groupPK, expiredAt)
}

// ListEffectThinSubjectGroups mocks base method.
func (m *MockGroupService) ListEffectThinSubjectGroups(systemID string, pks []int64) (map[int64][]types.ThinSubjectGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEffectThinSubjectGroups", systemID, pks)
	ret0, _ := ret[0].(map[int64][]types.ThinSubjectGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEffectThinSubjectGroups indicates an expected call of ListEffectThinSubjectGroups.
func (mr *MockGroupServiceMockRecorder) ListEffectThinSubjectGroups(systemID, pks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEffectThinSubjectGroups", reflect.TypeOf((*MockGroupService)(nil).ListEffectThinSubjectGroups), systemID, pks)
}

// ListEffectThinSubjectGroupsBySubjectPKs mocks base method.
func (m *MockGroupService) ListEffectThinSubjectGroupsBySubjectPKs(pks []int64) ([]types.ThinSubjectGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEffectThinSubjectGroupsBySubjectPKs", pks)
	ret0, _ := ret[0].([]types.ThinSubjectGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEffectThinSubjectGroupsBySubjectPKs indicates an expected call of ListEffectThinSubjectGroupsBySubjectPKs.
func (mr *MockGroupServiceMockRecorder) ListEffectThinSubjectGroupsBySubjectPKs(pks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEffectThinSubjectGroupsBySubjectPKs", reflect.TypeOf((*MockGroupService)(nil).ListEffectThinSubjectGroupsBySubjectPKs), pks)
}

// ListGroupAuthBySystemGroupPKs mocks base method.
func (m *MockGroupService) ListGroupAuthBySystemGroupPKs(systemID string, groupPKs []int64) ([]types.GroupAuthType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroupAuthBySystemGroupPKs", systemID, groupPKs)
	ret0, _ := ret[0].([]types.GroupAuthType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroupAuthBySystemGroupPKs indicates an expected call of ListGroupAuthBySystemGroupPKs.
func (mr *MockGroupServiceMockRecorder) ListGroupAuthBySystemGroupPKs(systemID, groupPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupAuthBySystemGroupPKs", reflect.TypeOf((*MockGroupService)(nil).ListGroupAuthBySystemGroupPKs), systemID, groupPKs)
}

// ListGroupAuthSystemIDs mocks base method.
func (m *MockGroupService) ListGroupAuthSystemIDs(groupPK int64) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroupAuthSystemIDs", groupPK)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroupAuthSystemIDs indicates an expected call of ListGroupAuthSystemIDs.
func (mr *MockGroupServiceMockRecorder) ListGroupAuthSystemIDs(groupPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupAuthSystemIDs", reflect.TypeOf((*MockGroupService)(nil).ListGroupAuthSystemIDs), groupPK)
}

// ListMember mocks base method.
func (m *MockGroupService) ListMember(groupPK int64) ([]types.GroupMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMember", groupPK)
	ret0, _ := ret[0].([]types.GroupMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMember indicates an expected call of ListMember.
func (mr *MockGroupServiceMockRecorder) ListMember(groupPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMember", reflect.TypeOf((*MockGroupService)(nil).ListMember), groupPK)
}

// ListPagingMember mocks base method.
func (m *MockGroupService) ListPagingMember(groupPK, limit, offset int64) ([]types.GroupMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingMember", groupPK, limit, offset)
	ret0, _ := ret[0].([]types.GroupMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingMember indicates an expected call of ListPagingMember.
func (mr *MockGroupServiceMockRecorder) ListPagingMember(groupPK, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingMember", reflect.TypeOf((*MockGroupService)(nil).ListPagingMember), groupPK, limit, offset)
}

// ListPagingMemberBeforeExpiredAt mocks base method.
func (m *MockGroupService) ListPagingMemberBeforeExpiredAt(groupPK, expiredAt, limit, offset int64) ([]types.GroupMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingMemberBeforeExpiredAt", groupPK, expiredAt, limit, offset)
	ret0, _ := ret[0].([]types.GroupMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingMemberBeforeExpiredAt indicates an expected call of ListPagingMemberBeforeExpiredAt.
func (mr *MockGroupServiceMockRecorder) ListPagingMemberBeforeExpiredAt(groupPK, expiredAt, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingMemberBeforeExpiredAt", reflect.TypeOf((*MockGroupService)(nil).ListPagingMemberBeforeExpiredAt), groupPK, expiredAt, limit, offset)
}

// ListSubjectGroups mocks base method.
func (m *MockGroupService) ListSubjectGroups(subjectPK, beforeExpiredAt int64) ([]types.SubjectGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSubjectGroups", subjectPK, beforeExpiredAt)
	ret0, _ := ret[0].([]types.SubjectGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubjectGroups indicates an expected call of ListSubjectGroups.
func (mr *MockGroupServiceMockRecorder) ListSubjectGroups(subjectPK, beforeExpiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubjectGroups", reflect.TypeOf((*MockGroupService)(nil).ListSubjectGroups), subjectPK, beforeExpiredAt)
}

// UpdateMembersExpiredAtWithTx mocks base method.
func (m *MockGroupService) UpdateMembersExpiredAtWithTx(tx *sqlx.Tx, groupPK int64, members []types.SubjectRelationForUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMembersExpiredAtWithTx", tx, groupPK, members)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMembersExpiredAtWithTx indicates an expected call of UpdateMembersExpiredAtWithTx.
func (mr *MockGroupServiceMockRecorder) UpdateMembersExpiredAtWithTx(tx, groupPK, members interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMembersExpiredAtWithTx", reflect.TypeOf((*MockGroupService)(nil).UpdateMembersExpiredAtWithTx), tx, groupPK, members)
}
