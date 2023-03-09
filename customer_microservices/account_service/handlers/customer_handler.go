package handlers

import (
	"labireen/customer_microservices/account_service/services"
	"labireen/customer_microservices/account_service/utilities/jwtx"
	"labireen/customer_microservices/account_service/utilities/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	svc services.CustomerService
}

func NewCustomerHandler(svc services.CustomerService) *customerHandler {
	return &customerHandler{svc}
}

func (cH *customerHandler) GetMe(ctx *gin.Context) {
	temp, exist := ctx.Get("currentUser")
	if !exist {
		log := response.ErrorLog{
			Field:   "token",
			Message: "Key token value does not exists",
		}
		response.Error(ctx, http.StatusNotFound, "Key error", log)
	}

	user := temp.(jwtx.CustomerClaims)

	userResp, err := cH.svc.GetCustomer(user.ID)
	if err != nil {
		log := response.ErrorLog{
			Field:   "ID",
			Message: "ID sent is not valid",
		}
		response.Error(ctx, http.StatusNotFound, "Cannot find requested data", log)
		return
	}

	response.Success(ctx, http.StatusOK, "success", userResp)
}
