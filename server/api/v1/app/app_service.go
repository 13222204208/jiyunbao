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

type AppServiceApi struct {
}

var appServiceService = service.ServiceGroupApp.AppServiceGroup.AppServiceService

// CreateAppService 创建AppService
// @Tags AppService
// @Summary 创建AppService
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppService true "创建AppService"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appService/createAppService [post]
func (appServiceApi *AppServiceApi) CreateAppService(c *gin.Context) {
	var appService app.AppService
	err := c.ShouldBindJSON(&appService)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Title": {utils.NotEmpty()},
		"Price": {utils.NotEmpty()},
		"Way":   {utils.NotEmpty()},
	}
	if err := utils.Verify(appService, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appServiceService.CreateAppService(appService); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppService 删除AppService
// @Tags AppService
// @Summary 删除AppService
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppService true "删除AppService"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appService/deleteAppService [delete]
func (appServiceApi *AppServiceApi) DeleteAppService(c *gin.Context) {
	var appService app.AppService
	err := c.ShouldBindJSON(&appService)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appServiceService.DeleteAppService(appService); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppServiceByIds 批量删除AppService
// @Tags AppService
// @Summary 批量删除AppService
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppService"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appService/deleteAppServiceByIds [delete]
func (appServiceApi *AppServiceApi) DeleteAppServiceByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appServiceService.DeleteAppServiceByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppService 更新AppService
// @Tags AppService
// @Summary 更新AppService
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppService true "更新AppService"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appService/updateAppService [put]
func (appServiceApi *AppServiceApi) UpdateAppService(c *gin.Context) {
	var appService app.AppService
	err := c.ShouldBindJSON(&appService)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Title": {utils.NotEmpty()},
		"Price": {utils.NotEmpty()},
		"Way":   {utils.NotEmpty()},
	}
	if err := utils.Verify(appService, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appServiceService.UpdateAppService(appService); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppService 用id查询AppService
// @Tags AppService
// @Summary 用id查询AppService
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppService true "用id查询AppService"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appService/findAppService [get]
func (appServiceApi *AppServiceApi) FindAppService(c *gin.Context) {
	var appService app.AppService
	err := c.ShouldBindQuery(&appService)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappService, err := appServiceService.GetAppService(appService.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappService": reappService}, c)
	}
}

// GetAppServiceList 分页获取AppService列表
// @Tags AppService
// @Summary 分页获取AppService列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.AppServiceSearch true "分页获取AppService列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appService/getAppServiceList [get]
func (appServiceApi *AppServiceApi) GetAppServiceList(c *gin.Context) {
	var pageInfo appReq.AppServiceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appServiceService.GetAppServiceInfoList(pageInfo); err != nil {
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

//获取服务设置的价格信息
func (appServiceApi *AppServiceApi) List(c *gin.Context) {
	if err, list := appServiceService.List(); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": list}, c)
	}
}
