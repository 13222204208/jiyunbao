package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppServiceRouter struct {
}

// InitAppServiceRouter 初始化 AppService 路由信息
func (s *AppServiceRouter) InitAppServiceRouter(Router *gin.RouterGroup) {
	appServiceRouter := Router.Group("appService").Use(middleware.OperationRecord())
	appServiceRouterWithoutRecord := Router.Group("appService")

	appServicePrivateRouter := Router.Group("service").Use(middleware.JWTAuthMiddleware())

	var appServiceApi = v1.ApiGroupApp.AppApiGroup.AppServiceApi
	{
		appServiceRouter.POST("createAppService", appServiceApi.CreateAppService)             // 新建AppService
		appServiceRouter.DELETE("deleteAppService", appServiceApi.DeleteAppService)           // 删除AppService
		appServiceRouter.DELETE("deleteAppServiceByIds", appServiceApi.DeleteAppServiceByIds) // 批量删除AppService
		appServiceRouter.PUT("updateAppService", appServiceApi.UpdateAppService)              // 更新AppService
	}
	{
		appServiceRouterWithoutRecord.GET("findAppService", appServiceApi.FindAppService)       // 根据ID获取AppService
		appServiceRouterWithoutRecord.GET("getAppServiceList", appServiceApi.GetAppServiceList) // 获取AppService列表
	}

	{
		//获取服务设置信息列表
		appServicePrivateRouter.GET("list", appServiceApi.List)
	}
}
