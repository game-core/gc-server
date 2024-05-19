// Code generated by MockGen. DO NOT EDIT.
// Source: ./account_service.go

// Package account is a generated GoMock package.
package account

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v9 "github.com/redis/go-redis/v9"
	gorm "gorm.io/gorm"
)

// MockAccountService is a mock of AccountService interface.
type MockAccountService struct {
	ctrl     *gomock.Controller
	recorder *MockAccountServiceMockRecorder
}

// MockAccountServiceMockRecorder is the mock recorder for MockAccountService.
type MockAccountServiceMockRecorder struct {
	mock *MockAccountService
}

// NewMockAccountService creates a new mock instance.
func NewMockAccountService(ctrl *gomock.Controller) *MockAccountService {
	mock := &MockAccountService{ctrl: ctrl}
	mock.recorder = &MockAccountServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountService) EXPECT() *MockAccountServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAccountService) Create(ctx context.Context, tx *gorm.DB, req *AccountCreateRequest) (*AccountCreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, tx, req)
	ret0, _ := ret[0].(*AccountCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAccountServiceMockRecorder) Create(ctx, tx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccountService)(nil).Create), ctx, tx, req)
}

// CreateUserId mocks base method.
func (m *MockAccountService) CreateUserId(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserId", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserId indicates an expected call of CreateUserId.
func (mr *MockAccountServiceMockRecorder) CreateUserId(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserId", reflect.TypeOf((*MockAccountService)(nil).CreateUserId), ctx)
}

// Get mocks base method.
func (m *MockAccountService) Get(ctx context.Context, req *AccountGetRequest) (*AccountGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, req)
	ret0, _ := ret[0].(*AccountGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAccountServiceMockRecorder) Get(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAccountService)(nil).Get), ctx, req)
}

// GetToken mocks base method.
func (m *MockAccountService) GetToken(ctx context.Context, req *AccountGetTokenRequest) (*AccountGetTokenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToken", ctx, req)
	ret0, _ := ret[0].(*AccountGetTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToken indicates an expected call of GetToken.
func (mr *MockAccountServiceMockRecorder) GetToken(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken", reflect.TypeOf((*MockAccountService)(nil).GetToken), ctx, req)
}

// Login mocks base method.
func (m *MockAccountService) Login(ctx context.Context, mtx *gorm.DB, rtx v9.Pipeliner, req *AccountLoginRequest) (*AccountLoginResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, mtx, rtx, req)
	ret0, _ := ret[0].(*AccountLoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAccountServiceMockRecorder) Login(ctx, mtx, rtx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAccountService)(nil).Login), ctx, mtx, rtx, req)
}
