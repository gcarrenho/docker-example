package orders

import "context"

type ordersRepository interface {
	getOrderByOrderNumber(ctx context.Context, orderNumber string) (OrdersResponse, error)
}
