package router

import (
	"file_flow/api"
	"github.com/gin-gonic/gin"
)

func shareSetup() {
	RegRoute(func(public *gin.RouterGroup, auth *gin.RouterGroup) {
		shareApi := api.NewShareApi()

		public.GET("/info/:id", shareApi.GetShareInfo)

		share := auth.Group("share")
		share.POST("/create", shareApi.CreateShare)
	})
}
