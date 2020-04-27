// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hotstone-seo/hotstone-seo/server/service (interfaces: CenterService)

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	repository "github.com/hotstone-seo/hotstone-seo/server/repository"
	service "github.com/hotstone-seo/hotstone-seo/server/service"
	reflect "reflect"
)

// MockCenterService is a mock of CenterService interface
type MockCenterService struct {
	ctrl     *gomock.Controller
	recorder *MockCenterServiceMockRecorder
}

// MockCenterServiceMockRecorder is the mock recorder for MockCenterService
type MockCenterServiceMockRecorder struct {
	mock *MockCenterService
}

// NewMockCenterService creates a new mock instance
func NewMockCenterService(ctrl *gomock.Controller) *MockCenterService {
	mock := &MockCenterService{ctrl: ctrl}
	mock.recorder = &MockCenterServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCenterService) EXPECT() *MockCenterServiceMockRecorder {
	return m.recorder
}

// AddBreadcrumbList mocks base method
func (m *MockCenterService) AddBreadcrumbList(arg0 context.Context, arg1 service.BreadcrumbListRequest) (*repository.StructuredData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBreadcrumbList", arg0, arg1)
	ret0, _ := ret[0].(*repository.StructuredData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBreadcrumbList indicates an expected call of AddBreadcrumbList
func (mr *MockCenterServiceMockRecorder) AddBreadcrumbList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBreadcrumbList", reflect.TypeOf((*MockCenterService)(nil).AddBreadcrumbList), arg0, arg1)
}

// AddCanonicalTag mocks base method
func (m *MockCenterService) AddCanonicalTag(arg0 context.Context, arg1 service.CanonicalTagRequest) (*repository.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCanonicalTag", arg0, arg1)
	ret0, _ := ret[0].(*repository.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCanonicalTag indicates an expected call of AddCanonicalTag
func (mr *MockCenterServiceMockRecorder) AddCanonicalTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCanonicalTag", reflect.TypeOf((*MockCenterService)(nil).AddCanonicalTag), arg0, arg1)
}

// AddFAQPage mocks base method
func (m *MockCenterService) AddFAQPage(arg0 context.Context, arg1 service.FAQPageRequest) (*repository.StructuredData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFAQPage", arg0, arg1)
	ret0, _ := ret[0].(*repository.StructuredData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddFAQPage indicates an expected call of AddFAQPage
func (mr *MockCenterServiceMockRecorder) AddFAQPage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFAQPage", reflect.TypeOf((*MockCenterService)(nil).AddFAQPage), arg0, arg1)
}

// AddMetaTag mocks base method
func (m *MockCenterService) AddMetaTag(arg0 context.Context, arg1 service.MetaTagRequest) (*repository.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMetaTag", arg0, arg1)
	ret0, _ := ret[0].(*repository.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMetaTag indicates an expected call of AddMetaTag
func (mr *MockCenterServiceMockRecorder) AddMetaTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMetaTag", reflect.TypeOf((*MockCenterService)(nil).AddMetaTag), arg0, arg1)
}

// AddScriptTag mocks base method
func (m *MockCenterService) AddScriptTag(arg0 context.Context, arg1 service.ScriptTagRequest) (*repository.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddScriptTag", arg0, arg1)
	ret0, _ := ret[0].(*repository.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddScriptTag indicates an expected call of AddScriptTag
func (mr *MockCenterServiceMockRecorder) AddScriptTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddScriptTag", reflect.TypeOf((*MockCenterService)(nil).AddScriptTag), arg0, arg1)
}

// AddTitleTag mocks base method
func (m *MockCenterService) AddTitleTag(arg0 context.Context, arg1 service.TitleTagRequest) (*repository.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTitleTag", arg0, arg1)
	ret0, _ := ret[0].(*repository.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTitleTag indicates an expected call of AddTitleTag
func (mr *MockCenterServiceMockRecorder) AddTitleTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTitleTag", reflect.TypeOf((*MockCenterService)(nil).AddTitleTag), arg0, arg1)
}

// UpdateBreadcrumbList mocks base method
func (m *MockCenterService) UpdateBreadcrumbList(arg0 context.Context, arg1 service.BreadcrumbListRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBreadcrumbList", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBreadcrumbList indicates an expected call of UpdateBreadcrumbList
func (mr *MockCenterServiceMockRecorder) UpdateBreadcrumbList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBreadcrumbList", reflect.TypeOf((*MockCenterService)(nil).UpdateBreadcrumbList), arg0, arg1)
}

// UpdateCanonicalTag mocks base method
func (m *MockCenterService) UpdateCanonicalTag(arg0 context.Context, arg1 service.CanonicalTagRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCanonicalTag", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCanonicalTag indicates an expected call of UpdateCanonicalTag
func (mr *MockCenterServiceMockRecorder) UpdateCanonicalTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCanonicalTag", reflect.TypeOf((*MockCenterService)(nil).UpdateCanonicalTag), arg0, arg1)
}

// UpdateFAQPage mocks base method
func (m *MockCenterService) UpdateFAQPage(arg0 context.Context, arg1 service.FAQPageRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFAQPage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFAQPage indicates an expected call of UpdateFAQPage
func (mr *MockCenterServiceMockRecorder) UpdateFAQPage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFAQPage", reflect.TypeOf((*MockCenterService)(nil).UpdateFAQPage), arg0, arg1)
}

// UpdateMetaTag mocks base method
func (m *MockCenterService) UpdateMetaTag(arg0 context.Context, arg1 service.MetaTagRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMetaTag", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMetaTag indicates an expected call of UpdateMetaTag
func (mr *MockCenterServiceMockRecorder) UpdateMetaTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMetaTag", reflect.TypeOf((*MockCenterService)(nil).UpdateMetaTag), arg0, arg1)
}

// UpdateScriptTag mocks base method
func (m *MockCenterService) UpdateScriptTag(arg0 context.Context, arg1 service.ScriptTagRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateScriptTag", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateScriptTag indicates an expected call of UpdateScriptTag
func (mr *MockCenterServiceMockRecorder) UpdateScriptTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScriptTag", reflect.TypeOf((*MockCenterService)(nil).UpdateScriptTag), arg0, arg1)
}

// UpdateTitleTag mocks base method
func (m *MockCenterService) UpdateTitleTag(arg0 context.Context, arg1 service.TitleTagRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTitleTag", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTitleTag indicates an expected call of UpdateTitleTag
func (mr *MockCenterServiceMockRecorder) UpdateTitleTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTitleTag", reflect.TypeOf((*MockCenterService)(nil).UpdateTitleTag), arg0, arg1)
}
