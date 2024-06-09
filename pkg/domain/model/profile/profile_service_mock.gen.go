// Code generated by MockGen. DO NOT EDIT.
// Source: ./profile_service.go

// Package profile is a generated GoMock package.
package profile

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockProfileService is a mock of ProfileService interface.
type MockProfileService struct {
	ctrl     *gomock.Controller
	recorder *MockProfileServiceMockRecorder
}

// MockProfileServiceMockRecorder is the mock recorder for MockProfileService.
type MockProfileServiceMockRecorder struct {
	mock *MockProfileService
}

// NewMockProfileService creates a new mock instance.
func NewMockProfileService(ctrl *gomock.Controller) *MockProfileService {
	mock := &MockProfileService{ctrl: ctrl}
	mock.recorder = &MockProfileServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProfileService) EXPECT() *MockProfileServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProfileService) Create(ctx context.Context, tx *gorm.DB, req *ProfileCreateRequest) (*ProfileCreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, tx, req)
	ret0, _ := ret[0].(*ProfileCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProfileServiceMockRecorder) Create(ctx, tx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProfileService)(nil).Create), ctx, tx, req)
}

// Get mocks base method.
func (m *MockProfileService) Get(ctx context.Context, req *ProfileGetRequest) (*ProfileGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, req)
	ret0, _ := ret[0].(*ProfileGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProfileServiceMockRecorder) Get(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProfileService)(nil).Get), ctx, req)
}

// Update mocks base method.
func (m *MockProfileService) Update(ctx context.Context, tx *gorm.DB, req *ProfileUpdateRequest) (*ProfileUpdateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, tx, req)
	ret0, _ := ret[0].(*ProfileUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockProfileServiceMockRecorder) Update(ctx, tx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProfileService)(nil).Update), ctx, tx, req)
}