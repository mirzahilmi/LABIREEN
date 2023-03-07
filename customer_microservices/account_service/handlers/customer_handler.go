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
	user := ctx.MustGet("currentUser").(jwtx.CustomerClaims)

	userResp, err := cH.svc.GetCustomer(user.Email)
	if err != nil {
		response.FailOrError(ctx, http.StatusUnauthorized, "Invalid email or password", nil)
		return
	}

	response.Success(ctx, http.StatusOK, "success", userResp)
}
