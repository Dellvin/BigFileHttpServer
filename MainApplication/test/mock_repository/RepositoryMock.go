// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/dellvin/Desktop/Работа/BigFileHttpServer/MainApplication/internal/files/repository/repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	model "HttpBigFilesServer/MainApplication/internal/files/model"
	io "io"
	os "os"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockInterfaceDataBase is a mock of InterfaceDataBase interface.
type MockInterfaceDataBase struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceDataBaseMockRecorder
}

// MockInterfaceDataBaseMockRecorder is the mock recorder for MockInterfaceDataBase.
type MockInterfaceDataBaseMockRecorder struct {
	mock *MockInterfaceDataBase
}

// NewMockInterfaceDataBase creates a new mock instance.
func NewMockInterfaceDataBase(ctrl *gomock.Controller) *MockInterfaceDataBase {
	mock := &MockInterfaceDataBase{ctrl: ctrl}
	mock.recorder = &MockInterfaceDataBaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterfaceDataBase) EXPECT() *MockInterfaceDataBaseMockRecorder {
	return m.recorder
}

// GenID mocks base method.
func (m *MockInterfaceDataBase) GenID() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenID")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenID indicates an expected call of GenID.
func (mr *MockInterfaceDataBaseMockRecorder) GenID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenID", reflect.TypeOf((*MockInterfaceDataBase)(nil).GenID))
}

// Get mocks base method.
func (m *MockInterfaceDataBase) Get(arg0 uint64) (model.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(model.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockInterfaceDataBaseMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterfaceDataBase)(nil).Get), arg0)
}

// Save mocks base method.
func (m *MockInterfaceDataBase) Save(arg0 model.File) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockInterfaceDataBaseMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockInterfaceDataBase)(nil).Save), arg0)
}

// MockInterfaceFile is a mock of InterfaceFile interface.
type MockInterfaceFile struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceFileMockRecorder
}

// MockInterfaceFileMockRecorder is the mock recorder for MockInterfaceFile.
type MockInterfaceFileMockRecorder struct {
	mock *MockInterfaceFile
}

// NewMockInterfaceFile creates a new mock instance.
func NewMockInterfaceFile(ctrl *gomock.Controller) *MockInterfaceFile {
	mock := &MockInterfaceFile{ctrl: ctrl}
	mock.recorder = &MockInterfaceFileMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterfaceFile) EXPECT() *MockInterfaceFileMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockInterfaceFile) Get(arg0, arg1 uint64) (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockInterfaceFileMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterfaceFile)(nil).Get), arg0, arg1)
}

// Remove mocks base method.
func (m *MockInterfaceFile) Remove(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Remove", arg0)
}

// Remove indicates an expected call of Remove.
func (mr *MockInterfaceFileMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockInterfaceFile)(nil).Remove), arg0)
}

// Save mocks base method.
func (m *MockInterfaceFile) Save(arg0 uint64, arg1 io.ReadCloser, arg2 uint64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockInterfaceFileMockRecorder) Save(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockInterfaceFile)(nil).Save), arg0, arg1, arg2)
}
