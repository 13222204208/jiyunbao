package mini

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MiniUserRouter struct {
}

// InitMiniUserRouter 初始化 MiniUser 路由信息
func (s *MiniUserRouter) InitMiniUserRouter(Router *gin.RouterGroup) {
	miniUserRouter := Router.Group("miniUser").Use(middleware.OperationRecord())
	miniUserRouterWithoutRecord := Router.Group("miniUser")

	miniUserPublicRouter := Router.Group("miniUser")
	miniUserPrivateRouter := Router.Group("miniUser").Use(middleware.JWTAuthMiddlewareMini())
	var miniUserApi = v1.ApiGroupApp.MiniApiGroup.MiniUserApi
	{
		miniUserRouter.POST("createMiniUser", miniUserApi.CreateMiniUser)             // 新建MiniUser
		miniUserRouter.DELETE("deleteMiniUser", miniUserApi.DeleteMiniUser)           // 删除MiniUser
		miniUserRouter.DELETE("deleteMiniUserByIds", miniUserApi.DeleteMiniUserByIds) // 批量删除MiniUser
		miniUserRouter.PUT("updateMiniUser", miniUserApi.UpdateMiniUser)              // 更新MiniUser
	}
	{
		miniUserRouterWithoutRecord.GET("findMiniUser", miniUserApi.FindMiniUser)       // 根据ID获取MiniUser
		miniUserRouterWithoutRecord.GET("getMiniUserList", miniUserApi.GetMiniUserList) // 获取MiniUser列表
	}

	{
		//小程序用户登陆
		miniUserPublicRouter.POST("login", miniUserApi.Login)

	}

	{
		//获取小程序用户信息
		miniUserPrivateRouter.GET("info", miniUserApi.Info)
		//获取小程序手机号
		miniUserPrivateRouter.POST("getPhone", miniUserApi.GetPhone)
	}
}
