// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/config/config.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockConfig is a mock of Config interface
type MockConfig struct {
	ctrl     *gomock.Controller
	recorder *MockConfigMockRecorder
}

// MockConfigMockRecorder is the mock recorder for MockConfig
type MockConfigMockRecorder struct {
	mock *MockConfig
}

// NewMockConfig creates a new mock instance
func NewMockConfig(ctrl *gomock.Controller) *MockConfig {
	mock := &MockConfig{ctrl: ctrl}
	mock.recorder = &MockConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfig) EXPECT() *MockConfigMockRecorder {
	return m.recorder
}

// GetBindAddress mocks base method
func (m *MockConfig) GetBindAddress() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBindAddress")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetBindAddress indicates an expected call of GetBindAddress
func (mr *MockConfigMockRecorder) GetBindAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBindAddress", reflect.TypeOf((*MockConfig)(nil).GetBindAddress))
}

// GetPort mocks base method
func (m *MockConfig) GetPort() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPort")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetPort indicates an expected call of GetPort
func (mr *MockConfigMockRecorder) GetPort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPort", reflect.TypeOf((*MockConfig)(nil).GetPort))
}

// GetLogLevel mocks base method
func (m *MockConfig) GetLogLevel() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogLevel")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetLogLevel indicates an expected call of GetLogLevel
func (mr *MockConfigMockRecorder) GetLogLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogLevel", reflect.TypeOf((*MockConfig)(nil).GetLogLevel))
}

// GetSFTPHost mocks base method
func (m *MockConfig) GetSFTPHost() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSFTPHost")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSFTPHost indicates an expected call of GetSFTPHost
func (mr *MockConfigMockRecorder) GetSFTPHost() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSFTPHost", reflect.TypeOf((*MockConfig)(nil).GetSFTPHost))
}

// GetSFTPPort mocks base method
func (m *MockConfig) GetSFTPPort() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSFTPPort")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetSFTPPort indicates an expected call of GetSFTPPort
func (mr *MockConfigMockRecorder) GetSFTPPort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSFTPPort", reflect.TypeOf((*MockConfig)(nil).GetSFTPPort))
}

// GetSFTPUser mocks base method
func (m *MockConfig) GetSFTPUser() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSFTPUser")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSFTPUser indicates an expected call of GetSFTPUser
func (mr *MockConfigMockRecorder) GetSFTPUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSFTPUser", reflect.TypeOf((*MockConfig)(nil).GetSFTPUser))
}

// GetSFTPPass mocks base method
func (m *MockConfig) GetSFTPPass() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSFTPPass")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSFTPPass indicates an expected call of GetSFTPPass
func (mr *MockConfigMockRecorder) GetSFTPPass() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSFTPPass", reflect.TypeOf((*MockConfig)(nil).GetSFTPPass))
}

// GetSFTPKey mocks base method
func (m *MockConfig) GetSFTPKey() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSFTPKey")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetSFTPKey indicates an expected call of GetSFTPKey
func (mr *MockConfigMockRecorder) GetSFTPKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSFTPKey", reflect.TypeOf((*MockConfig)(nil).GetSFTPKey))
}

// GetSFTPKeyFile mocks base method
func (m *MockConfig) GetSFTPKeyFile() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSFTPKeyFile")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSFTPKeyFile indicates an expected call of GetSFTPKeyFile
func (mr *MockConfigMockRecorder) GetSFTPKeyFile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSFTPKeyFile", reflect.TypeOf((*MockConfig)(nil).GetSFTPKeyFile))
}

// GetSFTPKeyPassphrase mocks base method
func (m *MockConfig) GetSFTPKeyPassphrase() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSFTPKeyPassphrase")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetSFTPKeyPassphrase indicates an expected call of GetSFTPKeyPassphrase
func (mr *MockConfigMockRecorder) GetSFTPKeyPassphrase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSFTPKeyPassphrase", reflect.TypeOf((*MockConfig)(nil).GetSFTPKeyPassphrase))
}

// GetSFTPPaths mocks base method
func (m *MockConfig) GetSFTPPaths() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSFTPPaths")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetSFTPPaths indicates an expected call of GetSFTPPaths
func (mr *MockConfigMockRecorder) GetSFTPPaths() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSFTPPaths", reflect.TypeOf((*MockConfig)(nil).GetSFTPPaths))
}
