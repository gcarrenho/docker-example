package orders

import (
	"context"
	"database/sql"
	"docker-example/internal/app/orders/model"
	errordetails "github.com/gcarrenho/common-libs/pkg/errordetails"
)

type mysqlOrdersRepository struct {
	db *sql.DB
}

func NewMYSQLOrdersRepository(db *sql.DB) *mysqlOrdersRepository {
	return &mysqlOrdersRepository{
		db: db,
	}
}

func (o *mysqlOrdersRepository) GetOrderByOrderNumber(ctx context.Context, orderNumber string) (model.OrdersResponse, error) {
	query := "SELECT orderNumber, currencyCode, amount, created_at FROM orders WHERE orderNumber = ?"
	row := o.db.QueryRowContext(ctx, query, orderNumber)

	var order model.OrdersResponse

	err := row.Scan(&order.OrderNumber, &order.Currency, &order.Amount, &order.CreatedAt)
	if err != nil {
		return model.OrdersResponse{}, errordetails.NewErrorDetails(err).Str("orderNumber", orderNumber).Msg("when scan query")
	}

	return order, nil
}
