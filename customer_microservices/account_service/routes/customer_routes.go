package routes

import (
	"labireen/customer_microservices/account_service/handlers"
	"labireen/customer_microservices/account_service/middleware"

	"github.com/gin-gonic/gin"
)

type CustomerRoutes struct {
	Router          *gin.Engine
	CustomerHandler handlers.CustomerHandler
}

func (r *CustomerRoutes) Register() {
	customer := r.Router.Group("customer")
	customer.GET("/profile", middleware.ValidateToken(), r.CustomerHandler.GetMe)
}
