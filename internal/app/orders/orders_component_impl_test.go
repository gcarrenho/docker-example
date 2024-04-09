package orders

import (
	"context"
	"testing"

	"github.com/huandu/go-assert"
	"go.uber.org/mock/gomock"
)

func Test_UpdateProductPreset_GivenAnOrganizationIdAndProductPresetAssigment(t *testing.T) {
	ctx := context.Background()

	type params struct {
		orderNumber string
	}

	type want struct {
		err error
	}

	type mock struct {
		ordersRepository *mocks.MockOrdersRepository
	}

	tests := []struct {
		name  string
		input params
		want  want
		mock  func(m mock)
	}{
		{
			name: "Successful",
			input: params{
				orderNumber: orderNumber,
			},
			want: want{
				err: nil,
			},
			mock: func(m mock) {
				m.orderSvc.
					EXPECT().
					GetOrderByOrderNumber(gomock.AssignableToTypeOf(ctx)).
					Return(Orders{}, nil)

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			m := mock{
				ordersRepository: mocks.NewMockOrdersRepository(mockCtrl),
			}
			tt.mock(m)

			orderService := NewOrdersComponentImpl(
				m.ordersRepository,
			)

			err := orderService.FindOrderByOrderNumber(ctx, tt.input.orderNumber)
			assert.Equal(t, tt.want.err, err)
		})
	}
}