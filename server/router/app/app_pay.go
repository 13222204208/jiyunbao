package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppPayRouter struct {
}

// InitAppPayRouter 初始化 AppPay 路由信息
func (s *AppPayRouter) InitAppPayRouter(Router *gin.RouterGroup) {
	appPayRouter := Router.Group("appPay").Use(middleware.OperationRecord())
	appPayRouterWithoutRecord := Router.Group("appPay")

	appPayPrivateRouter := Router.Group("pay").Use(middleware.JWTAuthMiddleware())
	var appPayApi = v1.ApiGroupApp.AppApiGroup.AppPayApi
	{
		appPayRouter.POST("createAppPay", appPayApi.CreateAppPay)             // 新建AppPay
		appPayRouter.DELETE("deleteAppPay", appPayApi.DeleteAppPay)           // 删除AppPay
		appPayRouter.DELETE("deleteAppPayByIds", appPayApi.DeleteAppPayByIds) // 批量删除AppPay
		appPayRouter.PUT("updateAppPay", appPayApi.UpdateAppPay)              // 更新AppPay
	}
	{
		appPayRouterWithoutRecord.GET("findAppPay", appPayApi.FindAppPay)       // 根据ID获取AppPay
		appPayRouterWithoutRecord.GET("getAppPayList", appPayApi.GetAppPayList) // 获取AppPay列表
	}

	{
		//支付认证
		appPayPrivateRouter.POST("attestation", appPayApi.Attestation)

		//编辑支付认证
		appPayPrivateRouter.PUT("edit", appPayApi.Edit)
	}
}
