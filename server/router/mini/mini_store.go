package mini

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type MiniStoreRouter struct {
}

// InitMiniSetRouter 初始化 MiniSet 路由信息
func (s *MiniSetRouter) InitMiniStoreRouter(Router *gin.RouterGroup) {
	miniStorePublicRouter := Router.Group("store")
	var miniStoreApi = v1.ApiGroupApp.MiniApiGroup.StoreApi
	{
		miniStorePublicRouter.GET("list", miniStoreApi.List)
	}
}
