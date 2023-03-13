package models

import "github.com/gin-gonic/gin"

type Paginate struct {
	Page     int `form:"page" binding:"required,gt=0"`
	PageSize int `form:"pageSize" binding:"required,gt=0"`
}

func GetPaginate(ctx *gin.Context) (Paginate, error) {
	var p Paginate
	err := ctx.ShouldBind(&p)
	if err != nil {
		return Paginate{}, err
	}
	return p, nil
}
