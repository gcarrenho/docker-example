package orders

type OrdersComponentImpl struct {
	ordersRepository mysqlOrdersRepository
}

func NewOrdersComponentImpl(ordersRepository mysqlOrdersRepository) *OrdersComponentImpl {
	return &OrdersComponentImpl{
		ordersRepository: ordersRepository,
	}
}

func (o *OrdersComponentImpl) FindOrderByOrderNumber(orderNumber string) (OrdersResponse, error) {
	return o.ordersRepository.getOrderByOrderNUmber(orderNumber)
}
