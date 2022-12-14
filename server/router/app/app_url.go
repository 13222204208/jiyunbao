package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppUrlRouter struct {
}

// InitAppUrlRouter 初始化 AppUrl 路由信息
func (s *AppUrlRouter) InitAppUrlRouter(Router *gin.RouterGroup) {
	appUrlRouter := Router.Group("appUrl").Use(middleware.OperationRecord())
	appUrlRouterWithoutRecord := Router.Group("appUrl")

	appUrlPublicRouter := Router.Group("url")
	var appUrlApi = v1.ApiGroupApp.AppApiGroup.AppUrlApi
	{
		appUrlRouter.POST("createAppUrl", appUrlApi.CreateAppUrl)             // 新建AppUrl
		appUrlRouter.DELETE("deleteAppUrl", appUrlApi.DeleteAppUrl)           // 删除AppUrl
		appUrlRouter.DELETE("deleteAppUrlByIds", appUrlApi.DeleteAppUrlByIds) // 批量删除AppUrl
		appUrlRouter.PUT("updateAppUrl", appUrlApi.UpdateAppUrl)              // 更新AppUrl
	}
	{
		appUrlRouterWithoutRecord.GET("findAppUrl", appUrlApi.FindAppUrl)       // 根据ID获取AppUrl
		appUrlRouterWithoutRecord.GET("getAppUrlList", appUrlApi.GetAppUrlList) // 获取AppUrl列表
	}

	{
		//回调方法
		appUrlPublicRouter.POST("notify", appUrlApi.Notify)
	}
}
