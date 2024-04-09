package orders

type OrdersRepository interface {
	getOrderByOrderNUmber(orderNumber string) (OrdersResponse, error)
}
