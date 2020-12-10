// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kastenhq/kubestr/pkg/csi (interfaces: DataValidator)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDataValidator is a mock of DataValidator interface
type MockDataValidator struct {
	ctrl     *gomock.Controller
	recorder *MockDataValidatorMockRecorder
}

// MockDataValidatorMockRecorder is the mock recorder for MockDataValidator
type MockDataValidatorMockRecorder struct {
	mock *MockDataValidator
}

// NewMockDataValidator creates a new mock instance
func NewMockDataValidator(ctrl *gomock.Controller) *MockDataValidator {
	mock := &MockDataValidator{ctrl: ctrl}
	mock.recorder = &MockDataValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataValidator) EXPECT() *MockDataValidatorMockRecorder {
	return m.recorder
}

// FetchPodData mocks base method
func (m *MockDataValidator) FetchPodData(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchPodData", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchPodData indicates an expected call of FetchPodData
func (mr *MockDataValidatorMockRecorder) FetchPodData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchPodData", reflect.TypeOf((*MockDataValidator)(nil).FetchPodData), arg0, arg1)
}