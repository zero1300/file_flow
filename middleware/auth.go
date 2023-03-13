package middleware

import (
	"context"
	"file_flow/common/jwt"
	"file_flow/global"
	"file_flow/global/constants"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

const TokenPrefix = "Bearer "
const TokenName = "Authorization"

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader(TokenName)
		if token == "" || !strings.HasPrefix(token, TokenPrefix) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		token = token[len(TokenPrefix):]
		user, err := jwt.ParseToken(token)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		id := user.ID
		err = global.Redis.Get(context.Background(), constants.TokenKey+strconv.Itoa(id)).Err()
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Set("id", id)
		ctx.Next()
	}
}
