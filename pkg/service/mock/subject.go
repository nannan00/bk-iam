// Code generated by MockGen. DO NOT EDIT.
// Source: subject.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	types "iam/pkg/service/types"
	reflect "reflect"
)

// MockSubjectService is a mock of SubjectService interface
type MockSubjectService struct {
	ctrl     *gomock.Controller
	recorder *MockSubjectServiceMockRecorder
}

// MockSubjectServiceMockRecorder is the mock recorder for MockSubjectService
type MockSubjectServiceMockRecorder struct {
	mock *MockSubjectService
}

// NewMockSubjectService creates a new mock instance
func NewMockSubjectService(ctrl *gomock.Controller) *MockSubjectService {
	mock := &MockSubjectService{ctrl: ctrl}
	mock.recorder = &MockSubjectServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSubjectService) EXPECT() *MockSubjectServiceMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockSubjectService) Get(pk int64) (types.Subject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", pk)
	ret0, _ := ret[0].(types.Subject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockSubjectServiceMockRecorder) Get(pk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSubjectService)(nil).Get), pk)
}

// GetPK mocks base method
func (m *MockSubjectService) GetPK(_type, id string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPK", _type, id)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPK indicates an expected call of GetPK
func (mr *MockSubjectServiceMockRecorder) GetPK(_type, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPK", reflect.TypeOf((*MockSubjectService)(nil).GetPK), _type, id)
}

// GetCount mocks base method
func (m *MockSubjectService) GetCount(_type string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount", _type)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCount indicates an expected call of GetCount
func (mr *MockSubjectServiceMockRecorder) GetCount(_type interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockSubjectService)(nil).GetCount), _type)
}

