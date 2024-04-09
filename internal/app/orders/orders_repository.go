package orders

import "context"

type OrdersRepository interface {
	getOrderByOrderNUmber(ctx context.Context, orderNumber string) (OrdersResponse, error)
}
