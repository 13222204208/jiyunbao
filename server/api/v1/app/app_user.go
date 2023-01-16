package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AppUserApi struct {
}

var appUserService = service.ServiceGroupApp.AppServiceGroup.AppUserService

// CreateAppUser 创建AppUser
// @Tags AppUser
// @Summary 创建AppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppUser true "创建AppUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appUser/createAppUser [post]
func (appUserApi *AppUserApi) CreateAppUser(c *gin.Context) {
	var appUser app.AppUser
	err := c.ShouldBindJSON(&appUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appUserService.CreateAppUser(appUser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppUser 删除AppUser
// @Tags AppUser
// @Summary 删除AppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppUser true "删除AppUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appUser/deleteAppUser [delete]
func (appUserApi *AppUserApi) DeleteAppUser(c *gin.Context) {
	var appUser app.AppUser
	err := c.ShouldBindJSON(&appUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appUserService.DeleteAppUser(appUser); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppUserByIds 批量删除AppUser
// @Tags AppUser
// @Summary 批量删除AppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appUser/deleteAppUserByIds [delete]
func (appUserApi *AppUserApi) DeleteAppUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appUserService.DeleteAppUserByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppUser 更新AppUser
// @Tags AppUser
// @Summary 更新AppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppUser true "更新AppUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appUser/updateAppUser [put]
func (appUserApi *AppUserApi) UpdateAppUser(c *gin.Context) {
	var appUser app.AppUser
	err := c.ShouldBindJSON(&appUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appUserService.UpdateAppUser(appUser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppUser 用id查询AppUser
// @Tags AppUser
// @Summary 用id查询AppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppUser true "用id查询AppUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appUser/findAppUser [get]
func (appUserApi *AppUserApi) FindAppUser(c *gin.Context) {
	var appUser app.AppUser
	err := c.ShouldBindQuery(&appUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappUser, err := appUserService.GetAppUser(appUser.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappUser": reappUser}, c)
	}
}

// GetAppUserList 分页获取AppUser列表
// @Tags AppUser
// @Summary 分页获取AppUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.AppUserSearch true "分页获取AppUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appUser/getAppUserList [get]
func (appUserApi *AppUserApi) GetAppUserList(c *gin.Context) {
	var pageInfo appReq.AppUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appUserService.GetAppUserInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

//注册
func (appUserApi *AppUserApi) Register(c *gin.Context) {
	var appUser app.AppUser
	err := c.ShouldBindJSON(&appUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := utils.Verify(appUser, utils.AppUserRegister); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := appUserService.Register(appUser); err != nil {

		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("注册成功", c)
	}
}

func (appUserApi *AppUserApi) Login(c *gin.Context) {
	var p appReq.Login
	err := c.ShouldBindJSON(&p)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(p, utils.UserCode); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	phone := p.Phone
	password := p.Password
	code := p.Code

	if err, info := appUserService.Login(phone, password, code); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(gin.H{"data": info}, c)
	}

}

//修改密码
func (appUserApi *AppUserApi) UpdatePassword(c *gin.Context) {
	var p appReq.UpdatePassword
	err := c.ShouldBindJSON(&p)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(p, utils.AppUserUpdatePassword); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	phone := c.MustGet("phone").(string)

	if err := appUserService.UpdatePassword(phone, p.OldPassword, p.NewPassword); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithMessage("修改成功", c)
	}

}

//忘记密码
func (appUserApi *AppUserApi) ForgetPassword(c *gin.Context) {
	var p appReq.ForgetPassword
	err := c.ShouldBindJSON(&p)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(p, utils.AppUserForgetPassword); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := appUserService.ForgetPassword(p); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

//修改昵称
func (appUserApi *AppUserApi) UpdateNickname(c *gin.Context) {
	var p appReq.UpdateNickname
	err := c.ShouldBindJSON(&p)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(p, utils.AppUserUpdateNickname); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	phone := c.MustGet("phone").(string)

	if err := appUserService.UpdateNickname(phone, p.Nickname); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

//修改头像
func (appUserApi *AppUserApi) UpdateAvatar(c *gin.Context) {
	var p appReq.UpdateAvatar
	err := c.ShouldBindJSON(&p)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(p, utils.AppUserUpdateAvatar); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	phone := c.MustGet("phone").(string)

	if err := appUserService.UpdateAvatar(phone, p.Avatar); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

//修改手机号
func (appUserApi *AppUserApi) UpdatePhone(c *gin.Context) {
	var p appReq.UpdatePhone
	err := c.ShouldBindJSON(&p)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(p, utils.AppUserUpdatePhone); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	phone := c.MustGet("phone").(string)
	if phone != p.OldPhone {
		response.FailWithMessage("不是当前用户手机号", c)
		return
	}

	if err := appUserService.UpdatePhone(p.OldPhone, p.OldCode, p.NewPhone, p.NewCode); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

//短信发送
func (appUserApi *AppUserApi) Code(c *gin.Context) {
	var p appReq.Phone
	err := c.ShouldBindJSON(&p)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := utils.Verify(p, utils.UserCode); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := appUserService.SendCode(p.Phone); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithMessage("发送成功", c)
	}
}

//获取用户信息
func (appUserApi *AppUserApi) Info(c *gin.Context) {
	uid := c.MustGet("id").(uint)
	if err, info, storeName := appUserService.Info(uid); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(gin.H{"info": info, "storeName": storeName}, c)
	}
}

//获取用户身份证信息
func (appUserApi *AppUserApi) IdImage(c *gin.Context) {
	var p appReq.IdImage
	err := c.ShouldBindJSON(&p)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := utils.Verify(p, utils.MiniIdImageVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err, info := appUserService.IdImage(p.Image, p.Side); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(gin.H{"info": info}, c)
	}
}

//获取银行卡信息
func (appUserApi *AppUserApi) BankCard(c *gin.Context) {
	var p appReq.BankCard
	err := c.ShouldBindJSON(&p)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := utils.Verify(p, utils.MiniBankCardVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err, info := appUserService.BankCard(p.Image); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(gin.H{"info": info}, c)
	}
}

//获取营业执照信息
func (appUserApi *AppUserApi) Bus(c *gin.Context) {
	var p appReq.Bus
	err := c.ShouldBindJSON(&p)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := utils.Verify(p, utils.MiniBusVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err, info := appUserService.Bus(p.Image); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(gin.H{"info": info}, c)
	}
}
