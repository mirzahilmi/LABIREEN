package middleware

import (
	"labireen/order_service/pkg/jwtx"
	"labireen/order_service/pkg/response"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.Request.Header.Get("Authorization")

		if !strings.HasPrefix(authorization, "Bearer ") {
			ctx.Abort()

			log := response.ErrorLog{
				Field:   "token",
				Message: "Cannot find a valid token on bearer",
			}

			response.Error(ctx, http.StatusNotFound, "token not found", log)
			return
		}

		tokenJwt := authorization[7:]
		claims := jwtx.MenuClaims{}
		jwtKey := os.Getenv("SECRET")

		if err := jwtx.DecodeToken(tokenJwt, &claims, jwtKey); err != nil {
			ctx.Abort()

			log := response.ErrorLog{
				Field:   "token",
				Message: err.Error(),
			}

			response.Error(ctx, http.StatusUnauthorized, "token error", log)
			return
		}

		ctx.Set("currentUser", claims)
		ctx.Next()
	}
}
