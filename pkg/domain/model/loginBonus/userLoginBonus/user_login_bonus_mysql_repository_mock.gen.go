// Code generated by MockGen. DO NOT EDIT.
// Source: ./user_login_bonus_mysql_repository.gen.go

// Package userLoginBonus is a generated GoMock package.
package userLoginBonus

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockUserLoginBonusMysqlRepository is a mock of UserLoginBonusMysqlRepository interface.
type MockUserLoginBonusMysqlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserLoginBonusMysqlRepositoryMockRecorder
}

// MockUserLoginBonusMysqlRepositoryMockRecorder is the mock recorder for MockUserLoginBonusMysqlRepository.
type MockUserLoginBonusMysqlRepositoryMockRecorder struct {
	mock *MockUserLoginBonusMysqlRepository
}

// NewMockUserLoginBonusMysqlRepository creates a new mock instance.
func NewMockUserLoginBonusMysqlRepository(ctrl *gomock.Controller) *MockUserLoginBonusMysqlRepository {
	mock := &MockUserLoginBonusMysqlRepository{ctrl: ctrl}
	mock.recorder = &MockUserLoginBonusMysqlRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserLoginBonusMysqlRepository) EXPECT() *MockUserLoginBonusMysqlRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockUserLoginBonusMysqlRepository) Create(ctx context.Context, tx *gorm.DB, m *UserLoginBonus) (*UserLoginBonus, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*UserLoginBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserLoginBonusMysqlRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserLoginBonusMysqlRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockUserLoginBonusMysqlRepository) CreateList(ctx context.Context, tx *gorm.DB, ms UserLoginBonuses) (UserLoginBonuses, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(UserLoginBonuses)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockUserLoginBonusMysqlRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockUserLoginBonusMysqlRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockUserLoginBonusMysqlRepository) Delete(ctx context.Context, tx *gorm.DB, m *UserLoginBonus) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserLoginBonusMysqlRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserLoginBonusMysqlRepository)(nil).Delete), ctx, tx, m)
}

// DeleteList mocks base method.
func (m *MockUserLoginBonusMysqlRepository) DeleteList(ctx context.Context, tx *gorm.DB, ms UserLoginBonuses) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteList", ctx, tx, ms)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteList indicates an expected call of DeleteList.
func (mr *MockUserLoginBonusMysqlRepositoryMockRecorder) DeleteList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteList", reflect.TypeOf((*MockUserLoginBonusMysqlRepository)(nil).DeleteList), ctx, tx, ms)
}

// Find mocks base method.
func (m *MockUserLoginBonusMysqlRepository) Find(ctx context.Context, userId string, masterLoginBonusId int64) (*UserLoginBonus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, userId, masterLoginBonusId)
	ret0, _ := ret[0].(*UserLoginBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockUserLoginBonusMysqlRepositoryMockRecorder) Find(ctx, userId, masterLoginBonusId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUserLoginBonusMysqlRepository)(nil).Find), ctx, userId, masterLoginBonusId)
}

// FindByUserId mocks base method.
func (m *MockUserLoginBonusMysqlRepository) FindByUserId(ctx context.Context, userId string) (*UserLoginBonus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserId", ctx, userId)
	ret0, _ := ret[0].(*UserLoginBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserId indicates an expected call of FindByUserId.
func (mr *MockUserLoginBonusMysqlRepositoryMockRecorder) FindByUserId(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserId", reflect.TypeOf((*MockUserLoginBonusMysqlRepository)(nil).FindByUserId), ctx, userId)
}

// FindByUserIdAndMasterLoginBonusId mocks base method.
func (m *MockUserLoginBonusMysqlRepository) FindByUserIdAndMasterLoginBonusId(ctx context.Context, userId string, masterLoginBonusId int64) (*UserLoginBonus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserIdAndMasterLoginBonusId", ctx, userId, masterLoginBonusId)
	ret0, _ := ret[0].(*UserLoginBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserIdAndMasterLoginBonusId indicates an expected call of FindByUserIdAndMasterLoginBonusId.
func (mr *MockUserLoginBonusMysqlRepositoryMockRecorder) FindByUserIdAndMasterLoginBonusId(ctx, userId, masterLoginBonusId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserIdAndMasterLoginBonusId", reflect.TypeOf((*MockUserLoginBonusMysqlRepository)(nil).FindByUserIdAndMasterLoginBonusId), ctx, userId, masterLoginBonusId)
}

// FindList mocks base method.
func (m *MockUserLoginBonusMysqlRepository) FindList(ctx context.Context, userId string) (UserLoginBonuses, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx, userId)
	ret0, _ := ret[0].(UserLoginBonuses)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockUserLoginBonusMysqlRepositoryMockRecorder) FindList(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockUserLoginBonusMysqlRepository)(nil).FindList), ctx, userId)
}

