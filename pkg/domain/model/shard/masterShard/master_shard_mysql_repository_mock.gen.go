// Code generated by MockGen. DO NOT EDIT.
// Source: ./master_shard_mysql_repository.gen.go

// Package masterShard is a generated GoMock package.
package masterShard

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockMasterShardMysqlRepository is a mock of MasterShardMysqlRepository interface.
type MockMasterShardMysqlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMasterShardMysqlRepositoryMockRecorder
}

// MockMasterShardMysqlRepositoryMockRecorder is the mock recorder for MockMasterShardMysqlRepository.
type MockMasterShardMysqlRepositoryMockRecorder struct {
	mock *MockMasterShardMysqlRepository
}

// NewMockMasterShardMysqlRepository creates a new mock instance.
func NewMockMasterShardMysqlRepository(ctrl *gomock.Controller) *MockMasterShardMysqlRepository {
	mock := &MockMasterShardMysqlRepository{ctrl: ctrl}
	mock.recorder = &MockMasterShardMysqlRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMasterShardMysqlRepository) EXPECT() *MockMasterShardMysqlRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockMasterShardMysqlRepository) Create(ctx context.Context, tx *gorm.DB, m *MasterShard) (*MasterShard, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*MasterShard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMasterShardMysqlRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMasterShardMysqlRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockMasterShardMysqlRepository) CreateList(ctx context.Context, tx *gorm.DB, ms MasterShards) (MasterShards, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(MasterShards)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockMasterShardMysqlRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockMasterShardMysqlRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockMasterShardMysqlRepository) Delete(ctx context.Context, tx *gorm.DB, m *MasterShard) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMasterShardMysqlRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMasterShardMysqlRepository)(nil).Delete), ctx, tx, m)
}

// Find mocks base method.
func (m *MockMasterShardMysqlRepository) Find(ctx context.Context, masterShardId int64) (*MasterShard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, masterShardId)
	ret0, _ := ret[0].(*MasterShard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMasterShardMysqlRepositoryMockRecorder) Find(ctx, masterShardId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMasterShardMysqlRepository)(nil).Find), ctx, masterShardId)
}

// FindByShardKey mocks base method.
func (m *MockMasterShardMysqlRepository) FindByShardKey(ctx context.Context, shardKey string) (*MasterShard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByShardKey", ctx, shardKey)
	ret0, _ := ret[0].(*MasterShard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByShardKey indicates an expected call of FindByShardKey.
func (mr *MockMasterShardMysqlRepositoryMockRecorder) FindByShardKey(ctx, shardKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByShardKey", reflect.TypeOf((*MockMasterShardMysqlRepository)(nil).FindByShardKey), ctx, shardKey)
}

// FindList mocks base method.
func (m *MockMasterShardMysqlRepository) FindList(ctx context.Context) (MasterShards, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx)
	ret0, _ := ret[0].(MasterShards)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockMasterShardMysqlRepositoryMockRecorder) FindList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockMasterShardMysqlRepository)(nil).FindList), ctx)
}

// FindListByShardKey mocks base method.
func (m *MockMasterShardMysqlRepository) FindListByShardKey(ctx context.Context, shardKey string) (MasterShards, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByShardKey", ctx, shardKey)
	ret0, _ := ret[0].(MasterShards)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByShardKey indicates an expected call of FindListByShardKey.
func (mr *MockMasterShardMysqlRepositoryMockRecorder) FindListByShardKey(ctx, shardKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByShardKey", reflect.TypeOf((*MockMasterShardMysqlRepository)(nil).FindListByShardKey), ctx, shardKey)
}

// FindOrNil mocks base method.
func (m *MockMasterShardMysqlRepository) FindOrNil(ctx context.Context, masterShardId int64) (*MasterShard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, masterShardId)
	ret0, _ := ret[0].(*MasterShard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockMasterShardMysqlRepositoryMockRecorder) FindOrNil(ctx, masterShardId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockMasterShardMysqlRepository)(nil).FindOrNil), ctx, masterShardId)
}

// FindOrNilByShardKey mocks base method.
func (m *MockMasterShardMysqlRepository) FindOrNilByShardKey(ctx context.Context, shardKey string) (*MasterShard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByShardKey", ctx, shardKey)
	ret0, _ := ret[0].(*MasterShard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByShardKey indicates an expected call of FindOrNilByShardKey.
func (mr *MockMasterShardMysqlRepositoryMockRecorder) FindOrNilByShardKey(ctx, shardKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByShardKey", reflect.TypeOf((*MockMasterShardMysqlRepository)(nil).FindOrNilByShardKey), ctx, shardKey)
}

// Update mocks base method.
func (m_2 *MockMasterShardMysqlRepository) Update(ctx context.Context, tx *gorm.DB, m *MasterShard) (*MasterShard, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*MasterShard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMasterShardMysqlRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMasterShardMysqlRepository)(nil).Update), ctx, tx, m)
}