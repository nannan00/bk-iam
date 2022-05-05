// Code generated by MockGen. DO NOT EDIT.
// Source: action.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	types "iam/pkg/service/types"
	reflect "reflect"
)

// MockActionService is a mock of ActionService interface
type MockActionService struct {
	ctrl     *gomock.Controller
	recorder *MockActionServiceMockRecorder
}

// MockActionServiceMockRecorder is the mock recorder for MockActionService
type MockActionServiceMockRecorder struct {
	mock *MockActionService
}

// NewMockActionService creates a new mock instance
func NewMockActionService(ctrl *gomock.Controller) *MockActionService {
	mock := &MockActionService{ctrl: ctrl}
	mock.recorder = &MockActionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockActionService) EXPECT() *MockActionServiceMockRecorder {
	return m.recorder
}

// GetActionPK mocks base method
func (m *MockActionService) GetActionPK(system, id string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActionPK", system, id)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActionPK indicates an expected call of GetActionPK
func (mr *MockActionServiceMockRecorder) GetActionPK(system, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActionPK", reflect.TypeOf((*MockActionService)(nil).GetActionPK), system, id)
}

// Get mocks base method
func (m *MockActionService) Get(system, id string) (types.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", system, id)
	ret0, _ := ret[0].(types.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockActionServiceMockRecorder) Get(system, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockActionService)(nil).Get), system, id)
}

// ListBySystem mocks base method
func (m *MockActionService) ListBySystem(system string) ([]types.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySystem", system)
	ret0, _ := ret[0].([]types.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySystem indicates an expected call of ListBySystem
func (mr *MockActionServiceMockRecorder) ListBySystem(system interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySystem", reflect.TypeOf((*MockActionService)(nil).ListBySystem), system)
}

// ListAllBySystem mocks base method
func (m *MockActionService) ListAllBySystem(system string) ([]types.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllBySystem", system)
	ret0, _ := ret[0].([]types.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllBySystem indicates an expected call of ListAllBySystem
func (mr *MockActionServiceMockRecorder) ListAllBySystem(system interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllBySystem", reflect.TypeOf((*MockActionService)(nil).ListAllBySystem), system)
}

// BulkCreate mocks base method
func (m *MockActionService) BulkCreate(system string, actions []types.Action) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreate", system, actions)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkCreate indicates an expected call of BulkCreate
func (mr *MockActionServiceMockRecorder) BulkCreate(system, actions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreate", reflect.TypeOf((*MockActionService)(nil).BulkCreate), system, actions)
}

// Update mocks base method
func (m *MockActionService) Update(system, actionID string, action types.Action) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", system, actionID, action)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockActionServiceMockRecorder) Update(system, actionID, action interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockActionService)(nil).Update), system, actionID, action)
}

// BulkDelete mocks base method
func (m *MockActionService) BulkDelete(system string, actionIDs []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDelete", system, actionIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDelete indicates an expected call of BulkDelete
func (mr *MockActionServiceMockRecorder) BulkDelete(system, actionIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDelete", reflect.TypeOf((*MockActionService)(nil).BulkDelete), system, actionIDs)
}

// GetThinActionByPK mocks base method
func (m *MockActionService) GetThinActionByPK(pk int64) (types.ThinAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetThinActionByPK", pk)
	ret0, _ := ret[0].(types.ThinAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetThinActionByPK indicates an expected call of GetThinActionByPK
func (mr *MockActionServiceMockRecorder) GetThinActionByPK(pk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThinActionByPK", reflect.TypeOf((*MockActionService)(nil).GetThinActionByPK), pk)
}

// ListThinActionByPKs mocks base method
func (m *MockActionService) ListThinActionByPKs(pks []int64) ([]types.ThinAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListThinActionByPKs", pks)
	ret0, _ := ret[0].([]types.ThinAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListThinActionByPKs indicates an expected call of ListThinActionByPKs
func (mr *MockActionServiceMockRecorder) ListThinActionByPKs(pks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListThinActionByPKs", reflect.TypeOf((*MockActionService)(nil).ListThinActionByPKs), pks)
}

// ListThinActionBySystem mocks base method
func (m *MockActionService) ListThinActionBySystem(system string) ([]types.ThinAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListThinActionBySystem", system)
	ret0, _ := ret[0].([]types.ThinAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListThinActionBySystem indicates an expected call of ListThinActionBySystem
func (mr *MockActionServiceMockRecorder) ListThinActionBySystem(system interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListThinActionBySystem", reflect.TypeOf((*MockActionService)(nil).ListThinActionBySystem), system)
}

// ListThinActionResourceTypes mocks base method
func (m *MockActionService) ListThinActionResourceTypes(system, actionID string) ([]types.ThinActionResourceType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListThinActionResourceTypes", system, actionID)
	ret0, _ := ret[0].([]types.ThinActionResourceType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListThinActionResourceTypes indicates an expected call of ListThinActionResourceTypes
func (mr *MockActionServiceMockRecorder) ListThinActionResourceTypes(system, actionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListThinActionResourceTypes", reflect.TypeOf((*MockActionService)(nil).ListThinActionResourceTypes), system, actionID)
}

// ListActionResourceTypeIDByResourceTypeSystem mocks base method
func (m *MockActionService) ListActionResourceTypeIDByResourceTypeSystem(resourceTypeSystem string) ([]types.ActionResourceTypeID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListActionResourceTypeIDByResourceTypeSystem", resourceTypeSystem)
	ret0, _ := ret[0].([]types.ActionResourceTypeID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListActionResourceTypeIDByResourceTypeSystem indicates an expected call of ListActionResourceTypeIDByResourceTypeSystem
func (mr *MockActionServiceMockRecorder) ListActionResourceTypeIDByResourceTypeSystem(resourceTypeSystem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListActionResourceTypeIDByResourceTypeSystem", reflect.TypeOf((*MockActionService)(nil).ListActionResourceTypeIDByResourceTypeSystem), resourceTypeSystem)
}

// ListActionResourceTypeIDByActionSystem mocks base method
func (m *MockActionService) ListActionResourceTypeIDByActionSystem(actionSystem string) ([]types.ActionResourceTypeID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListActionResourceTypeIDByActionSystem", actionSystem)
	ret0, _ := ret[0].([]types.ActionResourceTypeID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListActionResourceTypeIDByActionSystem indicates an expected call of ListActionResourceTypeIDByActionSystem
func (mr *MockActionServiceMockRecorder) ListActionResourceTypeIDByActionSystem(actionSystem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListActionResourceTypeIDByActionSystem", reflect.TypeOf((*MockActionService)(nil).ListActionResourceTypeIDByActionSystem), actionSystem)
}

// ListActionInstanceSelectionIDBySystem mocks base method
func (m *MockActionService) ListActionInstanceSelectionIDBySystem(system string) ([]types.ActionInstanceSelectionID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListActionInstanceSelectionIDBySystem", system)
	ret0, _ := ret[0].([]types.ActionInstanceSelectionID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListActionInstanceSelectionIDBySystem indicates an expected call of ListActionInstanceSelectionIDBySystem
func (mr *MockActionServiceMockRecorder) ListActionInstanceSelectionIDBySystem(system interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListActionInstanceSelectionIDBySystem", reflect.TypeOf((*MockActionService)(nil).ListActionInstanceSelectionIDBySystem), system)
}
