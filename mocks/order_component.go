// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/orders/orders_component.go
//
// Generated by this command:
//
//	mockgen -source=internal/app/orders/orders_component.go -destination=mocks/order_component.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	orders "docker-example/internal/app/orders"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockOrdersComponent is a mock of OrdersComponent interface.
type MockOrdersComponent struct {
	ctrl     *gomock.Controller
	recorder *MockOrdersComponentMockRecorder
}

// MockOrdersComponentMockRecorder is the mock recorder for MockOrdersComponent.
type MockOrdersComponentMockRecorder struct {
	mock *MockOrdersComponent
}

// NewMockOrdersComponent creates a new mock instance.
func NewMockOrdersComponent(ctrl *gomock.Controller) *MockOrdersComponent {
	mock := &MockOrdersComponent{ctrl: ctrl}
	mock.recorder = &MockOrdersComponentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrdersComponent) EXPECT() *MockOrdersComponentMockRecorder {
	return m.recorder
}

// FindOrderByOrderNumber mocks base method.
func (m *MockOrdersComponent) FindOrderByOrderNumber(ctx context.Context, orderNumber string) (orders.OrdersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrderByOrderNumber", ctx, orderNumber)
	ret0, _ := ret[0].(orders.OrdersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrderByOrderNumber indicates an expected call of FindOrderByOrderNumber.
func (mr *MockOrdersComponentMockRecorder) FindOrderByOrderNumber(ctx, orderNumber any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrderByOrderNumber", reflect.TypeOf((*MockOrdersComponent)(nil).FindOrderByOrderNumber), ctx, orderNumber)
}
