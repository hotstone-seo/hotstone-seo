// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hotstone-seo/hotstone-seo/server/repository (interfaces: HistoryRepo)

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	repository "github.com/hotstone-seo/hotstone-seo/server/repository"
	reflect "reflect"
)

// MockHistoryRepo is a mock of HistoryRepo interface
type MockHistoryRepo struct {
	ctrl     *gomock.Controller
	recorder *MockHistoryRepoMockRecorder
}

// MockHistoryRepoMockRecorder is the mock recorder for MockHistoryRepo
type MockHistoryRepoMockRecorder struct {
	mock *MockHistoryRepo
}

// NewMockHistoryRepo creates a new mock instance
func NewMockHistoryRepo(ctrl *gomock.Controller) *MockHistoryRepo {
	mock := &MockHistoryRepo{ctrl: ctrl}
	mock.recorder = &MockHistoryRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHistoryRepo) EXPECT() *MockHistoryRepoMockRecorder {
	return m.recorder
}

// Insert mocks base method
func (m *MockHistoryRepo) Insert(arg0 context.Context, arg1 repository.History) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockHistoryRepoMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockHistoryRepo)(nil).Insert), arg0, arg1)
}