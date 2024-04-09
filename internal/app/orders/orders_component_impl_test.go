package orders

/*
import "docker-example/mocks"

type mockOrdersComponent struct {
	ordersRepository *mocks.MockordersRepository
}

/*
func setupOrderComponentTest(t *testing.T, fn func()) (
	ordersComponent *OrdersComponentImpl,
	tearDown func(),
) {
	t.Helper()

	mockCtrl := gomock.NewController(t)

	mock := mockOrdersComponent{
		ordersRepository: mocks.NewMockordersRepository(mockCtrl),
	}

	ordersComponent = NewOrdersComponentImpl(
		//mock.ordersRepository,
	)

	return ordersComponent, func() {
		defer mockCtrl.Finish()
	}
}

func TestOrdersComponent(t *testing.T) {
	for scenario, fn := range map[string]func(
		t *testing.T,
		ordersComponent *OrdersComponentImpl,
	){
		"OrdersHandler": testOrdercompGetOrderByOrderNumber,
	} {
		t.Run(scenario, func(t *testing.T) {
			ordersComponent, tearDown := setupOrderComponentTest(t, nil)
			defer tearDown()
			fn(t, ordersComponent)
		})
	}
}

func testOrdercompGetOrderByOrderNumber(t *testing.T, ordersComponent *OrdersComponentImpl) {
	ctx := context.Background()

	type params struct {
		orderNumber string
	}

	type want struct {
		ordersResponse OrdersResponse
		err            error
	}

	tests := []struct {
		name  string
		input params
		want  want
		mock  func(m mockOrdersComponent)
	}{
		{
			name: "Successful",
			input: params{
				orderNumber: "1",
			},
			want: want{
				err: nil,
			},
			mock: func(m mockOrdersComponent) {


			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order, err := ordersComponent.FindOrderByOrderNumber(ctx, tt.input.orderNumber)
			assert.Equal(t, tt.want.err, err)
			assert.Equal(t, tt.want.ordersResponse, order)
		})
	}
}
*/
