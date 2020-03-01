// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hotstone-seo/hotstone-seo/server/urlstore (interfaces: URLSyncRepo)

// Package mock_urlstore is a generated GoMock package.
package mock_urlstore

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	urlstore "github.com/hotstone-seo/hotstone-seo/server/urlstore"
)

// MockURLSyncRepo is a mock of URLSyncRepo interface
type MockURLSyncRepo struct {
	ctrl     *gomock.Controller
	recorder *MockURLSyncRepoMockRecorder
}

// MockURLSyncRepoMockRecorder is the mock recorder for MockURLSyncRepo
type MockURLSyncRepoMockRecorder struct {
	mock *MockURLSyncRepo
}

// NewMockURLSyncRepo creates a new mock instance
func NewMockURLSyncRepo(ctrl *gomock.Controller) *MockURLSyncRepo {
	mock := &MockURLSyncRepo{ctrl: ctrl}
	mock.recorder = &MockURLSyncRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockURLSyncRepo) EXPECT() *MockURLSyncRepoMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockURLSyncRepo) Find(arg0 context.Context) ([]*urlstore.URLSync, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].([]*urlstore.URLSync)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockURLSyncRepoMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockURLSyncRepo)(nil).Find), arg0)
}

// FindOne mocks base method
func (m *MockURLSyncRepo) FindOne(arg0 context.Context, arg1 int64) (*urlstore.URLSync, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0, arg1)
	ret0, _ := ret[0].(*urlstore.URLSync)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne
func (mr *MockURLSyncRepoMockRecorder) FindOne(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockURLSyncRepo)(nil).FindOne), arg0, arg1)
}

// GetLatestVersion mocks base method
func (m *MockURLSyncRepo) GetLatestVersion(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestVersion", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestVersion indicates an expected call of GetLatestVersion
func (mr *MockURLSyncRepoMockRecorder) GetLatestVersion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestVersion", reflect.TypeOf((*MockURLSyncRepo)(nil).GetLatestVersion), arg0)
}

// GetListDiff mocks base method
func (m *MockURLSyncRepo) GetListDiff(arg0 context.Context, arg1 int64) ([]*urlstore.URLSync, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListDiff", arg0, arg1)
	ret0, _ := ret[0].([]*urlstore.URLSync)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListDiff indicates an expected call of GetListDiff
func (mr *MockURLSyncRepoMockRecorder) GetListDiff(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListDiff", reflect.TypeOf((*MockURLSyncRepo)(nil).GetListDiff), arg0, arg1)
}

// Insert mocks base method
func (m *MockURLSyncRepo) Insert(arg0 context.Context, arg1 urlstore.URLSync) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockURLSyncRepoMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockURLSyncRepo)(nil).Insert), arg0, arg1)
}
