// Code generated by MockGen. DO NOT EDIT.
// Source: ./master_resource_mysql_repository.gen.go

// Package masterResource is a generated GoMock package.
package masterResource

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockMasterResourceMysqlRepository is a mock of MasterResourceMysqlRepository interface.
type MockMasterResourceMysqlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMasterResourceMysqlRepositoryMockRecorder
}

// MockMasterResourceMysqlRepositoryMockRecorder is the mock recorder for MockMasterResourceMysqlRepository.
type MockMasterResourceMysqlRepositoryMockRecorder struct {
	mock *MockMasterResourceMysqlRepository
}

// NewMockMasterResourceMysqlRepository creates a new mock instance.
func NewMockMasterResourceMysqlRepository(ctrl *gomock.Controller) *MockMasterResourceMysqlRepository {
	mock := &MockMasterResourceMysqlRepository{ctrl: ctrl}
	mock.recorder = &MockMasterResourceMysqlRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMasterResourceMysqlRepository) EXPECT() *MockMasterResourceMysqlRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockMasterResourceMysqlRepository) Create(ctx context.Context, tx *gorm.DB, m *MasterResource) (*MasterResource, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*MasterResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMasterResourceMysqlRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMasterResourceMysqlRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockMasterResourceMysqlRepository) CreateList(ctx context.Context, tx *gorm.DB, ms MasterResources) (MasterResources, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(MasterResources)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockMasterResourceMysqlRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockMasterResourceMysqlRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockMasterResourceMysqlRepository) Delete(ctx context.Context, tx *gorm.DB, m *MasterResource) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMasterResourceMysqlRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMasterResourceMysqlRepository)(nil).Delete), ctx, tx, m)
}

// DeleteList mocks base method.
func (m *MockMasterResourceMysqlRepository) DeleteList(ctx context.Context, tx *gorm.DB, ms MasterResources) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteList", ctx, tx, ms)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteList indicates an expected call of DeleteList.
func (mr *MockMasterResourceMysqlRepositoryMockRecorder) DeleteList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteList", reflect.TypeOf((*MockMasterResourceMysqlRepository)(nil).DeleteList), ctx, tx, ms)
}

// Find mocks base method.
func (m *MockMasterResourceMysqlRepository) Find(ctx context.Context, masterResourceId int64) (*MasterResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, masterResourceId)
	ret0, _ := ret[0].(*MasterResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMasterResourceMysqlRepositoryMockRecorder) Find(ctx, masterResourceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMasterResourceMysqlRepository)(nil).Find), ctx, masterResourceId)
}

// FindByMasterResourceEnum mocks base method.
func (m *MockMasterResourceMysqlRepository) FindByMasterResourceEnum(ctx context.Context, masterResourceEnum MasterResourceEnum) (*MasterResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByMasterResourceEnum", ctx, masterResourceEnum)
	ret0, _ := ret[0].(*MasterResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByMasterResourceEnum indicates an expected call of FindByMasterResourceEnum.
func (mr *MockMasterResourceMysqlRepositoryMockRecorder) FindByMasterResourceEnum(ctx, masterResourceEnum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMasterResourceEnum", reflect.TypeOf((*MockMasterResourceMysqlRepository)(nil).FindByMasterResourceEnum), ctx, masterResourceEnum)
}

// FindList mocks base method.
func (m *MockMasterResourceMysqlRepository) FindList(ctx context.Context) (MasterResources, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx)
	ret0, _ := ret[0].(MasterResources)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockMasterResourceMysqlRepositoryMockRecorder) FindList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockMasterResourceMysqlRepository)(nil).FindList), ctx)
}

// FindListByMasterResourceEnum mocks base method.
func (m *MockMasterResourceMysqlRepository) FindListByMasterResourceEnum(ctx context.Context, masterResourceEnum MasterResourceEnum) (MasterResources, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByMasterResourceEnum", ctx, masterResourceEnum)
	ret0, _ := ret[0].(MasterResources)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByMasterResourceEnum indicates an expected call of FindListByMasterResourceEnum.
func (mr *MockMasterResourceMysqlRepositoryMockRecorder) FindListByMasterResourceEnum(ctx, masterResourceEnum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByMasterResourceEnum", reflect.TypeOf((*MockMasterResourceMysqlRepository)(nil).FindListByMasterResourceEnum), ctx, masterResourceEnum)
}

// FindOrNil mocks base method.
func (m *MockMasterResourceMysqlRepository) FindOrNil(ctx context.Context, masterResourceId int64) (*MasterResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, masterResourceId)
	ret0, _ := ret[0].(*MasterResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockMasterResourceMysqlRepositoryMockRecorder) FindOrNil(ctx, masterResourceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockMasterResourceMysqlRepository)(nil).FindOrNil), ctx, masterResourceId)
}

// FindOrNilByMasterResourceEnum mocks base method.
func (m *MockMasterResourceMysqlRepository) FindOrNilByMasterResourceEnum(ctx context.Context, masterResourceEnum MasterResourceEnum) (*MasterResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByMasterResourceEnum", ctx, masterResourceEnum)
	ret0, _ := ret[0].(*MasterResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByMasterResourceEnum indicates an expected call of FindOrNilByMasterResourceEnum.
func (mr *MockMasterResourceMysqlRepositoryMockRecorder) FindOrNilByMasterResourceEnum(ctx, masterResourceEnum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByMasterResourceEnum", reflect.TypeOf((*MockMasterResourceMysqlRepository)(nil).FindOrNilByMasterResourceEnum), ctx, masterResourceEnum)
}

// Update mocks base method.
func (m_2 *MockMasterResourceMysqlRepository) Update(ctx context.Context, tx *gorm.DB, m *MasterResource) (*MasterResource, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*MasterResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMasterResourceMysqlRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMasterResourceMysqlRepository)(nil).Update), ctx, tx, m)
}

// UpdateList mocks base method.
func (m *MockMasterResourceMysqlRepository) UpdateList(ctx context.Context, tx *gorm.DB, ms MasterResources) (MasterResources, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateList", ctx, tx, ms)
	ret0, _ := ret[0].(MasterResources)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateList indicates an expected call of UpdateList.
func (mr *MockMasterResourceMysqlRepositoryMockRecorder) UpdateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateList", reflect.TypeOf((*MockMasterResourceMysqlRepository)(nil).UpdateList), ctx, tx, ms)
}
