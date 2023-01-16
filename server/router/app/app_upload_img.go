package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppUploadImgRouter struct {
}

// InitAppUploadImgRouter 初始化 AppUploadImg 路由信息
func (s *AppUploadImgRouter) InitAppUploadImgRouter(Router *gin.RouterGroup) {
	appUploadImgRouter := Router.Group("appUploadImg").Use(middleware.OperationRecord())
	appUploadImgRouterWithoutRecord := Router.Group("appUploadImg")
	var appUploadImgApi = v1.ApiGroupApp.AppApiGroup.AppUploadImgApi
	{
		appUploadImgRouter.POST("createAppUploadImg", appUploadImgApi.CreateAppUploadImg)             // 新建AppUploadImg
		appUploadImgRouter.DELETE("deleteAppUploadImg", appUploadImgApi.DeleteAppUploadImg)           // 删除AppUploadImg
		appUploadImgRouter.DELETE("deleteAppUploadImgByIds", appUploadImgApi.DeleteAppUploadImgByIds) // 批量删除AppUploadImg
		appUploadImgRouter.PUT("updateAppUploadImg", appUploadImgApi.UpdateAppUploadImg)              // 更新AppUploadImg

		//图片上送
		appUploadImgRouter.POST("imgSubmit", appUploadImgApi.ImgSubmit)

	}
	{
		appUploadImgRouterWithoutRecord.GET("findAppUploadImg", appUploadImgApi.FindAppUploadImg)       // 根据ID获取AppUploadImg
		appUploadImgRouterWithoutRecord.GET("getAppUploadImgList", appUploadImgApi.GetAppUploadImgList) // 获取AppUploadImg列表
	}
}
