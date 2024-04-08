package orders

type OrdersComponentImpl struct {
	ordersRepository mysqlOrdersRepository
}

func NewOrdersComponentImpl(ordersRepository mysqlOrdersRepository) *OrdersComponentImpl {
	return &OrdersComponentImpl{
		ordersRepository: ordersRepository,
	}
}

func (o *OrdersComponentImpl) findOrderByID(id int64) (OrdersResponse, error) {
	return o.ordersRepository.getOrderByID(id)
}
