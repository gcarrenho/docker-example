package orders

type OrdersRepository interface {
	getOrderByID(id int64) (OrdersResponse, error)
}
