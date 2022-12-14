package mini

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MiniSetRouter struct {
}

// InitMiniSetRouter 初始化 MiniSet 路由信息
func (s *MiniSetRouter) InitMiniSetRouter(Router *gin.RouterGroup) {
	miniSetRouter := Router.Group("miniSet").Use(middleware.OperationRecord())
	miniSetRouterWithoutRecord := Router.Group("miniSet")
	var miniSetApi = v1.ApiGroupApp.MiniApiGroup.MiniSetApi
	{
		miniSetRouter.POST("createMiniSet", miniSetApi.CreateMiniSet)   // 新建MiniSet
		miniSetRouter.DELETE("deleteMiniSet", miniSetApi.DeleteMiniSet) // 删除MiniSet
		miniSetRouter.DELETE("deleteMiniSetByIds", miniSetApi.DeleteMiniSetByIds) // 批量删除MiniSet
		miniSetRouter.PUT("updateMiniSet", miniSetApi.UpdateMiniSet)    // 更新MiniSet
	}
	{
		miniSetRouterWithoutRecord.GET("findMiniSet", miniSetApi.FindMiniSet)        // 根据ID获取MiniSet
		miniSetRouterWithoutRecord.GET("getMiniSetList", miniSetApi.GetMiniSetList)  // 获取MiniSet列表
	}
}
