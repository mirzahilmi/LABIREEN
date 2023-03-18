package handlers

import (
	"labireen/entities"
	"labireen/pkg/jwtx"
	"labireen/pkg/response"
	"labireen/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler interface {
	CreateOrder(ctx *gin.Context)
}

type orderHandler struct {
	svc services.OrderService
}

func NewOrderHandler(svc services.OrderService) *orderHandler {
	return &orderHandler{svc}
}

func (hdl *orderHandler) CreateOrder(ctx *gin.Context) {
	var request entities.OrderRequestParams
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	temp, exist := ctx.Get("currentUser")
	if !exist {
		log := response.ErrorLog{
			Field:   "token",
			Message: "Key token value does not exists",
		}
		response.Error(ctx, http.StatusNotFound, "Key error", log)
	}

	user := temp.(jwtx.UserClaims)

	resp, err := hdl.svc.RegisterOrder(request, user.ID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to register order", err.Error())
		return
	}

	response.Success(ctx, http.StatusOK, "Order successfuly created", resp.StatusMessage)
}

func (hdl *orderHandler) UpdateOrder(ctx *gin.Context) {
	var order entities.OrderRequestParams
	if err := ctx.ShouldBindJSON(&order); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	if err := hdl.svc.EditOrder(order); err != nil {
		response.Error(ctx, http.StatusNotAcceptable, "Request Error", err.Error())
	}
}
