// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/service/sftp_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	model "github.com/arunvelsriram/sftp-exporter/pkg/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSFTPService is a mock of SFTPService interface
type MockSFTPService struct {
	ctrl     *gomock.Controller
	recorder *MockSFTPServiceMockRecorder
}

// MockSFTPServiceMockRecorder is the mock recorder for MockSFTPService
type MockSFTPServiceMockRecorder struct {
	mock *MockSFTPService
}

// NewMockSFTPService creates a new mock instance
func NewMockSFTPService(ctrl *gomock.Controller) *MockSFTPService {
	mock := &MockSFTPService{ctrl: ctrl}
	mock.recorder = &MockSFTPServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSFTPService) EXPECT() *MockSFTPServiceMockRecorder {
	return m.recorder
}

// Connect mocks base method
func (m *MockSFTPService) Connect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect
func (mr *MockSFTPServiceMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockSFTPService)(nil).Connect))
}

// Close mocks base method
func (m *MockSFTPService) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockSFTPServiceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSFTPService)(nil).Close))
}

// FSStats mocks base method
func (m *MockSFTPService) FSStats() model.FSStats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FSStats")
	ret0, _ := ret[0].(model.FSStats)
	return ret0
}

// FSStats indicates an expected call of FSStats
func (mr *MockSFTPServiceMockRecorder) FSStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FSStats", reflect.TypeOf((*MockSFTPService)(nil).FSStats))
}

// ObjectStats mocks base method
func (m *MockSFTPService) ObjectStats() model.ObjectStats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectStats")
	ret0, _ := ret[0].(model.ObjectStats)
	return ret0
}

// ObjectStats indicates an expected call of ObjectStats
func (mr *MockSFTPServiceMockRecorder) ObjectStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectStats", reflect.TypeOf((*MockSFTPService)(nil).ObjectStats))
}
