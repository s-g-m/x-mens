// Code generated by MockGen. DO NOT EDIT.
// Source: x-men/app/adapter (interfaces: HttpAdapter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHttpAdapter is a mock of HttpAdapter interface.
type MockHttpAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockHttpAdapterMockRecorder
}

// MockHttpAdapterMockRecorder is the mock recorder for MockHttpAdapter.
type MockHttpAdapterMockRecorder struct {
	mock *MockHttpAdapter
}

// NewMockHttpAdapter creates a new mock instance.
func NewMockHttpAdapter(ctrl *gomock.Controller) *MockHttpAdapter {
	mock := &MockHttpAdapter{ctrl: ctrl}
	mock.recorder = &MockHttpAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHttpAdapter) EXPECT() *MockHttpAdapterMockRecorder {
	return m.recorder
}

// GetBody mocks base method.
func (m *MockHttpAdapter) GetBody() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBody")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetBody indicates an expected call of GetBody.
func (mr *MockHttpAdapterMockRecorder) GetBody() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBody", reflect.TypeOf((*MockHttpAdapter)(nil).GetBody))
}

// SendResponse mocks base method.
func (m *MockHttpAdapter) SendResponse(arg0 int, arg1 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendResponse", arg0, arg1)
}

// SendResponse indicates an expected call of SendResponse.
func (mr *MockHttpAdapterMockRecorder) SendResponse(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendResponse", reflect.TypeOf((*MockHttpAdapter)(nil).SendResponse), arg0, arg1)
}
