// Code generated by MockGen. DO NOT EDIT.
// Source: ./exchange_service.go

// Package exchange is a generated GoMock package.
package exchange

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockExchangeService is a mock of ExchangeService interface.
type MockExchangeService struct {
	ctrl     *gomock.Controller
	recorder *MockExchangeServiceMockRecorder
}

// MockExchangeServiceMockRecorder is the mock recorder for MockExchangeService.
type MockExchangeServiceMockRecorder struct {
	mock *MockExchangeService
}

// NewMockExchangeService creates a new mock instance.
func NewMockExchangeService(ctrl *gomock.Controller) *MockExchangeService {
	mock := &MockExchangeService{ctrl: ctrl}
	mock.recorder = &MockExchangeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExchangeService) EXPECT() *MockExchangeServiceMockRecorder {
	return m.recorder
}

// Receive mocks base method.
func (m *MockExchangeService) Receive(ctx context.Context, tx *gorm.DB, now time.Time, req *ExchangeReceiveRequest) (*ExchangeReceiveResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Receive", ctx, tx, now, req)
	ret0, _ := ret[0].(*ExchangeReceiveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Receive indicates an expected call of Receive.
func (mr *MockExchangeServiceMockRecorder) Receive(ctx, tx, now, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*MockExchangeService)(nil).Receive), ctx, tx, now, req)
}

// Update mocks base method.
func (m *MockExchangeService) Update(ctx context.Context, tx *gorm.DB, now time.Time, req *ExchangeUpdateRequest) (*ExchangeUpdateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, tx, now, req)
	ret0, _ := ret[0].(*ExchangeUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockExchangeServiceMockRecorder) Update(ctx, tx, now, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockExchangeService)(nil).Update), ctx, tx, now, req)
}
