package router

import (
	v1 "auth/auth/api/v1"
	"github.com/gin-gonic/gin"
)

func initAuthRouter(Router *gin.RouterGroup) {
	AuthRouter := Router.Group("auth")
	{
		AuthRouter.POST("updateAuth", v1.UpdateAuth)
		AuthRouter.POST("getPolicyPathByAuthorityId", v1.GetPolicyPathByAuthorityId)
	}
}
