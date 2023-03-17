package routes

import (
	"labireen/handlers"
	"labireen/middleware"

	"github.com/gin-gonic/gin"
)

type MenuRoutes struct {
	Router      *gin.Engine
	MenuHandler handlers.MenuHandler
}

func (r *MenuRoutes) Register() {
	Menu := r.Router.Group("menu")
	Menu.POST("/create", r.MenuHandler.CreateMenu)
	Menu.GET("/view", r.MenuHandler.ViewMenu)
	Menu.GET("/view/:merchant-name", r.MenuHandler.ViewMenu)
	Menu.DELETE("/delete", middleware.ValidateToken(), r.MenuHandler.DeleteMenu)
}
