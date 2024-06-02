// Code generated by MockGen. DO NOT EDIT.
// Source: ./master_login_bonus_mysql_repository.gen.go

// Package masterLoginBonus is a generated GoMock package.
package masterLoginBonus

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockMasterLoginBonusMysqlRepository is a mock of MasterLoginBonusMysqlRepository interface.
type MockMasterLoginBonusMysqlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMasterLoginBonusMysqlRepositoryMockRecorder
}

// MockMasterLoginBonusMysqlRepositoryMockRecorder is the mock recorder for MockMasterLoginBonusMysqlRepository.
type MockMasterLoginBonusMysqlRepositoryMockRecorder struct {
	mock *MockMasterLoginBonusMysqlRepository
}

// NewMockMasterLoginBonusMysqlRepository creates a new mock instance.
func NewMockMasterLoginBonusMysqlRepository(ctrl *gomock.Controller) *MockMasterLoginBonusMysqlRepository {
	mock := &MockMasterLoginBonusMysqlRepository{ctrl: ctrl}
	mock.recorder = &MockMasterLoginBonusMysqlRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMasterLoginBonusMysqlRepository) EXPECT() *MockMasterLoginBonusMysqlRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockMasterLoginBonusMysqlRepository) Create(ctx context.Context, tx *gorm.DB, m *MasterLoginBonus) (*MasterLoginBonus, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*MasterLoginBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMasterLoginBonusMysqlRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMasterLoginBonusMysqlRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockMasterLoginBonusMysqlRepository) CreateList(ctx context.Context, tx *gorm.DB, ms MasterLoginBonuses) (MasterLoginBonuses, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(MasterLoginBonuses)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockMasterLoginBonusMysqlRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockMasterLoginBonusMysqlRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockMasterLoginBonusMysqlRepository) Delete(ctx context.Context, tx *gorm.DB, m *MasterLoginBonus) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMasterLoginBonusMysqlRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMasterLoginBonusMysqlRepository)(nil).Delete), ctx, tx, m)
}

// DeleteList mocks base method.
func (m *MockMasterLoginBonusMysqlRepository) DeleteList(ctx context.Context, tx *gorm.DB, ms MasterLoginBonuses) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteList", ctx, tx, ms)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteList indicates an expected call of DeleteList.
func (mr *MockMasterLoginBonusMysqlRepositoryMockRecorder) DeleteList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteList", reflect.TypeOf((*MockMasterLoginBonusMysqlRepository)(nil).DeleteList), ctx, tx, ms)
}

// Find mocks base method.
func (m *MockMasterLoginBonusMysqlRepository) Find(ctx context.Context, masterLoginBonusId int64) (*MasterLoginBonus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, masterLoginBonusId)
	ret0, _ := ret[0].(*MasterLoginBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMasterLoginBonusMysqlRepositoryMockRecorder) Find(ctx, masterLoginBonusId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMasterLoginBonusMysqlRepository)(nil).Find), ctx, masterLoginBonusId)
}

// FindByMasterEventId mocks base method.
func (m *MockMasterLoginBonusMysqlRepository) FindByMasterEventId(ctx context.Context, masterEventId int64) (*MasterLoginBonus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByMasterEventId", ctx, masterEventId)
	ret0, _ := ret[0].(*MasterLoginBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByMasterEventId indicates an expected call of FindByMasterEventId.
func (mr *MockMasterLoginBonusMysqlRepositoryMockRecorder) FindByMasterEventId(ctx, masterEventId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMasterEventId", reflect.TypeOf((*MockMasterLoginBonusMysqlRepository)(nil).FindByMasterEventId), ctx, masterEventId)
}

// FindList mocks base method.
func (m *MockMasterLoginBonusMysqlRepository) FindList(ctx context.Context) (MasterLoginBonuses, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx)
	ret0, _ := ret[0].(MasterLoginBonuses)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockMasterLoginBonusMysqlRepositoryMockRecorder) FindList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockMasterLoginBonusMysqlRepository)(nil).FindList), ctx)
}

// FindListByMasterEventId mocks base method.
func (m *MockMasterLoginBonusMysqlRepository) FindListByMasterEventId(ctx context.Context, masterEventId int64) (MasterLoginBonuses, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByMasterEventId", ctx, masterEventId)
	ret0, _ := ret[0].(MasterLoginBonuses)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByMasterEventId indicates an expected call of FindListByMasterEventId.
func (mr *MockMasterLoginBonusMysqlRepositoryMockRecorder) FindListByMasterEventId(ctx, masterEventId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByMasterEventId", reflect.TypeOf((*MockMasterLoginBonusMysqlRepository)(nil).FindListByMasterEventId), ctx, masterEventId)
}

// FindOrNil mocks base method.
func (m *MockMasterLoginBonusMysqlRepository) FindOrNil(ctx context.Context, masterLoginBonusId int64) (*MasterLoginBonus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, masterLoginBonusId)
	ret0, _ := ret[0].(*MasterLoginBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockMasterLoginBonusMysqlRepositoryMockRecorder) FindOrNil(ctx, masterLoginBonusId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockMasterLoginBonusMysqlRepository)(nil).FindOrNil), ctx, masterLoginBonusId)
}

// FindOrNilByMasterEventId mocks base method.
func (m *MockMasterLoginBonusMysqlRepository) FindOrNilByMasterEventId(ctx context.Context, masterEventId int64) (*MasterLoginBonus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByMasterEventId", ctx, masterEventId)
	ret0, _ := ret[0].(*MasterLoginBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByMasterEventId indicates an expected call of FindOrNilByMasterEventId.
func (mr *MockMasterLoginBonusMysqlRepositoryMockRecorder) FindOrNilByMasterEventId(ctx, masterEventId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByMasterEventId", reflect.TypeOf((*MockMasterLoginBonusMysqlRepository)(nil).FindOrNilByMasterEventId), ctx, masterEventId)
}

// Update mocks base method.
func (m_2 *MockMasterLoginBonusMysqlRepository) Update(ctx context.Context, tx *gorm.DB, m *MasterLoginBonus) (*MasterLoginBonus, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*MasterLoginBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMasterLoginBonusMysqlRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMasterLoginBonusMysqlRepository)(nil).Update), ctx, tx, m)
}

// UpdateList mocks base method.
func (m *MockMasterLoginBonusMysqlRepository) UpdateList(ctx context.Context, tx *gorm.DB, ms MasterLoginBonuses) (MasterLoginBonuses, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateList", ctx, tx, ms)
	ret0, _ := ret[0].(MasterLoginBonuses)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateList indicates an expected call of UpdateList.
func (mr *MockMasterLoginBonusMysqlRepositoryMockRecorder) UpdateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateList", reflect.TypeOf((*MockMasterLoginBonusMysqlRepository)(nil).UpdateList), ctx, tx, ms)
}
