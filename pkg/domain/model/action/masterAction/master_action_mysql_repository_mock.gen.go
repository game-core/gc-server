// Code generated by MockGen. DO NOT EDIT.
// Source: ./master_action_mysql_repository.gen.go

// Package masterAction is a generated GoMock package.
package masterAction

import (
	context "context"
	reflect "reflect"

	masterActionStep "github.com/game-core/gc-server/pkg/domain/model/action/masterActionStep"
	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockMasterActionMysqlRepository is a mock of MasterActionMysqlRepository interface.
type MockMasterActionMysqlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMasterActionMysqlRepositoryMockRecorder
}

// MockMasterActionMysqlRepositoryMockRecorder is the mock recorder for MockMasterActionMysqlRepository.
type MockMasterActionMysqlRepositoryMockRecorder struct {
	mock *MockMasterActionMysqlRepository
}

// NewMockMasterActionMysqlRepository creates a new mock instance.
func NewMockMasterActionMysqlRepository(ctrl *gomock.Controller) *MockMasterActionMysqlRepository {
	mock := &MockMasterActionMysqlRepository{ctrl: ctrl}
	mock.recorder = &MockMasterActionMysqlRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMasterActionMysqlRepository) EXPECT() *MockMasterActionMysqlRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockMasterActionMysqlRepository) Create(ctx context.Context, tx *gorm.DB, m *MasterAction) (*MasterAction, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*MasterAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockMasterActionMysqlRepository) CreateList(ctx context.Context, tx *gorm.DB, ms MasterActions) (MasterActions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(MasterActions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockMasterActionMysqlRepository) Delete(ctx context.Context, tx *gorm.DB, m *MasterAction) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).Delete), ctx, tx, m)
}

// Find mocks base method.
func (m *MockMasterActionMysqlRepository) Find(ctx context.Context, masterActionId int64) (*MasterAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, masterActionId)
	ret0, _ := ret[0].(*MasterAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) Find(ctx, masterActionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).Find), ctx, masterActionId)
}

