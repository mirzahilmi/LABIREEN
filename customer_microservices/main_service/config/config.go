package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	User     string
	Password string
	Name     string
}

// GetDB returns a connection to the database
func GetDB() (*gorm.DB, error) {
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	return db, nil
}

func ConnectDB() {

}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate()
}
