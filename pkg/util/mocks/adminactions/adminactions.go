// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/frontend/adminactions (interfaces: KubeActions,AzureActions)

// Package mock_adminactions is a generated GoMock package.
package mock_adminactions

import (
	context "context"
	http "net/http"
	reflect "reflect"

	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-06-01/compute"
	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// MockKubeActions is a mock of KubeActions interface.
type MockKubeActions struct {
	ctrl     *gomock.Controller
	recorder *MockKubeActionsMockRecorder
}

// MockKubeActionsMockRecorder is the mock recorder for MockKubeActions.
type MockKubeActionsMockRecorder struct {
	mock *MockKubeActions
}

// NewMockKubeActions creates a new mock instance.
func NewMockKubeActions(ctrl *gomock.Controller) *MockKubeActions {
	mock := &MockKubeActions{ctrl: ctrl}
	mock.recorder = &MockKubeActionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubeActions) EXPECT() *MockKubeActionsMockRecorder {
	return m.recorder
}

// ApproveAllCsrs mocks base method.
func (m *MockKubeActions) ApproveAllCsrs(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApproveAllCsrs", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApproveAllCsrs indicates an expected call of ApproveAllCsrs.
func (mr *MockKubeActionsMockRecorder) ApproveAllCsrs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApproveAllCsrs", reflect.TypeOf((*MockKubeActions)(nil).ApproveAllCsrs), arg0)
}

// ApproveCsr mocks base method.
func (m *MockKubeActions) ApproveCsr(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApproveCsr", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApproveCsr indicates an expected call of ApproveCsr.
func (mr *MockKubeActionsMockRecorder) ApproveCsr(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApproveCsr", reflect.TypeOf((*MockKubeActions)(nil).ApproveCsr), arg0, arg1)
}

// CordonNode mocks base method.
func (m *MockKubeActions) CordonNode(arg0 context.Context, arg1 string, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CordonNode", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CordonNode indicates an expected call of CordonNode.
func (mr *MockKubeActionsMockRecorder) CordonNode(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CordonNode", reflect.TypeOf((*MockKubeActions)(nil).CordonNode), arg0, arg1, arg2)
}

// DrainNode mocks base method.
func (m *MockKubeActions) DrainNode(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DrainNode", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DrainNode indicates an expected call of DrainNode.
func (mr *MockKubeActionsMockRecorder) DrainNode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DrainNode", reflect.TypeOf((*MockKubeActions)(nil).DrainNode), arg0, arg1)
}

// KubeCreateOrUpdate mocks base method.
func (m *MockKubeActions) KubeCreateOrUpdate(arg0 context.Context, arg1 *unstructured.Unstructured) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KubeCreateOrUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// KubeCreateOrUpdate indicates an expected call of KubeCreateOrUpdate.
func (mr *MockKubeActionsMockRecorder) KubeCreateOrUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KubeCreateOrUpdate", reflect.TypeOf((*MockKubeActions)(nil).KubeCreateOrUpdate), arg0, arg1)
}

// KubeDelete mocks base method.
func (m *MockKubeActions) KubeDelete(arg0 context.Context, arg1, arg2, arg3 string, arg4 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KubeDelete", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// KubeDelete indicates an expected call of KubeDelete.
func (mr *MockKubeActionsMockRecorder) KubeDelete(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KubeDelete", reflect.TypeOf((*MockKubeActions)(nil).KubeDelete), arg0, arg1, arg2, arg3, arg4)
}

// KubeGet mocks base method.
func (m *MockKubeActions) KubeGet(arg0 context.Context, arg1, arg2, arg3 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KubeGet", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KubeGet indicates an expected call of KubeGet.
func (mr *MockKubeActionsMockRecorder) KubeGet(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KubeGet", reflect.TypeOf((*MockKubeActions)(nil).KubeGet), arg0, arg1, arg2, arg3)
}

// KubeGetPodLogs mocks base method.
func (m *MockKubeActions) KubeGetPodLogs(arg0 context.Context, arg1, arg2, arg3 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KubeGetPodLogs", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KubeGetPodLogs indicates an expected call of KubeGetPodLogs.
func (mr *MockKubeActionsMockRecorder) KubeGetPodLogs(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KubeGetPodLogs", reflect.TypeOf((*MockKubeActions)(nil).KubeGetPodLogs), arg0, arg1, arg2, arg3)
}

// KubeList mocks base method.
func (m *MockKubeActions) KubeList(arg0 context.Context, arg1, arg2 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KubeList", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KubeList indicates an expected call of KubeList.
func (mr *MockKubeActionsMockRecorder) KubeList(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KubeList", reflect.TypeOf((*MockKubeActions)(nil).KubeList), arg0, arg1, arg2)
}

// Upgrade mocks base method.
func (m *MockKubeActions) Upgrade(arg0 context.Context, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upgrade", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upgrade indicates an expected call of Upgrade.
func (mr *MockKubeActionsMockRecorder) Upgrade(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrade", reflect.TypeOf((*MockKubeActions)(nil).Upgrade), arg0, arg1)
}

// MockAzureActions is a mock of AzureActions interface.
type MockAzureActions struct {
	ctrl     *gomock.Controller
	recorder *MockAzureActionsMockRecorder
}

// MockAzureActionsMockRecorder is the mock recorder for MockAzureActions.
type MockAzureActionsMockRecorder struct {
	mock *MockAzureActions
}

// NewMockAzureActions creates a new mock instance.
func NewMockAzureActions(ctrl *gomock.Controller) *MockAzureActions {
	mock := &MockAzureActions{ctrl: ctrl}
	mock.recorder = &MockAzureActionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAzureActions) EXPECT() *MockAzureActionsMockRecorder {
	return m.recorder
}

// NICReconcileFailedState mocks base method.
func (m *MockAzureActions) NICReconcileFailedState(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NICReconcileFailedState", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// NICReconcileFailedState indicates an expected call of NICReconcileFailedState.
func (mr *MockAzureActionsMockRecorder) NICReconcileFailedState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NICReconcileFailedState", reflect.TypeOf((*MockAzureActions)(nil).NICReconcileFailedState), arg0, arg1)
}

// ResourcesList mocks base method.
func (m *MockAzureActions) ResourcesList(arg0 context.Context) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourcesList", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResourcesList indicates an expected call of ResourcesList.
func (mr *MockAzureActionsMockRecorder) ResourcesList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourcesList", reflect.TypeOf((*MockAzureActions)(nil).ResourcesList), arg0)
}

// VMRedeployAndWait mocks base method.
func (m *MockAzureActions) VMRedeployAndWait(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VMRedeployAndWait", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// VMRedeployAndWait indicates an expected call of VMRedeployAndWait.
func (mr *MockAzureActionsMockRecorder) VMRedeployAndWait(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VMRedeployAndWait", reflect.TypeOf((*MockAzureActions)(nil).VMRedeployAndWait), arg0, arg1)
}

// VMResize mocks base method.
func (m *MockAzureActions) VMResize(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VMResize", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// VMResize indicates an expected call of VMResize.
func (mr *MockAzureActionsMockRecorder) VMResize(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VMResize", reflect.TypeOf((*MockAzureActions)(nil).VMResize), arg0, arg1, arg2)
}

// VMSerialConsole mocks base method.
func (m *MockAzureActions) VMSerialConsole(arg0 context.Context, arg1 http.ResponseWriter, arg2 *logrus.Entry, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VMSerialConsole", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// VMSerialConsole indicates an expected call of VMSerialConsole.
func (mr *MockAzureActionsMockRecorder) VMSerialConsole(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VMSerialConsole", reflect.TypeOf((*MockAzureActions)(nil).VMSerialConsole), arg0, arg1, arg2, arg3)
}

// VMSizeList mocks base method.
func (m *MockAzureActions) VMSizeList(arg0 context.Context) ([]compute.ResourceSku, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VMSizeList", arg0)
	ret0, _ := ret[0].([]compute.ResourceSku)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VMSizeList indicates an expected call of VMSizeList.
func (mr *MockAzureActionsMockRecorder) VMSizeList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VMSizeList", reflect.TypeOf((*MockAzureActions)(nil).VMSizeList), arg0)
}

// VMStartAndWait mocks base method.
func (m *MockAzureActions) VMStartAndWait(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VMStartAndWait", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// VMStartAndWait indicates an expected call of VMStartAndWait.
func (mr *MockAzureActionsMockRecorder) VMStartAndWait(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VMStartAndWait", reflect.TypeOf((*MockAzureActions)(nil).VMStartAndWait), arg0, arg1)
}

// VMStopAndWait mocks base method.
func (m *MockAzureActions) VMStopAndWait(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VMStopAndWait", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// VMStopAndWait indicates an expected call of VMStopAndWait.
func (mr *MockAzureActionsMockRecorder) VMStopAndWait(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VMStopAndWait", reflect.TypeOf((*MockAzureActions)(nil).VMStopAndWait), arg0, arg1)
}
