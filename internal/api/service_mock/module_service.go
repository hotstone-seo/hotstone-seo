// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hotstone-seo/hotstone-seo/internal/api/service (interfaces: ModuleService)

// Package service_mock is a generated GoMock package.
package service_mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	repository "github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	service "github.com/hotstone-seo/hotstone-seo/internal/api/service"
	reflect "reflect"
)

// MockModuleService is a mock of ModuleService interface
type MockModuleService struct {
	ctrl     *gomock.Controller
	recorder *MockModuleServiceMockRecorder
}

// MockModuleServiceMockRecorder is the mock recorder for MockModuleService
type MockModuleServiceMockRecorder struct {
	mock *MockModuleService
}

// NewMockModuleService creates a new mock instance
func NewMockModuleService(ctrl *gomock.Controller) *MockModuleService {
	mock := &MockModuleService{ctrl: ctrl}
	mock.recorder = &MockModuleServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModuleService) EXPECT() *MockModuleServiceMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockModuleService) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockModuleServiceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockModuleService)(nil).Delete), arg0, arg1)
}

// Find mocks base method
func (m *MockModuleService) Find(arg0 context.Context, arg1 repository.PaginationParam) ([]*repository.Module, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].([]*repository.Module)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockModuleServiceMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockModuleService)(nil).Find), arg0, arg1)
}

// FindOne mocks base method
func (m *MockModuleService) FindOne(arg0 context.Context, arg1 int64) (*repository.Module, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0, arg1)
	ret0, _ := ret[0].(*repository.Module)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne
func (mr *MockModuleServiceMockRecorder) FindOne(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockModuleService)(nil).FindOne), arg0, arg1)
}

// Insert mocks base method
func (m *MockModuleService) Insert(arg0 context.Context, arg1 service.ModuleRequest) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockModuleServiceMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockModuleService)(nil).Insert), arg0, arg1)
}

// Update mocks base method
func (m *MockModuleService) Update(arg0 context.Context, arg1 service.ModuleRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockModuleServiceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockModuleService)(nil).Update), arg0, arg1)
}
