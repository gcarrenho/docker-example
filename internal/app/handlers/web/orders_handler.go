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

func NewOrdersHandlers(router *gin.RouterGroup, ordersComponent orders.OrdersComponent) {
	ordersInternalHandler := OrdersHandlers{
		ordersComponent: ordersComponent,
	}

	router.GET("orders/:orderNumber", ordersInternalHandler.getOrder)
}

func (oh *OrdersHandlers) getOrder(c *gin.Context) {
	orderNumber := c.Param("orderNumber")

	orderResponse, err := oh.ordersComponent.FindOrderByOrderNumber(c, orderNumber)
	if err != nil {
		log.Error().Err(err).Msg("")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "can not find order",
		})
		return
	}

	c.JSON(http.StatusOK, orderResponse)
}
