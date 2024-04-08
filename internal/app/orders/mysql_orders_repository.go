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

func (o *mysqlOrdersRepository) getOrderByID(id int64) (OrdersResponse, error) {
	// Ejecutar la consulta SQL para obtener la orden por ID
	query := "SELECT * FROM orders WHERE id = ?"
	row := o.db.QueryRow(query, id)

	var order OrdersResponse

	err := row.Scan(&order.OrderNumber)
	if err != nil {
		// Manejar el error si ocurre
		return OrdersResponse{}, err
	}

	return order, nil
}
