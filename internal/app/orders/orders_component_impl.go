package orders

import "context"

type OrdersComponentImpl struct {
	ordersRepository mysqlOrdersRepository
}

func NewOrdersComponentImpl(ordersRepository mysqlOrdersRepository) *OrdersComponentImpl {
	return &OrdersComponentImpl{
		ordersRepository: ordersRepository,
	}
}

func (o *OrdersComponentImpl) FindOrderByOrderNumber(ctx context.Context, orderNumber string) (OrdersResponse, error) {
	return o.ordersRepository.getOrderByOrderNumber(ctx, orderNumber)
}
