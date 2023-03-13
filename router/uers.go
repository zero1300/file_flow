package router

import (
	"file_flow/api"
	"github.com/gin-gonic/gin"
)

func userSetup() {
	RegRoute(func(public *gin.RouterGroup, auth *gin.RouterGroup) {

		userApi := api.NewUserApi()

		public.POST("/login", userApi.Login)
		public.POST("/reg", userApi.Register)

		user := auth.Group("user")
		user.GET("/:id", userApi.GetUserById)
		user.POST("/list", userApi.UserList)
		user.POST("/update", userApi.UpdateUser)
		user.DELETE("/:id", userApi.DelUser)
		//user.GET("", func(c *gin.Context) {
		//
		//})
	})
}
