package response

import "github.com/gin-gonic/gin"

type response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Body    any    `json:"body"`
}

func Success(c *gin.Context, httpCode int, msg string, body interface{}) {
	switch httpCode / 100 {
	case 2:
		c.JSON(httpCode, response{
			Status:  "success",
			Message: msg,
			Body:    body,
		})

	default:
		c.JSON(500, response{
			Status:  "error",
			Message: "RESPONSE ERROR",
			Body:    nil,
		})
	}
}

func FailOrError(c *gin.Context, httpCode int, msg string, err error) {
	switch httpCode / 100 {
	case 4:
		c.JSON(httpCode, response{
			Status:  "fail",
			Message: msg,
			Body:    err,
		})

	case 5:
		c.JSON(httpCode, response{
			Status:  "error",
			Message: msg,
			Body:    nil,
		})

	default:
		c.JSON(500, response{
			Status:  "error",
			Message: "RESPONSE ERROR",
			Body:    err,
		})
	}
}
