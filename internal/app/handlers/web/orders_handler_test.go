package web

import (
	"docker-example/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"docker-example/internal/app/orders/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type mockOrdersHandler struct {
	ordersComponent *mocks.MockOrdersComponent
}

func setupTest(t *testing.T, fn func()) (
	mock mockOrdersHandler,
	responseRecorder *httptest.ResponseRecorder,
	router *gin.Engine,
	tearDown func(),
) {
	t.Helper()

	gin.SetMode(gin.TestMode)

	responseRecorder = httptest.NewRecorder()
	router = gin.Default()
	v1Group := router.Group("/")

	mockCtrl := gomock.NewController(t)

	mock = mockOrdersHandler{
		ordersComponent: mocks.NewMockOrdersComponent(mockCtrl),
	}

	NewOrdersHandlers(v1Group, mock.ordersComponent)

	return mock, responseRecorder, router, func() {
		defer mockCtrl.Finish()
	}
}

func TestOrdersHandler(t *testing.T) {
	for scenario, fn := range map[string]func(
		t *testing.T,
		mock mockOrdersHandler,
		responseRecorder *httptest.ResponseRecorder,
		route *gin.Engine,
	){
		"OrdersHandler": testGetOrder,
	} {
		t.Run(scenario, func(t *testing.T) {
			mock, responseRecorder, router, tearDown := setupTest(t, nil)
			defer tearDown()
			fn(t, mock, responseRecorder, router)
		})
	}
}

func testGetOrder(t *testing.T, mock mockOrdersHandler, responseRecorder *httptest.ResponseRecorder, router *gin.Engine,
) {
	type want struct {
		expected string
		httpCode int
	}

	tests := []struct {
		name  string
		want  want
		mocks func(m mockOrdersHandler)
	}{
		{
			name: "When FindOrderByOrderNumber fails",

			want: want{expected: `{"error":"can not find the order"}`, httpCode: http.StatusInternalServerError},
			mocks: func(m mockOrdersHandler) {
				m.ordersComponent.EXPECT().FindOrderByOrderNumber(gomock.Any(), "1").Return(model.OrdersResponse{}, errors.New("some error"))
			},
		},
		{
			name: "When FindOrderByOrderNumber is Succes",

			want: want{expected: `{"order_number":"","currency":"","amount":0,"creates_at":"0001-01-01T00:00:00Z"}`, httpCode: http.StatusOK},
			mocks: func(m mockOrdersHandler) {
				m.ordersComponent.EXPECT().FindOrderByOrderNumber(gomock.Any(), "1").Return(model.OrdersResponse{}, nil)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mocks(mock)
			responseRecorder := httptest.NewRecorder()

			req, err := http.NewRequest(http.MethodGet, "/orders/1", nil)
			require.NoError(t, err)

			router.ServeHTTP(responseRecorder, req)

			assert.Equal(t, tc.want.httpCode, responseRecorder.Code)
			require.JSONEq(t, tc.want.expected, responseRecorder.Body.String())
		})
	}
}
