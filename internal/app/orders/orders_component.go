package orders

type OrdersComponent interface {
	FindOrderByOrderNumber(orderNumber string) (OrdersResponse, error)
}
