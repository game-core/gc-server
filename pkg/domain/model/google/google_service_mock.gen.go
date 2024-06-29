// Code generated by MockGen. DO NOT EDIT.
// Source: ./google_service.go

// Package google is a generated GoMock package.
package google

import (
	context "context"
	reflect "reflect"

	adminGoogle "github.com/game-core/gc-server/pkg/domain/model/google/adminGoogle"
	gomock "github.com/golang/mock/gomock"
)

// MockGoogleService is a mock of GoogleService interface.
type MockGoogleService struct {
	ctrl     *gomock.Controller
	recorder *MockGoogleServiceMockRecorder
}

// MockGoogleServiceMockRecorder is the mock recorder for MockGoogleService.
type MockGoogleServiceMockRecorder struct {
	mock *MockGoogleService
}

// NewMockGoogleService creates a new mock instance.
func NewMockGoogleService(ctrl *gomock.Controller) *MockGoogleService {
	mock := &MockGoogleService{ctrl: ctrl}
	mock.recorder = &MockGoogleServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGoogleService) EXPECT() *MockGoogleServiceMockRecorder {
	return m.recorder
}

// GetAdminGoogleToken mocks base method.
func (m *MockGoogleService) GetAdminGoogleToken(ctx context.Context, code string) (*adminGoogle.AdminGoogleToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdminGoogleToken", ctx, code)
	ret0, _ := ret[0].(*adminGoogle.AdminGoogleToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdminGoogleToken indicates an expected call of GetAdminGoogleToken.
func (mr *MockGoogleServiceMockRecorder) GetAdminGoogleToken(ctx, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdminGoogleToken", reflect.TypeOf((*MockGoogleService)(nil).GetAdminGoogleToken), ctx, code)
}

// GetAdminGoogleTokenInfo mocks base method.
func (m *MockGoogleService) GetAdminGoogleTokenInfo(ctx context.Context, accessToken string) (*adminGoogle.AdminGoogleTokenInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdminGoogleTokenInfo", ctx, accessToken)
	ret0, _ := ret[0].(*adminGoogle.AdminGoogleTokenInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdminGoogleTokenInfo indicates an expected call of GetAdminGoogleTokenInfo.
func (mr *MockGoogleServiceMockRecorder) GetAdminGoogleTokenInfo(ctx, accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdminGoogleTokenInfo", reflect.TypeOf((*MockGoogleService)(nil).GetAdminGoogleTokenInfo), ctx, accessToken)
}

// GetAdminGoogleUrl mocks base method.
func (m *MockGoogleService) GetAdminGoogleUrl() (*adminGoogle.AdminGoogleURL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdminGoogleUrl")
	ret0, _ := ret[0].(*adminGoogle.AdminGoogleURL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdminGoogleUrl indicates an expected call of GetAdminGoogleUrl.
func (mr *MockGoogleServiceMockRecorder) GetAdminGoogleUrl() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdminGoogleUrl", reflect.TypeOf((*MockGoogleService)(nil).GetAdminGoogleUrl))
}
