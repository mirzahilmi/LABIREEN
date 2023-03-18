package entities

import "github.com/google/uuid"

type OrderRequestParams struct {
	OrderRequests OrderRequest `json:"order" binding:"required"`
}

type OrderRequest struct {
	MerchantID uuid.UUID          `json:"merchant_id" binding:"required"`
	CustomerID uuid.UUID          `json:"customer_id" binding:"required"`
	OrderItems []OrderItemRequest `json:"order_items" binding:"required"`
}

type OrderItemRequest struct {
	MenuItemID uuid.UUID `json:"menu_item_id" binding:"required"`
	Name       string    `json:"name" binding:"required"`
	Quantity   uint      `json:"quantity" binding:"required"`
	Price      int64     `json:"price" binding:"required"`
	Comment    string    `json:"comment,omitempty"`
}
