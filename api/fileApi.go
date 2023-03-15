package api

import (
	"file_flow/common/resp"
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
	file, err := ctx.FormFile("file")
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}
	err = f.fileService.UploadFile(file)
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}
}
