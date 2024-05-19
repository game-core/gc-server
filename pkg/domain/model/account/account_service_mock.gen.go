// Code generated by MockGen. DO NOT EDIT.
// Source: ./account_service.go

// Package account is a generated GoMock package.
package account

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
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

// CheckToken mocks base method.
func (m *MockAccountService) CheckToken(ctx context.Context, req *AccountCheckTokenRequest) (*AccountCheckTokenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckToken", ctx, req)
	ret0, _ := ret[0].(*AccountCheckTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckToken indicates an expected call of CheckToken.
func (mr *MockAccountServiceMockRecorder) CheckToken(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckToken", reflect.TypeOf((*MockAccountService)(nil).CheckToken), ctx, req)
}
