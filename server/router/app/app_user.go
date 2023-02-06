package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppUserRouter struct {
}

// InitAppUserRouter 初始化 AppUser 路由信息
func (s *AppUserRouter) InitAppUserRouter(Router *gin.RouterGroup) {
	appUserRouter := Router.Group("appUser").Use(middleware.OperationRecord())
	appUserRouterWithoutRecord := Router.Group("appUser")

	exaFileUploadAndDownloadApi := v1.ApiGroupApp.ExampleApiGroup.FileUploadAndDownloadApi

	appUserPublicRouter := Router.Group("user")
	appUserPrivateRouter := Router.Group("user").Use(middleware.JWTAuthMiddleware())
	var appUserApi = v1.ApiGroupApp.AppApiGroup.AppUserApi
	{
		appUserRouter.POST("createAppUser", appUserApi.CreateAppUser)             // 新建AppUser
		appUserRouter.DELETE("deleteAppUser", appUserApi.DeleteAppUser)           // 删除AppUser
		appUserRouter.DELETE("deleteAppUserByIds", appUserApi.DeleteAppUserByIds) // 批量删除AppUser
		appUserRouter.PUT("updateAppUser", appUserApi.UpdateAppUser)              // 更新AppUser
	}
	{
		appUserRouterWithoutRecord.GET("findAppUser", appUserApi.FindAppUser)       // 根据ID获取AppUser
		appUserRouterWithoutRecord.GET("getAppUserList", appUserApi.GetAppUserList) // 获取AppUser列表
	}

	{
		//我要开店 注册
		appUserPublicRouter.POST("register", appUserApi.Register)
		//登陆
		appUserPublicRouter.POST("login", appUserApi.Login)
		//机构列表
		appUserPublicRouter.GET("institutions", appUserApi.Institutions)
	}

	{
		//图片上传
		appUserPrivateRouter.POST("upload", exaFileUploadAndDownloadApi.UploadFile)

		//修改密码
		appUserPrivateRouter.PUT("password", appUserApi.UpdatePassword)

		//忘记密码
		appUserPublicRouter.PUT("forgetPassword", appUserApi.ForgetPassword)

		//修改昵称
		appUserPrivateRouter.PUT("nickname", appUserApi.UpdateNickname)

		//修改头像
		appUserPrivateRouter.PUT("avatar", appUserApi.UpdateAvatar)

		//修改手机号
		appUserPrivateRouter.PUT("phone", appUserApi.UpdatePhone)

		//获取用户信息
		appUserPrivateRouter.GET("info", appUserApi.Info)

		//身份证识别
		appUserPrivateRouter.POST("idImage", appUserApi.IdImage)

		//银行卡识别
		appUserPrivateRouter.POST("bankCard", appUserApi.BankCard)

		//营业执照识别
		appUserPrivateRouter.POST("bus", appUserApi.Bus)

		//发送手机验证码
		appUserPublicRouter.POST("code", appUserApi.Code)
	}
}