// FindListByUserId mocks base method.
func (m *MockUserLoginBonusMysqlRepository) FindListByUserId(ctx context.Context, userId string) (UserLoginBonuses, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByUserId", ctx, userId)
	ret0, _ := ret[0].(UserLoginBonuses)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByUserId indicates an expected call of FindListByUserId.
func (mr *MockUserLoginBonusMysqlRepositoryMockRecorder) FindListByUserId(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByUserId", reflect.TypeOf((*MockUserLoginBonusMysqlRepository)(nil).FindListByUserId), ctx, userId)
}

// FindListByUserIdAndMasterLoginBonusId mocks base method.
func (m *MockUserLoginBonusMysqlRepository) FindListByUserIdAndMasterLoginBonusId(ctx context.Context, userId string, masterLoginBonusId int64) (UserLoginBonuses, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByUserIdAndMasterLoginBonusId", ctx, userId, masterLoginBonusId)
	ret0, _ := ret[0].(UserLoginBonuses)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByUserIdAndMasterLoginBonusId indicates an expected call of FindListByUserIdAndMasterLoginBonusId.
func (mr *MockUserLoginBonusMysqlRepositoryMockRecorder) FindListByUserIdAndMasterLoginBonusId(ctx, userId, masterLoginBonusId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByUserIdAndMasterLoginBonusId", reflect.TypeOf((*MockUserLoginBonusMysqlRepository)(nil).FindListByUserIdAndMasterLoginBonusId), ctx, userId, masterLoginBonusId)
}

// FindOrNil mocks base method.
func (m *MockUserLoginBonusMysqlRepository) FindOrNil(ctx context.Context, userId string, masterLoginBonusId int64) (*UserLoginBonus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, userId, masterLoginBonusId)
	ret0, _ := ret[0].(*UserLoginBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockUserLoginBonusMysqlRepositoryMockRecorder) FindOrNil(ctx, userId, masterLoginBonusId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockUserLoginBonusMysqlRepository)(nil).FindOrNil), ctx, userId, masterLoginBonusId)
}

// FindOrNilByUserId mocks base method.
func (m *MockUserLoginBonusMysqlRepository) FindOrNilByUserId(ctx context.Context, userId string) (*UserLoginBonus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByUserId", ctx, userId)
	ret0, _ := ret[0].(*UserLoginBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByUserId indicates an expected call of FindOrNilByUserId.
func (mr *MockUserLoginBonusMysqlRepositoryMockRecorder) FindOrNilByUserId(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByUserId", reflect.TypeOf((*MockUserLoginBonusMysqlRepository)(nil).FindOrNilByUserId), ctx, userId)
}

// FindOrNilByUserIdAndMasterLoginBonusId mocks base method.
func (m *MockUserLoginBonusMysqlRepository) FindOrNilByUserIdAndMasterLoginBonusId(ctx context.Context, userId string, masterLoginBonusId int64) (*UserLoginBonus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByUserIdAndMasterLoginBonusId", ctx, userId, masterLoginBonusId)
	ret0, _ := ret[0].(*UserLoginBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByUserIdAndMasterLoginBonusId indicates an expected call of FindOrNilByUserIdAndMasterLoginBonusId.
func (mr *MockUserLoginBonusMysqlRepositoryMockRecorder) FindOrNilByUserIdAndMasterLoginBonusId(ctx, userId, masterLoginBonusId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByUserIdAndMasterLoginBonusId", reflect.TypeOf((*MockUserLoginBonusMysqlRepository)(nil).FindOrNilByUserIdAndMasterLoginBonusId), ctx, userId, masterLoginBonusId)
}

// Update mocks base method.
func (m_2 *MockUserLoginBonusMysqlRepository) Update(ctx context.Context, tx *gorm.DB, m *UserLoginBonus) (*UserLoginBonus, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*UserLoginBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUserLoginBonusMysqlRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserLoginBonusMysqlRepository)(nil).Update), ctx, tx, m)
}

// UpdateList mocks base method.
func (m *MockUserLoginBonusMysqlRepository) UpdateList(ctx context.Context, tx *gorm.DB, ms UserLoginBonuses) (UserLoginBonuses, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateList", ctx, tx, ms)
	ret0, _ := ret[0].(UserLoginBonuses)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateList indicates an expected call of UpdateList.
func (mr *MockUserLoginBonusMysqlRepositoryMockRecorder) UpdateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateList", reflect.TypeOf((*MockUserLoginBonusMysqlRepository)(nil).UpdateList), ctx, tx, ms)
}