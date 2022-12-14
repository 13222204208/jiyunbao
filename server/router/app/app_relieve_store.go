package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppRelieveStoreRouter struct {
}

// InitAppRelieveStoreRouter 初始化 AppRelieveStore 路由信息
func (s *AppRelieveStoreRouter) InitAppRelieveStoreRouter(Router *gin.RouterGroup) {
	appRelieveStoreRouter := Router.Group("appRelieveStore").Use(middleware.OperationRecord())
	appRelieveStoreRouterWithoutRecord := Router.Group("appRelieveStore")

	appRelieveStorePrivateRouter := Router.Group("relieveStore").Use(middleware.JWTAuthMiddleware())

	var appRelieveStoreApi = v1.ApiGroupApp.AppApiGroup.AppRelieveStoreApi
	{
		appRelieveStoreRouter.POST("createAppRelieveStore", appRelieveStoreApi.CreateAppRelieveStore)             // 新建AppRelieveStore
		appRelieveStoreRouter.DELETE("deleteAppRelieveStore", appRelieveStoreApi.DeleteAppRelieveStore)           // 删除AppRelieveStore
		appRelieveStoreRouter.DELETE("deleteAppRelieveStoreByIds", appRelieveStoreApi.DeleteAppRelieveStoreByIds) // 批量删除AppRelieveStore
		appRelieveStoreRouter.PUT("updateAppRelieveStore", appRelieveStoreApi.UpdateAppRelieveStore)              // 更新AppRelieveStore
	}
	{
		appRelieveStoreRouterWithoutRecord.GET("findAppRelieveStore", appRelieveStoreApi.FindAppRelieveStore)       // 根据ID获取AppRelieveStore
		appRelieveStoreRouterWithoutRecord.GET("getAppRelieveStoreList", appRelieveStoreApi.GetAppRelieveStoreList) // 获取AppRelieveStore列表
	}

	{
		//申请解除合作
		appRelieveStorePrivateRouter.POST("apply", appRelieveStoreApi.Apply)
	}
}
