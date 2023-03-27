package api

import (
	"file_flow/common/helper"
	"file_flow/common/resp"
	"file_flow/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ShareApi struct {
	shareService *service.ShareService
}

func NewShareApi() ShareApi {
	return ShareApi{
		shareService: service.NewShareService(),
	}
}

func (s ShareApi) CreateShare(ctx *gin.Context) {
	id, err := helper.GetUserIdByToken(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	var form struct {
		Id         int `form:"id"`
		Expiration int `form:"expiration"`
	}
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	shareId, err := s.shareService.CreateShare(form.Id, form.Expiration, id)
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"shareId": shareId,
	})
}

func (s ShareApi) GetShareInfo(ctx *gin.Context) {
	idString := ctx.Param("id")
	shareId, err := strconv.Atoi(idString)
	if err != nil {
		resp.Fail(ctx, "shareId 解析失败")
		return
	}
	info, err := s.shareService.GetShareInfo(shareId)
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}
	resp.Success(ctx, info)
}
