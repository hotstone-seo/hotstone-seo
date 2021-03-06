// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hotstone-seo/hotstone-seo/internal/urlstore (interfaces: SyncRepo)

// Package urlstore_mock is a generated GoMock package.
package urlstore_mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	urlstore "github.com/hotstone-seo/hotstone-seo/internal/urlstore"
	reflect "reflect"
)

// MockSyncRepo is a mock of SyncRepo interface
type MockSyncRepo struct {
	ctrl     *gomock.Controller
	recorder *MockSyncRepoMockRecorder
}

// MockSyncRepoMockRecorder is the mock recorder for MockSyncRepo
type MockSyncRepoMockRecorder struct {
	mock *MockSyncRepo
}

// NewMockSyncRepo creates a new mock instance
func NewMockSyncRepo(ctrl *gomock.Controller) *MockSyncRepo {
	mock := &MockSyncRepo{ctrl: ctrl}
	mock.recorder = &MockSyncRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSyncRepo) EXPECT() *MockSyncRepoMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockSyncRepo) Find(arg0 context.Context) ([]*urlstore.Sync, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].([]*urlstore.Sync)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockSyncRepoMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockSyncRepo)(nil).Find), arg0)
}

// FindOne mocks base method
func (m *MockSyncRepo) FindOne(arg0 context.Context, arg1 int64) (*urlstore.Sync, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0, arg1)
	ret0, _ := ret[0].(*urlstore.Sync)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne
func (mr *MockSyncRepoMockRecorder) FindOne(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockSyncRepo)(nil).FindOne), arg0, arg1)
}

// FindRule mocks base method
func (m *MockSyncRepo) FindRule(arg0 context.Context, arg1 int64) (*urlstore.Sync, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindRule", arg0, arg1)
	ret0, _ := ret[0].(*urlstore.Sync)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRule indicates an expected call of FindRule
func (mr *MockSyncRepoMockRecorder) FindRule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRule", reflect.TypeOf((*MockSyncRepo)(nil).FindRule), arg0, arg1)
}

// GetLatestVersion mocks base method
func (m *MockSyncRepo) GetLatestVersion(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestVersion", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestVersion indicates an expected call of GetLatestVersion
func (mr *MockSyncRepoMockRecorder) GetLatestVersion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestVersion", reflect.TypeOf((*MockSyncRepo)(nil).GetLatestVersion), arg0)
}

// GetListDiff mocks base method
func (m *MockSyncRepo) GetListDiff(arg0 context.Context, arg1 int64) ([]*urlstore.Sync, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListDiff", arg0, arg1)
	ret0, _ := ret[0].([]*urlstore.Sync)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListDiff indicates an expected call of GetListDiff
func (mr *MockSyncRepoMockRecorder) GetListDiff(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListDiff", reflect.TypeOf((*MockSyncRepo)(nil).GetListDiff), arg0, arg1)
}

// Insert mocks base method
func (m *MockSyncRepo) Insert(arg0 context.Context, arg1 urlstore.Sync) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockSyncRepoMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockSyncRepo)(nil).Insert), arg0, arg1)
}
