// Code generated by MockGen. DO NOT EDIT.
// Source: resource_type.go

// Package mock is a generated GoMock package.
package mock

import (
	types "iam/pkg/service/types"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockResourceTypeService is a mock of ResourceTypeService interface.
type MockResourceTypeService struct {
	ctrl     *gomock.Controller
	recorder *MockResourceTypeServiceMockRecorder
}

// MockResourceTypeServiceMockRecorder is the mock recorder for MockResourceTypeService.
type MockResourceTypeServiceMockRecorder struct {
	mock *MockResourceTypeService
}

// NewMockResourceTypeService creates a new mock instance.
func NewMockResourceTypeService(ctrl *gomock.Controller) *MockResourceTypeService {
	mock := &MockResourceTypeService{ctrl: ctrl}
	mock.recorder = &MockResourceTypeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceTypeService) EXPECT() *MockResourceTypeServiceMockRecorder {
	return m.recorder
}

// BulkCreate mocks base method.
func (m *MockResourceTypeService) BulkCreate(system string, resourceTypes []types.ResourceType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreate", system, resourceTypes)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkCreate indicates an expected call of BulkCreate.
func (mr *MockResourceTypeServiceMockRecorder) BulkCreate(system, resourceTypes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreate", reflect.TypeOf((*MockResourceTypeService)(nil).BulkCreate), system, resourceTypes)
}

// BulkDelete mocks base method.
func (m *MockResourceTypeService) BulkDelete(system string, resourceTypeIDs []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDelete", system, resourceTypeIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDelete indicates an expected call of BulkDelete.
func (mr *MockResourceTypeServiceMockRecorder) BulkDelete(system, resourceTypeIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDelete", reflect.TypeOf((*MockResourceTypeService)(nil).BulkDelete), system, resourceTypeIDs)
}

// Get mocks base method.
func (m *MockResourceTypeService) Get(system, id string) (types.ResourceType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", system, id)
	ret0, _ := ret[0].(types.ResourceType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockResourceTypeServiceMockRecorder) Get(system, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockResourceTypeService)(nil).Get), system, id)
}

// GetPK mocks base method.
func (m *MockResourceTypeService) GetPK(system, name string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPK", system, name)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPK indicates an expected call of GetPK.
func (mr *MockResourceTypeServiceMockRecorder) GetPK(system, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPK", reflect.TypeOf((*MockResourceTypeService)(nil).GetPK), system, name)
}

// ListBySystem mocks base method.
func (m *MockResourceTypeService) ListBySystem(system string) ([]types.ResourceType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySystem", system)
	ret0, _ := ret[0].([]types.ResourceType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySystem indicates an expected call of ListBySystem.
func (mr *MockResourceTypeServiceMockRecorder) ListBySystem(system interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySystem", reflect.TypeOf((*MockResourceTypeService)(nil).ListBySystem), system)
}

// Update mocks base method.
func (m *MockResourceTypeService) Update(system, resourceTypeID string, resourceType types.ResourceType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", system, resourceTypeID, resourceType)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockResourceTypeServiceMockRecorder) Update(system, resourceTypeID, resourceType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockResourceTypeService)(nil).Update), system, resourceTypeID, resourceType)
}
