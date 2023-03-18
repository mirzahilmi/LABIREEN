package entities

import "github.com/google/uuid"

type MenuRequestParams struct {
	MenuRequests MenuRequest `json:"menu" binding:"required"`
}

type MenuRequest struct {
	MerchantID uuid.UUID          `json:"merchant_id,omitempty"`
	Name       string             `json:"menu_name" binding:"required"`
	MenuGroups []MenuGroupRequest `json:"menu_group" binding:"required"`
}

type MenuGroupRequest struct {
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	MenuItems   []MenuItemRequest `json:"menu_item" binding:"required"`
}

type MenuItemRequest struct {
	Name        string `json:"name,omitempty"`
	Price       int64  `json:"price,omitempty"`
	Description string `json:"description,omitempty" binding:"max=100"`
	Stock       uint   `json:"stock,omitempty"`
	Photo       string `json:"photo,omitempty" binding:"url"`
}

type MenuRegisterParams struct {
	MenuRegister MenuRegister `json:"menu" binding:"required"`
}

type MenuRegister struct {
	MerchantID uuid.UUID           `json:"merchant_id,omitempty"`
	Name       string              `json:"menu_name" binding:"required"`
	MenuGroups []MenuGroupRegister `json:"menu_group" binding:"required"`
}

type MenuGroupRegister struct {
	Name        string             `json:"name" binding:"required"`
	Description string             `json:"description,omitempty" binding:"max=100"`
	MenuItems   []MenuItemRegister `json:"menu_item" binding:"required"`
}

type MenuItemRegister struct {
	Name        string `json:"name" binding:"required"`
	Price       int64  `json:"price" binding:"required"`
	Description string `json:"description,omitempty" binding:"max=100"`
	Stock       uint   `json:"stock,omitempty"`
	Photo       string `json:"photo,omitempty" binding:"url"`
}
