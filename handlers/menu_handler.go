package handlers

import (
	"labireen/entities"
	"labireen/pkg/jwtx"
	"labireen/pkg/response"
	"labireen/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MenuHandler interface {
	CreateMenu(ctx *gin.Context)
	ViewMenu(ctx *gin.Context)
	DeleteMenu(ctx *gin.Context)
}

type menuHandler struct {
	svc services.MenuService
}

func NewMenuHandler(svc services.MenuService) *menuHandler {
	return &menuHandler{svc}
}

func (hdl *menuHandler) CreateMenu(ctx *gin.Context) {
	var request entities.MenuRequestParams
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	if err := hdl.svc.CreateMenu(request); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to register menu", err.Error())
		return
	}

	response.Success(ctx, http.StatusOK, "Menu successfuly created", request)
}

func (hdl *menuHandler) ViewMenu(ctx *gin.Context) {
	param := ctx.Param("merchant-name")

	if param != "" {
		resp, err := hdl.svc.GetMenu(param)
		if err != nil {
			response.Error(ctx, http.StatusNotFound, "Cannot find requested data", err.Error())
			return
		}

		response.Success(ctx, http.StatusOK, "success", resp)
		return
	}

	resp, err := hdl.svc.GetAllMenu()
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Cannot find requested data", err.Error())
		return
	}

	response.Success(ctx, http.StatusOK, "success", resp)
}

func (hdl *menuHandler) DeleteMenu(ctx *gin.Context) {
	temp, exist := ctx.Get("currentUser")
	if !exist {
		log := response.ErrorLog{
			Field:   "token",
			Message: "Key token value does not exists",
		}
		response.Error(ctx, http.StatusNotFound, "Key error", log)
	}

	user := temp.(jwtx.UserClaims)

	if err := hdl.svc.DeleteMenu(user.ID); err != nil {
		response.Error(ctx, http.StatusNotFound, "Record not found", err.Error())
		return
	}

	response.Success(ctx, http.StatusOK, "Record successfuly deleted", nil)

}
