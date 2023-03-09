package main

import (
	"labireen/customer_microservices/account_service/config"
	"labireen/customer_microservices/account_service/handlers"
	"labireen/customer_microservices/account_service/middleware"
	"labireen/customer_microservices/account_service/repositories"
	"labireen/customer_microservices/account_service/services"
	"labireen/customer_microservices/account_service/utilities/mail"
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

	emailService := mail.NewGmailSender(os.Getenv("EMAIL_SENDER_NAME"), os.Getenv("EMAIL_SENDER_ADDRESS"), os.Getenv("EMAIL_SENDER_PASSWORD"))
	authService := services.NewAuthService(repositories.NewCustomerRepository(db))
	authHandler := handlers.NewAuthHandler(authService, emailService)

	customerService := services.NewCustomerService(repositories.NewCustomerRepository(db))
	customerHandler := handlers.NewCustomerHandler(customerService)

	r := gin.Default()

	auth := r.Group("auth")
	auth.POST("/register", authHandler.RegisterCustomer)
	auth.POST("/login", authHandler.LoginCustomer)
	auth.GET("/verify/:verification-code", authHandler.VerifyEmail)

	customer := r.Group("customer")
	customer.GET("/profile", middleware.ValidateToken(), customerHandler.GetMe)

	r.Run(":" + os.Getenv("PORT"))
}
