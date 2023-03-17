package main

import (
	"labireen/config"
	"labireen/handlers"
	"labireen/repositories"
	"labireen/routes"
	"labireen/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
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

	cr := coreapi.Client{}
	cr.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	menuService := services.NewMenuService(repositories.NewMenuRepository(db))
	menuHandler := handlers.NewMenuHandler(menuService)

	orderService := services.NewOrderService(cr, repositories.NewOrderRepository(db))
	orderHandler := handlers.NewOrderHandler(orderService)

	app := gin.Default()

	// Register menu routes
	menuRoutes := routes.MenuRoutes{
		Router:      app,
		MenuHandler: menuHandler,
	}
	menuRoutes.Register()

	// Register order routes
	orderRoutes := routes.OrderRoutes{
		Router:       app,
		OrderHandler: orderHandler,
	}
	orderRoutes.Register()

	app.Run("127.0.0.2:8080")
}
