package config

import (
	"fmt"
	"labireen/entities"
	"os"

	"github.com/midtrans/midtrans-go"
	"gorm.io/driver/mysql"

	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

// GetDB returns a connection to the database
func GetDB() (*gorm.DB, error) {
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}

	// MySQL Configuration
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Name)

	// PostgreSQL Configuration
	// dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
	// 	dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func SetMidtrans(key string) {
	midtrans.ServerKey = key
	midtrans.Environment = midtrans.Sandbox
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entities.Menu{},
		&entities.MenuGroup{},
		&entities.MenuItem{},
		&entities.Order{},
		&entities.OrderItem{},
		&entities.OrderStatus{},
	)
}
