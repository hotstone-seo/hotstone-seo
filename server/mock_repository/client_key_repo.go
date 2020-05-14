// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hotstone-seo/hotstone-seo/server/repository (interfaces: ClientKeyRepo)

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	repository "github.com/hotstone-seo/hotstone-seo/server/repository"
	dbkit "github.com/typical-go/typical-rest-server/pkg/dbkit"
	reflect "reflect"
)

// MockClientKeyRepo is a mock of ClientKeyRepo interface
type MockClientKeyRepo struct {
	ctrl     *gomock.Controller
	recorder *MockClientKeyRepoMockRecorder
}

// MockClientKeyRepoMockRecorder is the mock recorder for MockClientKeyRepo
type MockClientKeyRepoMockRecorder struct {
	mock *MockClientKeyRepo
}

// NewMockClientKeyRepo creates a new mock instance
func NewMockClientKeyRepo(ctrl *gomock.Controller) *MockClientKeyRepo {
	mock := &MockClientKeyRepo{ctrl: ctrl}
	mock.recorder = &MockClientKeyRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientKeyRepo) EXPECT() *MockClientKeyRepoMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockClientKeyRepo) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockClientKeyRepoMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClientKeyRepo)(nil).Delete), arg0, arg1)
}

// Find mocks base method
func (m *MockClientKeyRepo) Find(arg0 context.Context, arg1 ...dbkit.FindOption) ([]*repository.ClientKey, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].([]*repository.ClientKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockClientKeyRepoMockRecorder) Find(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockClientKeyRepo)(nil).Find), varargs...)
}

// FindOne mocks base method
func (m *MockClientKeyRepo) FindOne(arg0 context.Context, arg1 int64) (*repository.ClientKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0, arg1)
	ret0, _ := ret[0].(*repository.ClientKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne
func (mr *MockClientKeyRepoMockRecorder) FindOne(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockClientKeyRepo)(nil).FindOne), arg0, arg1)
}

// Insert mocks base method
func (m *MockClientKeyRepo) Insert(arg0 context.Context, arg1 repository.ClientKey) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockClientKeyRepoMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockClientKeyRepo)(nil).Insert), arg0, arg1)
}

// Update mocks base method
func (m *MockClientKeyRepo) Update(arg0 context.Context, arg1 repository.ClientKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockClientKeyRepoMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockClientKeyRepo)(nil).Update), arg0, arg1)
}
