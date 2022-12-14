package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppMchRouter struct {
}

// InitAppMchRouter 初始化 AppMch 路由信息
func (s *AppMchRouter) InitAppMchRouter(Router *gin.RouterGroup) {
	appMchRouter := Router.Group("appMch").Use(middleware.OperationRecord())
	appMchRouterWithoutRecord := Router.Group("appMch")

	appMchPrivateRouter := Router.Group("mch").Use(middleware.JWTAuthMiddleware())
	var appMchApi = v1.ApiGroupApp.AppApiGroup.AppMchApi
	{
		appMchRouter.POST("createAppMch", appMchApi.CreateAppMch)             // 新建AppMch
		appMchRouter.DELETE("deleteAppMch", appMchApi.DeleteAppMch)           // 删除AppMch
		appMchRouter.DELETE("deleteAppMchByIds", appMchApi.DeleteAppMchByIds) // 批量删除AppMch
		appMchRouter.PUT("updateAppMch", appMchApi.UpdateAppMch)              // 更新AppMch
	}
	{
		appMchRouterWithoutRecord.GET("findAppMch", appMchApi.FindAppMch)       // 根据ID获取AppMch
		appMchRouterWithoutRecord.GET("getAppMchList", appMchApi.GetAppMchList) // 获取AppMch列表
	}

	{
		//商户认证
		appMchPrivateRouter.POST("attestation", appMchApi.Attestation)

		//商户认证编辑
		appMchPrivateRouter.PUT("edit", appMchApi.Edit)
	}
}
