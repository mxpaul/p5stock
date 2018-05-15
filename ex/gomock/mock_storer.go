// Code generated by MockGen. DO NOT EDIT.
// Source: main.go

// Package main is a generated GoMock package.
package main

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStorer is a mock of Storer interface
type MockStorer struct {
	ctrl     *gomock.Controller
	recorder *MockStorerMockRecorder
}

// MockStorerMockRecorder is the mock recorder for MockStorer
type MockStorerMockRecorder struct {
	mock *MockStorer
}

// NewMockStorer creates a new mock instance
func NewMockStorer(ctrl *gomock.Controller) *MockStorer {
	mock := &MockStorer{ctrl: ctrl}
	mock.recorder = &MockStorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorer) EXPECT() *MockStorerMockRecorder {
	return m.recorder
}

// Upsert mocks base method
func (m *MockStorer) Upsert(arg0 interface{}) (int, error) {
	ret := m.ctrl.Call(m, "Upsert", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert
func (mr *MockStorerMockRecorder) Upsert(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockStorer)(nil).Upsert), arg0)
}
