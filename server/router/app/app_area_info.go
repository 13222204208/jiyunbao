package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppAreaInfoRouter struct {
}

// InitAppAreaInfoRouter 初始化 AppAreaInfo 路由信息
func (s *AppAreaInfoRouter) InitAppAreaInfoRouter(Router *gin.RouterGroup) {
	appAreaInfoRouter := Router.Group("appAreaInfo").Use(middleware.OperationRecord())
	appAreaInfoRouterWithoutRecord := Router.Group("appAreaInfo")
	var appAreaInfoApi = v1.ApiGroupApp.AppApiGroup.AppAreaInfoApi
	{
		appAreaInfoRouter.POST("createAppAreaInfo", appAreaInfoApi.CreateAppAreaInfo)   // 新建AppAreaInfo
		appAreaInfoRouter.DELETE("deleteAppAreaInfo", appAreaInfoApi.DeleteAppAreaInfo) // 删除AppAreaInfo
		appAreaInfoRouter.DELETE("deleteAppAreaInfoByIds", appAreaInfoApi.DeleteAppAreaInfoByIds) // 批量删除AppAreaInfo
		appAreaInfoRouter.PUT("updateAppAreaInfo", appAreaInfoApi.UpdateAppAreaInfo)    // 更新AppAreaInfo
	}
	{
		appAreaInfoRouterWithoutRecord.GET("findAppAreaInfo", appAreaInfoApi.FindAppAreaInfo)        // 根据ID获取AppAreaInfo
		appAreaInfoRouterWithoutRecord.GET("getAppAreaInfoList", appAreaInfoApi.GetAppAreaInfoList)  // 获取AppAreaInfo列表
	}
}
