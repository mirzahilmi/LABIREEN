package main

import (
	"labireen/server/internal/authentication/config"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("failed to load env file")
	}
	port := os.Getenv("PORT")

	// Initialize database connection
	db, err := config.GetDB()
	if err != nil {
		log.Fatalln("database initialization failed")
	}

	// Auto migrate entities
	if err := config.Migrate(db); err != nil {
		log.Fatalln("auto migrate error,", err)
	}

	r := gin.Default()

	r.Run(":" + port)
}
