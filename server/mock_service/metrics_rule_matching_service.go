// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hotstone-seo/hotstone-seo/server/service (interfaces: MetricsRuleMatchingService)

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	repository "github.com/hotstone-seo/hotstone-seo/server/repository"
	url "net/url"
	reflect "reflect"
)

// MockMetricsRuleMatchingService is a mock of MetricsRuleMatchingService interface
type MockMetricsRuleMatchingService struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsRuleMatchingServiceMockRecorder
}

// MockMetricsRuleMatchingServiceMockRecorder is the mock recorder for MockMetricsRuleMatchingService
type MockMetricsRuleMatchingServiceMockRecorder struct {
	mock *MockMetricsRuleMatchingService
}

// NewMockMetricsRuleMatchingService creates a new mock instance
func NewMockMetricsRuleMatchingService(ctrl *gomock.Controller) *MockMetricsRuleMatchingService {
	mock := &MockMetricsRuleMatchingService{ctrl: ctrl}
	mock.recorder = &MockMetricsRuleMatchingServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMetricsRuleMatchingService) EXPECT() *MockMetricsRuleMatchingServiceMockRecorder {
	return m.recorder
}

// CountMatched mocks base method
func (m *MockMetricsRuleMatchingService) CountMatched(arg0 context.Context, arg1 url.Values) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountMatched", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountMatched indicates an expected call of CountMatched
func (mr *MockMetricsRuleMatchingServiceMockRecorder) CountMatched(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountMatched", reflect.TypeOf((*MockMetricsRuleMatchingService)(nil).CountMatched), arg0, arg1)
}

// CountUniquePage mocks base method
func (m *MockMetricsRuleMatchingService) CountUniquePage(arg0 context.Context, arg1 url.Values) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountUniquePage", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountUniquePage indicates an expected call of CountUniquePage
func (mr *MockMetricsRuleMatchingServiceMockRecorder) CountUniquePage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountUniquePage", reflect.TypeOf((*MockMetricsRuleMatchingService)(nil).CountUniquePage), arg0, arg1)
}

// Insert mocks base method
func (m *MockMetricsRuleMatchingService) Insert(arg0 context.Context, arg1 repository.MetricsRuleMatching) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockMetricsRuleMatchingServiceMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockMetricsRuleMatchingService)(nil).Insert), arg0, arg1)
}

// ListCountHitPerDay mocks base method
func (m *MockMetricsRuleMatchingService) ListCountHitPerDay(arg0 context.Context, arg1, arg2, arg3 string) ([]*repository.MetricsCountHitPerDay, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCountHitPerDay", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*repository.MetricsCountHitPerDay)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCountHitPerDay indicates an expected call of ListCountHitPerDay
func (mr *MockMetricsRuleMatchingServiceMockRecorder) ListCountHitPerDay(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCountHitPerDay", reflect.TypeOf((*MockMetricsRuleMatchingService)(nil).ListCountHitPerDay), arg0, arg1, arg2, arg3)
}

// ListMismatchedCount mocks base method
func (m *MockMetricsRuleMatchingService) ListMismatchedCount(arg0 context.Context, arg1 repository.PaginationParam) ([]*repository.MetricsMismatchedCount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMismatchedCount", arg0, arg1)
	ret0, _ := ret[0].([]*repository.MetricsMismatchedCount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMismatchedCount indicates an expected call of ListMismatchedCount
func (mr *MockMetricsRuleMatchingServiceMockRecorder) ListMismatchedCount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMismatchedCount", reflect.TypeOf((*MockMetricsRuleMatchingService)(nil).ListMismatchedCount), arg0, arg1)
}

// SetMatched mocks base method
func (m *MockMetricsRuleMatchingService) SetMatched(arg0 *repository.MetricsRuleMatching, arg1 string, arg2 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMatched", arg0, arg1, arg2)
}

// SetMatched indicates an expected call of SetMatched
func (mr *MockMetricsRuleMatchingServiceMockRecorder) SetMatched(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMatched", reflect.TypeOf((*MockMetricsRuleMatchingService)(nil).SetMatched), arg0, arg1, arg2)
}

// SetMismatched mocks base method
func (m *MockMetricsRuleMatchingService) SetMismatched(arg0 *repository.MetricsRuleMatching, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMismatched", arg0, arg1)
}

// SetMismatched indicates an expected call of SetMismatched
func (mr *MockMetricsRuleMatchingServiceMockRecorder) SetMismatched(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMismatched", reflect.TypeOf((*MockMetricsRuleMatchingService)(nil).SetMismatched), arg0, arg1)
}
