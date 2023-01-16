package mini

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MiniClassifyRouter struct {
}

// InitMiniClassifyRouter 初始化 MiniClassify 路由信息
func (s *MiniClassifyRouter) InitMiniClassifyRouter(Router *gin.RouterGroup) {
	miniClassifyRouter := Router.Group("miniClassify").Use(middleware.OperationRecord())
	miniClassifyRouterWithoutRecord := Router.Group("miniClassify")

	miniClassifyPublicRouter := Router.Group("miniClassify")

	var miniClassifyApi = v1.ApiGroupApp.MiniApiGroup.MiniClassifyApi
	{
		miniClassifyRouter.POST("createMiniClassify", miniClassifyApi.CreateMiniClassify)             // 新建MiniClassify
		miniClassifyRouter.DELETE("deleteMiniClassify", miniClassifyApi.DeleteMiniClassify)           // 删除MiniClassify
		miniClassifyRouter.DELETE("deleteMiniClassifyByIds", miniClassifyApi.DeleteMiniClassifyByIds) // 批量删除MiniClassify
		miniClassifyRouter.PUT("updateMiniClassify", miniClassifyApi.UpdateMiniClassify)              // 更新MiniClassify
	}
	{
		miniClassifyRouterWithoutRecord.GET("findMiniClassify", miniClassifyApi.FindMiniClassify)       // 根据ID获取MiniClassify
		miniClassifyRouterWithoutRecord.GET("getMiniClassifyList", miniClassifyApi.GetMiniClassifyList) // 获取MiniClassify列表
	}

	{
		miniClassifyPublicRouter.GET("list", miniClassifyApi.List)
	}
}
