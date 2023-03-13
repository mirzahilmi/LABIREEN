package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Order struct {
	ID            uuid.UUID       `gorm:"primaryKey;autoIncrement:false"`
	OrderItemID   []OrderItem     `gorm:"not null"`
	OrderStatusID OrderStatus     `gorm:"not null"`
	MerchantID    uuid.UUID       `gorm:"not null"`
	Gross         decimal.Decimal `gorm:"type:decimal(10,3);not null"`
	OrderPlaced   time.Time       `gorm:"autoCreateTime"`
	OrderPaid     time.Time
}

type OrderItem struct {
	ID         uuid.UUID       `gorm:"primaryKey;autoIncrement:false"`
	MenuItemID MenuItem        `gorm:"not null"`
	Quantity   uint            `gorm:"not null"`
	Price      decimal.Decimal `gorm:"type:decimal(10,3);not null"`
	Comments   string          `gorm:"not null;type:varchar(50)"`
}

type OrderStatus struct {
	ID        uuid.UUID `gorm:"primaryKey;autoIncrement:false"`
	Prepared  bool      `gorm:"default:false"`
	Finished  bool      `gorm:"default:false"`
	Delivered bool      `gorm:"default:false"`
	Cancelled bool      `gorm:"default:false"`
}
