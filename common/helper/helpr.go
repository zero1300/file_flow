package helper

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func GetUserIdByToken(ctx *gin.Context) (int, error) {
	get, exists := ctx.Get("id")
	id := get.(int)
	if exists {
		return id, nil
	}
	return -1, errors.New("")
}
