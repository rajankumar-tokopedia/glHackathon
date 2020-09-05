// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rajankumar549/glHackathon/src/interfaces/apperror (interfaces: AppError)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAppError is a mock of AppError interface
type MockAppError struct {
	ctrl     *gomock.Controller
	recorder *MockAppErrorMockRecorder
}

// MockAppErrorMockRecorder is the mock recorder for MockAppError
type MockAppErrorMockRecorder struct {
	mock *MockAppError
}

// NewMockAppError creates a new mock instance
func NewMockAppError(ctrl *gomock.Controller) *MockAppError {
	mock := &MockAppError{ctrl: ctrl}
	mock.recorder = &MockAppErrorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppError) EXPECT() *MockAppErrorMockRecorder {
	return m.recorder
}

// ErrorHandler mocks base method
func (m *MockAppError) ErrorHandler(arg0 error) (string, string, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ErrorHandler", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(int)
	return ret0, ret1, ret2
}

// ErrorHandler indicates an expected call of ErrorHandler
func (mr *MockAppErrorMockRecorder) ErrorHandler(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ErrorHandler", reflect.TypeOf((*MockAppError)(nil).ErrorHandler), arg0)
}