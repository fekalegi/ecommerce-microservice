// Code generated by MockGen. DO NOT EDIT.
// Source: repository_interface.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	domain "ecommerce-microservice/order/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockOrderRepository is a mock of OrderRepository interface.
type MockOrderRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepositoryMockRecorder
}

// MockOrderRepositoryMockRecorder is the mock recorder for MockOrderRepository.
type MockOrderRepositoryMockRecorder struct {
	mock *MockOrderRepository
}

// NewMockOrderRepository creates a new mock instance.
func NewMockOrderRepository(ctrl *gomock.Controller) *MockOrderRepository {
	mock := &MockOrderRepository{ctrl: ctrl}
	mock.recorder = &MockOrderRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepository) EXPECT() *MockOrderRepositoryMockRecorder {
	return m.recorder
}

// CreateInitialOrder mocks base method.
func (m *MockOrderRepository) CreateInitialOrder(order domain.Order, product domain.OrderProduct) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInitialOrder", order, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateInitialOrder indicates an expected call of CreateInitialOrder.
func (mr *MockOrderRepositoryMockRecorder) CreateInitialOrder(order, product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInitialOrder", reflect.TypeOf((*MockOrderRepository)(nil).CreateInitialOrder), order, product)
}

// CreateOrder mocks base method.
func (m *MockOrderRepository) CreateOrder(order domain.Order) (*domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", order)
	ret0, _ := ret[0].(*domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrderRepositoryMockRecorder) CreateOrder(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrderRepository)(nil).CreateOrder), order)
}

// CreateOrderProduct mocks base method.
func (m *MockOrderRepository) CreateOrderProduct(product domain.OrderProduct) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrderProduct", product)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrderProduct indicates an expected call of CreateOrderProduct.
func (mr *MockOrderRepositoryMockRecorder) CreateOrderProduct(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrderProduct", reflect.TypeOf((*MockOrderRepository)(nil).CreateOrderProduct), product)
}

// CreateRatingSeller mocks base method.
func (m *MockOrderRepository) CreateRatingSeller(seller domain.SellerRating) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRatingSeller", seller)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRatingSeller indicates an expected call of CreateRatingSeller.
func (mr *MockOrderRepositoryMockRecorder) CreateRatingSeller(seller interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRatingSeller", reflect.TypeOf((*MockOrderRepository)(nil).CreateRatingSeller), seller)
}

