package entities

import (
	"time"

	"github.com/google/uuid"
)

type Menu struct {
	ID         uuid.UUID   `gorm:"primaryKey;autoIncrement:false"`
	MerchantID uuid.UUID   `gorm:"not null;size:36"`
	NMID       string      `gorm:"not null"`
	Name       string      `gorm:"not null"`
	CreatedAt  time.Time   `gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `gorm:"autoUpdateTime"`
	MenuGroups []MenuGroup `gorm:"foreignKey:MenuID"`
}

type MenuGroup struct {
	ID          uuid.UUID  `gorm:"primaryKey;autoIncrement:false"`
	Name        string     `gorm:"not null"`
	Description string     `gorm:"type:varchar(100)"`
	MenuID      uuid.UUID  `gorm:"not null;type:uuid;size:36"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime"`
	MenuItems   []MenuItem `gorm:"foreignKey:MenuGroupID"`
}

type MenuItem struct {
	ID          uuid.UUID   `gorm:"primaryKey;autoIncrement:false"`
	Name        string      `gorm:"not null"`
	Price       int64       `gorm:"not null"`
	Description string      `gorm:"type:varchar(100)"`
	Stock       uint        `gorm:"check:stock >= 0"`
	Photo       string      `gorm:"not null"`
	MenuGroupID uuid.UUID   `gorm:"not null;type:uuid;size:36"`
	CreatedAt   time.Time   `gorm:"autoCreateTime"`
	UpdatedAt   time.Time   `gorm:"autoUpdateTime"`
	OrderItems  []OrderItem `gorm:"foreignKey:MenuItemID"`
}
