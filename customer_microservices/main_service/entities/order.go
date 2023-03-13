package entities

import "github.com/google/uuid"

type OrderParam struct {
	ID uuid.UUID
	CustomerID uuid.UUID  `gorm:"foreignKey"`
}
