package main

import (
	"labireen/customer_microservices/account_service/config"
	"labireen/customer_microservices/account_service/handlers"
	"labireen/customer_microservices/account_service/repositories"
	"labireen/customer_microservices/account_service/routes"
	"labireen/customer_microservices/account_service/services"
	"labireen/customer_microservices/account_service/utilities/mail"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	app *gin.Engine
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

	emailService := mail.NewGmailSender(os.Getenv("EMAIL_SENDER_NAME"), os.Getenv("EMAIL_SENDER_ADDRESS"), os.Getenv("EMAIL_SENDER_PASSWORD"))
	authService := services.NewAuthService(repositories.NewCustomerRepository(db))
	customerService := services.NewCustomerService(repositories.NewCustomerRepository(db))

	authHandler := handlers.NewAuthHandler(authService, emailService)
	customerHandler := handlers.NewCustomerHandler(customerService)

	// Register auth routes
	authRoutes := routes.AuthRoutes{
		Router:      app,
		AuthHandler: authHandler,
	}
	authRoutes.Register()

	// Register customer routes
	customerRoutes := routes.CustomerRoutes{
		Router:          app,
		CustomerHandler: customerHandler,
	}
	customerRoutes.Register()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
