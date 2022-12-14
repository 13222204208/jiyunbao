package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppFeedbackRouter struct {
}

// InitAppFeedbackRouter 初始化 AppFeedback 路由信息
func (s *AppFeedbackRouter) InitAppFeedbackRouter(Router *gin.RouterGroup) {
	appFeedbackRouter := Router.Group("appFeedback").Use(middleware.OperationRecord())
	appFeedbackRouterWithoutRecord := Router.Group("appFeedback")

	appFeedbackPrivateRouter := Router.Group("feedback").Use(middleware.JWTAuthMiddleware())

	var appFeedbackApi = v1.ApiGroupApp.AppApiGroup.AppFeedbackApi
	{
		appFeedbackRouter.POST("createAppFeedback", appFeedbackApi.CreateAppFeedback)             // 新建AppFeedback
		appFeedbackRouter.DELETE("deleteAppFeedback", appFeedbackApi.DeleteAppFeedback)           // 删除AppFeedback
		appFeedbackRouter.DELETE("deleteAppFeedbackByIds", appFeedbackApi.DeleteAppFeedbackByIds) // 批量删除AppFeedback
		appFeedbackRouter.PUT("updateAppFeedback", appFeedbackApi.UpdateAppFeedback)              // 更新AppFeedback
	}
	{
		appFeedbackRouterWithoutRecord.GET("findAppFeedback", appFeedbackApi.FindAppFeedback)       // 根据ID获取AppFeedback
		appFeedbackRouterWithoutRecord.GET("getAppFeedbackList", appFeedbackApi.GetAppFeedbackList) // 获取AppFeedback列表
	}

	{
		//提交反馈内容
		appFeedbackPrivateRouter.POST("create", appFeedbackApi.Create)
	}
}
