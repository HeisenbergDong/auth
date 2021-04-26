package router

import (
	"auth/middleware"
	"github.com/gin-gonic/gin"
)

// 初始化总路由
func InitRouters() *gin.Engine {
	Router := gin.Default()
	// 跨域
	Router.Use(middleware.Cors())
	AuthGroup := Router.Group("")
	initAuthRouter(AuthGroup)
	return Router
}
