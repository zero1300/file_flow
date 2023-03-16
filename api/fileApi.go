package api

import (
	"file_flow/common/helper"
	"file_flow/common/resp"
	"file_flow/models"
	"file_flow/service"
	"github.com/gin-gonic/gin"
)

type FileApi struct {
	fileService *service.FileService
}

func NewFileApi() FileApi {
	fileService := service.NewFileService()
	return FileApi{
		fileService: fileService,
	}
}

func (f FileApi) FileUpload(ctx *gin.Context) {
	uid, err := helper.GetUserIdByToken(ctx)
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}
	file, err := ctx.FormFile("file")
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}
	err = f.fileService.UploadFile(file, uid)
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}
}

func (f FileApi) UserFiles(ctx *gin.Context) {

	var paginateExpand struct {
		models.Paginate
		parentId int `form:"parentId"`
	}

	id, err2 := helper.GetUserIdByToken(ctx)
	if err2 != nil {
		resp.Fail(ctx, err2.Error())
		return
	}
	err := ctx.ShouldBind(&paginateExpand)
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}
	files, count, err := f.fileService.GetUserFile(id, paginateExpand.parentId, paginateExpand.Paginate)
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}
	resp.Success(ctx, gin.H{
		"files": files,
		"total": count,
	})

}
