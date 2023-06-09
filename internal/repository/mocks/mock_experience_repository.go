// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/interfaces/experience_repository_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	model "github.com/aabdullahgungor/personal-resume-api/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockIExperienceRepository is a mock of IExperienceRepository interface.
type MockIExperienceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIExperienceRepositoryMockRecorder
}

// MockIExperienceRepositoryMockRecorder is the mock recorder for MockIExperienceRepository.
type MockIExperienceRepositoryMockRecorder struct {
	mock *MockIExperienceRepository
}

// NewMockIExperienceRepository creates a new mock instance.
func NewMockIExperienceRepository(ctrl *gomock.Controller) *MockIExperienceRepository {
	mock := &MockIExperienceRepository{ctrl: ctrl}
	mock.recorder = &MockIExperienceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIExperienceRepository) EXPECT() *MockIExperienceRepositoryMockRecorder {
	return m.recorder
}

// CreateExperience mocks base method.
func (m *MockIExperienceRepository) CreateExperience(experience *model.Experience) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExperience", experience)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateExperience indicates an expected call of CreateExperience.
func (mr *MockIExperienceRepositoryMockRecorder) CreateExperience(experience interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExperience", reflect.TypeOf((*MockIExperienceRepository)(nil).CreateExperience), experience)
}

// DeleteExperience mocks base method.
func (m *MockIExperienceRepository) DeleteExperience(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExperience", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExperience indicates an expected call of DeleteExperience.
func (mr *MockIExperienceRepositoryMockRecorder) DeleteExperience(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExperience", reflect.TypeOf((*MockIExperienceRepository)(nil).DeleteExperience), id)
}

// EditExperience mocks base method.
func (m *MockIExperienceRepository) EditExperience(experience *model.Experience) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditExperience", experience)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditExperience indicates an expected call of EditExperience.
func (mr *MockIExperienceRepositoryMockRecorder) EditExperience(experience interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditExperience", reflect.TypeOf((*MockIExperienceRepository)(nil).EditExperience), experience)
}

// GetAllExperiences mocks base method.
func (m *MockIExperienceRepository) GetAllExperiences() ([]model.Experience, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllExperiences")
	ret0, _ := ret[0].([]model.Experience)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllExperiences indicates an expected call of GetAllExperiences.
func (mr *MockIExperienceRepositoryMockRecorder) GetAllExperiences() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllExperiences", reflect.TypeOf((*MockIExperienceRepository)(nil).GetAllExperiences))
}

// GetExperienceById mocks base method.
func (m *MockIExperienceRepository) GetExperienceById(id int) (model.Experience, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExperienceById", id)
	ret0, _ := ret[0].(model.Experience)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExperienceById indicates an expected call of GetExperienceById.
func (mr *MockIExperienceRepositoryMockRecorder) GetExperienceById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExperienceById", reflect.TypeOf((*MockIExperienceRepository)(nil).GetExperienceById), id)
}