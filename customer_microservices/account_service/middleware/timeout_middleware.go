package middleware

import (
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func respondTimeout(c *gin.Context) {
	c.String(http.StatusRequestTimeout, "Request time out")
}

func TimeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(5*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(respondTimeout),
	)
}
