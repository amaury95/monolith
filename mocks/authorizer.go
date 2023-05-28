// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/amaury95/monolith (interfaces: IAuthorizer)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIAuthorizer is a mock of IAuthorizer interface.
type MockIAuthorizer struct {
	ctrl     *gomock.Controller
	recorder *MockIAuthorizerMockRecorder
}

// MockIAuthorizerMockRecorder is the mock recorder for MockIAuthorizer.
type MockIAuthorizerMockRecorder struct {
	mock *MockIAuthorizer
}

// NewMockIAuthorizer creates a new mock instance.
func NewMockIAuthorizer(ctrl *gomock.Controller) *MockIAuthorizer {
	mock := &MockIAuthorizer{ctrl: ctrl}
	mock.recorder = &MockIAuthorizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAuthorizer) EXPECT() *MockIAuthorizerMockRecorder {
	return m.recorder
}

// Authorize mocks base method.
func (m *MockIAuthorizer) Authorize(arg0, arg1, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorize", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Authorize indicates an expected call of Authorize.
func (mr *MockIAuthorizerMockRecorder) Authorize(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorize", reflect.TypeOf((*MockIAuthorizer)(nil).Authorize), arg0, arg1, arg2)
}
