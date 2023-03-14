package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Menu struct {
	ID         uuid.UUID   `gorm:"primaryKey;autoIncrement:false"`
	MerchantID uuid.UUID   `gorm:"primaryKey;autoIncrement:false"`
	Name       string      `gorm:"not null"`
	MenuGroups []MenuGroup `gorm:"foreignKey:MenuID"`
	CreatedAt  time.Time   `gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `gorm:"autoUpdateTime"`
}

type MenuGroup struct {
	ID          uuid.UUID  `gorm:"primaryKey;autoIncrement:false"`
	Name        string     `gorm:"not null"`
	Description string     `gorm:"type:varchar(100)"`
	MenuID      uuid.UUID  `gorm:"not null;type:uuid;size:36"`
	MenuItems   []MenuItem `gorm:"foreignKey:MenuGroupID"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime"`
}

type MenuItem struct {
	ID          uuid.UUID       `gorm:"primaryKey;autoIncrement:false"`
	Name        string          `gorm:"not null"`
	Price       decimal.Decimal `gorm:"type:decimal(10,3)"`
	Description string          `gorm:"type:varchar(100)"`
	Stock       uint
	Photo       string    `gorm:"not null"`
	MenuGroupID uuid.UUID `gorm:"not null;type:uuid;size:36"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

type MenuRequestParams struct {
	MenuRequests MenuRequest `json:"menu" binding:"required"`
}

type MenuRequest struct {
	MerchantID uuid.UUID          `json:"merchant_id" binding:"required"`
	Name       string             `json:"menu_name" binding:"required"`
	MenuGroups []MenuGroupRequest `json:"menu_group" binding:"required"`
}

type MenuGroupRequest struct {
	Name        string            `json:"name" binding:"required"`
	Description string            `json:"description,omitempty" binding:"max=100"`
	MenuItems   []MenuItemRequest `json:"menu_item" binding:"required"`
}

type MenuItemRequest struct {
	Name        string          `json:"name" binding:"required"`
	Price       decimal.Decimal `json:"price" binding:"required"`
	Description string          `json:"description,omitempty" binding:"max=100"`
	Stock       uint            `json:"stock,omitempty"`
	Photo       string          `json:"photo,omitempty" binding:"url"`
}

func (u *MenuItem) BeforeCreate(tx *gorm.DB) error {
	if u.Photo == "" {
		u.Photo = "https://kagqhqxcmsnypptkpxqj.supabase.co/storage/v1/object/sign/menu-photo/image-placeholder.svg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJtZW51LXBob3RvL2ltYWdlLXBsYWNlaG9sZGVyLnN2ZyIsImlhdCI6MTY3ODUxMDcxNywiZXhwIjoxNzEwMDQ2NzE3fQ.ozW-gvApLPeIpxMr22mBwbYNdq0wbJ36w5Dci8goi0M&t=2023-03-11T04%3A58%3A38.086Z"
	}
	return nil
}
