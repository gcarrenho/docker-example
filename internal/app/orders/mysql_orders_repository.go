package orders

import (
	"database/sql"
)

type mysqlOrdersRepository struct {
	db *sql.DB
}

func NewMYSQLOrdersRepository(db *sql.DB) *mysqlOrdersRepository {
	return &mysqlOrdersRepository{
		db: db,
	}
}

func (o *mysqlOrdersRepository) getOrderByOrderNUmber(orderNumber string) (OrdersResponse, error) {
	query := "SELECT orderNumber, currencyCode, amount, created_at FROM orders WHERE orderNumber = ?"
	row := o.db.QueryRow(query, orderNumber)

	var order OrdersResponse

	err := row.Scan(&order.OrderNumber, &order.Currency, &order.Amount, &order.CreatedAt)
	if err != nil {
		return OrdersResponse{}, err
	}

	return order, nil
}
