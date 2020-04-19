// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hotstone-seo/hotstone-seo/server/metric (interfaces: RuleMatchingRepo)

// Package mock_metric is a generated GoMock package.
package mock_metric

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	metric "github.com/hotstone-seo/hotstone-seo/server/metric"
	reflect "reflect"
)

// MockRuleMatchingRepo is a mock of RuleMatchingRepo interface
type MockRuleMatchingRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRuleMatchingRepoMockRecorder
}

// MockRuleMatchingRepoMockRecorder is the mock recorder for MockRuleMatchingRepo
type MockRuleMatchingRepoMockRecorder struct {
	mock *MockRuleMatchingRepo
}

// NewMockRuleMatchingRepo creates a new mock instance
func NewMockRuleMatchingRepo(ctrl *gomock.Controller) *MockRuleMatchingRepo {
	mock := &MockRuleMatchingRepo{ctrl: ctrl}
	mock.recorder = &MockRuleMatchingRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRuleMatchingRepo) EXPECT() *MockRuleMatchingRepoMockRecorder {
	return m.recorder
}

// Insert mocks base method
func (m *MockRuleMatchingRepo) Insert(arg0 context.Context, arg1 metric.RuleMatching) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockRuleMatchingRepoMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRuleMatchingRepo)(nil).Insert), arg0, arg1)
}