// FindByMasterActionStepEnum mocks base method.
func (m *MockMasterActionMysqlRepository) FindByMasterActionStepEnum(ctx context.Context, masterActionStepEnum masterActionStep.MasterActionStepEnum) (*MasterAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByMasterActionStepEnum", ctx, masterActionStepEnum)
	ret0, _ := ret[0].(*MasterAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByMasterActionStepEnum indicates an expected call of FindByMasterActionStepEnum.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindByMasterActionStepEnum(ctx, masterActionStepEnum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMasterActionStepEnum", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindByMasterActionStepEnum), ctx, masterActionStepEnum)
}

// FindByMasterActionStepEnumAndTargetId mocks base method.
func (m *MockMasterActionMysqlRepository) FindByMasterActionStepEnumAndTargetId(ctx context.Context, masterActionStepEnum masterActionStep.MasterActionStepEnum, targetId *int64) (*MasterAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByMasterActionStepEnumAndTargetId", ctx, masterActionStepEnum, targetId)
	ret0, _ := ret[0].(*MasterAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByMasterActionStepEnumAndTargetId indicates an expected call of FindByMasterActionStepEnumAndTargetId.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindByMasterActionStepEnumAndTargetId(ctx, masterActionStepEnum, targetId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMasterActionStepEnumAndTargetId", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindByMasterActionStepEnumAndTargetId), ctx, masterActionStepEnum, targetId)
}

// FindByTargetId mocks base method.
func (m *MockMasterActionMysqlRepository) FindByTargetId(ctx context.Context, targetId *int64) (*MasterAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByTargetId", ctx, targetId)
	ret0, _ := ret[0].(*MasterAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByTargetId indicates an expected call of FindByTargetId.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindByTargetId(ctx, targetId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByTargetId", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindByTargetId), ctx, targetId)
}

// FindList mocks base method.
func (m *MockMasterActionMysqlRepository) FindList(ctx context.Context) (MasterActions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx)
	ret0, _ := ret[0].(MasterActions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindList), ctx)
}

// FindListByMasterActionStepEnum mocks base method.
func (m *MockMasterActionMysqlRepository) FindListByMasterActionStepEnum(ctx context.Context, masterActionStepEnum masterActionStep.MasterActionStepEnum) (MasterActions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByMasterActionStepEnum", ctx, masterActionStepEnum)
	ret0, _ := ret[0].(MasterActions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByMasterActionStepEnum indicates an expected call of FindListByMasterActionStepEnum.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindListByMasterActionStepEnum(ctx, masterActionStepEnum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByMasterActionStepEnum", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindListByMasterActionStepEnum), ctx, masterActionStepEnum)
}

// FindListByMasterActionStepEnumAndTargetId mocks base method.
func (m *MockMasterActionMysqlRepository) FindListByMasterActionStepEnumAndTargetId(ctx context.Context, masterActionStepEnum masterActionStep.MasterActionStepEnum, targetId *int64) (MasterActions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByMasterActionStepEnumAndTargetId", ctx, masterActionStepEnum, targetId)
	ret0, _ := ret[0].(MasterActions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByMasterActionStepEnumAndTargetId indicates an expected call of FindListByMasterActionStepEnumAndTargetId.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindListByMasterActionStepEnumAndTargetId(ctx, masterActionStepEnum, targetId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByMasterActionStepEnumAndTargetId", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindListByMasterActionStepEnumAndTargetId), ctx, masterActionStepEnum, targetId)
}

// FindListByTargetId mocks base method.
func (m *MockMasterActionMysqlRepository) FindListByTargetId(ctx context.Context, targetId *int64) (MasterActions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByTargetId", ctx, targetId)
	ret0, _ := ret[0].(MasterActions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByTargetId indicates an expected call of FindListByTargetId.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindListByTargetId(ctx, targetId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByTargetId", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindListByTargetId), ctx, targetId)
}

// FindOrNil mocks base method.
func (m *MockMasterActionMysqlRepository) FindOrNil(ctx context.Context, masterActionId int64) (*MasterAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, masterActionId)
	ret0, _ := ret[0].(*MasterAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindOrNil(ctx, masterActionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindOrNil), ctx, masterActionId)
}

// FindOrNilByMasterActionStepEnum mocks base method.
func (m *MockMasterActionMysqlRepository) FindOrNilByMasterActionStepEnum(ctx context.Context, masterActionStepEnum masterActionStep.MasterActionStepEnum) (*MasterAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByMasterActionStepEnum", ctx, masterActionStepEnum)
	ret0, _ := ret[0].(*MasterAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByMasterActionStepEnum indicates an expected call of FindOrNilByMasterActionStepEnum.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindOrNilByMasterActionStepEnum(ctx, masterActionStepEnum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByMasterActionStepEnum", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindOrNilByMasterActionStepEnum), ctx, masterActionStepEnum)
}

// FindOrNilByMasterActionStepEnumAndTargetId mocks base method.
func (m *MockMasterActionMysqlRepository) FindOrNilByMasterActionStepEnumAndTargetId(ctx context.Context, masterActionStepEnum masterActionStep.MasterActionStepEnum, targetId *int64) (*MasterAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByMasterActionStepEnumAndTargetId", ctx, masterActionStepEnum, targetId)
	ret0, _ := ret[0].(*MasterAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByMasterActionStepEnumAndTargetId indicates an expected call of FindOrNilByMasterActionStepEnumAndTargetId.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindOrNilByMasterActionStepEnumAndTargetId(ctx, masterActionStepEnum, targetId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByMasterActionStepEnumAndTargetId", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindOrNilByMasterActionStepEnumAndTargetId), ctx, masterActionStepEnum, targetId)
}

// FindOrNilByTargetId mocks base method.
func (m *MockMasterActionMysqlRepository) FindOrNilByTargetId(ctx context.Context, targetId *int64) (*MasterAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByTargetId", ctx, targetId)
	ret0, _ := ret[0].(*MasterAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByTargetId indicates an expected call of FindOrNilByTargetId.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindOrNilByTargetId(ctx, targetId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByTargetId", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindOrNilByTargetId), ctx, targetId)
}

// Update mocks base method.
func (m_2 *MockMasterActionMysqlRepository) Update(ctx context.Context, tx *gorm.DB, m *MasterAction) (*MasterAction, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*MasterAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).Update), ctx, tx, m)
}
