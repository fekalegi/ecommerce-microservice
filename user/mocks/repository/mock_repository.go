// Code generated by MockGen. DO NOT EDIT.
// Source: repository_interface.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	domain "ecommerce-microservice/user/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockUserRepository) AddUser(user domain.User) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", user)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser.
func (mr *MockUserRepositoryMockRecorder) AddUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockUserRepository)(nil).AddUser), user)
}

// CheckLogin mocks base method.
func (m *MockUserRepository) CheckLogin(username, password string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckLogin", username, password)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckLogin indicates an expected call of CheckLogin.
func (mr *MockUserRepositoryMockRecorder) CheckLogin(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckLogin", reflect.TypeOf((*MockUserRepository)(nil).CheckLogin), username, password)
}

// DeleteUser mocks base method.
func (m *MockUserRepository) DeleteUser(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserRepositoryMockRecorder) DeleteUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserRepository)(nil).DeleteUser), id)
}

// FindAllUser mocks base method.
func (m *MockUserRepository) FindAllUser(offset, limit int, filter string) ([]domain.User, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllUser", offset, limit, filter)
	ret0, _ := ret[0].([]domain.User)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindAllUser indicates an expected call of FindAllUser.
func (mr *MockUserRepositoryMockRecorder) FindAllUser(offset, limit, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllUser", reflect.TypeOf((*MockUserRepository)(nil).FindAllUser), offset, limit, filter)
}

// FindUserByID mocks base method.
func (m *MockUserRepository) FindUserByID(id int) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByID", id)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByID indicates an expected call of FindUserByID.
func (mr *MockUserRepositoryMockRecorder) FindUserByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByID", reflect.TypeOf((*MockUserRepository)(nil).FindUserByID), id)
}

// FindUserByIDAndAuthID mocks base method.
func (m *MockUserRepository) FindUserByIDAndAuthID(id int, authID uuid.UUID) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByIDAndAuthID", id, authID)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByIDAndAuthID indicates an expected call of FindUserByIDAndAuthID.
func (mr *MockUserRepositoryMockRecorder) FindUserByIDAndAuthID(id, authID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByIDAndAuthID", reflect.TypeOf((*MockUserRepository)(nil).FindUserByIDAndAuthID), id, authID)
}

// UpdateAuthUUID mocks base method.
func (m *MockUserRepository) UpdateAuthUUID(id int, authID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAuthUUID", id, authID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAuthUUID indicates an expected call of UpdateAuthUUID.
func (mr *MockUserRepositoryMockRecorder) UpdateAuthUUID(id, authID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthUUID", reflect.TypeOf((*MockUserRepository)(nil).UpdateAuthUUID), id, authID)
}

// UpdateUser mocks base method.
func (m *MockUserRepository) UpdateUser(request domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", request)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserRepositoryMockRecorder) UpdateUser(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserRepository)(nil).UpdateUser), request)
}

// MockRoleRepository is a mock of RoleRepository interface.
type MockRoleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRoleRepositoryMockRecorder
}

// MockRoleRepositoryMockRecorder is the mock recorder for MockRoleRepository.
type MockRoleRepositoryMockRecorder struct {
	mock *MockRoleRepository
}

// NewMockRoleRepository creates a new mock instance.
func NewMockRoleRepository(ctrl *gomock.Controller) *MockRoleRepository {
	mock := &MockRoleRepository{ctrl: ctrl}
	mock.recorder = &MockRoleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleRepository) EXPECT() *MockRoleRepositoryMockRecorder {
	return m.recorder
}

// CreateRole mocks base method.
func (m *MockRoleRepository) CreateRole(role domain.Role) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRole", role)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRole indicates an expected call of CreateRole.
func (mr *MockRoleRepositoryMockRecorder) CreateRole(role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRole", reflect.TypeOf((*MockRoleRepository)(nil).CreateRole), role)
}

// DeleteRole mocks base method.
func (m *MockRoleRepository) DeleteRole(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRole", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRole indicates an expected call of DeleteRole.
func (mr *MockRoleRepositoryMockRecorder) DeleteRole(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRole", reflect.TypeOf((*MockRoleRepository)(nil).DeleteRole), id)
}

// FetchAllRole mocks base method.
func (m *MockRoleRepository) FetchAllRole() ([]domain.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAllRole")
	ret0, _ := ret[0].([]domain.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAllRole indicates an expected call of FetchAllRole.
func (mr *MockRoleRepositoryMockRecorder) FetchAllRole() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAllRole", reflect.TypeOf((*MockRoleRepository)(nil).FetchAllRole))
}

// FetchRoleByID mocks base method.
func (m *MockRoleRepository) FetchRoleByID(id int) (*domain.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchRoleByID", id)
	ret0, _ := ret[0].(*domain.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchRoleByID indicates an expected call of FetchRoleByID.
func (mr *MockRoleRepositoryMockRecorder) FetchRoleByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchRoleByID", reflect.TypeOf((*MockRoleRepository)(nil).FetchRoleByID), id)
}

// UpdateRole mocks base method.
func (m *MockRoleRepository) UpdateRole(role domain.Role) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRole", role)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRole indicates an expected call of UpdateRole.
func (mr *MockRoleRepositoryMockRecorder) UpdateRole(role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRole", reflect.TypeOf((*MockRoleRepository)(nil).UpdateRole), role)
}