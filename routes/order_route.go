package routes

import (
	"labireen/handlers"
	"labireen/middleware"

	"github.com/gin-gonic/gin"
)

type OrderRoutes struct {
	Router       *gin.Engine
	OrderHandler handlers.OrderHandler
}

func (r *OrderRoutes) Register() {
	Order := r.Router.Group("order")
	Order.POST("/create", middleware.ValidateToken(), r.OrderHandler.CreateOrder)
}
