// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/services/logs/interface.go

// Package logs is a generated GoMock package.
package logs

import (
	context "context"
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
func (m *MockLogService) ComponentContainerLog(ctx context.Context, appName, envName, componentName, replicaName, containerId string, options *LogOptions) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComponentContainerLog", ctx, appName, envName, componentName, replicaName, containerId, options)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComponentContainerLog indicates an expected call of ComponentContainerLog.
func (mr *MockLogServiceMockRecorder) ComponentContainerLog(ctx, appName, envName, componentName, replicaName, containerId, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComponentContainerLog", reflect.TypeOf((*MockLogService)(nil).ComponentContainerLog), ctx, appName, envName, componentName, replicaName, containerId, options)
}

// ComponentInventory mocks base method.
func (m *MockLogService) ComponentInventory(ctx context.Context, appName, envName, componentName string, options *InventoryOptions) ([]Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComponentInventory", ctx, appName, envName, componentName, options)
	ret0, _ := ret[0].([]Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComponentInventory indicates an expected call of ComponentInventory.
func (mr *MockLogServiceMockRecorder) ComponentInventory(ctx, appName, envName, componentName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComponentInventory", reflect.TypeOf((*MockLogService)(nil).ComponentInventory), ctx, appName, envName, componentName, options)
}

// ComponentLog mocks base method.
func (m *MockLogService) ComponentLog(ctx context.Context, appName, envName, componentName string, options *LogOptions) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComponentLog", ctx, appName, envName, componentName, options)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComponentLog indicates an expected call of ComponentLog.
func (mr *MockLogServiceMockRecorder) ComponentLog(ctx, appName, envName, componentName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComponentLog", reflect.TypeOf((*MockLogService)(nil).ComponentLog), ctx, appName, envName, componentName, options)
}

// ComponentPodLog mocks base method.
func (m *MockLogService) ComponentPodLog(ctx context.Context, appName, envName, componentName, replicaName string, options *LogOptions) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComponentPodLog", ctx, appName, envName, componentName, replicaName, options)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComponentPodLog indicates an expected call of ComponentPodLog.
func (mr *MockLogServiceMockRecorder) ComponentPodLog(ctx, appName, envName, componentName, replicaName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComponentPodLog", reflect.TypeOf((*MockLogService)(nil).ComponentPodLog), ctx, appName, envName, componentName, replicaName, options)
}

// JobContainerLog mocks base method.
func (m *MockLogService) JobContainerLog(ctx context.Context, appName, envName, jobComponentName, jobName, replicaName, containerId string, options *LogOptions) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JobContainerLog", ctx, appName, envName, jobComponentName, jobName, replicaName, containerId, options)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// JobContainerLog indicates an expected call of JobContainerLog.
func (mr *MockLogServiceMockRecorder) JobContainerLog(ctx, appName, envName, jobComponentName, jobName, replicaName, containerId, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JobContainerLog", reflect.TypeOf((*MockLogService)(nil).JobContainerLog), ctx, appName, envName, jobComponentName, jobName, replicaName, containerId, options)
}

// JobInventory mocks base method.
func (m *MockLogService) JobInventory(ctx context.Context, appName, envName, jobComponentName, jobName string, options *InventoryOptions) ([]Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JobInventory", ctx, appName, envName, jobComponentName, jobName, options)
	ret0, _ := ret[0].([]Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// JobInventory indicates an expected call of JobInventory.
func (mr *MockLogServiceMockRecorder) JobInventory(ctx, appName, envName, jobComponentName, jobName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JobInventory", reflect.TypeOf((*MockLogService)(nil).JobInventory), ctx, appName, envName, jobComponentName, jobName, options)
}

// JobLog mocks base method.
func (m *MockLogService) JobLog(ctx context.Context, appName, envName, jobComponentName, jobName string, options *LogOptions) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JobLog", ctx, appName, envName, jobComponentName, jobName, options)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// JobLog indicates an expected call of JobLog.
func (mr *MockLogServiceMockRecorder) JobLog(ctx, appName, envName, jobComponentName, jobName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JobLog", reflect.TypeOf((*MockLogService)(nil).JobLog), ctx, appName, envName, jobComponentName, jobName, options)
}

// JobPodLog mocks base method.
func (m *MockLogService) JobPodLog(ctx context.Context, appName, envName, jobComponentName, jobName, replicaName string, options *LogOptions) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JobPodLog", ctx, appName, envName, jobComponentName, jobName, replicaName, options)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// JobPodLog indicates an expected call of JobPodLog.
func (mr *MockLogServiceMockRecorder) JobPodLog(ctx, appName, envName, jobComponentName, jobName, replicaName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JobPodLog", reflect.TypeOf((*MockLogService)(nil).JobPodLog), ctx, appName, envName, jobComponentName, jobName, replicaName, options)
}

// PipelineJobContainerLog mocks base method.
func (m *MockLogService) PipelineJobContainerLog(ctx context.Context, appName, pipelineJobName, replicaName, containerId string, options *LogOptions) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PipelineJobContainerLog", ctx, appName, pipelineJobName, replicaName, containerId, options)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PipelineJobContainerLog indicates an expected call of PipelineJobContainerLog.
func (mr *MockLogServiceMockRecorder) PipelineJobContainerLog(ctx, appName, pipelineJobName, replicaName, containerId, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PipelineJobContainerLog", reflect.TypeOf((*MockLogService)(nil).PipelineJobContainerLog), ctx, appName, pipelineJobName, replicaName, containerId, options)
}

// PipelineJobInventory mocks base method.
func (m *MockLogService) PipelineJobInventory(ctx context.Context, appName, pipelineJobName string, options *InventoryOptions) ([]Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PipelineJobInventory", ctx, appName, pipelineJobName, options)
	ret0, _ := ret[0].([]Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PipelineJobInventory indicates an expected call of PipelineJobInventory.
func (mr *MockLogServiceMockRecorder) PipelineJobInventory(ctx, appName, pipelineJobName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PipelineJobInventory", reflect.TypeOf((*MockLogService)(nil).PipelineJobInventory), ctx, appName, pipelineJobName, options)
}
