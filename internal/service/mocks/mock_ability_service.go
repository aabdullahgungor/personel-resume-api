// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/interfaces/ability_service_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	model "github.com/aabdullahgungor/personal-resume-api/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockIAbilityService is a mock of IAbilityService interface.
type MockIAbilityService struct {
	ctrl     *gomock.Controller
	recorder *MockIAbilityServiceMockRecorder
}

// MockIAbilityServiceMockRecorder is the mock recorder for MockIAbilityService.
type MockIAbilityServiceMockRecorder struct {
	mock *MockIAbilityService
}

// NewMockIAbilityService creates a new mock instance.
func NewMockIAbilityService(ctrl *gomock.Controller) *MockIAbilityService {
	mock := &MockIAbilityService{ctrl: ctrl}
	mock.recorder = &MockIAbilityServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAbilityService) EXPECT() *MockIAbilityServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIAbilityService) Create(ability *model.Ability) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ability)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIAbilityServiceMockRecorder) Create(ability interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIAbilityService)(nil).Create), ability)
}

// Delete mocks base method.
func (m *MockIAbilityService) Delete(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIAbilityServiceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIAbilityService)(nil).Delete), id)
}

// Edit mocks base method.
func (m *MockIAbilityService) Edit(ability *model.Ability) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edit", ability)
	ret0, _ := ret[0].(error)
	return ret0
}

// Edit indicates an expected call of Edit.
func (mr *MockIAbilityServiceMockRecorder) Edit(ability interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edit", reflect.TypeOf((*MockIAbilityService)(nil).Edit), ability)
}

// GetAll mocks base method.
func (m *MockIAbilityService) GetAll() ([]model.Ability, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]model.Ability)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIAbilityServiceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIAbilityService)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockIAbilityService) GetById(id string) (model.Ability, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(model.Ability)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockIAbilityServiceMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockIAbilityService)(nil).GetById), id)
}