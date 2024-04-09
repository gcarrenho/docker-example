package orders

import (
	"context"
	"docker-example/mocks"
	"errors"
	"testing"

	"docker-example/internal/app/orders/model"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type mockOrdersComponent struct {
	ordersRepository *mocks.MockordersRepository
}

func setupOrderComponentTest(t *testing.T, fn func()) (
	ordersComponent *OrdersComponentImpl,
	mock mockOrdersComponent,
	tearDown func(),
) {
	t.Helper()

	mockCtrl := gomock.NewController(t)

	mock = mockOrdersComponent{
		ordersRepository: mocks.NewMockordersRepository(mockCtrl),
	}

	ordersComponent = NewOrdersComponentImpl(
		mock.ordersRepository,
	)

	return ordersComponent, mock, func() {
		defer mockCtrl.Finish()
	}
}

func TestOrdersComponent(t *testing.T) {
	for scenario, fn := range map[string]func(
		t *testing.T,
		ordersComponent *OrdersComponentImpl,
		mock mockOrdersComponent,
	){
		"OrdersComponent": testFindOrderByOrderNumber,
	} {
		t.Run(scenario, func(t *testing.T) {
			ordersComponent, mock, tearDown := setupOrderComponentTest(t, nil)
			defer tearDown()
			fn(t, ordersComponent, mock)
		})
	}
}

func testFindOrderByOrderNumber(t *testing.T, ordersComponent *OrdersComponentImpl, mock mockOrdersComponent) {
	ctx := context.Background()

	type params struct {
		orderNumber string
	}

	type want struct {
		ordersResponse model.OrdersResponse
		err            error
	}

	tests := []struct {
		name  string
		input params
		want  want
		mock  func(m mockOrdersComponent)
	}{
		{
			name: "When FindOrderByOrderNumber fails",
			input: params{
				orderNumber: "1",
			},
			want: want{
				err: errors.New("some error"),
			},
			mock: func(m mockOrdersComponent) {
				m.ordersRepository.EXPECT().GetOrderByOrderNumber(ctx, "1").Return(model.OrdersResponse{}, errors.New("some error"))
			},
		},
		{
			name: "When FindOrderByOrderNumber is Successful",
			input: params{
				orderNumber: "1",
			},
			want: want{
				ordersResponse: model.OrdersResponse{
					OrderNumber: "1",
				},
				err: nil,
			},
			mock: func(m mockOrdersComponent) {
				m.ordersRepository.EXPECT().GetOrderByOrderNumber(ctx, "1").Return(model.OrdersResponse{OrderNumber: "1"}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(mock)
			order, err := ordersComponent.FindOrderByOrderNumber(ctx, tt.input.orderNumber)
			assert.Equal(t, tt.want.err, err)
			assert.Equal(t, tt.want.ordersResponse, order)
		})
	}
}
