package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppRateRouter struct {
}

// InitAppRateRouter 初始化 AppRate 路由信息
func (s *AppRateRouter) InitAppRateRouter(Router *gin.RouterGroup) {
	appRateRouter := Router.Group("appRate").Use(middleware.OperationRecord())
	appRateRouterWithoutRecord := Router.Group("appRate")
	var appRateApi = v1.ApiGroupApp.AppApiGroup.AppRateApi
	{
		appRateRouter.POST("createAppRate", appRateApi.CreateAppRate)             // 新建AppRate
		appRateRouter.DELETE("deleteAppRate", appRateApi.DeleteAppRate)           // 删除AppRate
		appRateRouter.DELETE("deleteAppRateByIds", appRateApi.DeleteAppRateByIds) // 批量删除AppRate
		appRateRouter.PUT("updateAppRate", appRateApi.UpdateAppRate)              // 更新AppRate
	}
	{
		appRateRouterWithoutRecord.GET("findAppRate", appRateApi.FindAppRate)       // 根据ID获取AppRate
		appRateRouterWithoutRecord.GET("getAppRateList", appRateApi.GetAppRateList) // 获取AppRate列表

		appRateRouterWithoutRecord.GET("list", appRateApi.List) // 获取AppRate列表
	}
}
