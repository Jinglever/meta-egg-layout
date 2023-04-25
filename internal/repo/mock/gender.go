// Code generated by MockGen. DO NOT EDIT.
// Source: meta-egg-layout/internal/repo (interfaces: GenderRepo)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	model "meta-egg-layout/gen/model"
	option "meta-egg-layout/gen/repo/option"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockGenderRepo is a mock of GenderRepo interface.
type MockGenderRepo struct {
	ctrl     *gomock.Controller
	recorder *MockGenderRepoMockRecorder
}

// MockGenderRepoMockRecorder is the mock recorder for MockGenderRepo.
type MockGenderRepoMockRecorder struct {
	mock *MockGenderRepo
}

// NewMockGenderRepo creates a new mock instance.
func NewMockGenderRepo(ctrl *gomock.Controller) *MockGenderRepo {
	mock := &MockGenderRepo{ctrl: ctrl}
	mock.recorder = &MockGenderRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenderRepo) EXPECT() *MockGenderRepoMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockGenderRepo) Count(arg0 context.Context, arg1 ...option.Option) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Count", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockGenderRepoMockRecorder) Count(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockGenderRepo)(nil).Count), varargs...)
}

// GetByID mocks base method.
func (m *MockGenderRepo) GetByID(arg0 context.Context, arg1 uint64, arg2 ...option.Option) (*model.Gender, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByID", varargs...)
	ret0, _ := ret[0].(*model.Gender)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockGenderRepoMockRecorder) GetByID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockGenderRepo)(nil).GetByID), varargs...)
}

// GetByIDs mocks base method.
func (m *MockGenderRepo) GetByIDs(arg0 context.Context, arg1 []uint64, arg2 ...option.Option) ([]*model.Gender, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByIDs", varargs...)
	ret0, _ := ret[0].([]*model.Gender)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDs indicates an expected call of GetByIDs.
func (mr *MockGenderRepoMockRecorder) GetByIDs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDs", reflect.TypeOf((*MockGenderRepo)(nil).GetByIDs), varargs...)
}

// GetSematicByID mocks base method.
func (m *MockGenderRepo) GetSematicByID(arg0 uint64) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSematicByID", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSematicByID indicates an expected call of GetSematicByID.
func (mr *MockGenderRepoMockRecorder) GetSematicByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSematicByID", reflect.TypeOf((*MockGenderRepo)(nil).GetSematicByID), arg0)
}

// GetTX mocks base method.
func (m *MockGenderRepo) GetTX(arg0 context.Context) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTX", arg0)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// GetTX indicates an expected call of GetTX.
func (mr *MockGenderRepoMockRecorder) GetTX(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTX", reflect.TypeOf((*MockGenderRepo)(nil).GetTX), arg0)
}

// Gets mocks base method.
func (m *MockGenderRepo) Gets(arg0 context.Context, arg1 ...option.Option) ([]*model.Gender, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Gets", varargs...)
	ret0, _ := ret[0].([]*model.Gender)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Gets indicates an expected call of Gets.
func (mr *MockGenderRepoMockRecorder) Gets(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gets", reflect.TypeOf((*MockGenderRepo)(nil).Gets), varargs...)
}
