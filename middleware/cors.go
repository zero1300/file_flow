package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	config := cors.Config{
		AllowOriginFunc: func(origin string) bool {
			// 运行所有站点
			return true
		},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "OPTION"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "Origin", "Accept", "Content-Length"},
	}

	return cors.New(config)
}
