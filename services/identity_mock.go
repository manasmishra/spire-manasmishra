// Code generated by MockGen. DO NOT EDIT.
// Source: identity.go

package services

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	common "github.com/spiffe/sri/pkg/common"
)

// MockIdentity is a mock of Identity interface
type MockIdentity struct {
	ctrl     *gomock.Controller
	recorder *MockIdentityMockRecorder
}

// MockIdentityMockRecorder is the mock recorder for MockIdentity
type MockIdentityMockRecorder struct {
	mock *MockIdentity
}

// NewMockIdentity creates a new mock instance
func NewMockIdentity(ctrl *gomock.Controller) *MockIdentity {
	mock := &MockIdentity{ctrl: ctrl}
	mock.recorder = &MockIdentityMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockIdentity) EXPECT() *MockIdentityMockRecorder {
	return _m.recorder
}

// Resolve mocks base method
func (_m *MockIdentity) Resolve(baseSpiffeIDs []string) (map[string]*common.Selectors, error) {
	ret := _m.ctrl.Call(_m, "Resolve", baseSpiffeIDs)
	ret0, _ := ret[0].(map[string]*common.Selectors)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Resolve indicates an expected call of Resolve
func (_mr *MockIdentityMockRecorder) Resolve(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Resolve", reflect.TypeOf((*MockIdentity)(nil).Resolve), arg0)
}

// CreateEntry mocks base method
func (_m *MockIdentity) CreateEntry(baseSpiffeID string, selector *common.Selector) error {
	ret := _m.ctrl.Call(_m, "CreateEntry", baseSpiffeID, selector)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEntry indicates an expected call of CreateEntry
func (_mr *MockIdentityMockRecorder) CreateEntry(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateEntry", reflect.TypeOf((*MockIdentity)(nil).CreateEntry), arg0, arg1)
}
