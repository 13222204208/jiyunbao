package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppCustInfoRouter struct {
}

// InitAppCustInfoRouter 初始化 AppCustInfo 路由信息
func (s *AppCustInfoRouter) InitAppCustInfoRouter(Router *gin.RouterGroup) {
	appCustInfoRouter := Router.Group("appCustInfo").Use(middleware.OperationRecord())
	appCustInfoRouterWithoutRecord := Router.Group("appCustInfo")
	var appCustInfoApi = v1.ApiGroupApp.AppApiGroup.AppCustInfoApi
	{
		appCustInfoRouter.POST("createAppCustInfo", appCustInfoApi.CreateAppCustInfo)             // 新建AppCustInfo
		appCustInfoRouter.DELETE("deleteAppCustInfo", appCustInfoApi.DeleteAppCustInfo)           // 删除AppCustInfo
		appCustInfoRouter.DELETE("deleteAppCustInfoByIds", appCustInfoApi.DeleteAppCustInfoByIds) // 批量删除AppCustInfo
		appCustInfoRouter.PUT("updateAppCustInfo", appCustInfoApi.UpdateAppCustInfo)              // 更新AppCustInfo

		//后台资料上送
		appCustInfoRouter.POST("addCustInfoApply", appCustInfoApi.AddCustInfoApply)

		//资料确认
		appCustInfoRouter.POST("auditCustInfoApply", appCustInfoApi.AuditCustInfoApply)

		//商户状态查询
		appCustInfoRouter.POST("queryCustApply", appCustInfoApi.QueryCustApply)

		//基本信息变更申请
		appCustInfoRouter.POST("changeMercBaseInfo", appCustInfoApi.ChangeMercBaseInfo)
	}
	{
		appCustInfoRouterWithoutRecord.GET("findAppCustInfo", appCustInfoApi.FindAppCustInfo)       // 根据ID获取AppCustInfo
		appCustInfoRouterWithoutRecord.GET("getAppCustInfoList", appCustInfoApi.GetAppCustInfoList) // 获取AppCustInfo列表

	}
}
