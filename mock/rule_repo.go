// Code generated by MockGen. DO NOT EDIT.
// Source: app/repository/rule_repo.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	sql "database/sql"
	gomock "github.com/golang/mock/gomock"
	repository "github.com/hotstone-seo/hotstone-server/app/repository"
	reflect "reflect"
)

// MockRuleRepo is a mock of RuleRepo interface
type MockRuleRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRuleRepoMockRecorder
}

// MockRuleRepoMockRecorder is the mock recorder for MockRuleRepo
type MockRuleRepoMockRecorder struct {
	mock *MockRuleRepo
}

// NewMockRuleRepo creates a new mock instance
func NewMockRuleRepo(ctrl *gomock.Controller) *MockRuleRepo {
	mock := &MockRuleRepo{ctrl: ctrl}
	mock.recorder = &MockRuleRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRuleRepo) EXPECT() *MockRuleRepoMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockRuleRepo) Find(ctx context.Context, tx *sql.Tx, id int64) (*repository.Rule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, tx, id)
	ret0, _ := ret[0].(*repository.Rule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockRuleRepoMockRecorder) Find(ctx, tx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRuleRepo)(nil).Find), ctx, tx, id)
}

// List mocks base method
func (m *MockRuleRepo) List(ctx context.Context, tx *sql.Tx) ([]*repository.Rule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, tx)
	ret0, _ := ret[0].([]*repository.Rule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockRuleRepoMockRecorder) List(ctx, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRuleRepo)(nil).List), ctx, tx)
}

// Insert mocks base method
func (m *MockRuleRepo) Insert(ctx context.Context, tx *sql.Tx, rule repository.Rule) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, tx, rule)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockRuleRepoMockRecorder) Insert(ctx, tx, rule interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRuleRepo)(nil).Insert), ctx, tx, rule)
}

// Delete mocks base method
func (m *MockRuleRepo) Delete(ctx context.Context, tx *sql.Tx, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, tx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockRuleRepoMockRecorder) Delete(ctx, tx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRuleRepo)(nil).Delete), ctx, tx, id)
}

// Update mocks base method
func (m *MockRuleRepo) Update(ctx context.Context, tx *sql.Tx, rule repository.Rule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, tx, rule)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockRuleRepoMockRecorder) Update(ctx, tx, rule interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRuleRepo)(nil).Update), ctx, tx, rule)
}

// DB mocks base method
func (m *MockRuleRepo) DB() *sql.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DB")
	ret0, _ := ret[0].(*sql.DB)
	return ret0
}

// DB indicates an expected call of DB
func (mr *MockRuleRepoMockRecorder) DB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockRuleRepo)(nil).DB))
}
