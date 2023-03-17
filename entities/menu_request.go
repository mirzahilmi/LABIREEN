package entities

import "github.com/google/uuid"

type MenuRequestParams struct {
	MenuRequests MenuRequest `json:"menu" binding:"required"`
}

type MenuRequest struct {
	MerchantID uuid.UUID          `json:"merchant_id,omitempty"`
	NMID       string             `json:"nmid,omitempty"`
	Name       string             `json:"menu_name" binding:"required"`
	MenuGroups []MenuGroupRequest `json:"menu_group" binding:"required"`
}

type MenuGroupRequest struct {
	Name        string            `json:"name" binding:"required"`
	Description string            `json:"description,omitempty" binding:"max=100"`
	MenuItems   []MenuItemRequest `json:"menu_item" binding:"required"`
}

type MenuItemRequest struct {
	Name        string `json:"name" binding:"required"`
	Price       int64  `json:"price" binding:"required"`
	Description string `json:"description,omitempty" binding:"max=100"`
	Stock       uint   `json:"stock,omitempty"`
	Photo       string `json:"photo,omitempty" binding:"url"`
}
