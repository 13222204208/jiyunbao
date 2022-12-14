package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppNoticeRouter struct {
}

// InitAppNoticeRouter 初始化 AppNotice 路由信息
func (s *AppNoticeRouter) InitAppNoticeRouter(Router *gin.RouterGroup) {
	appNoticeRouter := Router.Group("appNotice").Use(middleware.OperationRecord())
	appNoticeRouterWithoutRecord := Router.Group("appNotice")

	appNoticePrivateRouter := Router.Group("notice").Use(middleware.JWTAuthMiddleware())
	var appNoticeApi = v1.ApiGroupApp.AppApiGroup.AppNoticeApi
	{
		appNoticeRouter.POST("createAppNotice", appNoticeApi.CreateAppNotice)             // 新建AppNotice
		appNoticeRouter.DELETE("deleteAppNotice", appNoticeApi.DeleteAppNotice)           // 删除AppNotice
		appNoticeRouter.DELETE("deleteAppNoticeByIds", appNoticeApi.DeleteAppNoticeByIds) // 批量删除AppNotice
		appNoticeRouter.PUT("updateAppNotice", appNoticeApi.UpdateAppNotice)              // 更新AppNotice
	}
	{
		appNoticeRouterWithoutRecord.GET("findAppNotice", appNoticeApi.FindAppNotice)       // 根据ID获取AppNotice
		appNoticeRouterWithoutRecord.GET("getAppNoticeList", appNoticeApi.GetAppNoticeList) // 获取AppNotice列表
	}

	{
		//获取公告的内容
		appNoticePrivateRouter.GET("info", appNoticeApi.Info)
	}
}
