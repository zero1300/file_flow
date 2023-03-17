package router

import (
	"file_flow/api"
	"github.com/gin-gonic/gin"
)

func fileSetup() {
	RegRoute(func(public *gin.RouterGroup, auth *gin.RouterGroup) {
		fileApi := api.NewFileApi()

		file := auth.Group("file")
		file.POST("/upload", fileApi.FileUpload)
		file.POST("/getUserFiles", fileApi.UserFiles)
		file.POST("/newFolder", fileApi.NewFolder)
	})
}
