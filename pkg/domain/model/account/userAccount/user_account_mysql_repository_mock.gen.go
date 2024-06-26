// Code generated by MockGen. DO NOT EDIT.
// Source: ./user_account_mysql_repository.gen.go

// Package userAccount is a generated GoMock package.
package userAccount

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockUserAccountMysqlRepository is a mock of UserAccountMysqlRepository interface.
type MockUserAccountMysqlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserAccountMysqlRepositoryMockRecorder
}

// MockUserAccountMysqlRepositoryMockRecorder is the mock recorder for MockUserAccountMysqlRepository.
type MockUserAccountMysqlRepositoryMockRecorder struct {
	mock *MockUserAccountMysqlRepository
}

// NewMockUserAccountMysqlRepository creates a new mock instance.
func NewMockUserAccountMysqlRepository(ctrl *gomock.Controller) *MockUserAccountMysqlRepository {
	mock := &MockUserAccountMysqlRepository{ctrl: ctrl}
	mock.recorder = &MockUserAccountMysqlRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserAccountMysqlRepository) EXPECT() *MockUserAccountMysqlRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockUserAccountMysqlRepository) Create(ctx context.Context, tx *gorm.DB, m *UserAccount) (*UserAccount, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*UserAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserAccountMysqlRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserAccountMysqlRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockUserAccountMysqlRepository) CreateList(ctx context.Context, tx *gorm.DB, ms UserAccounts) (UserAccounts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(UserAccounts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockUserAccountMysqlRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockUserAccountMysqlRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockUserAccountMysqlRepository) Delete(ctx context.Context, tx *gorm.DB, m *UserAccount) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserAccountMysqlRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserAccountMysqlRepository)(nil).Delete), ctx, tx, m)
}

// DeleteList mocks base method.
func (m *MockUserAccountMysqlRepository) DeleteList(ctx context.Context, tx *gorm.DB, ms UserAccounts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteList", ctx, tx, ms)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteList indicates an expected call of DeleteList.
func (mr *MockUserAccountMysqlRepositoryMockRecorder) DeleteList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteList", reflect.TypeOf((*MockUserAccountMysqlRepository)(nil).DeleteList), ctx, tx, ms)
}

// Find mocks base method.
func (m *MockUserAccountMysqlRepository) Find(ctx context.Context, userId string) (*UserAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, userId)
	ret0, _ := ret[0].(*UserAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockUserAccountMysqlRepositoryMockRecorder) Find(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUserAccountMysqlRepository)(nil).Find), ctx, userId)
}

// FindList mocks base method.
func (m *MockUserAccountMysqlRepository) FindList(ctx context.Context, userId string) (UserAccounts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx, userId)
	ret0, _ := ret[0].(UserAccounts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockUserAccountMysqlRepositoryMockRecorder) FindList(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockUserAccountMysqlRepository)(nil).FindList), ctx, userId)
}

// FindOrNil mocks base method.
func (m *MockUserAccountMysqlRepository) FindOrNil(ctx context.Context, userId string) (*UserAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, userId)
	ret0, _ := ret[0].(*UserAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockUserAccountMysqlRepositoryMockRecorder) FindOrNil(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockUserAccountMysqlRepository)(nil).FindOrNil), ctx, userId)
}

// Update mocks base method.
func (m_2 *MockUserAccountMysqlRepository) Update(ctx context.Context, tx *gorm.DB, m *UserAccount) (*UserAccount, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*UserAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUserAccountMysqlRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserAccountMysqlRepository)(nil).Update), ctx, tx, m)
}

// UpdateList mocks base method.
func (m *MockUserAccountMysqlRepository) UpdateList(ctx context.Context, tx *gorm.DB, ms UserAccounts) (UserAccounts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateList", ctx, tx, ms)
	ret0, _ := ret[0].(UserAccounts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateList indicates an expected call of UpdateList.
func (mr *MockUserAccountMysqlRepositoryMockRecorder) UpdateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateList", reflect.TypeOf((*MockUserAccountMysqlRepository)(nil).UpdateList), ctx, tx, ms)
}
