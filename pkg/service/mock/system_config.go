// Code generated by MockGen. DO NOT EDIT.
// Source: system_config.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSystemConfigService is a mock of SystemConfigService interface.
type MockSystemConfigService struct {
	ctrl     *gomock.Controller
	recorder *MockSystemConfigServiceMockRecorder
}

// MockSystemConfigServiceMockRecorder is the mock recorder for MockSystemConfigService.
type MockSystemConfigServiceMockRecorder struct {
	mock *MockSystemConfigService
}

// NewMockSystemConfigService creates a new mock instance.
func NewMockSystemConfigService(ctrl *gomock.Controller) *MockSystemConfigService {
	mock := &MockSystemConfigService{ctrl: ctrl}
	mock.recorder = &MockSystemConfigServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSystemConfigService) EXPECT() *MockSystemConfigServiceMockRecorder {
	return m.recorder
}

// CreateOrUpdateActionGroups mocks base method.
func (m *MockSystemConfigService) CreateOrUpdateActionGroups(system string, actionGroup []interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateActionGroups", system, actionGroup)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateActionGroups indicates an expected call of CreateOrUpdateActionGroups.
func (mr *MockSystemConfigServiceMockRecorder) CreateOrUpdateActionGroups(system, actionGroup interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateActionGroups", reflect.TypeOf((*MockSystemConfigService)(nil).CreateOrUpdateActionGroups), system, actionGroup)
}

// CreateOrUpdateCommonActions mocks base method.
func (m *MockSystemConfigService) CreateOrUpdateCommonActions(system string, actions []interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateCommonActions", system, actions)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateCommonActions indicates an expected call of CreateOrUpdateCommonActions.
func (mr *MockSystemConfigServiceMockRecorder) CreateOrUpdateCommonActions(system, actions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateCommonActions", reflect.TypeOf((*MockSystemConfigService)(nil).CreateOrUpdateCommonActions), system, actions)
}

// CreateOrUpdateFeatureShieldRules mocks base method.
func (m *MockSystemConfigService) CreateOrUpdateFeatureShieldRules(system string, featureShieldRules []interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateFeatureShieldRules", system, featureShieldRules)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateFeatureShieldRules indicates an expected call of CreateOrUpdateFeatureShieldRules.
func (mr *MockSystemConfigServiceMockRecorder) CreateOrUpdateFeatureShieldRules(system, featureShieldRules interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateFeatureShieldRules", reflect.TypeOf((*MockSystemConfigService)(nil).CreateOrUpdateFeatureShieldRules), system, featureShieldRules)
}

// CreateOrUpdateResourceCreatorActions mocks base method.
func (m *MockSystemConfigService) CreateOrUpdateResourceCreatorActions(system string, resourceCreatorAction map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateResourceCreatorActions", system, resourceCreatorAction)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateResourceCreatorActions indicates an expected call of CreateOrUpdateResourceCreatorActions.
func (mr *MockSystemConfigServiceMockRecorder) CreateOrUpdateResourceCreatorActions(system, resourceCreatorAction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateResourceCreatorActions", reflect.TypeOf((*MockSystemConfigService)(nil).CreateOrUpdateResourceCreatorActions), system, resourceCreatorAction)
}

// CreateOrUpdateSystemManagers mocks base method.
func (m *MockSystemConfigService) CreateOrUpdateSystemManagers(system string, systemManagers []interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateSystemManagers", system, systemManagers)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateSystemManagers indicates an expected call of CreateOrUpdateSystemManagers.
func (mr *MockSystemConfigServiceMockRecorder) CreateOrUpdateSystemManagers(system, systemManagers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateSystemManagers", reflect.TypeOf((*MockSystemConfigService)(nil).CreateOrUpdateSystemManagers), system, systemManagers)
}

// Exists mocks base method.
func (m *MockSystemConfigService) Exists(system, key string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", system, key)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exists indicates an expected call of Exists.
func (mr *MockSystemConfigServiceMockRecorder) Exists(system, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockSystemConfigService)(nil).Exists), system, key)
}

// GetActionGroups mocks base method.
func (m *MockSystemConfigService) GetActionGroups(system string) ([]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActionGroups", system)
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActionGroups indicates an expected call of GetActionGroups.
func (mr *MockSystemConfigServiceMockRecorder) GetActionGroups(system interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActionGroups", reflect.TypeOf((*MockSystemConfigService)(nil).GetActionGroups), system)
}

// GetCommonActions mocks base method.
func (m *MockSystemConfigService) GetCommonActions(system string) ([]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommonActions", system)
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommonActions indicates an expected call of GetCommonActions.
func (mr *MockSystemConfigServiceMockRecorder) GetCommonActions(system interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommonActions", reflect.TypeOf((*MockSystemConfigService)(nil).GetCommonActions), system)
}

// GetFeatureShieldRules mocks base method.
func (m *MockSystemConfigService) GetFeatureShieldRules(system string) ([]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeatureShieldRules", system)
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeatureShieldRules indicates an expected call of GetFeatureShieldRules.
func (mr *MockSystemConfigServiceMockRecorder) GetFeatureShieldRules(system interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeatureShieldRules", reflect.TypeOf((*MockSystemConfigService)(nil).GetFeatureShieldRules), system)
}

// GetResourceCreatorActions mocks base method.
func (m *MockSystemConfigService) GetResourceCreatorActions(system string) (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceCreatorActions", system)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceCreatorActions indicates an expected call of GetResourceCreatorActions.
func (mr *MockSystemConfigServiceMockRecorder) GetResourceCreatorActions(system interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceCreatorActions", reflect.TypeOf((*MockSystemConfigService)(nil).GetResourceCreatorActions), system)
}

// GetSystemManagers mocks base method.
func (m *MockSystemConfigService) GetSystemManagers(system string) ([]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSystemManagers", system)
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSystemManagers indicates an expected call of GetSystemManagers.
func (mr *MockSystemConfigServiceMockRecorder) GetSystemManagers(system interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSystemManagers", reflect.TypeOf((*MockSystemConfigService)(nil).GetSystemManagers), system)
}

// get mocks base method.
func (m *MockSystemConfigService) get(system, key string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "get", system, key)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// get indicates an expected call of get.
func (mr *MockSystemConfigServiceMockRecorder) get(system, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "get", reflect.TypeOf((*MockSystemConfigService)(nil).get), system, key)
}
