// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hotstone-seo/hotstone-seo/server/service (interfaces: RoleTypeService)

// Package service_mock is a generated GoMock package.
package service_mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	repository "github.com/hotstone-seo/hotstone-seo/server/repository"
	service "github.com/hotstone-seo/hotstone-seo/server/service"
	reflect "reflect"
)

// MockRoleTypeService is a mock of RoleTypeService interface
type MockRoleTypeService struct {
	ctrl     *gomock.Controller
	recorder *MockRoleTypeServiceMockRecorder
}

// MockRoleTypeServiceMockRecorder is the mock recorder for MockRoleTypeService
type MockRoleTypeServiceMockRecorder struct {
	mock *MockRoleTypeService
}

// NewMockRoleTypeService creates a new mock instance
func NewMockRoleTypeService(ctrl *gomock.Controller) *MockRoleTypeService {
	mock := &MockRoleTypeService{ctrl: ctrl}
	mock.recorder = &MockRoleTypeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRoleTypeService) EXPECT() *MockRoleTypeServiceMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockRoleTypeService) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockRoleTypeServiceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRoleTypeService)(nil).Delete), arg0, arg1)
}

// Find mocks base method
func (m *MockRoleTypeService) Find(arg0 context.Context, arg1 repository.PaginationParam) ([]*repository.RoleType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].([]*repository.RoleType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockRoleTypeServiceMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRoleTypeService)(nil).Find), arg0, arg1)
}

// FindOne mocks base method
func (m *MockRoleTypeService) FindOne(arg0 context.Context, arg1 int64) (*repository.RoleType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0, arg1)
	ret0, _ := ret[0].(*repository.RoleType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne
func (mr *MockRoleTypeServiceMockRecorder) FindOne(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockRoleTypeService)(nil).FindOne), arg0, arg1)
}

// FindOneByName mocks base method
func (m *MockRoleTypeService) FindOneByName(arg0 context.Context, arg1 string) (*repository.RoleType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByName", arg0, arg1)
	ret0, _ := ret[0].(*repository.RoleType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByName indicates an expected call of FindOneByName
func (mr *MockRoleTypeServiceMockRecorder) FindOneByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByName", reflect.TypeOf((*MockRoleTypeService)(nil).FindOneByName), arg0, arg1)
}

// Insert mocks base method
func (m *MockRoleTypeService) Insert(arg0 context.Context, arg1 service.RoleTypeRequest) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockRoleTypeServiceMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRoleTypeService)(nil).Insert), arg0, arg1)
}

// Update mocks base method
func (m *MockRoleTypeService) Update(arg0 context.Context, arg1 repository.RoleType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockRoleTypeServiceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRoleTypeService)(nil).Update), arg0, arg1)
}
