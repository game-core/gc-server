// Code generated by MockGen. DO NOT EDIT.
// Source: ./admin_health_mysql_repository.gen.go

// Package adminHealth is a generated GoMock package.
package adminHealth

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockAdminHealthMysqlRepository is a mock of AdminHealthMysqlRepository interface.
type MockAdminHealthMysqlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAdminHealthMysqlRepositoryMockRecorder
}

// MockAdminHealthMysqlRepositoryMockRecorder is the mock recorder for MockAdminHealthMysqlRepository.
type MockAdminHealthMysqlRepositoryMockRecorder struct {
	mock *MockAdminHealthMysqlRepository
}

// NewMockAdminHealthMysqlRepository creates a new mock instance.
func NewMockAdminHealthMysqlRepository(ctrl *gomock.Controller) *MockAdminHealthMysqlRepository {
	mock := &MockAdminHealthMysqlRepository{ctrl: ctrl}
	mock.recorder = &MockAdminHealthMysqlRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminHealthMysqlRepository) EXPECT() *MockAdminHealthMysqlRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockAdminHealthMysqlRepository) Create(ctx context.Context, tx *gorm.DB, m *AdminHealth) (*AdminHealth, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*AdminHealth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAdminHealthMysqlRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAdminHealthMysqlRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockAdminHealthMysqlRepository) CreateList(ctx context.Context, tx *gorm.DB, ms AdminHealths) (AdminHealths, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(AdminHealths)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockAdminHealthMysqlRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockAdminHealthMysqlRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockAdminHealthMysqlRepository) Delete(ctx context.Context, tx *gorm.DB, m *AdminHealth) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAdminHealthMysqlRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAdminHealthMysqlRepository)(nil).Delete), ctx, tx, m)
}

// DeleteList mocks base method.
func (m *MockAdminHealthMysqlRepository) DeleteList(ctx context.Context, tx *gorm.DB, ms AdminHealths) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteList", ctx, tx, ms)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteList indicates an expected call of DeleteList.
func (mr *MockAdminHealthMysqlRepositoryMockRecorder) DeleteList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteList", reflect.TypeOf((*MockAdminHealthMysqlRepository)(nil).DeleteList), ctx, tx, ms)
}

// Find mocks base method.
func (m *MockAdminHealthMysqlRepository) Find(ctx context.Context, healthId int64) (*AdminHealth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, healthId)
	ret0, _ := ret[0].(*AdminHealth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockAdminHealthMysqlRepositoryMockRecorder) Find(ctx, healthId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockAdminHealthMysqlRepository)(nil).Find), ctx, healthId)
}

// FindList mocks base method.
func (m *MockAdminHealthMysqlRepository) FindList(ctx context.Context) (AdminHealths, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx)
	ret0, _ := ret[0].(AdminHealths)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockAdminHealthMysqlRepositoryMockRecorder) FindList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockAdminHealthMysqlRepository)(nil).FindList), ctx)
}

// FindOrNil mocks base method.
func (m *MockAdminHealthMysqlRepository) FindOrNil(ctx context.Context, healthId int64) (*AdminHealth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, healthId)
	ret0, _ := ret[0].(*AdminHealth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockAdminHealthMysqlRepositoryMockRecorder) FindOrNil(ctx, healthId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockAdminHealthMysqlRepository)(nil).FindOrNil), ctx, healthId)
}

// Update mocks base method.
func (m_2 *MockAdminHealthMysqlRepository) Update(ctx context.Context, tx *gorm.DB, m *AdminHealth) (*AdminHealth, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*AdminHealth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAdminHealthMysqlRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAdminHealthMysqlRepository)(nil).Update), ctx, tx, m)
}

// UpdateList mocks base method.
func (m *MockAdminHealthMysqlRepository) UpdateList(ctx context.Context, tx *gorm.DB, ms AdminHealths) (AdminHealths, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateList", ctx, tx, ms)
	ret0, _ := ret[0].(AdminHealths)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateList indicates an expected call of UpdateList.
func (mr *MockAdminHealthMysqlRepositoryMockRecorder) UpdateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateList", reflect.TypeOf((*MockAdminHealthMysqlRepository)(nil).UpdateList), ctx, tx, ms)
}
