package response

import "github.com/gin-gonic/gin"

type response struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorLog struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func Success(ctx *gin.Context, httpCode int, msg string, data interface{}) {
	ctx.JSON(httpCode, response{
		Status:  "success",
		Code:    httpCode,
		Message: msg,
		Data:    data,
	})
}

func Error(ctx *gin.Context, httpCode int, msg string, data interface{}) {
	ctx.JSON(httpCode, response{
		Status:  "error",
		Code:    httpCode,
		Message: msg,
		Data:    data,
	})
}
