// Code generated by MockGen. DO NOT EDIT.
// Source: subject_group.go

// Package mock is a generated GoMock package.
package mock

import (
	dao "iam/pkg/database/dao"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockSubjectGroupManager is a mock of SubjectGroupManager interface.
type MockSubjectGroupManager struct {
	ctrl     *gomock.Controller
	recorder *MockSubjectGroupManagerMockRecorder
}

// MockSubjectGroupManagerMockRecorder is the mock recorder for MockSubjectGroupManager.
type MockSubjectGroupManagerMockRecorder struct {
	mock *MockSubjectGroupManager
}

// NewMockSubjectGroupManager creates a new mock instance.
func NewMockSubjectGroupManager(ctrl *gomock.Controller) *MockSubjectGroupManager {
	mock := &MockSubjectGroupManager{ctrl: ctrl}
	mock.recorder = &MockSubjectGroupManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubjectGroupManager) EXPECT() *MockSubjectGroupManagerMockRecorder {
	return m.recorder
}

// BulkCreateWithTx mocks base method.
func (m *MockSubjectGroupManager) BulkCreateWithTx(tx *sqlx.Tx, relations []dao.SubjectRelation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreateWithTx", tx, relations)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkCreateWithTx indicates an expected call of BulkCreateWithTx.
func (mr *MockSubjectGroupManagerMockRecorder) BulkCreateWithTx(tx, relations interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreateWithTx", reflect.TypeOf((*MockSubjectGroupManager)(nil).BulkCreateWithTx), tx, relations)
}

// BulkDeleteByGroupMembersWithTx mocks base method.
func (m *MockSubjectGroupManager) BulkDeleteByGroupMembersWithTx(tx *sqlx.Tx, groupPK int64, subjectPKs []int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteByGroupMembersWithTx", tx, groupPK, subjectPKs)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkDeleteByGroupMembersWithTx indicates an expected call of BulkDeleteByGroupMembersWithTx.
func (mr *MockSubjectGroupManagerMockRecorder) BulkDeleteByGroupMembersWithTx(tx, groupPK, subjectPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteByGroupMembersWithTx", reflect.TypeOf((*MockSubjectGroupManager)(nil).BulkDeleteByGroupMembersWithTx), tx, groupPK, subjectPKs)
}

// BulkDeleteByGroupPKs mocks base method.
func (m *MockSubjectGroupManager) BulkDeleteByGroupPKs(tx *sqlx.Tx, groupPKs []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteByGroupPKs", tx, groupPKs)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDeleteByGroupPKs indicates an expected call of BulkDeleteByGroupPKs.
func (mr *MockSubjectGroupManagerMockRecorder) BulkDeleteByGroupPKs(tx, groupPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteByGroupPKs", reflect.TypeOf((*MockSubjectGroupManager)(nil).BulkDeleteByGroupPKs), tx, groupPKs)
}

// BulkDeleteBySubjectPKs mocks base method.
func (m *MockSubjectGroupManager) BulkDeleteBySubjectPKs(tx *sqlx.Tx, subjectPKs []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteBySubjectPKs", tx, subjectPKs)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDeleteBySubjectPKs indicates an expected call of BulkDeleteBySubjectPKs.
func (mr *MockSubjectGroupManagerMockRecorder) BulkDeleteBySubjectPKs(tx, subjectPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteBySubjectPKs", reflect.TypeOf((*MockSubjectGroupManager)(nil).BulkDeleteBySubjectPKs), tx, subjectPKs)
}

// FilterGroupPKsHasMemberBeforeExpiredAt mocks base method.
func (m *MockSubjectGroupManager) FilterGroupPKsHasMemberBeforeExpiredAt(groupPKs []int64, expiredAt int64) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterGroupPKsHasMemberBeforeExpiredAt", groupPKs, expiredAt)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterGroupPKsHasMemberBeforeExpiredAt indicates an expected call of FilterGroupPKsHasMemberBeforeExpiredAt.
func (mr *MockSubjectGroupManagerMockRecorder) FilterGroupPKsHasMemberBeforeExpiredAt(groupPKs, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterGroupPKsHasMemberBeforeExpiredAt", reflect.TypeOf((*MockSubjectGroupManager)(nil).FilterGroupPKsHasMemberBeforeExpiredAt), groupPKs, expiredAt)
}

// GetExpiredAtBySubjectGroup mocks base method.
func (m *MockSubjectGroupManager) GetExpiredAtBySubjectGroup(subjectPK, groupPK int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExpiredAtBySubjectGroup", subjectPK, groupPK)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExpiredAtBySubjectGroup indicates an expected call of GetExpiredAtBySubjectGroup.
func (mr *MockSubjectGroupManagerMockRecorder) GetExpiredAtBySubjectGroup(subjectPK, groupPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExpiredAtBySubjectGroup", reflect.TypeOf((*MockSubjectGroupManager)(nil).GetExpiredAtBySubjectGroup), subjectPK, groupPK)
}

// GetGroupMemberCount mocks base method.
func (m *MockSubjectGroupManager) GetGroupMemberCount(groupPK int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupMemberCount", groupPK)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupMemberCount indicates an expected call of GetGroupMemberCount.
func (mr *MockSubjectGroupManagerMockRecorder) GetGroupMemberCount(groupPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupMemberCount", reflect.TypeOf((*MockSubjectGroupManager)(nil).GetGroupMemberCount), groupPK)
}

// GetGroupMemberCountBeforeExpiredAt mocks base method.
func (m *MockSubjectGroupManager) GetGroupMemberCountBeforeExpiredAt(groupPK, expiredAt int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupMemberCountBeforeExpiredAt", groupPK, expiredAt)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupMemberCountBeforeExpiredAt indicates an expected call of GetGroupMemberCountBeforeExpiredAt.
func (mr *MockSubjectGroupManagerMockRecorder) GetGroupMemberCountBeforeExpiredAt(groupPK, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupMemberCountBeforeExpiredAt", reflect.TypeOf((*MockSubjectGroupManager)(nil).GetGroupMemberCountBeforeExpiredAt), groupPK, expiredAt)
}

// GetGroupSubjectCountBeforeExpiredAt mocks base method.
func (m *MockSubjectGroupManager) GetGroupSubjectCountBeforeExpiredAt(expiredAt int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupSubjectCountBeforeExpiredAt", expiredAt)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupSubjectCountBeforeExpiredAt indicates an expected call of GetGroupSubjectCountBeforeExpiredAt.
func (mr *MockSubjectGroupManagerMockRecorder) GetGroupSubjectCountBeforeExpiredAt(expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupSubjectCountBeforeExpiredAt", reflect.TypeOf((*MockSubjectGroupManager)(nil).GetGroupSubjectCountBeforeExpiredAt), expiredAt)
}

// GetSubjectGroupCount mocks base method.
func (m *MockSubjectGroupManager) GetSubjectGroupCount(subjectPK int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubjectGroupCount", subjectPK)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubjectGroupCount indicates an expected call of GetSubjectGroupCount.
func (mr *MockSubjectGroupManagerMockRecorder) GetSubjectGroupCount(subjectPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubjectGroupCount", reflect.TypeOf((*MockSubjectGroupManager)(nil).GetSubjectGroupCount), subjectPK)
}

// GetSubjectGroupCountBeforeExpiredAt mocks base method.
func (m *MockSubjectGroupManager) GetSubjectGroupCountBeforeExpiredAt(subjectPK, expiredAt int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubjectGroupCountBeforeExpiredAt", subjectPK, expiredAt)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubjectGroupCountBeforeExpiredAt indicates an expected call of GetSubjectGroupCountBeforeExpiredAt.
func (mr *MockSubjectGroupManagerMockRecorder) GetSubjectGroupCountBeforeExpiredAt(subjectPK, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubjectGroupCountBeforeExpiredAt", reflect.TypeOf((*MockSubjectGroupManager)(nil).GetSubjectGroupCountBeforeExpiredAt), subjectPK, expiredAt)
}

// GetSubjectSystemGroupCount mocks base method.
func (m *MockSubjectGroupManager) GetSubjectSystemGroupCount(subjectPK int64, systemID string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubjectSystemGroupCount", subjectPK, systemID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubjectSystemGroupCount indicates an expected call of GetSubjectSystemGroupCount.
func (mr *MockSubjectGroupManagerMockRecorder) GetSubjectSystemGroupCount(subjectPK, systemID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubjectSystemGroupCount", reflect.TypeOf((*MockSubjectGroupManager)(nil).GetSubjectSystemGroupCount), subjectPK, systemID)
}

// GetSubjectSystemGroupCountBeforeExpiredAt mocks base method.
func (m *MockSubjectGroupManager) GetSubjectSystemGroupCountBeforeExpiredAt(subjectPK int64, systemID string, expiredAt int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubjectSystemGroupCountBeforeExpiredAt", subjectPK, systemID, expiredAt)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubjectSystemGroupCountBeforeExpiredAt indicates an expected call of GetSubjectSystemGroupCountBeforeExpiredAt.
func (mr *MockSubjectGroupManagerMockRecorder) GetSubjectSystemGroupCountBeforeExpiredAt(subjectPK, systemID, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubjectSystemGroupCountBeforeExpiredAt", reflect.TypeOf((*MockSubjectGroupManager)(nil).GetSubjectSystemGroupCountBeforeExpiredAt), subjectPK, systemID, expiredAt)
}

// ListGroupMember mocks base method.
func (m *MockSubjectGroupManager) ListGroupMember(groupPK int64) ([]dao.SubjectRelation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroupMember", groupPK)
	ret0, _ := ret[0].([]dao.SubjectRelation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroupMember indicates an expected call of ListGroupMember.
func (mr *MockSubjectGroupManagerMockRecorder) ListGroupMember(groupPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupMember", reflect.TypeOf((*MockSubjectGroupManager)(nil).ListGroupMember), groupPK)
}

// ListPagingGroupMember mocks base method.
func (m *MockSubjectGroupManager) ListPagingGroupMember(groupPK, limit, offset int64) ([]dao.SubjectRelation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingGroupMember", groupPK, limit, offset)
	ret0, _ := ret[0].([]dao.SubjectRelation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingGroupMember indicates an expected call of ListPagingGroupMember.
func (mr *MockSubjectGroupManagerMockRecorder) ListPagingGroupMember(groupPK, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingGroupMember", reflect.TypeOf((*MockSubjectGroupManager)(nil).ListPagingGroupMember), groupPK, limit, offset)
}

// ListPagingGroupMemberBeforeExpiredAt mocks base method.
func (m *MockSubjectGroupManager) ListPagingGroupMemberBeforeExpiredAt(groupPK, expiredAt, limit, offset int64) ([]dao.SubjectRelation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingGroupMemberBeforeExpiredAt", groupPK, expiredAt, limit, offset)
	ret0, _ := ret[0].([]dao.SubjectRelation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingGroupMemberBeforeExpiredAt indicates an expected call of ListPagingGroupMemberBeforeExpiredAt.
func (mr *MockSubjectGroupManagerMockRecorder) ListPagingGroupMemberBeforeExpiredAt(groupPK, expiredAt, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingGroupMemberBeforeExpiredAt", reflect.TypeOf((*MockSubjectGroupManager)(nil).ListPagingGroupMemberBeforeExpiredAt), groupPK, expiredAt, limit, offset)
}

// ListPagingGroupSubjectBeforeExpiredAt mocks base method.
func (m *MockSubjectGroupManager) ListPagingGroupSubjectBeforeExpiredAt(expiredAt, limit, offset int64) ([]dao.ThinSubjectRelation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingGroupSubjectBeforeExpiredAt", expiredAt, limit, offset)
	ret0, _ := ret[0].([]dao.ThinSubjectRelation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingGroupSubjectBeforeExpiredAt indicates an expected call of ListPagingGroupSubjectBeforeExpiredAt.
func (mr *MockSubjectGroupManagerMockRecorder) ListPagingGroupSubjectBeforeExpiredAt(expiredAt, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingGroupSubjectBeforeExpiredAt", reflect.TypeOf((*MockSubjectGroupManager)(nil).ListPagingGroupSubjectBeforeExpiredAt), expiredAt, limit, offset)
}

// ListPagingSubjectGroupBeforeExpiredAt mocks base method.
func (m *MockSubjectGroupManager) ListPagingSubjectGroupBeforeExpiredAt(subjectPK, expiredAt, limit, offset int64) ([]dao.SubjectRelation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingSubjectGroupBeforeExpiredAt", subjectPK, expiredAt, limit, offset)
	ret0, _ := ret[0].([]dao.SubjectRelation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingSubjectGroupBeforeExpiredAt indicates an expected call of ListPagingSubjectGroupBeforeExpiredAt.
func (mr *MockSubjectGroupManagerMockRecorder) ListPagingSubjectGroupBeforeExpiredAt(subjectPK, expiredAt, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingSubjectGroupBeforeExpiredAt", reflect.TypeOf((*MockSubjectGroupManager)(nil).ListPagingSubjectGroupBeforeExpiredAt), subjectPK, expiredAt, limit, offset)
}

// ListPagingSubjectGroups mocks base method.
func (m *MockSubjectGroupManager) ListPagingSubjectGroups(subjectPK, limit, offset int64) ([]dao.SubjectRelation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingSubjectGroups", subjectPK, limit, offset)
	ret0, _ := ret[0].([]dao.SubjectRelation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingSubjectGroups indicates an expected call of ListPagingSubjectGroups.
func (mr *MockSubjectGroupManagerMockRecorder) ListPagingSubjectGroups(subjectPK, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingSubjectGroups", reflect.TypeOf((*MockSubjectGroupManager)(nil).ListPagingSubjectGroups), subjectPK, limit, offset)
}

// ListPagingSubjectSystemGroupBeforeExpiredAt mocks base method.
func (m *MockSubjectGroupManager) ListPagingSubjectSystemGroupBeforeExpiredAt(subjectPK int64, systemID string, expiredAt, limit, offset int64) ([]dao.SubjectRelation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingSubjectSystemGroupBeforeExpiredAt", subjectPK, systemID, expiredAt, limit, offset)
	ret0, _ := ret[0].([]dao.SubjectRelation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingSubjectSystemGroupBeforeExpiredAt indicates an expected call of ListPagingSubjectSystemGroupBeforeExpiredAt.
func (mr *MockSubjectGroupManagerMockRecorder) ListPagingSubjectSystemGroupBeforeExpiredAt(subjectPK, systemID, expiredAt, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingSubjectSystemGroupBeforeExpiredAt", reflect.TypeOf((*MockSubjectGroupManager)(nil).ListPagingSubjectSystemGroupBeforeExpiredAt), subjectPK, systemID, expiredAt, limit, offset)
}

// ListPagingSubjectSystemGroups mocks base method.
func (m *MockSubjectGroupManager) ListPagingSubjectSystemGroups(subjectPK int64, systemID string, limit, offset int64) ([]dao.SubjectRelation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingSubjectSystemGroups", subjectPK, systemID, limit, offset)
	ret0, _ := ret[0].([]dao.SubjectRelation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingSubjectSystemGroups indicates an expected call of ListPagingSubjectSystemGroups.
func (mr *MockSubjectGroupManagerMockRecorder) ListPagingSubjectSystemGroups(subjectPK, systemID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingSubjectSystemGroups", reflect.TypeOf((*MockSubjectGroupManager)(nil).ListPagingSubjectSystemGroups), subjectPK, systemID, limit, offset)
}

// ListThinRelationAfterExpiredAtBySubjectPKs mocks base method.
func (m *MockSubjectGroupManager) ListThinRelationAfterExpiredAtBySubjectPKs(subjectPKs []int64, expiredAt int64) ([]dao.ThinSubjectRelation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListThinRelationAfterExpiredAtBySubjectPKs", subjectPKs, expiredAt)
	ret0, _ := ret[0].([]dao.ThinSubjectRelation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListThinRelationAfterExpiredAtBySubjectPKs indicates an expected call of ListThinRelationAfterExpiredAtBySubjectPKs.
func (mr *MockSubjectGroupManagerMockRecorder) ListThinRelationAfterExpiredAtBySubjectPKs(subjectPKs, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListThinRelationAfterExpiredAtBySubjectPKs", reflect.TypeOf((*MockSubjectGroupManager)(nil).ListThinRelationAfterExpiredAtBySubjectPKs), subjectPKs, expiredAt)
}

// ListThinRelationBySubjectPKGroupPKs mocks base method.
func (m *MockSubjectGroupManager) ListThinRelationBySubjectPKGroupPKs(subjectPK int64, groupPKs []int64) ([]dao.ThinSubjectRelation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListThinRelationBySubjectPKGroupPKs", subjectPK, groupPKs)
	ret0, _ := ret[0].([]dao.ThinSubjectRelation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListThinRelationBySubjectPKGroupPKs indicates an expected call of ListThinRelationBySubjectPKGroupPKs.
func (mr *MockSubjectGroupManagerMockRecorder) ListThinRelationBySubjectPKGroupPKs(subjectPK, groupPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListThinRelationBySubjectPKGroupPKs", reflect.TypeOf((*MockSubjectGroupManager)(nil).ListThinRelationBySubjectPKGroupPKs), subjectPK, groupPKs)
}

// UpdateExpiredAtWithTx mocks base method.
func (m *MockSubjectGroupManager) UpdateExpiredAtWithTx(tx *sqlx.Tx, relations []dao.SubjectRelationForUpdateExpiredAt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExpiredAtWithTx", tx, relations)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExpiredAtWithTx indicates an expected call of UpdateExpiredAtWithTx.
func (mr *MockSubjectGroupManagerMockRecorder) UpdateExpiredAtWithTx(tx, relations interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExpiredAtWithTx", reflect.TypeOf((*MockSubjectGroupManager)(nil).UpdateExpiredAtWithTx), tx, relations)
}
