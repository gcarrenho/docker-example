package web

import (
	"docker-example/internal/app/orders"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrdersHandlers struct {
	ordersComponent orders.OrdersComponent
}

func NewOrdersHandlers(router *gin.Engine, ordersComponent orders.OrdersComponent) {
	ordersInternalHandler := OrdersHandlers{
		ordersComponent: ordersComponent,
	}

	router.GET("/orders/:id", ordersInternalHandler.getOrder)
}

func (oh *OrdersHandlers) getOrder(c *gin.Context) {
	ID := c.Param("id")

	strID, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid id")
		return
	}

	fmt.Printf("oh.ordersComponent: %v\n", strID)
}
