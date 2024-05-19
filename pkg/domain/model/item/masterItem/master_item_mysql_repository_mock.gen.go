// Code generated by MockGen. DO NOT EDIT.
// Source: ./master_item_mysql_repository.gen.go

// Package masterItem is a generated GoMock package.
package masterItem

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockMasterItemMysqlRepository is a mock of MasterItemMysqlRepository interface.
type MockMasterItemMysqlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMasterItemMysqlRepositoryMockRecorder
}

// MockMasterItemMysqlRepositoryMockRecorder is the mock recorder for MockMasterItemMysqlRepository.
type MockMasterItemMysqlRepositoryMockRecorder struct {
	mock *MockMasterItemMysqlRepository
}

// NewMockMasterItemMysqlRepository creates a new mock instance.
func NewMockMasterItemMysqlRepository(ctrl *gomock.Controller) *MockMasterItemMysqlRepository {
	mock := &MockMasterItemMysqlRepository{ctrl: ctrl}
	mock.recorder = &MockMasterItemMysqlRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMasterItemMysqlRepository) EXPECT() *MockMasterItemMysqlRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockMasterItemMysqlRepository) Create(ctx context.Context, tx *gorm.DB, m *MasterItem) (*MasterItem, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*MasterItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMasterItemMysqlRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMasterItemMysqlRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockMasterItemMysqlRepository) CreateList(ctx context.Context, tx *gorm.DB, ms MasterItems) (MasterItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(MasterItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockMasterItemMysqlRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockMasterItemMysqlRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockMasterItemMysqlRepository) Delete(ctx context.Context, tx *gorm.DB, m *MasterItem) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMasterItemMysqlRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMasterItemMysqlRepository)(nil).Delete), ctx, tx, m)
}

// DeleteList mocks base method.
func (m *MockMasterItemMysqlRepository) DeleteList(ctx context.Context, tx *gorm.DB, ms MasterItems) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteList", ctx, tx, ms)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteList indicates an expected call of DeleteList.
func (mr *MockMasterItemMysqlRepositoryMockRecorder) DeleteList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteList", reflect.TypeOf((*MockMasterItemMysqlRepository)(nil).DeleteList), ctx, tx, ms)
}

// Find mocks base method.
func (m *MockMasterItemMysqlRepository) Find(ctx context.Context, masterItemId int64) (*MasterItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, masterItemId)
	ret0, _ := ret[0].(*MasterItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMasterItemMysqlRepositoryMockRecorder) Find(ctx, masterItemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMasterItemMysqlRepository)(nil).Find), ctx, masterItemId)
}

// FindList mocks base method.
func (m *MockMasterItemMysqlRepository) FindList(ctx context.Context) (MasterItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx)
	ret0, _ := ret[0].(MasterItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockMasterItemMysqlRepositoryMockRecorder) FindList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockMasterItemMysqlRepository)(nil).FindList), ctx)
}

// FindOrNil mocks base method.
func (m *MockMasterItemMysqlRepository) FindOrNil(ctx context.Context, masterItemId int64) (*MasterItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, masterItemId)
	ret0, _ := ret[0].(*MasterItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockMasterItemMysqlRepositoryMockRecorder) FindOrNil(ctx, masterItemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockMasterItemMysqlRepository)(nil).FindOrNil), ctx, masterItemId)
}

// Update mocks base method.
func (m_2 *MockMasterItemMysqlRepository) Update(ctx context.Context, tx *gorm.DB, m *MasterItem) (*MasterItem, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*MasterItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMasterItemMysqlRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMasterItemMysqlRepository)(nil).Update), ctx, tx, m)
}

// UpdateList mocks base method.
func (m *MockMasterItemMysqlRepository) UpdateList(ctx context.Context, tx *gorm.DB, ms MasterItems) (MasterItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateList", ctx, tx, ms)
	ret0, _ := ret[0].(MasterItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateList indicates an expected call of UpdateList.
func (mr *MockMasterItemMysqlRepositoryMockRecorder) UpdateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateList", reflect.TypeOf((*MockMasterItemMysqlRepository)(nil).UpdateList), ctx, tx, ms)
}
