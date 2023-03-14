package handlers

import (
	"labireen/entities"
	"labireen/pkg/jwtx"
	"labireen/pkg/response"
	"labireen/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	temp, exist := ctx.Get("currentUser")
	if !exist {
		log := response.ErrorLog{
			Field:   "token",
			Message: "Key token value does not exists",
		}
		response.Error(ctx, http.StatusNotFound, "Key error", log)
	}

	user := temp.(jwtx.MenuClaims)

	if err := hdl.svc.CreateMenu(request, user.ID); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to register menu", err.Error())
		return
	}

	response.Success(ctx, http.StatusOK, "Menu successfuly created", request)
}

func (hdl *menuHandler) ViewMenu(ctx *gin.Context) {
	temp, exist := ctx.Get("currentUser")
	if !exist {
		log := response.ErrorLog{
			Field:   "token",
			Message: "Key token value does not exists",
		}
		response.Error(ctx, http.StatusNotFound, "Key error", log)
	}

	user := temp.(jwtx.MenuClaims)

	menuRsp, err := hdl.svc.GetMenu(user.ID)
	if err != nil {
		log := response.ErrorLog{
			Field:   "ID",
			Message: "ID sent is not valid",
		}
		response.Error(ctx, http.StatusNotFound, "Cannot find requested data", log)
		return
	}

	response.Success(ctx, http.StatusOK, "success", menuRsp)
}
