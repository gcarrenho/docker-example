package web

import (
	"docker-example/internal/app/orders"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type OrdersHandlers struct {
	ordersComponent orders.OrdersComponent
}

func NewOrdersHandlers(router *gin.Engine, ordersComponent orders.OrdersComponent) {
	ordersInternalHandler := OrdersHandlers{
		ordersComponent: ordersComponent,
	}

	router.GET("/orders/:orderNumber", ordersInternalHandler.getOrder)
}

func (oh *OrdersHandlers) getOrder(c *gin.Context) {
	orderNumber := c.Param("orderNumber")

	orderResponse, err := oh.ordersComponent.FindOrderByOrderNumber(c, orderNumber)
	if err != nil {
		log.Error().Msg("Finding order by order number: " + err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, "can not find the order")
		return
	}

	c.JSON(http.StatusOK, orderResponse)
}
