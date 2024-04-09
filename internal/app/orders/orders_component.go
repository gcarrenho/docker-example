package orders

import "context"

type OrdersComponent interface {
	FindOrderByOrderNumber(ctx context.Context, orderNumber string) (OrdersResponse, error)
}
