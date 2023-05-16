// Code generated by MockGen. DO NOT EDIT.
// Source: ./services/logs/interface.go

// Package logs is a generated GoMock package.
package logs

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLogService is a mock of Service interface.
type MockLogService struct {
	ctrl     *gomock.Controller
	recorder *MockLogServiceMockRecorder
}

// MockLogServiceMockRecorder is the mock recorder for MockLogService.
type MockLogServiceMockRecorder struct {
	mock *MockLogService
}

// NewMockLogService creates a new mock instance.
func NewMockLogService(ctrl *gomock.Controller) *MockLogService {
	mock := &MockLogService{ctrl: ctrl}
	mock.recorder = &MockLogServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogService) EXPECT() *MockLogServiceMockRecorder {
	return m.recorder
}

// ComponentContainerLog mocks base method.
func (m *MockLogService) ComponentContainerLog(appName, envName, componentName, replicaName, containerId string, options *LogOptions) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComponentContainerLog", appName, envName, componentName, replicaName, containerId, options)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComponentContainerLog indicates an expected call of ComponentContainerLog.
func (mr *MockLogServiceMockRecorder) ComponentContainerLog(appName, envName, componentName, replicaName, containerId, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComponentContainerLog", reflect.TypeOf((*MockLogService)(nil).ComponentContainerLog), appName, envName, componentName, replicaName, containerId, options)
}

// ComponentInventory mocks base method.
func (m *MockLogService) ComponentInventory(appName, envName, componentName string, options *ComponentPodInventoryOptions) ([]Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComponentInventory", appName, envName, componentName, options)
	ret0, _ := ret[0].([]Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComponentInventory indicates an expected call of ComponentInventory.
func (mr *MockLogServiceMockRecorder) ComponentInventory(appName, envName, componentName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComponentInventory", reflect.TypeOf((*MockLogService)(nil).ComponentInventory), appName, envName, componentName, options)
}

// ComponentLog mocks base method.
func (m *MockLogService) ComponentLog(appName, envName, componentName string, options *LogOptions) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComponentLog", appName, envName, componentName, options)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComponentLog indicates an expected call of ComponentLog.
func (mr *MockLogServiceMockRecorder) ComponentLog(appName, envName, componentName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComponentLog", reflect.TypeOf((*MockLogService)(nil).ComponentLog), appName, envName, componentName, options)
}

// ComponentPodLog mocks base method.
func (m *MockLogService) ComponentPodLog(appName, envName, componentName, replicaName string, options *LogOptions) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComponentPodLog", appName, envName, componentName, replicaName, options)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComponentPodLog indicates an expected call of ComponentPodLog.
func (mr *MockLogServiceMockRecorder) ComponentPodLog(appName, envName, componentName, replicaName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComponentPodLog", reflect.TypeOf((*MockLogService)(nil).ComponentPodLog), appName, envName, componentName, replicaName, options)
}