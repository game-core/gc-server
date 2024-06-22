// Code generated by MockGen. DO NOT EDIT.
// Source: ./user_exchange_mysql_repository.gen.go

// Package userExchange is a generated GoMock package.
package userExchange

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockUserExchangeMysqlRepository is a mock of UserExchangeMysqlRepository interface.
type MockUserExchangeMysqlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserExchangeMysqlRepositoryMockRecorder
}

// MockUserExchangeMysqlRepositoryMockRecorder is the mock recorder for MockUserExchangeMysqlRepository.
type MockUserExchangeMysqlRepositoryMockRecorder struct {
	mock *MockUserExchangeMysqlRepository
}

// NewMockUserExchangeMysqlRepository creates a new mock instance.
func NewMockUserExchangeMysqlRepository(ctrl *gomock.Controller) *MockUserExchangeMysqlRepository {
	mock := &MockUserExchangeMysqlRepository{ctrl: ctrl}
	mock.recorder = &MockUserExchangeMysqlRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserExchangeMysqlRepository) EXPECT() *MockUserExchangeMysqlRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockUserExchangeMysqlRepository) Create(ctx context.Context, tx *gorm.DB, m *UserExchange) (*UserExchange, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*UserExchange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserExchangeMysqlRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserExchangeMysqlRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockUserExchangeMysqlRepository) CreateList(ctx context.Context, tx *gorm.DB, ms UserExchanges) (UserExchanges, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(UserExchanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockUserExchangeMysqlRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockUserExchangeMysqlRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockUserExchangeMysqlRepository) Delete(ctx context.Context, tx *gorm.DB, m *UserExchange) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserExchangeMysqlRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserExchangeMysqlRepository)(nil).Delete), ctx, tx, m)
}

// DeleteList mocks base method.
func (m *MockUserExchangeMysqlRepository) DeleteList(ctx context.Context, tx *gorm.DB, ms UserExchanges) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteList", ctx, tx, ms)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteList indicates an expected call of DeleteList.
func (mr *MockUserExchangeMysqlRepositoryMockRecorder) DeleteList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteList", reflect.TypeOf((*MockUserExchangeMysqlRepository)(nil).DeleteList), ctx, tx, ms)
}

// Find mocks base method.
func (m *MockUserExchangeMysqlRepository) Find(ctx context.Context, userId string, masterExchangeId int64) (*UserExchange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, userId, masterExchangeId)
	ret0, _ := ret[0].(*UserExchange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockUserExchangeMysqlRepositoryMockRecorder) Find(ctx, userId, masterExchangeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUserExchangeMysqlRepository)(nil).Find), ctx, userId, masterExchangeId)
}

// FindList mocks base method.
func (m *MockUserExchangeMysqlRepository) FindList(ctx context.Context, userId string) (UserExchanges, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx, userId)
	ret0, _ := ret[0].(UserExchanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockUserExchangeMysqlRepositoryMockRecorder) FindList(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockUserExchangeMysqlRepository)(nil).FindList), ctx, userId)
}

// FindOrNil mocks base method.
func (m *MockUserExchangeMysqlRepository) FindOrNil(ctx context.Context, userId string, masterExchangeId int64) (*UserExchange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, userId, masterExchangeId)
	ret0, _ := ret[0].(*UserExchange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockUserExchangeMysqlRepositoryMockRecorder) FindOrNil(ctx, userId, masterExchangeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockUserExchangeMysqlRepository)(nil).FindOrNil), ctx, userId, masterExchangeId)
}

// Update mocks base method.
func (m_2 *MockUserExchangeMysqlRepository) Update(ctx context.Context, tx *gorm.DB, m *UserExchange) (*UserExchange, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*UserExchange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUserExchangeMysqlRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserExchangeMysqlRepository)(nil).Update), ctx, tx, m)
}

// UpdateList mocks base method.
func (m *MockUserExchangeMysqlRepository) UpdateList(ctx context.Context, tx *gorm.DB, ms UserExchanges) (UserExchanges, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateList", ctx, tx, ms)
	ret0, _ := ret[0].(UserExchanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateList indicates an expected call of UpdateList.
func (mr *MockUserExchangeMysqlRepositoryMockRecorder) UpdateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateList", reflect.TypeOf((*MockUserExchangeMysqlRepository)(nil).UpdateList), ctx, tx, ms)
}