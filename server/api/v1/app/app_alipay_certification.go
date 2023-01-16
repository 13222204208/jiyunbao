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

type AppAlipayCertificationApi struct {
}

var appAlipayCertificationService = service.ServiceGroupApp.AppServiceGroup.AppAlipayCertificationService

// CreateAppAlipayCertification 创建AppAlipayCertification
// @Tags AppAlipayCertification
// @Summary 创建AppAlipayCertification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppAlipayCertification true "创建AppAlipayCertification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appAlipayCertification/createAppAlipayCertification [post]
func (appAlipayCertificationApi *AppAlipayCertificationApi) CreateAppAlipayCertification(c *gin.Context) {
	var appAlipayCertification app.AppAlipayCertification
	err := c.ShouldBindJSON(&appAlipayCertification)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"MercId": {utils.NotEmpty()},
	}
	if err := utils.Verify(appAlipayCertification, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appAlipayCertificationService.CreateAppAlipayCertification(appAlipayCertification); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppAlipayCertification 删除AppAlipayCertification
// @Tags AppAlipayCertification
// @Summary 删除AppAlipayCertification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppAlipayCertification true "删除AppAlipayCertification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appAlipayCertification/deleteAppAlipayCertification [delete]
func (appAlipayCertificationApi *AppAlipayCertificationApi) DeleteAppAlipayCertification(c *gin.Context) {
	var appAlipayCertification app.AppAlipayCertification
	err := c.ShouldBindJSON(&appAlipayCertification)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appAlipayCertificationService.DeleteAppAlipayCertification(appAlipayCertification); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppAlipayCertificationByIds 批量删除AppAlipayCertification
// @Tags AppAlipayCertification
// @Summary 批量删除AppAlipayCertification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppAlipayCertification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appAlipayCertification/deleteAppAlipayCertificationByIds [delete]
func (appAlipayCertificationApi *AppAlipayCertificationApi) DeleteAppAlipayCertificationByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appAlipayCertificationService.DeleteAppAlipayCertificationByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppAlipayCertification 更新AppAlipayCertification
// @Tags AppAlipayCertification
// @Summary 更新AppAlipayCertification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppAlipayCertification true "更新AppAlipayCertification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appAlipayCertification/updateAppAlipayCertification [put]
func (appAlipayCertificationApi *AppAlipayCertificationApi) UpdateAppAlipayCertification(c *gin.Context) {
	var appAlipayCertification app.AppAlipayCertification
	err := c.ShouldBindJSON(&appAlipayCertification)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"MercId": {utils.NotEmpty()},
	}
	if err := utils.Verify(appAlipayCertification, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appAlipayCertificationService.UpdateAppAlipayCertification(appAlipayCertification); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppAlipayCertification 用id查询AppAlipayCertification
// @Tags AppAlipayCertification
// @Summary 用id查询AppAlipayCertification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppAlipayCertification true "用id查询AppAlipayCertification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appAlipayCertification/findAppAlipayCertification [get]
func (appAlipayCertificationApi *AppAlipayCertificationApi) FindAppAlipayCertification(c *gin.Context) {
	var appAlipayCertification app.AppAlipayCertification
	err := c.ShouldBindQuery(&appAlipayCertification)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappAlipayCertification, err := appAlipayCertificationService.GetAppAlipayCertification(appAlipayCertification.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappAlipayCertification": reappAlipayCertification}, c)
	}
}

// GetAppAlipayCertificationList 分页获取AppAlipayCertification列表
// @Tags AppAlipayCertification
// @Summary 分页获取AppAlipayCertification列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.AppAlipayCertificationSearch true "分页获取AppAlipayCertification列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appAlipayCertification/getAppAlipayCertificationList [get]
func (appAlipayCertificationApi *AppAlipayCertificationApi) GetAppAlipayCertificationList(c *gin.Context) {
	var pageInfo appReq.AppAlipayCertificationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appAlipayCertificationService.GetAppAlipayCertificationInfoList(pageInfo); err != nil {
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

//支付宝实名认证
func (appAlipayCertificationApi *AppAlipayCertificationApi) AlipayCertification(c *gin.Context) {
	var a app.AppAlipayCertification
	err := c.ShouldBindJSON(&a)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	if err := utils.Verify(a, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appAlipayCertificationService.AlipayCertification(a.ID); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("上送成功", c)
	}
}

//支付宝实名认证状态查询
func (appAlipayCertificationApi *AppAlipayCertificationApi) AlipayCertificationState(c *gin.Context) {
	var a app.AppAlipayCertification
	err := c.ShouldBindQuery(&a)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appAlipayCertificationService.AlipayCertificationState(a.MercId, a.SendChannelApplyNo); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("审核通过，请扫描返回的小程序码在小程序端完成授权流程", c)
	}
}

//微信实名认证授权状态查询
func (appAlipayCertificationApi *AppAlipayCertificationApi) GetAlipayAuthState(c *gin.Context) {
	var a app.AppAlipayCertification
	err := c.ShouldBindQuery(&a)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if a.MercId == "" {
		response.FailWithMessage("商户号不能为空", c)
		return
	}
	if err, statusMsg := appAlipayCertificationService.GetAlipayAuthState(a.MercId); err != nil {
		global.GVA_LOG.Error("查询授权状态失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage(statusMsg, c)
	}
}
