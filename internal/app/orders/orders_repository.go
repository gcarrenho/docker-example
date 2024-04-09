package orders

import (
	"context"
	"docker-example/internal/app/orders/model"
)

type ordersRepository interface {
	GetOrderByOrderNumber(ctx context.Context, orderNumber string) (model.OrdersResponse, error)
}
