package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppFacilitatingAgencyRouter struct {
}

// InitAppFacilitatingAgencyRouter 初始化 AppFacilitatingAgency 路由信息
func (s *AppFacilitatingAgencyRouter) InitAppFacilitatingAgencyRouter(Router *gin.RouterGroup) {
	appFacilitatingAgencyRouter := Router.Group("appFacilitatingAgency").Use(middleware.OperationRecord())
	appFacilitatingAgencyRouterWithoutRecord := Router.Group("appFacilitatingAgency")
	var appFacilitatingAgencyApi = v1.ApiGroupApp.AppApiGroup.AppFacilitatingAgencyApi
	{
		appFacilitatingAgencyRouter.POST("createAppFacilitatingAgency", appFacilitatingAgencyApi.CreateAppFacilitatingAgency)   // 新建AppFacilitatingAgency
		appFacilitatingAgencyRouter.DELETE("deleteAppFacilitatingAgency", appFacilitatingAgencyApi.DeleteAppFacilitatingAgency) // 删除AppFacilitatingAgency
		appFacilitatingAgencyRouter.DELETE("deleteAppFacilitatingAgencyByIds", appFacilitatingAgencyApi.DeleteAppFacilitatingAgencyByIds) // 批量删除AppFacilitatingAgency
		appFacilitatingAgencyRouter.PUT("updateAppFacilitatingAgency", appFacilitatingAgencyApi.UpdateAppFacilitatingAgency)    // 更新AppFacilitatingAgency
	}
	{
		appFacilitatingAgencyRouterWithoutRecord.GET("findAppFacilitatingAgency", appFacilitatingAgencyApi.FindAppFacilitatingAgency)        // 根据ID获取AppFacilitatingAgency
		appFacilitatingAgencyRouterWithoutRecord.GET("getAppFacilitatingAgencyList", appFacilitatingAgencyApi.GetAppFacilitatingAgencyList)  // 获取AppFacilitatingAgency列表
	}
}
