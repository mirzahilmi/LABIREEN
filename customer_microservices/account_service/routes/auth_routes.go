package routes

import (
	"labireen/customer_microservices/account_service/handlers"

	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	Router      *gin.Engine
	AuthHandler handlers.AuthHandler
}

func (r *AuthRoutes) Register() {
	auth := r.Router.Group("auth")
	auth.POST("/register", r.AuthHandler.RegisterCustomer)
	auth.POST("/login", r.AuthHandler.LoginCustomer)
	auth.GET("/verify/:verification-code", r.AuthHandler.VerifyEmail)
}
