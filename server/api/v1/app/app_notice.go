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

type AppNoticeApi struct {
}

var appNoticeService = service.ServiceGroupApp.AppServiceGroup.AppNoticeService

// CreateAppNotice 创建AppNotice
// @Tags AppNotice
// @Summary 创建AppNotice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppNotice true "创建AppNotice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appNotice/createAppNotice [post]
func (appNoticeApi *AppNoticeApi) CreateAppNotice(c *gin.Context) {
	var appNotice app.AppNotice
	err := c.ShouldBindJSON(&appNotice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Contents": {utils.NotEmpty()},
		"Way":      {utils.NotEmpty()},
	}
	if err := utils.Verify(appNotice, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appNoticeService.CreateAppNotice(appNotice); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppNotice 删除AppNotice
// @Tags AppNotice
// @Summary 删除AppNotice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppNotice true "删除AppNotice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appNotice/deleteAppNotice [delete]
func (appNoticeApi *AppNoticeApi) DeleteAppNotice(c *gin.Context) {
	var appNotice app.AppNotice
	err := c.ShouldBindJSON(&appNotice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appNoticeService.DeleteAppNotice(appNotice); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppNoticeByIds 批量删除AppNotice
// @Tags AppNotice
// @Summary 批量删除AppNotice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppNotice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appNotice/deleteAppNoticeByIds [delete]
func (appNoticeApi *AppNoticeApi) DeleteAppNoticeByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appNoticeService.DeleteAppNoticeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppNotice 更新AppNotice
// @Tags AppNotice
// @Summary 更新AppNotice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppNotice true "更新AppNotice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appNotice/updateAppNotice [put]
func (appNoticeApi *AppNoticeApi) UpdateAppNotice(c *gin.Context) {
	var appNotice app.AppNotice
	err := c.ShouldBindJSON(&appNotice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Contents": {utils.NotEmpty()},
		"Way":      {utils.NotEmpty()},
	}
	if err := utils.Verify(appNotice, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appNoticeService.UpdateAppNotice(appNotice); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppNotice 用id查询AppNotice
// @Tags AppNotice
// @Summary 用id查询AppNotice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppNotice true "用id查询AppNotice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appNotice/findAppNotice [get]
func (appNoticeApi *AppNoticeApi) FindAppNotice(c *gin.Context) {
	var appNotice app.AppNotice
	err := c.ShouldBindQuery(&appNotice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappNotice, err := appNoticeService.GetAppNotice(appNotice.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappNotice": reappNotice}, c)
	}
}

// GetAppNoticeList 分页获取AppNotice列表
// @Tags AppNotice
// @Summary 分页获取AppNotice列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.AppNoticeSearch true "分页获取AppNotice列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appNotice/getAppNoticeList [get]
func (appNoticeApi *AppNoticeApi) GetAppNoticeList(c *gin.Context) {
	var pageInfo appReq.AppNoticeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appNoticeService.GetAppNoticeInfoList(pageInfo); err != nil {
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

//获取公告内容
func (appNoticeApi *AppNoticeApi) Info(c *gin.Context) {
	var appNotice app.AppNotice
	err := c.ShouldBindQuery(&appNotice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	verify := utils.Rules{
		"Way": {utils.NotEmpty()},
	}
	if err := utils.Verify(appNotice, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if info, err := appNoticeService.Info(appNotice); err != nil {
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"info": info}, c)
	}
}
