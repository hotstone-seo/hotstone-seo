// Code generated by MockGen. DO NOT EDIT.
// Source: app/repository/tag_repo.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	repository "github.com/hotstone-seo/hotstone-server/app/repository"
	reflect "reflect"
)

// MockTagRepo is a mock of TagRepo interface
type MockTagRepo struct {
	ctrl     *gomock.Controller
	recorder *MockTagRepoMockRecorder
}

// MockTagRepoMockRecorder is the mock recorder for MockTagRepo
type MockTagRepoMockRecorder struct {
	mock *MockTagRepo
}

// NewMockTagRepo creates a new mock instance
func NewMockTagRepo(ctrl *gomock.Controller) *MockTagRepo {
	mock := &MockTagRepo{ctrl: ctrl}
	mock.recorder = &MockTagRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTagRepo) EXPECT() *MockTagRepoMockRecorder {
	return m.recorder
}

// FindOne mocks base method
func (m *MockTagRepo) FindOne(arg0 context.Context, arg1 int64) (*repository.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0, arg1)
	ret0, _ := ret[0].(*repository.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne
func (mr *MockTagRepoMockRecorder) FindOne(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockTagRepo)(nil).FindOne), arg0, arg1)
}

// Find mocks base method
func (m *MockTagRepo) Find(arg0 context.Context) ([]*repository.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].([]*repository.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockTagRepoMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockTagRepo)(nil).Find), arg0)
}

// FindByRuleAndLocale mocks base method
func (m *MockTagRepo) FindByRuleAndLocale(ctx context.Context, ruleID, localeID int64) ([]*repository.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByRuleAndLocale", ctx, ruleID, localeID)
	ret0, _ := ret[0].([]*repository.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByRuleAndLocale indicates an expected call of FindByRuleAndLocale
func (mr *MockTagRepoMockRecorder) FindByRuleAndLocale(ctx, ruleID, localeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByRuleAndLocale", reflect.TypeOf((*MockTagRepo)(nil).FindByRuleAndLocale), ctx, ruleID, localeID)
}

// Insert mocks base method
func (m *MockTagRepo) Insert(arg0 context.Context, arg1 repository.Tag) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockTagRepoMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTagRepo)(nil).Insert), arg0, arg1)
}

// Delete mocks base method
func (m *MockTagRepo) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockTagRepoMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTagRepo)(nil).Delete), arg0, arg1)
}

// Update mocks base method
func (m *MockTagRepo) Update(arg0 context.Context, arg1 repository.Tag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockTagRepoMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTagRepo)(nil).Update), arg0, arg1)
}
