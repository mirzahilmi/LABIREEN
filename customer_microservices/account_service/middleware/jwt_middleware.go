package middleware

import (
	"errors"
	"labireen/customer_microservices/account_service/utilities/jwtx"
	"labireen/customer_microservices/account_service/utilities/response"
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
			msg := "token not found"
			response.FailOrError(ctx, http.StatusForbidden, msg, errors.New(msg))
			return
		}

		tokenJwt := authorization[7:]
		claims := jwtx.CustomerClaims{}
		jwtKey := os.Getenv("SECRET")

		if err := jwtx.DecodeToken(tokenJwt, &claims, jwtKey); err != nil {
			ctx.Abort()
			response.FailOrError(ctx, http.StatusUnauthorized, "unauthorized", err)
			return
		}

		ctx.Set("currentUser", claims)
		ctx.Next()
	}
}
