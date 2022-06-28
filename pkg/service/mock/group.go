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
func (m *MockGroupService) BulkCreateGroupMembersWithTx(tx *sqlx.Tx, parentPK int64, relations []types.SubjectRelation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreateGroupMembersWithTx", tx, parentPK, relations)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkCreateGroupMembersWithTx indicates an expected call of BulkCreateGroupMembersWithTx.
func (mr *MockGroupServiceMockRecorder) BulkCreateGroupMembersWithTx(tx, parentPK, relations interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreateGroupMembersWithTx", reflect.TypeOf((*MockGroupService)(nil).BulkCreateGroupMembersWithTx), tx, parentPK, relations)
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
func (m *MockGroupService) BulkDeleteGroupMembers(parentPK int64, userPKs, departmentPKs []int64) (map[string]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteGroupMembers", parentPK, userPKs, departmentPKs)
	ret0, _ := ret[0].(map[string]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkDeleteGroupMembers indicates an expected call of BulkDeleteGroupMembers.
func (mr *MockGroupServiceMockRecorder) BulkDeleteGroupMembers(parentPK, userPKs, departmentPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteGroupMembers", reflect.TypeOf((*MockGroupService)(nil).BulkDeleteGroupMembers), parentPK, userPKs, departmentPKs)
}

// GetMemberCount mocks base method.
func (m *MockGroupService) GetMemberCount(parentPK int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMemberCount", parentPK)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMemberCount indicates an expected call of GetMemberCount.
func (mr *MockGroupServiceMockRecorder) GetMemberCount(parentPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMemberCount", reflect.TypeOf((*MockGroupService)(nil).GetMemberCount), parentPK)
}

// GetMemberCountBeforeExpiredAt mocks base method.
func (m *MockGroupService) GetMemberCountBeforeExpiredAt(parentPK, expiredAt int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMemberCountBeforeExpiredAt", parentPK, expiredAt)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMemberCountBeforeExpiredAt indicates an expected call of GetMemberCountBeforeExpiredAt.
func (mr *MockGroupServiceMockRecorder) GetMemberCountBeforeExpiredAt(parentPK, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMemberCountBeforeExpiredAt", reflect.TypeOf((*MockGroupService)(nil).GetMemberCountBeforeExpiredAt), parentPK, expiredAt)
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

// ListExistSubjectsBeforeExpiredAt mocks base method.
func (m *MockGroupService) ListExistSubjectsBeforeExpiredAt(parentPKs []int64, expiredAt int64) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListExistSubjectsBeforeExpiredAt", parentPKs, expiredAt)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListExistSubjectsBeforeExpiredAt indicates an expected call of ListExistSubjectsBeforeExpiredAt.
func (mr *MockGroupServiceMockRecorder) ListExistSubjectsBeforeExpiredAt(parentPKs, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListExistSubjectsBeforeExpiredAt", reflect.TypeOf((*MockGroupService)(nil).ListExistSubjectsBeforeExpiredAt), parentPKs, expiredAt)
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
func (m *MockGroupService) ListMember(parentPK int64) ([]types.GroupMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMember", parentPK)
	ret0, _ := ret[0].([]types.GroupMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMember indicates an expected call of ListMember.
func (mr *MockGroupServiceMockRecorder) ListMember(parentPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMember", reflect.TypeOf((*MockGroupService)(nil).ListMember), parentPK)
}

// ListPagingMember mocks base method.
func (m *MockGroupService) ListPagingMember(parentPK, limit, offset int64) ([]types.GroupMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingMember", parentPK, limit, offset)
	ret0, _ := ret[0].([]types.GroupMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingMember indicates an expected call of ListPagingMember.
func (mr *MockGroupServiceMockRecorder) ListPagingMember(parentPK, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingMember", reflect.TypeOf((*MockGroupService)(nil).ListPagingMember), parentPK, limit, offset)
}

// ListPagingMemberBeforeExpiredAt mocks base method.
func (m *MockGroupService) ListPagingMemberBeforeExpiredAt(parentPK, expiredAt, limit, offset int64) ([]types.GroupMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingMemberBeforeExpiredAt", parentPK, expiredAt, limit, offset)
	ret0, _ := ret[0].([]types.GroupMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingMemberBeforeExpiredAt indicates an expected call of ListPagingMemberBeforeExpiredAt.
func (mr *MockGroupServiceMockRecorder) ListPagingMemberBeforeExpiredAt(parentPK, expiredAt, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingMemberBeforeExpiredAt", reflect.TypeOf((*MockGroupService)(nil).ListPagingMemberBeforeExpiredAt), parentPK, expiredAt, limit, offset)
}

// ListSubjectAllGroupPKs mocks base method.
func (m *MockGroupService) ListSubjectAllGroupPKs(subjectPK int64) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSubjectAllGroupPKs", subjectPK)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubjectAllGroupPKs indicates an expected call of ListSubjectAllGroupPKs.
func (mr *MockGroupServiceMockRecorder) ListSubjectAllGroupPKs(subjectPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubjectAllGroupPKs", reflect.TypeOf((*MockGroupService)(nil).ListSubjectAllGroupPKs), subjectPK)
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
func (m *MockGroupService) UpdateMembersExpiredAtWithTx(tx *sqlx.Tx, parentPK int64, members []types.SubjectRelationPKPolicyExpiredAt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMembersExpiredAtWithTx", tx, parentPK, members)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMembersExpiredAtWithTx indicates an expected call of UpdateMembersExpiredAtWithTx.
func (mr *MockGroupServiceMockRecorder) UpdateMembersExpiredAtWithTx(tx, parentPK, members interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMembersExpiredAtWithTx", reflect.TypeOf((*MockGroupService)(nil).UpdateMembersExpiredAtWithTx), tx, parentPK, members)
}
