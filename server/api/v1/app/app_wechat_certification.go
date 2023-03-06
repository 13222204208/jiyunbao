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

type AppWechatCertificationApi struct {
}

var appWechatCertificationService = service.ServiceGroupApp.AppServiceGroup.AppWechatCertificationService

// CreateAppWechatCertification 创建AppWechatCertification
// @Tags AppWechatCertification
// @Summary 创建AppWechatCertification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppWechatCertification true "创建AppWechatCertification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appWechatCertification/createAppWechatCertification [post]
func (appWechatCertificationApi *AppWechatCertificationApi) CreateAppWechatCertification(c *gin.Context) {
	var appWechatCertification app.AppWechatCertification
	err := c.ShouldBindJSON(&appWechatCertification)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"MercId": {utils.NotEmpty()},
	}
	if err := utils.Verify(appWechatCertification, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appWechatCertificationService.CreateAppWechatCertification(appWechatCertification); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppWechatCertification 删除AppWechatCertification
// @Tags AppWechatCertification
// @Summary 删除AppWechatCertification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppWechatCertification true "删除AppWechatCertification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appWechatCertification/deleteAppWechatCertification [delete]
func (appWechatCertificationApi *AppWechatCertificationApi) DeleteAppWechatCertification(c *gin.Context) {
	var appWechatCertification app.AppWechatCertification
	err := c.ShouldBindJSON(&appWechatCertification)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appWechatCertificationService.DeleteAppWechatCertification(appWechatCertification); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppWechatCertificationByIds 批量删除AppWechatCertification
// @Tags AppWechatCertification
// @Summary 批量删除AppWechatCertification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppWechatCertification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appWechatCertification/deleteAppWechatCertificationByIds [delete]
func (appWechatCertificationApi *AppWechatCertificationApi) DeleteAppWechatCertificationByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appWechatCertificationService.DeleteAppWechatCertificationByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppWechatCertification 更新AppWechatCertification
// @Tags AppWechatCertification
// @Summary 更新AppWechatCertification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppWechatCertification true "更新AppWechatCertification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appWechatCertification/updateAppWechatCertification [put]
func (appWechatCertificationApi *AppWechatCertificationApi) UpdateAppWechatCertification(c *gin.Context) {
	var appWechatCertification app.AppWechatCertification
	err := c.ShouldBindJSON(&appWechatCertification)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"MercId": {utils.NotEmpty()},
	}
	if err := utils.Verify(appWechatCertification, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appWechatCertificationService.UpdateAppWechatCertification(appWechatCertification); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppWechatCertification 用id查询AppWechatCertification
// @Tags AppWechatCertification
// @Summary 用id查询AppWechatCertification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppWechatCertification true "用id查询AppWechatCertification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appWechatCertification/findAppWechatCertification [get]
func (appWechatCertificationApi *AppWechatCertificationApi) FindAppWechatCertification(c *gin.Context) {
	var appWechatCertification app.AppWechatCertification
	err := c.ShouldBindQuery(&appWechatCertification)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappWechatCertification, err := appWechatCertificationService.GetAppWechatCertification(appWechatCertification.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappWechatCertification": reappWechatCertification}, c)
	}
}

// GetAppWechatCertificationList 分页获取AppWechatCertification列表
// @Tags AppWechatCertification
// @Summary 分页获取AppWechatCertification列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.AppWechatCertificationSearch true "分页获取AppWechatCertification列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appWechatCertification/getAppWechatCertificationList [get]
func (appWechatCertificationApi *AppWechatCertificationApi) GetAppWechatCertificationList(c *gin.Context) {
	var pageInfo appReq.AppWechatCertificationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appWechatCertificationService.GetAppWechatCertificationInfoList(pageInfo); err != nil {
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

//微信实名认证
func (appWechatCertificationApi *AppWechatCertificationApi) WechatCertification(c *gin.Context) {
	var appWechatCertification app.AppWechatCertification
	err := c.ShouldBindJSON(&appWechatCertification)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	if err := utils.Verify(appWechatCertification, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appWechatCertificationService.WechatCertification(appWechatCertification.ID); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("上送成功", c)
	}
}

//微信实名认证状态查询
func (appWechatCertificationApi *AppWechatCertificationApi) CertificationState(c *gin.Context) {
	var appWechatCertification app.AppWechatCertification
	err := c.ShouldBindQuery(&appWechatCertification)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appWechatCertificationService.CertificationState(appWechatCertification.ApplyNo); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("审核通过，请扫描微信支付返回的小程序码在小程序端完成授权流程", c)
	}
}

//微信实名认证明细查询
func (appWechatCertificationApi *AppWechatCertificationApi) GetAuthMessagesByMercId(c *gin.Context) {
	var appWechatCertification app.AppWechatCertification
	err := c.ShouldBindQuery(&appWechatCertification)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if appWechatCertification.MercId == "" {
		response.FailWithMessage("商户号不能为空", c)
		return
	}
	if err, statusMsg := appWechatCertificationService.GetAuthMessagesByMercId(appWechatCertification.MercId); err != nil {
		global.GVA_LOG.Error("查询明细失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage(statusMsg, c)
	}
}

//微信实名认证授权状态查询
func (appWechatCertificationApi *AppWechatCertificationApi) GetAuthState(c *gin.Context) {
	var appWechatCertification app.AppWechatCertification
	err := c.ShouldBindQuery(&appWechatCertification)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if appWechatCertification.MercId == "" {
		response.FailWithMessage("商户号不能为空", c)
		return
	}
	if err, statusMsg := appWechatCertificationService.GetAuthState(appWechatCertification.MercId); err != nil {
		global.GVA_LOG.Error("查询授权状态失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage(statusMsg, c)
	}
}

func (appWechatCertificationApi *AppWechatCertificationApi) Authentication(c *gin.Context) {

	uid := c.MustGet("id").(uint)
	if err, qrcode := appWechatCertificationService.Authentication(uid); err != nil {

		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(gin.H{"qrcode": qrcode}, c)
	}
}
