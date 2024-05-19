// Code generated by MockGen. DO NOT EDIT.
// Source: ./user_account_token_mysql_repository.gen.go

// Package userAccountToken is a generated GoMock package.
package userAccountToken

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockUserAccountTokenMysqlRepository is a mock of UserAccountTokenMysqlRepository interface.
type MockUserAccountTokenMysqlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserAccountTokenMysqlRepositoryMockRecorder
}

// MockUserAccountTokenMysqlRepositoryMockRecorder is the mock recorder for MockUserAccountTokenMysqlRepository.
type MockUserAccountTokenMysqlRepositoryMockRecorder struct {
	mock *MockUserAccountTokenMysqlRepository
}

// NewMockUserAccountTokenMysqlRepository creates a new mock instance.
func NewMockUserAccountTokenMysqlRepository(ctrl *gomock.Controller) *MockUserAccountTokenMysqlRepository {
	mock := &MockUserAccountTokenMysqlRepository{ctrl: ctrl}
	mock.recorder = &MockUserAccountTokenMysqlRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserAccountTokenMysqlRepository) EXPECT() *MockUserAccountTokenMysqlRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockUserAccountTokenMysqlRepository) Create(ctx context.Context, tx *gorm.DB, m *UserAccountToken) (*UserAccountToken, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*UserAccountToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserAccountTokenMysqlRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserAccountTokenMysqlRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockUserAccountTokenMysqlRepository) CreateList(ctx context.Context, tx *gorm.DB, ms UserAccountTokens) (UserAccountTokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(UserAccountTokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockUserAccountTokenMysqlRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockUserAccountTokenMysqlRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockUserAccountTokenMysqlRepository) Delete(ctx context.Context, tx *gorm.DB, m *UserAccountToken) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserAccountTokenMysqlRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserAccountTokenMysqlRepository)(nil).Delete), ctx, tx, m)
}

// DeleteList mocks base method.
func (m *MockUserAccountTokenMysqlRepository) DeleteList(ctx context.Context, tx *gorm.DB, ms UserAccountTokens) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteList", ctx, tx, ms)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteList indicates an expected call of DeleteList.
func (mr *MockUserAccountTokenMysqlRepositoryMockRecorder) DeleteList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteList", reflect.TypeOf((*MockUserAccountTokenMysqlRepository)(nil).DeleteList), ctx, tx, ms)
}

// Find mocks base method.
func (m *MockUserAccountTokenMysqlRepository) Find(ctx context.Context, userId string) (*UserAccountToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, userId)
	ret0, _ := ret[0].(*UserAccountToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockUserAccountTokenMysqlRepositoryMockRecorder) Find(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUserAccountTokenMysqlRepository)(nil).Find), ctx, userId)
}

// FindList mocks base method.
func (m *MockUserAccountTokenMysqlRepository) FindList(ctx context.Context, userId string) (UserAccountTokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx, userId)
	ret0, _ := ret[0].(UserAccountTokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockUserAccountTokenMysqlRepositoryMockRecorder) FindList(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockUserAccountTokenMysqlRepository)(nil).FindList), ctx, userId)
}

// FindOrNil mocks base method.
func (m *MockUserAccountTokenMysqlRepository) FindOrNil(ctx context.Context, userId string) (*UserAccountToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, userId)
	ret0, _ := ret[0].(*UserAccountToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockUserAccountTokenMysqlRepositoryMockRecorder) FindOrNil(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockUserAccountTokenMysqlRepository)(nil).FindOrNil), ctx, userId)
}

// Update mocks base method.
func (m_2 *MockUserAccountTokenMysqlRepository) Update(ctx context.Context, tx *gorm.DB, m *UserAccountToken) (*UserAccountToken, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*UserAccountToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUserAccountTokenMysqlRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserAccountTokenMysqlRepository)(nil).Update), ctx, tx, m)
}

// UpdateList mocks base method.
func (m *MockUserAccountTokenMysqlRepository) UpdateList(ctx context.Context, tx *gorm.DB, ms UserAccountTokens) (UserAccountTokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateList", ctx, tx, ms)
	ret0, _ := ret[0].(UserAccountTokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateList indicates an expected call of UpdateList.
func (mr *MockUserAccountTokenMysqlRepositoryMockRecorder) UpdateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateList", reflect.TypeOf((*MockUserAccountTokenMysqlRepository)(nil).UpdateList), ctx, tx, ms)
}
