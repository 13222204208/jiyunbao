package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppSignContractRouter struct {
}

// InitAppSignContractRouter 初始化 AppSignContract 路由信息
func (s *AppSignContractRouter) InitAppSignContractRouter(Router *gin.RouterGroup) {
	appSignContractRouter := Router.Group("appSignContract").Use(middleware.OperationRecord())
	appSignContractRouterWithoutRecord := Router.Group("appSignContract")
	var appSignContractApi = v1.ApiGroupApp.AppApiGroup.AppSignContractApi
	{
		appSignContractRouter.POST("createAppSignContract", appSignContractApi.CreateAppSignContract)             // 新建AppSignContract
		appSignContractRouter.DELETE("deleteAppSignContract", appSignContractApi.DeleteAppSignContract)           // 删除AppSignContract
		appSignContractRouter.DELETE("deleteAppSignContractByIds", appSignContractApi.DeleteAppSignContractByIds) // 批量删除AppSignContract
		appSignContractRouter.PUT("updateAppSignContract", appSignContractApi.UpdateAppSignContract)              // 更新AppSignContract

		appSignContractRouter.POST("queryAuthInfo", appSignContractApi.QueryAuthInfo)
	}
	{
		appSignContractRouterWithoutRecord.GET("findAppSignContract", appSignContractApi.FindAppSignContract)       // 根据ID获取AppSignContract
		appSignContractRouterWithoutRecord.GET("getAppSignContractList", appSignContractApi.GetAppSignContractList) // 获取AppSignContract列表
	}
}