// DeleteOrder mocks base method.
func (m *MockOrderRepository) DeleteOrder(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrder", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrder indicates an expected call of DeleteOrder.
func (mr *MockOrderRepositoryMockRecorder) DeleteOrder(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrder", reflect.TypeOf((*MockOrderRepository)(nil).DeleteOrder), id)
}

// DeleteOrderProduct mocks base method.
func (m *MockOrderRepository) DeleteOrderProduct(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrderProduct", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrderProduct indicates an expected call of DeleteOrderProduct.
func (mr *MockOrderRepositoryMockRecorder) DeleteOrderProduct(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrderProduct", reflect.TypeOf((*MockOrderRepository)(nil).DeleteOrderProduct), id)
}

// FetchAllOrderByUserID mocks base method.
func (m *MockOrderRepository) FetchAllOrderByUserID(userID, offset, limit, status int) ([]domain.Order, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAllOrderByUserID", userID, offset, limit, status)
	ret0, _ := ret[0].([]domain.Order)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FetchAllOrderByUserID indicates an expected call of FetchAllOrderByUserID.
func (mr *MockOrderRepositoryMockRecorder) FetchAllOrderByUserID(userID, offset, limit, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAllOrderByUserID", reflect.TypeOf((*MockOrderRepository)(nil).FetchAllOrderByUserID), userID, offset, limit, status)
}

// FetchOrderByID mocks base method.
func (m *MockOrderRepository) FetchOrderByID(id int64) (*domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchOrderByID", id)
	ret0, _ := ret[0].(*domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchOrderByID indicates an expected call of FetchOrderByID.
func (mr *MockOrderRepositoryMockRecorder) FetchOrderByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchOrderByID", reflect.TypeOf((*MockOrderRepository)(nil).FetchOrderByID), id)
}

// FetchOrderProductID mocks base method.
func (m *MockOrderRepository) FetchOrderProductID(orderProductID int64) (*domain.OrderProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchOrderProductID", orderProductID)
	ret0, _ := ret[0].(*domain.OrderProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchOrderProductID indicates an expected call of FetchOrderProductID.
func (mr *MockOrderRepositoryMockRecorder) FetchOrderProductID(orderProductID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchOrderProductID", reflect.TypeOf((*MockOrderRepository)(nil).FetchOrderProductID), orderProductID)
}

// FindOrderProductByOrderIDAndProductID mocks base method.
func (m *MockOrderRepository) FindOrderProductByOrderIDAndProductID(orderID int64, productID uuid.UUID) (*domain.OrderProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrderProductByOrderIDAndProductID", orderID, productID)
	ret0, _ := ret[0].(*domain.OrderProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrderProductByOrderIDAndProductID indicates an expected call of FindOrderProductByOrderIDAndProductID.
func (mr *MockOrderRepositoryMockRecorder) FindOrderProductByOrderIDAndProductID(orderID, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrderProductByOrderIDAndProductID", reflect.TypeOf((*MockOrderRepository)(nil).FindOrderProductByOrderIDAndProductID), orderID, productID)
}

// FindRatingSellerBySellerID mocks base method.
func (m *MockOrderRepository) FindRatingSellerBySellerID(sellerID int) (domain.SellerRating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindRatingSellerBySellerID", sellerID)
	ret0, _ := ret[0].(domain.SellerRating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRatingSellerBySellerID indicates an expected call of FindRatingSellerBySellerID.
func (mr *MockOrderRepositoryMockRecorder) FindRatingSellerBySellerID(sellerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRatingSellerBySellerID", reflect.TypeOf((*MockOrderRepository)(nil).FindRatingSellerBySellerID), sellerID)
}

// FindRatingSellerID mocks base method.
func (m *MockOrderRepository) FindRatingSellerID(sellerID int) (*domain.SellerRating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindRatingSellerID", sellerID)
	ret0, _ := ret[0].(*domain.SellerRating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRatingSellerID indicates an expected call of FindRatingSellerID.
func (mr *MockOrderRepositoryMockRecorder) FindRatingSellerID(sellerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRatingSellerID", reflect.TypeOf((*MockOrderRepository)(nil).FindRatingSellerID), sellerID)
}

// UpdateOrder mocks base method.
func (m *MockOrderRepository) UpdateOrder(id int64, request domain.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrder", id, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrder indicates an expected call of UpdateOrder.
func (mr *MockOrderRepositoryMockRecorder) UpdateOrder(id, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrder", reflect.TypeOf((*MockOrderRepository)(nil).UpdateOrder), id, request)
}

// UpdateOrderProduct mocks base method.
func (m *MockOrderRepository) UpdateOrderProduct(id int64, product domain.OrderProduct) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderProduct", id, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrderProduct indicates an expected call of UpdateOrderProduct.
func (mr *MockOrderRepositoryMockRecorder) UpdateOrderProduct(id, product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderProduct", reflect.TypeOf((*MockOrderRepository)(nil).UpdateOrderProduct), id, product)
}

// UpdateOrderStatus mocks base method.
func (m *MockOrderRepository) UpdateOrderStatus(id int64, status int, history domain.OrderHistory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderStatus", id, status, history)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrderStatus indicates an expected call of UpdateOrderStatus.
func (mr *MockOrderRepositoryMockRecorder) UpdateOrderStatus(id, status, history interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderStatus", reflect.TypeOf((*MockOrderRepository)(nil).UpdateOrderStatus), id, status, history)
}

// UpdateRatingSeller mocks base method.
func (m *MockOrderRepository) UpdateRatingSeller(seller domain.SellerRating) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRatingSeller", seller)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRatingSeller indicates an expected call of UpdateRatingSeller.
func (mr *MockOrderRepositoryMockRecorder) UpdateRatingSeller(seller interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRatingSeller", reflect.TypeOf((*MockOrderRepository)(nil).UpdateRatingSeller), seller)
}
