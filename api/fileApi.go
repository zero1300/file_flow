package api

import (
	"file_flow/common/helper"
	"file_flow/common/resp"
	"file_flow/models"
	"file_flow/service"
	"github.com/gin-gonic/gin"
	"strconv"
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
	var form struct {
		ParentId int `form:"parentId" binding:"required"`
	}
	err = ctx.ShouldBind(&form)
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}
	err = f.fileService.UploadFile(file, uid, form.ParentId)
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}
}

func (f FileApi) UserFiles(ctx *gin.Context) {

	var paginateExpand struct {
		models.Paginate
		ParentId int `form:"parentId"`
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
	files, count, err := f.fileService.GetUserFile(id, paginateExpand.ParentId, paginateExpand.Paginate)
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}
	resp.Success(ctx, gin.H{
		"files": files,
		"total": count,
	})

}

func (f FileApi) NewFolder(ctx *gin.Context) {
	id, err := helper.GetUserIdByToken(ctx)
	if err != nil {
		resp.Fail(ctx, "token 无效")
		return
	}
	var form struct {
		Name     string `form:"name"`
		ParentId int    `form:"parentId"`
	}
	err = ctx.ShouldBind(&form)
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}
	err = f.fileService.NewFolder(form.Name, form.ParentId, id)
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}
	resp.Success(ctx, nil)
}

func (f FileApi) DelFile(ctx *gin.Context) {
	uid, err := helper.GetUserIdByToken(ctx)
	if err != nil {
		resp.Fail(ctx, "token 无效")
		return
	}
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	err = f.fileService.DelUserFile(uid, id)
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}
	resp.Success(ctx, nil)

}
