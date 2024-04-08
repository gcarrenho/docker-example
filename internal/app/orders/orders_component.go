package orders

type OrdersComponent interface {
	findOrderByID(id int64) (OrdersResponse, error)
}
