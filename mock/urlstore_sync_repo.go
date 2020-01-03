// Code generated by MockGen. DO NOT EDIT.
// Source: app/repository/urlstore_sync_repo.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	repository "github.com/hotstone-seo/hotstone-server/app/repository"
	reflect "reflect"
)

// MockURLStoreSyncRepo is a mock of URLStoreSyncRepo interface
type MockURLStoreSyncRepo struct {
	ctrl     *gomock.Controller
	recorder *MockURLStoreSyncRepoMockRecorder
}

// MockURLStoreSyncRepoMockRecorder is the mock recorder for MockURLStoreSyncRepo
type MockURLStoreSyncRepoMockRecorder struct {
	mock *MockURLStoreSyncRepo
}

// NewMockURLStoreSyncRepo creates a new mock instance
func NewMockURLStoreSyncRepo(ctrl *gomock.Controller) *MockURLStoreSyncRepo {
	mock := &MockURLStoreSyncRepo{ctrl: ctrl}
	mock.recorder = &MockURLStoreSyncRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockURLStoreSyncRepo) EXPECT() *MockURLStoreSyncRepoMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockURLStoreSyncRepo) Find(ctx context.Context, id int64) (*repository.URLStoreSync, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, id)
	ret0, _ := ret[0].(*repository.URLStoreSync)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockURLStoreSyncRepoMockRecorder) Find(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockURLStoreSyncRepo)(nil).Find), ctx, id)
}

// List mocks base method
func (m *MockURLStoreSyncRepo) List(ctx context.Context) ([]*repository.URLStoreSync, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]*repository.URLStoreSync)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockURLStoreSyncRepoMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockURLStoreSyncRepo)(nil).List), ctx)
}

// Insert mocks base method
func (m *MockURLStoreSyncRepo) Insert(ctx context.Context, URLStoreSync repository.URLStoreSync) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, URLStoreSync)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockURLStoreSyncRepoMockRecorder) Insert(ctx, URLStoreSync interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockURLStoreSyncRepo)(nil).Insert), ctx, URLStoreSync)
}

// GetLatestVersion mocks base method
func (m *MockURLStoreSyncRepo) GetLatestVersion(ctx context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestVersion", ctx)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestVersion indicates an expected call of GetLatestVersion
func (mr *MockURLStoreSyncRepoMockRecorder) GetLatestVersion(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestVersion", reflect.TypeOf((*MockURLStoreSyncRepo)(nil).GetLatestVersion), ctx)
}

// GetListDiff mocks base method
func (m *MockURLStoreSyncRepo) GetListDiff(ctx context.Context, offsetVersion int64) ([]*repository.URLStoreSync, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListDiff", ctx, offsetVersion)
	ret0, _ := ret[0].([]*repository.URLStoreSync)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListDiff indicates an expected call of GetListDiff
func (mr *MockURLStoreSyncRepoMockRecorder) GetListDiff(ctx, offsetVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListDiff", reflect.TypeOf((*MockURLStoreSyncRepo)(nil).GetListDiff), ctx, offsetVersion)
}