// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hotstone-seo/hotstone-seo/server/service (interfaces: MetricService)

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	analyt "github.com/hotstone-seo/hotstone-seo/analyt"
	repository "github.com/hotstone-seo/hotstone-seo/server/repository"
	url "net/url"
	reflect "reflect"
)

// MockMetricService is a mock of MetricService interface
type MockMetricService struct {
	ctrl     *gomock.Controller
	recorder *MockMetricServiceMockRecorder
}

// MockMetricServiceMockRecorder is the mock recorder for MockMetricService
type MockMetricServiceMockRecorder struct {
	mock *MockMetricService
}

// NewMockMetricService creates a new mock instance
func NewMockMetricService(ctrl *gomock.Controller) *MockMetricService {
	mock := &MockMetricService{ctrl: ctrl}
	mock.recorder = &MockMetricServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMetricService) EXPECT() *MockMetricServiceMockRecorder {
	return m.recorder
}

// CountMatched mocks base method
func (m *MockMetricService) CountMatched(arg0 context.Context, arg1 url.Values) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountMatched", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountMatched indicates an expected call of CountMatched
func (mr *MockMetricServiceMockRecorder) CountMatched(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountMatched", reflect.TypeOf((*MockMetricService)(nil).CountMatched), arg0, arg1)
}

// CountUniquePage mocks base method
func (m *MockMetricService) CountUniquePage(arg0 context.Context, arg1 url.Values) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountUniquePage", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountUniquePage indicates an expected call of CountUniquePage
func (mr *MockMetricServiceMockRecorder) CountUniquePage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountUniquePage", reflect.TypeOf((*MockMetricService)(nil).CountUniquePage), arg0, arg1)
}

// DailyReports mocks base method
func (m *MockMetricService) DailyReports(arg0 context.Context, arg1, arg2, arg3 string) ([]*analyt.DailyReport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DailyReports", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*analyt.DailyReport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DailyReports indicates an expected call of DailyReports
func (mr *MockMetricServiceMockRecorder) DailyReports(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DailyReports", reflect.TypeOf((*MockMetricService)(nil).DailyReports), arg0, arg1, arg2, arg3)
}

// MismatchReports mocks base method
func (m *MockMetricService) MismatchReports(arg0 context.Context, arg1 repository.PaginationParam) ([]*analyt.MismatchReport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MismatchReports", arg0, arg1)
	ret0, _ := ret[0].([]*analyt.MismatchReport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MismatchReports indicates an expected call of MismatchReports
func (mr *MockMetricServiceMockRecorder) MismatchReports(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MismatchReports", reflect.TypeOf((*MockMetricService)(nil).MismatchReports), arg0, arg1)
}
