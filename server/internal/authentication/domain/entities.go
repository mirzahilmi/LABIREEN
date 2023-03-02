package domain

import (
	"time"
)

type Customer struct {
	ID          uint      `gorm:"primaryKey" json:"customer_id"`
	Name        string    `gorm:"type:varchar(50);not null" json:"name"`
	Email       string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"email"`
	Password    string    `gorm:"type:varchar(255);not null" json:"-"`
	PhoneNumber string    `gorm:"type:varchar(15);uniqueIndex;not null" json:"phone_number"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
