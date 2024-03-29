// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/api/usecase/user/interactor.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	context "context"
	entity "dataflow/pkg/domain/entity/user"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockInteractor is a mock of Interactor interface
type MockInteractor struct {
	ctrl     *gomock.Controller
	recorder *MockInteractorMockRecorder
}

// MockInteractorMockRecorder is the mock recorder for MockInteractor
type MockInteractorMockRecorder struct {
	mock *MockInteractor
}

// NewMockInteractor creates a new mock instance
func NewMockInteractor(ctrl *gomock.Controller) *MockInteractor {
	mock := &MockInteractor{ctrl: ctrl}
	mock.recorder = &MockInteractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInteractor) EXPECT() *MockInteractorMockRecorder {
	return m.recorder
}

// CreateNewUser mocks base method
func (m *MockInteractor) CreateNewUser(ctx context.Context, uid, name, thumbnail string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewUser", ctx, uid, name, thumbnail)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewUser indicates an expected call of CreateNewUser
func (mr *MockInteractorMockRecorder) CreateNewUser(ctx, uid, name, thumbnail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewUser", reflect.TypeOf((*MockInteractor)(nil).CreateNewUser), ctx, uid, name, thumbnail)
}

// GetUserProfile mocks base method
func (m *MockInteractor) GetUserProfile(ctx context.Context, uid string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserProfile", ctx, uid)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProfile indicates an expected call of GetUserProfile
func (mr *MockInteractorMockRecorder) GetUserProfile(ctx, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProfile", reflect.TypeOf((*MockInteractor)(nil).GetUserProfile), ctx, uid)
}

// GetAll mocks base method
func (m *MockInteractor) GetAll(ctx context.Context) (entity.UserSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].(entity.UserSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockInteractorMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockInteractor)(nil).GetAll), ctx)
}
