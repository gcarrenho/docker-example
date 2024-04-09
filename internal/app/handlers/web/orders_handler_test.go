package web

import (
	"docker-example/internal/app/orders"
	"docker-example/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type mockOrdersHandler struct {
	ordersComponent *mocks.MockOrdersComponent
}

func setupTest(t *testing.T, fn func()) (
	//client *http,
	mock mockOrdersHandler,
	responseRecorder *httptest.ResponseRecorder,
	router *gin.Engine,

	tearDown func(),
) {
	t.Helper()

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
		"OrdersHandler": testGetOrderByOrderNumber,
	} {
		t.Run(scenario, func(t *testing.T) {
			mock, responseRecorder, router, tearDown := setupTest(t, nil)
			defer tearDown()
			fn(t, mock, responseRecorder, router)
		})
	}
}

func testGetOrderByOrderNumber(t *testing.T, mock mockOrdersHandler, responseRecorder *httptest.ResponseRecorder, router *gin.Engine) {

	gin.SetMode(gin.TestMode)

	type want struct {
		expected string
		httpCode int
		err      error
	}

	tests := []struct {
		name  string
		want  want
		mocks func(m mockOrdersHandler)
	}{
		{
			name: "Succes",

			want: want{expected: ``, httpCode: http.StatusOK},
			mocks: func(m mockOrdersHandler) {
				m.ordersComponent.EXPECT().FindOrderByOrderNumber(gomock.Any(), "1").Return(orders.OrdersResponse{}, nil)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mocks(mock)

			req, err := http.NewRequest(http.MethodGet, "/orders/1", nil)
			require.NoError(t, err)

			router.ServeHTTP(responseRecorder, req)

			assert.Equal(t, tc.want.httpCode, responseRecorder.Code)
			//require.JSONEq(t, tc.want.expected, w.Body.String())
		})
	}
}
