package orders

import (
	"context"
	"docker-example/internal/app/orders/model"
)

type OrdersComponent interface {
	FindOrderByOrderNumber(ctx context.Context, orderNumber string) (model.OrdersResponse, error)
}
