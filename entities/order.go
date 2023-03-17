package entities

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID            uuid.UUID   `gorm:"primaryKey;autoIncrement:false"`
	OrderStatuses OrderStatus `gorm:"foreignKey:OrderID"`
	MerchantID    uuid.UUID   `gorm:"not null"`
	CustomerID    uuid.UUID   `gorm:"not null"`
	NMID          string      `gorm:"not null"`
	Gross         int64       `gorm:"not null"`
	Paid          bool        `gorm:"default:false"`
	OrderPlaced   time.Time   `gorm:"autoCreateTime"`
	OrderItems    []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID         uuid.UUID `gorm:"primaryKey;autoIncrement:false"`
	OrderID    uuid.UUID `gorm:"not null;type:uuid;size:36"`
	MenuItemID uuid.UUID `gorm:"not null;type:uuid;size:36"`
	Name       string    `gorm:"not null"`
	Quantity   uint      `gorm:"not null"`
	Price      int64     `gorm:"not null"`
	Comment    string    `gorm:"not null;type:varchar(50)"`
}

type OrderStatus struct {
	ID        uuid.UUID `gorm:"primaryKey;autoIncrement:false"`
	OrderID   uuid.UUID `gorm:"not null;type:uuid;size:36"`
	Prepared  bool      `gorm:"default:false"`
	Finished  bool      `gorm:"default:false"`
	Delivered bool      `gorm:"default:false"`
	Cancelled bool      `gorm:"default:false"`
}
