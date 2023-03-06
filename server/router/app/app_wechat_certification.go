package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppWechatCertificationRouter struct {
}

// InitAppWechatCertificationRouter 初始化 AppWechatCertification 路由信息
func (s *AppWechatCertificationRouter) InitAppWechatCertificationRouter(Router *gin.RouterGroup) {
	appWechatCertificationRouter := Router.Group("appWechatCertification").Use(middleware.OperationRecord())
	appWechatCertificationRouterWithoutRecord := Router.Group("appWechatCertification")

	appWechatRouter := Router.Group("wechat").Use(middleware.JWTAuthMiddleware())

	var appWechatCertificationApi = v1.ApiGroupApp.AppApiGroup.AppWechatCertificationApi
	{
		appWechatCertificationRouter.POST("createAppWechatCertification", appWechatCertificationApi.CreateAppWechatCertification)             // 新建AppWechatCertification
		appWechatCertificationRouter.DELETE("deleteAppWechatCertification", appWechatCertificationApi.DeleteAppWechatCertification)           // 删除AppWechatCertification
		appWechatCertificationRouter.DELETE("deleteAppWechatCertificationByIds", appWechatCertificationApi.DeleteAppWechatCertificationByIds) // 批量删除AppWechatCertification
		appWechatCertificationRouter.PUT("updateAppWechatCertification", appWechatCertificationApi.UpdateAppWechatCertification)              // 更新AppWechatCertification

		//微信实名认证上送
		appWechatCertificationRouter.POST("wechatCertification", appWechatCertificationApi.WechatCertification)

		//微信实名认证状态查询
		appWechatCertificationRouter.GET("certificationState", appWechatCertificationApi.CertificationState)

		//微信实名认证明细查询
		appWechatCertificationRouter.GET("getAuthMessagesByMercId", appWechatCertificationApi.GetAuthMessagesByMercId)

		//微信实名认证授权信息
		appWechatCertificationRouter.GET("getAuthState", appWechatCertificationApi.GetAuthState)
	}
	{
		appWechatCertificationRouterWithoutRecord.GET("findAppWechatCertification", appWechatCertificationApi.FindAppWechatCertification)       // 根据ID获取AppWechatCertification
		appWechatCertificationRouterWithoutRecord.GET("getAppWechatCertificationList", appWechatCertificationApi.GetAppWechatCertificationList) // 获取AppWechatCertification列表
	}

	{
		appWechatRouter.GET("authentication", appWechatCertificationApi.Authentication)
	}
}