// ListPaging mocks base method
func (m *MockSubjectService) ListPaging(_type string, limit, offset int64) ([]types.Subject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPaging", _type, limit, offset)
	ret0, _ := ret[0].([]types.Subject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPaging indicates an expected call of ListPaging
func (mr *MockSubjectServiceMockRecorder) ListPaging(_type, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPaging", reflect.TypeOf((*MockSubjectService)(nil).ListPaging), _type, limit, offset)
}

// ListPKsBySubjects mocks base method
func (m *MockSubjectService) ListPKsBySubjects(subjects []types.Subject) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPKsBySubjects", subjects)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPKsBySubjects indicates an expected call of ListPKsBySubjects
func (mr *MockSubjectServiceMockRecorder) ListPKsBySubjects(subjects interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPKsBySubjects", reflect.TypeOf((*MockSubjectService)(nil).ListPKsBySubjects), subjects)
}

// ListByPKs mocks base method
func (m *MockSubjectService) ListByPKs(pks []int64) ([]types.Subject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByPKs", pks)
	ret0, _ := ret[0].([]types.Subject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByPKs indicates an expected call of ListByPKs
func (mr *MockSubjectServiceMockRecorder) ListByPKs(pks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByPKs", reflect.TypeOf((*MockSubjectService)(nil).ListByPKs), pks)
}

// BulkCreate mocks base method
func (m *MockSubjectService) BulkCreate(subjects []types.Subject) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreate", subjects)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkCreate indicates an expected call of BulkCreate
func (mr *MockSubjectServiceMockRecorder) BulkCreate(subjects interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreate", reflect.TypeOf((*MockSubjectService)(nil).BulkCreate), subjects)
}

// BulkDelete mocks base method
func (m *MockSubjectService) BulkDelete(subjects []types.Subject) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDelete", subjects)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkDelete indicates an expected call of BulkDelete
func (mr *MockSubjectServiceMockRecorder) BulkDelete(subjects interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDelete", reflect.TypeOf((*MockSubjectService)(nil).BulkDelete), subjects)
}

// BulkUpdateName mocks base method
func (m *MockSubjectService) BulkUpdateName(subjects []types.Subject) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkUpdateName", subjects)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkUpdateName indicates an expected call of BulkUpdateName
func (mr *MockSubjectServiceMockRecorder) BulkUpdateName(subjects interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkUpdateName", reflect.TypeOf((*MockSubjectService)(nil).BulkUpdateName), subjects)
}

// GetEffectThinSubjectGroups mocks base method
func (m *MockSubjectService) GetEffectThinSubjectGroups(pk int64) ([]types.ThinSubjectGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEffectThinSubjectGroups", pk)
	ret0, _ := ret[0].([]types.ThinSubjectGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEffectThinSubjectGroups indicates an expected call of GetEffectThinSubjectGroups
func (mr *MockSubjectServiceMockRecorder) GetEffectThinSubjectGroups(pk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEffectThinSubjectGroups", reflect.TypeOf((*MockSubjectService)(nil).GetEffectThinSubjectGroups), pk)
}

// ListEffectThinSubjectGroups mocks base method
func (m *MockSubjectService) ListEffectThinSubjectGroups(pks []int64) (map[int64][]types.ThinSubjectGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEffectThinSubjectGroups", pks)
	ret0, _ := ret[0].(map[int64][]types.ThinSubjectGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEffectThinSubjectGroups indicates an expected call of ListEffectThinSubjectGroups
func (mr *MockSubjectServiceMockRecorder) ListEffectThinSubjectGroups(pks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEffectThinSubjectGroups", reflect.TypeOf((*MockSubjectService)(nil).ListEffectThinSubjectGroups), pks)
}

// ListSubjectGroups mocks base method
func (m *MockSubjectService) ListSubjectGroups(_type, id string, beforeExpiredAt int64) ([]types.SubjectGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSubjectGroups", _type, id, beforeExpiredAt)
	ret0, _ := ret[0].([]types.SubjectGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubjectGroups indicates an expected call of ListSubjectGroups
func (mr *MockSubjectServiceMockRecorder) ListSubjectGroups(_type, id, beforeExpiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubjectGroups", reflect.TypeOf((*MockSubjectService)(nil).ListSubjectGroups), _type, id, beforeExpiredAt)
}

// GetMemberCount mocks base method
func (m *MockSubjectService) GetMemberCount(_type, id string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMemberCount", _type, id)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMemberCount indicates an expected call of GetMemberCount
func (mr *MockSubjectServiceMockRecorder) GetMemberCount(_type, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMemberCount", reflect.TypeOf((*MockSubjectService)(nil).GetMemberCount), _type, id)
}

// GetMemberCountBeforeExpiredAt mocks base method
func (m *MockSubjectService) GetMemberCountBeforeExpiredAt(_type, id string, expiredAt int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMemberCountBeforeExpiredAt", _type, id, expiredAt)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMemberCountBeforeExpiredAt indicates an expected call of GetMemberCountBeforeExpiredAt
func (mr *MockSubjectServiceMockRecorder) GetMemberCountBeforeExpiredAt(_type, id, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMemberCountBeforeExpiredAt", reflect.TypeOf((*MockSubjectService)(nil).GetMemberCountBeforeExpiredAt), _type, id, expiredAt)
}

// ListPagingMember mocks base method
func (m *MockSubjectService) ListPagingMember(_type, id string, limit, offset int64) ([]types.SubjectMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingMember", _type, id, limit, offset)
	ret0, _ := ret[0].([]types.SubjectMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingMember indicates an expected call of ListPagingMember
func (mr *MockSubjectServiceMockRecorder) ListPagingMember(_type, id, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingMember", reflect.TypeOf((*MockSubjectService)(nil).ListPagingMember), _type, id, limit, offset)
}

// ListPagingMemberBeforeExpiredAt mocks base method
func (m *MockSubjectService) ListPagingMemberBeforeExpiredAt(_type, id string, expiredAt, limit, offset int64) ([]types.SubjectMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingMemberBeforeExpiredAt", _type, id, expiredAt, limit, offset)
	ret0, _ := ret[0].([]types.SubjectMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingMemberBeforeExpiredAt indicates an expected call of ListPagingMemberBeforeExpiredAt
func (mr *MockSubjectServiceMockRecorder) ListPagingMemberBeforeExpiredAt(_type, id, expiredAt, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingMemberBeforeExpiredAt", reflect.TypeOf((*MockSubjectService)(nil).ListPagingMemberBeforeExpiredAt), _type, id, expiredAt, limit, offset)
}

// ListExistSubjectsBeforeExpiredAt mocks base method
func (m *MockSubjectService) ListExistSubjectsBeforeExpiredAt(subjects []types.Subject, expiredAt int64) ([]types.Subject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListExistSubjectsBeforeExpiredAt", subjects, expiredAt)
	ret0, _ := ret[0].([]types.Subject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListExistSubjectsBeforeExpiredAt indicates an expected call of ListExistSubjectsBeforeExpiredAt
func (mr *MockSubjectServiceMockRecorder) ListExistSubjectsBeforeExpiredAt(subjects, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListExistSubjectsBeforeExpiredAt", reflect.TypeOf((*MockSubjectService)(nil).ListExistSubjectsBeforeExpiredAt), subjects, expiredAt)
}

// ListMember mocks base method
func (m *MockSubjectService) ListMember(_type, id string) ([]types.SubjectMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMember", _type, id)
	ret0, _ := ret[0].([]types.SubjectMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMember indicates an expected call of ListMember
func (mr *MockSubjectServiceMockRecorder) ListMember(_type, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMember", reflect.TypeOf((*MockSubjectService)(nil).ListMember), _type, id)
}

// UpdateMembersExpiredAt mocks base method
func (m *MockSubjectService) UpdateMembersExpiredAt(members []types.SubjectMember) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMembersExpiredAt", members)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMembersExpiredAt indicates an expected call of UpdateMembersExpiredAt
func (mr *MockSubjectServiceMockRecorder) UpdateMembersExpiredAt(members interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMembersExpiredAt", reflect.TypeOf((*MockSubjectService)(nil).UpdateMembersExpiredAt), members)
}

// BulkDeleteSubjectMembers mocks base method
func (m *MockSubjectService) BulkDeleteSubjectMembers(_type, id string, members []types.Subject) (map[string]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteSubjectMembers", _type, id, members)
	ret0, _ := ret[0].(map[string]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkDeleteSubjectMembers indicates an expected call of BulkDeleteSubjectMembers
func (mr *MockSubjectServiceMockRecorder) BulkDeleteSubjectMembers(_type, id, members interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteSubjectMembers", reflect.TypeOf((*MockSubjectService)(nil).BulkDeleteSubjectMembers), _type, id, members)
}

// BulkCreateSubjectMembers mocks base method
func (m *MockSubjectService) BulkCreateSubjectMembers(_type, id string, members []types.Subject, policyExpiredAt int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreateSubjectMembers", _type, id, members, policyExpiredAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkCreateSubjectMembers indicates an expected call of BulkCreateSubjectMembers
func (mr *MockSubjectServiceMockRecorder) BulkCreateSubjectMembers(_type, id, members, policyExpiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreateSubjectMembers", reflect.TypeOf((*MockSubjectService)(nil).BulkCreateSubjectMembers), _type, id, members, policyExpiredAt)
}

// GetSubjectDepartmentPKs mocks base method
func (m *MockSubjectService) GetSubjectDepartmentPKs(subjectPK int64) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubjectDepartmentPKs", subjectPK)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubjectDepartmentPKs indicates an expected call of GetSubjectDepartmentPKs
func (mr *MockSubjectServiceMockRecorder) GetSubjectDepartmentPKs(subjectPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubjectDepartmentPKs", reflect.TypeOf((*MockSubjectService)(nil).GetSubjectDepartmentPKs), subjectPK)
}

// GetSubjectDepartmentCount mocks base method
func (m *MockSubjectService) GetSubjectDepartmentCount() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubjectDepartmentCount")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubjectDepartmentCount indicates an expected call of GetSubjectDepartmentCount
func (mr *MockSubjectServiceMockRecorder) GetSubjectDepartmentCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubjectDepartmentCount", reflect.TypeOf((*MockSubjectService)(nil).GetSubjectDepartmentCount))
}

// ListPagingSubjectDepartment mocks base method
func (m *MockSubjectService) ListPagingSubjectDepartment(limit, offset int64) ([]types.SubjectDepartment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingSubjectDepartment", limit, offset)
	ret0, _ := ret[0].([]types.SubjectDepartment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingSubjectDepartment indicates an expected call of ListPagingSubjectDepartment
func (mr *MockSubjectServiceMockRecorder) ListPagingSubjectDepartment(limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingSubjectDepartment", reflect.TypeOf((*MockSubjectService)(nil).ListPagingSubjectDepartment), limit, offset)
}

// BulkCreateSubjectDepartments mocks base method
func (m *MockSubjectService) BulkCreateSubjectDepartments(subjectDepartments []types.SubjectDepartment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreateSubjectDepartments", subjectDepartments)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkCreateSubjectDepartments indicates an expected call of BulkCreateSubjectDepartments
func (mr *MockSubjectServiceMockRecorder) BulkCreateSubjectDepartments(subjectDepartments interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreateSubjectDepartments", reflect.TypeOf((*MockSubjectService)(nil).BulkCreateSubjectDepartments), subjectDepartments)
}

// BulkUpdateSubjectDepartments mocks base method
func (m *MockSubjectService) BulkUpdateSubjectDepartments(subjectDepartments []types.SubjectDepartment) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkUpdateSubjectDepartments", subjectDepartments)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkUpdateSubjectDepartments indicates an expected call of BulkUpdateSubjectDepartments
func (mr *MockSubjectServiceMockRecorder) BulkUpdateSubjectDepartments(subjectDepartments interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkUpdateSubjectDepartments", reflect.TypeOf((*MockSubjectService)(nil).BulkUpdateSubjectDepartments), subjectDepartments)
}

// BulkDeleteSubjectDepartments mocks base method
func (m *MockSubjectService) BulkDeleteSubjectDepartments(subjectIDs []string) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteSubjectDepartments", subjectIDs)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkDeleteSubjectDepartments indicates an expected call of BulkDeleteSubjectDepartments
func (mr *MockSubjectServiceMockRecorder) BulkDeleteSubjectDepartments(subjectIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteSubjectDepartments", reflect.TypeOf((*MockSubjectService)(nil).BulkDeleteSubjectDepartments), subjectIDs)
}

// ListSubjectPKByRole mocks base method
func (m *MockSubjectService) ListSubjectPKByRole(roleType, system string) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSubjectPKByRole", roleType, system)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubjectPKByRole indicates an expected call of ListSubjectPKByRole
func (mr *MockSubjectServiceMockRecorder) ListSubjectPKByRole(roleType, system interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubjectPKByRole", reflect.TypeOf((*MockSubjectService)(nil).ListSubjectPKByRole), roleType, system)
}

// ListRoleSystemIDBySubjectPK mocks base method
func (m *MockSubjectService) ListRoleSystemIDBySubjectPK(pk int64) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoleSystemIDBySubjectPK", pk)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoleSystemIDBySubjectPK indicates an expected call of ListRoleSystemIDBySubjectPK
func (mr *MockSubjectServiceMockRecorder) ListRoleSystemIDBySubjectPK(pk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoleSystemIDBySubjectPK", reflect.TypeOf((*MockSubjectService)(nil).ListRoleSystemIDBySubjectPK), pk)
}

// BulkCreateSubjectRoles mocks base method
func (m *MockSubjectService) BulkCreateSubjectRoles(roleType, system string, subjects []types.Subject) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreateSubjectRoles", roleType, system, subjects)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkCreateSubjectRoles indicates an expected call of BulkCreateSubjectRoles
func (mr *MockSubjectServiceMockRecorder) BulkCreateSubjectRoles(roleType, system, subjects interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreateSubjectRoles", reflect.TypeOf((*MockSubjectService)(nil).BulkCreateSubjectRoles), roleType, system, subjects)
}

// BulkDeleteSubjectRoles mocks base method
func (m *MockSubjectService) BulkDeleteSubjectRoles(roleType, system string, subjects []types.Subject) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteSubjectRoles", roleType, system, subjects)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDeleteSubjectRoles indicates an expected call of BulkDeleteSubjectRoles
func (mr *MockSubjectServiceMockRecorder) BulkDeleteSubjectRoles(roleType, system, subjects interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteSubjectRoles", reflect.TypeOf((*MockSubjectService)(nil).BulkDeleteSubjectRoles), roleType, system, subjects)
}
