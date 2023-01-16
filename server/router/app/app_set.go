package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppSetRouter struct {
}

// InitAppSetRouter 初始化 AppSet 路由信息
func (s *AppSetRouter) InitAppSetRouter(Router *gin.RouterGroup) {
	appSetRouter := Router.Group("appSet").Use(middleware.OperationRecord())
	appSetRouterWithoutRecord := Router.Group("appSet")

	appSetPrivateRouter := Router.Group("set").Use(middleware.JWTAuthMiddleware())

	var appSetApi = v1.ApiGroupApp.AppApiGroup.AppSetApi
	{
		appSetRouter.POST("createAppSet", appSetApi.CreateAppSet)             // 新建AppSet
		appSetRouter.DELETE("deleteAppSet", appSetApi.DeleteAppSet)           // 删除AppSet
		appSetRouter.DELETE("deleteAppSetByIds", appSetApi.DeleteAppSetByIds) // 批量删除AppSet
		appSetRouter.PUT("updateAppSet", appSetApi.UpdateAppSet)              // 更新AppSet
	}
	{
		appSetRouterWithoutRecord.GET("findAppSet", appSetApi.FindAppSet)       // 根据ID获取AppSet
		appSetRouterWithoutRecord.GET("getAppSetList", appSetApi.GetAppSetList) // 获取AppSet列表
	}

	{
		//获取我的 二维码图片
		appSetPrivateRouter.GET("qrcode", appSetApi.Qrcode)
	}
}
