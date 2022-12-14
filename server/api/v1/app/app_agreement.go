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

type AppAgreementApi struct {
}

var appAgreementService = service.ServiceGroupApp.AppServiceGroup.AppAgreementService

// CreateAppAgreement 创建AppAgreement
// @Tags AppAgreement
// @Summary 创建AppAgreement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppAgreement true "创建AppAgreement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appAgreement/createAppAgreement [post]
func (appAgreementApi *AppAgreementApi) CreateAppAgreement(c *gin.Context) {
	var appAgreement app.AppAgreement
	err := c.ShouldBindJSON(&appAgreement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Way": {utils.NotEmpty()},
	}
	if err := utils.Verify(appAgreement, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appAgreementService.CreateAppAgreement(appAgreement); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppAgreement 删除AppAgreement
// @Tags AppAgreement
// @Summary 删除AppAgreement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppAgreement true "删除AppAgreement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appAgreement/deleteAppAgreement [delete]
func (appAgreementApi *AppAgreementApi) DeleteAppAgreement(c *gin.Context) {
	var appAgreement app.AppAgreement
	err := c.ShouldBindJSON(&appAgreement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appAgreementService.DeleteAppAgreement(appAgreement); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppAgreementByIds 批量删除AppAgreement
// @Tags AppAgreement
// @Summary 批量删除AppAgreement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppAgreement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appAgreement/deleteAppAgreementByIds [delete]
func (appAgreementApi *AppAgreementApi) DeleteAppAgreementByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appAgreementService.DeleteAppAgreementByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppAgreement 更新AppAgreement
// @Tags AppAgreement
// @Summary 更新AppAgreement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppAgreement true "更新AppAgreement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appAgreement/updateAppAgreement [put]
func (appAgreementApi *AppAgreementApi) UpdateAppAgreement(c *gin.Context) {
	var appAgreement app.AppAgreement
	err := c.ShouldBindJSON(&appAgreement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Way": {utils.NotEmpty()},
	}
	if err := utils.Verify(appAgreement, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appAgreementService.UpdateAppAgreement(appAgreement); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppAgreement 用id查询AppAgreement
// @Tags AppAgreement
// @Summary 用id查询AppAgreement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppAgreement true "用id查询AppAgreement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appAgreement/findAppAgreement [get]
func (appAgreementApi *AppAgreementApi) FindAppAgreement(c *gin.Context) {
	var appAgreement app.AppAgreement
	err := c.ShouldBindQuery(&appAgreement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappAgreement, err := appAgreementService.GetAppAgreement(appAgreement.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappAgreement": reappAgreement}, c)
	}
}

// GetAppAgreementList 分页获取AppAgreement列表
// @Tags AppAgreement
// @Summary 分页获取AppAgreement列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.AppAgreementSearch true "分页获取AppAgreement列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appAgreement/getAppAgreementList [get]
func (appAgreementApi *AppAgreementApi) GetAppAgreementList(c *gin.Context) {
	var pageInfo appReq.AppAgreementSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appAgreementService.GetAppAgreementInfoList(pageInfo); err != nil {
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

//获取协议，用户协议，隐私政策
func (appAgreementApi *AppAgreementApi) Info(c *gin.Context) {
	var a app.AppAgreement
	err := c.ShouldBindQuery(&a)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if info, err := appAgreementService.Info(a); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"info": info}, c)
	}
}
