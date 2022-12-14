package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppAgreementRouter struct {
}

// InitAppAgreementRouter 初始化 AppAgreement 路由信息
func (s *AppAgreementRouter) InitAppAgreementRouter(Router *gin.RouterGroup) {
	appAgreementRouter := Router.Group("appAgreement").Use(middleware.OperationRecord())
	appAgreementRouterWithoutRecord := Router.Group("appAgreement")

	appAgreementPublicRouter := Router.Group("agreement")

	var appAgreementApi = v1.ApiGroupApp.AppApiGroup.AppAgreementApi
	{
		appAgreementRouter.POST("createAppAgreement", appAgreementApi.CreateAppAgreement)             // 新建AppAgreement
		appAgreementRouter.DELETE("deleteAppAgreement", appAgreementApi.DeleteAppAgreement)           // 删除AppAgreement
		appAgreementRouter.DELETE("deleteAppAgreementByIds", appAgreementApi.DeleteAppAgreementByIds) // 批量删除AppAgreement
		appAgreementRouter.PUT("updateAppAgreement", appAgreementApi.UpdateAppAgreement)              // 更新AppAgreement
	}
	{
		appAgreementRouterWithoutRecord.GET("findAppAgreement", appAgreementApi.FindAppAgreement)       // 根据ID获取AppAgreement
		appAgreementRouterWithoutRecord.GET("getAppAgreementList", appAgreementApi.GetAppAgreementList) // 获取AppAgreement列表
	}

	{
		appAgreementPublicRouter.GET("info", appAgreementApi.Info)
	}
}
