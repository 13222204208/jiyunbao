package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppAlipayCertificationRouter struct {
}

// InitAppAlipayCertificationRouter 初始化 AppAlipayCertification 路由信息
func (s *AppAlipayCertificationRouter) InitAppAlipayCertificationRouter(Router *gin.RouterGroup) {
	appAlipayCertificationRouter := Router.Group("appAlipayCertification").Use(middleware.OperationRecord())
	appAlipayCertificationRouterWithoutRecord := Router.Group("appAlipayCertification")
	var appAlipayCertificationApi = v1.ApiGroupApp.AppApiGroup.AppAlipayCertificationApi
	{
		appAlipayCertificationRouter.POST("createAppAlipayCertification", appAlipayCertificationApi.CreateAppAlipayCertification)             // 新建AppAlipayCertification
		appAlipayCertificationRouter.DELETE("deleteAppAlipayCertification", appAlipayCertificationApi.DeleteAppAlipayCertification)           // 删除AppAlipayCertification
		appAlipayCertificationRouter.DELETE("deleteAppAlipayCertificationByIds", appAlipayCertificationApi.DeleteAppAlipayCertificationByIds) // 批量删除AppAlipayCertification
		appAlipayCertificationRouter.PUT("updateAppAlipayCertification", appAlipayCertificationApi.UpdateAppAlipayCertification)              // 更新AppAlipayCertification

		//实名认证
		appAlipayCertificationRouter.POST("alipayCertification", appAlipayCertificationApi.AlipayCertification)

		//实证认证状态查询
		appAlipayCertificationRouter.GET("alipayCertificationState", appAlipayCertificationApi.AlipayCertificationState)

		//实名认证授权查询
		appAlipayCertificationRouter.GET("getAlipayAuthState", appAlipayCertificationApi.GetAlipayAuthState)
	}
	{
		appAlipayCertificationRouterWithoutRecord.GET("findAppAlipayCertification", appAlipayCertificationApi.FindAppAlipayCertification)       // 根据ID获取AppAlipayCertification
		appAlipayCertificationRouterWithoutRecord.GET("getAppAlipayCertificationList", appAlipayCertificationApi.GetAppAlipayCertificationList) // 获取AppAlipayCertification列表
	}
}
