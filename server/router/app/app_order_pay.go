package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppOrderPayRouter struct {
}

// InitAppOrderPayRouter 初始化 AppOrderPay 路由信息
func (s *AppOrderPayRouter) InitAppOrderPayRouter(Router *gin.RouterGroup) {
	appOrderPayRouter := Router.Group("appOrderPay").Use(middleware.OperationRecord())

	appOrderPayPublicRouter := Router.Group("order")
	appOrderPayRouterWithoutRecord := Router.Group("appOrderPay")
	var appOrderPayApi = v1.ApiGroupApp.AppApiGroup.AppOrderPayApi
	{
		appOrderPayRouter.POST("createAppOrderPay", appOrderPayApi.CreateAppOrderPay)             // 新建AppOrderPay
		appOrderPayRouter.DELETE("deleteAppOrderPay", appOrderPayApi.DeleteAppOrderPay)           // 删除AppOrderPay
		appOrderPayRouter.DELETE("deleteAppOrderPayByIds", appOrderPayApi.DeleteAppOrderPayByIds) // 批量删除AppOrderPay
		appOrderPayRouter.PUT("updateAppOrderPay", appOrderPayApi.UpdateAppOrderPay)              // 更新AppOrderPay
	}
	{
		appOrderPayRouterWithoutRecord.GET("findAppOrderPay", appOrderPayApi.FindAppOrderPay)       // 根据ID获取AppOrderPay
		appOrderPayRouterWithoutRecord.GET("getAppOrderPayList", appOrderPayApi.GetAppOrderPayList) // 获取AppOrderPay列表
	}

	{
		appOrderPayPublicRouter.POST("pay", appOrderPayApi.Pay)
		appOrderPayPublicRouter.POST("notify", appOrderPayApi.Notify)
	}
}
