package services

import (
	"fmt"
	"labireen/entities"
	"labireen/repositories"
	"time"

	"github.com/dchest/uniuri"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type OrderService interface {
	RegisterOrder(order entities.OrderRequestParams, id uuid.UUID) (coreapi.ChargeResponse, error)
	FindOrder(param string, id uuid.UUID) (entities.Order, error)
	EditOrder(order entities.OrderRequestParams) error
	DeleteOrder(id uuid.UUID) error
}

type orderServiceImpl struct {
	cr coreapi.Client
	rp repositories.OrderRepository
}

func NewOrderService(cr coreapi.Client, rp repositories.OrderRepository) OrderService {
	return &orderServiceImpl{cr, rp}
}

func (svc *orderServiceImpl) RegisterOrder(order entities.OrderRequestParams, id uuid.UUID) (coreapi.ChargeResponse, error) {
	newOrder := entities.Order{
		ID:            uuid.New(),
		OrderStatuses: entities.OrderStatus{ID: uuid.New()},
		MerchantID:    order.OrderRequests.MerchantID,
		CustomerID:    id,
		Gross:         sumOrderGross(order),
		Paid:          false,
		OrderPlaced:   time.Now(),
		OrderItems:    make([]entities.OrderItem, len(order.OrderRequests.OrderItems)),
	}

	for i, item := range order.OrderRequests.OrderItems {
		newOrder.OrderItems[i] = entities.OrderItem{
			ID:         uuid.New(),
			MenuItemID: item.MenuItemID,
			Name:       item.Name,
			Quantity:   item.Quantity,
			Price:      item.Price,
			Comment:    item.Comment,
		}
	}

	itemDetails := make([]midtrans.ItemDetails, len(newOrder.OrderItems))

	for i, item := range newOrder.OrderItems {
		itemDetails[i] = midtrans.ItemDetails{
			ID:    uniuri.NewLen(50),
			Price: item.Price,
			Qty:   int32(item.Quantity),
			Name:  item.Name,
		}
	}

	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeGopay,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  uniuri.NewLen(50),
			GrossAmt: sumOrderGross(order),
		},
		Items: &itemDetails,
	}

	fmt.Println("Total midtrans :", chargeReq.TransactionDetails.GrossAmt)
	var summ int64
	for _, item := range itemDetails {
		summ += item.Price
	}
	fmt.Println("Total methods :", summ)

	coreApiRes, err := svc.cr.ChargeTransaction(chargeReq)
	if err != nil {
		return coreapi.ChargeResponse{}, err.RawError
	}

	if err := svc.rp.Create(&newOrder); err != nil {
		return coreapi.ChargeResponse{}, err
	}

	return *coreApiRes, err
}
func (svc *orderServiceImpl) FindOrder(param string, id uuid.UUID) (entities.Order, error) {
	order, err := svc.rp.GetByID(param, id)
	if err != nil {
		return entities.Order{}, err
	}
	return *order, nil
}
func (svc *orderServiceImpl) EditOrder(order entities.OrderRequestParams) error {
	newOrder := entities.Order{
		ID:            uuid.New(),
		OrderStatuses: entities.OrderStatus{ID: uuid.New()},
		MerchantID:    order.OrderRequests.MerchantID,
		CustomerID:    order.OrderRequests.CustomerID,
		Gross:         sumOrderGross(order),
		Paid:          false,
		OrderPlaced:   time.Now(),
		OrderItems:    make([]entities.OrderItem, len(order.OrderRequests.OrderItems)),
	}

	for i, item := range order.OrderRequests.OrderItems {
		newOrder.OrderItems[i] = entities.OrderItem{
			ID:         uuid.New(),
			MenuItemID: item.MenuItemID,
			Quantity:   item.Quantity,
			Price:      item.Price,
			Comment:    item.Comment,
		}
	}

	if err := svc.rp.Update(&newOrder); err != nil {
		return err
	}

	return nil
}
func (svc *orderServiceImpl) DeleteOrder(id uuid.UUID) error {
	order, err := svc.rp.GetByID("customer_id", id)
	if err != nil {
		return err
	}

	if err := svc.rp.Delete(order); err != nil {
		return err
	}

	return nil
}

func sumOrderGross(order entities.OrderRequestParams) int64 {
	var sum int64
	for _, item := range order.OrderRequests.OrderItems {
		sum += item.Price * int64(item.Quantity)
	}
	return sum
}
