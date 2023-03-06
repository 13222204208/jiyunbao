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

type AppRateApi struct {
}

var appRateService = service.ServiceGroupApp.AppServiceGroup.AppRateService

// CreateAppRate 创建AppRate
// @Tags AppRate
// @Summary 创建AppRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppRate true "创建AppRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appRate/createAppRate [post]
func (appRateApi *AppRateApi) CreateAppRate(c *gin.Context) {
	var appRate app.AppRate
	err := c.ShouldBindJSON(&appRate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Title":  {utils.NotEmpty()},
		"Rate":   {utils.NotEmpty()},
		"Status": {utils.NotEmpty()},
	}
	if err := utils.Verify(appRate, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appRateService.CreateAppRate(appRate); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppRate 删除AppRate
// @Tags AppRate
// @Summary 删除AppRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppRate true "删除AppRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appRate/deleteAppRate [delete]
func (appRateApi *AppRateApi) DeleteAppRate(c *gin.Context) {
	var appRate app.AppRate
	err := c.ShouldBindJSON(&appRate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appRateService.DeleteAppRate(appRate); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppRateByIds 批量删除AppRate
// @Tags AppRate
// @Summary 批量删除AppRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appRate/deleteAppRateByIds [delete]
func (appRateApi *AppRateApi) DeleteAppRateByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appRateService.DeleteAppRateByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppRate 更新AppRate
// @Tags AppRate
// @Summary 更新AppRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppRate true "更新AppRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appRate/updateAppRate [put]
func (appRateApi *AppRateApi) UpdateAppRate(c *gin.Context) {
	var appRate app.AppRate
	err := c.ShouldBindJSON(&appRate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Title":  {utils.NotEmpty()},
		"Rate":   {utils.NotEmpty()},
		"Status": {utils.NotEmpty()},
	}
	if err := utils.Verify(appRate, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appRateService.UpdateAppRate(appRate); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppRate 用id查询AppRate
// @Tags AppRate
// @Summary 用id查询AppRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppRate true "用id查询AppRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appRate/findAppRate [get]
func (appRateApi *AppRateApi) FindAppRate(c *gin.Context) {
	var appRate app.AppRate
	err := c.ShouldBindQuery(&appRate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappRate, err := appRateService.GetAppRate(appRate.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappRate": reappRate}, c)
	}
}

// GetAppRateList 分页获取AppRate列表
// @Tags AppRate
// @Summary 分页获取AppRate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.AppRateSearch true "分页获取AppRate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appRate/getAppRateList [get]
func (appRateApi *AppRateApi) GetAppRateList(c *gin.Context) {
	var pageInfo appReq.AppRateSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appRateService.GetAppRateInfoList(pageInfo); err != nil {
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

func (appRateApi *AppRateApi) List(c *gin.Context) {
	if list, err := appRateService.List(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": list}, c)
	}
}
