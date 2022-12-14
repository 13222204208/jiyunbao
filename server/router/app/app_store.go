package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppStoreRouter struct {
}

// InitAppStoreRouter 初始化 AppStore 路由信息
func (s *AppStoreRouter) InitAppStoreRouter(Router *gin.RouterGroup) {
	appStoreRouter := Router.Group("appStore").Use(middleware.OperationRecord())
	appStoreRouterWithoutRecord := Router.Group("appStore")

	appStorePublicRouter := Router.Group("store").Use(middleware.JWTAuthMiddleware())
	var appStoreApi = v1.ApiGroupApp.AppApiGroup.AppStoreApi
	{
		appStoreRouter.POST("createAppStore", appStoreApi.CreateAppStore)             // 新建AppStore
		appStoreRouter.DELETE("deleteAppStore", appStoreApi.DeleteAppStore)           // 删除AppStore
		appStoreRouter.DELETE("deleteAppStoreByIds", appStoreApi.DeleteAppStoreByIds) // 批量删除AppStore
		appStoreRouter.PUT("updateAppStore", appStoreApi.UpdateAppStore)              // 更新AppStore
	}
	{
		appStoreRouterWithoutRecord.GET("findAppStore", appStoreApi.FindAppStore)       // 根据ID获取AppStore
		appStoreRouterWithoutRecord.GET("getAppStoreList", appStoreApi.GetAppStoreList) // 获取AppStore列表
	}
	//门店设置
	{
		appStorePublicRouter.PUT("edit", appStoreApi.Edit)
	}
}
