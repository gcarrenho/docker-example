package orders

import (
	"context"
	"docker-example/internal/app/orders/model"
)

type OrdersComponentImpl struct {
	ordersRepository ordersRepository
}

func NewOrdersComponentImpl(ordersRepository ordersRepository) *OrdersComponentImpl {
	return &OrdersComponentImpl{
		ordersRepository: ordersRepository,
	}
}

func (o *OrdersComponentImpl) FindOrderByOrderNumber(ctx context.Context, orderNumber string) (model.OrdersResponse, error) {
	return o.ordersRepository.GetOrderByOrderNumber(ctx, orderNumber)
}
