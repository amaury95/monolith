// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/amaury95/monolith (interfaces: ISession)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	jwt "github.com/golang-jwt/jwt/v4"
	gomock "github.com/golang/mock/gomock"
)

// MockISession is a mock of ISession interface.
type MockISession struct {
	ctrl     *gomock.Controller
	recorder *MockISessionMockRecorder
}

// MockISessionMockRecorder is the mock recorder for MockISession.
type MockISessionMockRecorder struct {
	mock *MockISession
}

// NewMockISession creates a new mock instance.
func NewMockISession(ctrl *gomock.Controller) *MockISession {
	mock := &MockISession{ctrl: ctrl}
	mock.recorder = &MockISessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISession) EXPECT() *MockISessionMockRecorder {
	return m.recorder
}

// Claims mocks base method.
func (m *MockISession) Claims(arg0 string) (*jwt.RegisteredClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Claims", arg0)
	ret0, _ := ret[0].(*jwt.RegisteredClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Claims indicates an expected call of Claims.
func (mr *MockISessionMockRecorder) Claims(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Claims", reflect.TypeOf((*MockISession)(nil).Claims), arg0)
}
