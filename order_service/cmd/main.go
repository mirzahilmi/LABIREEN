package main

import (
	"labireen/order_service/config"
	"labireen/order_service/handlers"
	"labireen/order_service/middleware"
	"labireen/order_service/repositories"
	"labireen/order_service/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(".env file loading failed")
	}

	// Initialize database connection
	db, err := config.GetDB()
	if err != nil {
		log.Fatalln("Database initialization failed")
	}

	// Auto migrate entities
	if err := config.Migrate(db); err != nil {
		log.Fatalln("Auto Migration failed")
	}

	menuService := services.NewMenuService(repositories.NewMenuRepository(db))
	menuHandler := handlers.NewMenuHandler(menuService)

	r := gin.Default()

	menu := r.Group("menu")
	menu.POST("create", middleware.ValidateToken(), menuHandler.CreateMenu)
	menu.GET("view", middleware.ValidateToken(), menuHandler.ViewMenu)

	r.Run(":" + os.Getenv("PORT"))
}
