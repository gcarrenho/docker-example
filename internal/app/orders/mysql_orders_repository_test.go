package orders

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T, fn func()) (
	client *mysqlOrdersRepository,
	sqlMock sqlmock.Sqlmock,
	tearDown func(),
) {
	t.Helper()

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error'%s' was not expected when opening a stub database connection", err)
	}

	client = NewMYSQLOrdersRepository(db)

	return client, sqlMock, func() {
		defer db.Close()
	}
}

func TestOrderRepository(t *testing.T) {
	for scenario, fn := range map[string]func(
		t *testing.T,
		client *mysqlOrdersRepository,
		sqlMock sqlmock.Sqlmock,
	){
		"GetOrderByOrderNumber": testGetOrderByOrderNumber,
	} {
		t.Run(scenario, func(t *testing.T) {
			client, sqlMock, tearDown := setupTest(t, nil)
			defer tearDown()
			fn(t, client, sqlMock)
		})
	}
}

func testGetOrderByOrderNumber(t *testing.T, client *mysqlOrdersRepository, sqlMock sqlmock.Sqlmock) {
	type want struct {
		err           error
		orderResponse OrdersResponse
	}

	ctx := context.Background()
	orderDate, _ := time.Parse(time.RFC3339Nano, "2022-12-21T00:18:44.8656209+01:00")

	tests := []struct {
		name        string
		orderNumber string
		want        want
		mocks       func(m sqlmock.Sqlmock)
	}{
		{

			name:        "Get order number fails",
			orderNumber: "1",
			want:        want{err: errors.New("some error"), orderResponse: OrdersResponse{}},
			mocks: func(m sqlmock.Sqlmock) {
				sqlMock.ExpectQuery("SELECT orderNumber, currencyCode, amount, created_at FROM orders WHERE orderNumber = ?").WillReturnError(errors.New("some error"))
			},
		},
		{

			name:        "Get order number Successful",
			orderNumber: "1",
			want: want{err: nil, orderResponse: OrdersResponse{
				OrderNumber: "1",
				Currency:    "USD",
				Amount:      10,
				CreatedAt:   orderDate,
			}},
			mocks: func(m sqlmock.Sqlmock) {
				rowsFetch := sqlmock.NewRows([]string{"orderNumber", "currencyCode", "amount", "created_at"}).
					AddRow("1", "USD", "10", orderDate)

				sqlMock.ExpectQuery("SELECT orderNumber, currencyCode, amount, created_at FROM orders WHERE orderNumber = ?").WithArgs(
					"1",
				).WillReturnRows(rowsFetch)
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			tc.mocks(sqlMock)

			orderResponse, err := client.getOrderByOrderNumber(ctx, tc.orderNumber)

			require.Equal(t, tc.want.err, err)
			require.Equal(t, tc.want.orderResponse, orderResponse)
		})
	}

}
