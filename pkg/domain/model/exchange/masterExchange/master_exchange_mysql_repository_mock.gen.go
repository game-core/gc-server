// Code generated by MockGen. DO NOT EDIT.
// Source: ./master_exchange_mysql_repository.gen.go

// Package masterExchange is a generated GoMock package.
package masterExchange

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockMasterExchangeMysqlRepository is a mock of MasterExchangeMysqlRepository interface.
type MockMasterExchangeMysqlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMasterExchangeMysqlRepositoryMockRecorder
}

// MockMasterExchangeMysqlRepositoryMockRecorder is the mock recorder for MockMasterExchangeMysqlRepository.
type MockMasterExchangeMysqlRepositoryMockRecorder struct {
	mock *MockMasterExchangeMysqlRepository
}

// NewMockMasterExchangeMysqlRepository creates a new mock instance.
func NewMockMasterExchangeMysqlRepository(ctrl *gomock.Controller) *MockMasterExchangeMysqlRepository {
	mock := &MockMasterExchangeMysqlRepository{ctrl: ctrl}
	mock.recorder = &MockMasterExchangeMysqlRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMasterExchangeMysqlRepository) EXPECT() *MockMasterExchangeMysqlRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockMasterExchangeMysqlRepository) Create(ctx context.Context, tx *gorm.DB, m *MasterExchange) (*MasterExchange, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*MasterExchange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMasterExchangeMysqlRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMasterExchangeMysqlRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockMasterExchangeMysqlRepository) CreateList(ctx context.Context, tx *gorm.DB, ms MasterExchanges) (MasterExchanges, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(MasterExchanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockMasterExchangeMysqlRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockMasterExchangeMysqlRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockMasterExchangeMysqlRepository) Delete(ctx context.Context, tx *gorm.DB, m *MasterExchange) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMasterExchangeMysqlRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMasterExchangeMysqlRepository)(nil).Delete), ctx, tx, m)
}

// DeleteList mocks base method.
func (m *MockMasterExchangeMysqlRepository) DeleteList(ctx context.Context, tx *gorm.DB, ms MasterExchanges) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteList", ctx, tx, ms)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteList indicates an expected call of DeleteList.
func (mr *MockMasterExchangeMysqlRepositoryMockRecorder) DeleteList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteList", reflect.TypeOf((*MockMasterExchangeMysqlRepository)(nil).DeleteList), ctx, tx, ms)
}

// Find mocks base method.
func (m *MockMasterExchangeMysqlRepository) Find(ctx context.Context, masterExchangeId int64) (*MasterExchange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, masterExchangeId)
	ret0, _ := ret[0].(*MasterExchange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMasterExchangeMysqlRepositoryMockRecorder) Find(ctx, masterExchangeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMasterExchangeMysqlRepository)(nil).Find), ctx, masterExchangeId)
}

// FindByMasterEventId mocks base method.
func (m *MockMasterExchangeMysqlRepository) FindByMasterEventId(ctx context.Context, masterEventId int64) (*MasterExchange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByMasterEventId", ctx, masterEventId)
	ret0, _ := ret[0].(*MasterExchange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByMasterEventId indicates an expected call of FindByMasterEventId.
func (mr *MockMasterExchangeMysqlRepositoryMockRecorder) FindByMasterEventId(ctx, masterEventId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMasterEventId", reflect.TypeOf((*MockMasterExchangeMysqlRepository)(nil).FindByMasterEventId), ctx, masterEventId)
}

// FindList mocks base method.
func (m *MockMasterExchangeMysqlRepository) FindList(ctx context.Context) (MasterExchanges, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx)
	ret0, _ := ret[0].(MasterExchanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockMasterExchangeMysqlRepositoryMockRecorder) FindList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockMasterExchangeMysqlRepository)(nil).FindList), ctx)
}

// FindListByMasterEventId mocks base method.
func (m *MockMasterExchangeMysqlRepository) FindListByMasterEventId(ctx context.Context, masterEventId int64) (MasterExchanges, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByMasterEventId", ctx, masterEventId)
	ret0, _ := ret[0].(MasterExchanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByMasterEventId indicates an expected call of FindListByMasterEventId.
func (mr *MockMasterExchangeMysqlRepositoryMockRecorder) FindListByMasterEventId(ctx, masterEventId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByMasterEventId", reflect.TypeOf((*MockMasterExchangeMysqlRepository)(nil).FindListByMasterEventId), ctx, masterEventId)
}

// FindOrNil mocks base method.
func (m *MockMasterExchangeMysqlRepository) FindOrNil(ctx context.Context, masterExchangeId int64) (*MasterExchange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, masterExchangeId)
	ret0, _ := ret[0].(*MasterExchange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockMasterExchangeMysqlRepositoryMockRecorder) FindOrNil(ctx, masterExchangeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockMasterExchangeMysqlRepository)(nil).FindOrNil), ctx, masterExchangeId)
}

// FindOrNilByMasterEventId mocks base method.
func (m *MockMasterExchangeMysqlRepository) FindOrNilByMasterEventId(ctx context.Context, masterEventId int64) (*MasterExchange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByMasterEventId", ctx, masterEventId)
	ret0, _ := ret[0].(*MasterExchange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByMasterEventId indicates an expected call of FindOrNilByMasterEventId.
func (mr *MockMasterExchangeMysqlRepositoryMockRecorder) FindOrNilByMasterEventId(ctx, masterEventId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByMasterEventId", reflect.TypeOf((*MockMasterExchangeMysqlRepository)(nil).FindOrNilByMasterEventId), ctx, masterEventId)
}

// Update mocks base method.
func (m_2 *MockMasterExchangeMysqlRepository) Update(ctx context.Context, tx *gorm.DB, m *MasterExchange) (*MasterExchange, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*MasterExchange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMasterExchangeMysqlRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMasterExchangeMysqlRepository)(nil).Update), ctx, tx, m)
}

// UpdateList mocks base method.
func (m *MockMasterExchangeMysqlRepository) UpdateList(ctx context.Context, tx *gorm.DB, ms MasterExchanges) (MasterExchanges, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateList", ctx, tx, ms)
	ret0, _ := ret[0].(MasterExchanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateList indicates an expected call of UpdateList.
func (mr *MockMasterExchangeMysqlRepositoryMockRecorder) UpdateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateList", reflect.TypeOf((*MockMasterExchangeMysqlRepository)(nil).UpdateList), ctx, tx, ms)
}