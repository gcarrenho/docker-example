package web

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/huandu/go-assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type mockOrderHandler struct {
	orderSvc *mocks.MockOrdersSvc
}

/*w := httptest.NewRecorder()
router := gin.Default()
v1Group := router.Group("/", func(c *gin.Context) {
	if tc.ctxUser != nil {
		c.Set(model.CtxKeyUser, *tc.ctxUser)
	}

	c.Next()
})

mockCtrl := gomock.NewController(t)
defer mockCtrl.Finish()

m := mockPresetHandler{
	jsonSchemaValidator: mocks.NewMockJsonSchemaValidator(mockCtrl),
	presetSvc:           mocks.NewMockPresetServices(mockCtrl),
}

tc.mocks(m)

NewPresetHandler(v1Group, m.jsonSchemaValidator, m.presetSvc, zerolog.Nop())*/

/*func setupTest(t *testing.T, fn func()) (
	//client *http,
	mockOrdersHandler mockOrderHandler,
	rgroup *gin.RouterGroup,

	tearDown func(),
) {
	t.Helper()

	mockOrdersHandler = mockOrderHandler{
		ordersSvc: mocks.NewMockOrderComponent(mockCtrl),
			}

	//client = NewMYSQLOrdersRepository(db)

	return mockOrdersHandler, rgroup, func() {
		defer db.Close()
	}
}*/

func TestOrderHandler(t *testing.T) {
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

func TestGetOrdersHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	type want struct {
		expected string
		httpCode int
		err      error
	}

	tests := []struct {
		name  string
		want  want
		mocks func(m mockPresetHandler)
	}{
		{
			name: "Succes",

			want: want{expected: ``, httpCode: http.StatusOK},
			mocks: func(m mockPresetHandler) {

				m.presetSvc.EXPECT().GetOrde(gomock.Any(), user.OrganizationId).Return(mockOrder, nil)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			router := gin.Default()
			v1Group := router.Group("/", func(c *gin.Context) {
				if tc.ctxUser != nil {
					c.Set(model.CtxKeyUser, *tc.ctxUser)
				}

				c.Next()
			})

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			m := mockOrdersHandler{
				ordersSvc: mocks.NewMockOrdersSvc(mockCtrl),
			}

			tc.mocks(m)

			NewOrdersHandler(v1Group, m.ordersSvc)

			req, err := http.NewRequest(http.MethodGet, tc.url, nil)
			require.NoError(t, err)

			router.ServeHTTP(w, req)

			assert.Equal(t, tc.want.httpCode, w.Code)
			require.JSONEq(t, tc.want.expected, w.Body.String())
		})
	}
}
