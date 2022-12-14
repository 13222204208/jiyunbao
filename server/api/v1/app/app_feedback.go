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

type AppFeedbackApi struct {
}

var appFeedbackService = service.ServiceGroupApp.AppServiceGroup.AppFeedbackService

// CreateAppFeedback 创建AppFeedback
// @Tags AppFeedback
// @Summary 创建AppFeedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppFeedback true "创建AppFeedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appFeedback/createAppFeedback [post]
func (appFeedbackApi *AppFeedbackApi) CreateAppFeedback(c *gin.Context) {
	var appFeedback app.AppFeedback
	err := c.ShouldBindJSON(&appFeedback)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appFeedbackService.CreateAppFeedback(appFeedback); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppFeedback 删除AppFeedback
// @Tags AppFeedback
// @Summary 删除AppFeedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppFeedback true "删除AppFeedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appFeedback/deleteAppFeedback [delete]
func (appFeedbackApi *AppFeedbackApi) DeleteAppFeedback(c *gin.Context) {
	var appFeedback app.AppFeedback
	err := c.ShouldBindJSON(&appFeedback)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appFeedbackService.DeleteAppFeedback(appFeedback); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppFeedbackByIds 批量删除AppFeedback
// @Tags AppFeedback
// @Summary 批量删除AppFeedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppFeedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appFeedback/deleteAppFeedbackByIds [delete]
func (appFeedbackApi *AppFeedbackApi) DeleteAppFeedbackByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appFeedbackService.DeleteAppFeedbackByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppFeedback 更新AppFeedback
// @Tags AppFeedback
// @Summary 更新AppFeedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppFeedback true "更新AppFeedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appFeedback/updateAppFeedback [put]
func (appFeedbackApi *AppFeedbackApi) UpdateAppFeedback(c *gin.Context) {
	var appFeedback app.AppFeedback
	err := c.ShouldBindJSON(&appFeedback)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appFeedbackService.UpdateAppFeedback(appFeedback); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppFeedback 用id查询AppFeedback
// @Tags AppFeedback
// @Summary 用id查询AppFeedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppFeedback true "用id查询AppFeedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appFeedback/findAppFeedback [get]
func (appFeedbackApi *AppFeedbackApi) FindAppFeedback(c *gin.Context) {
	var appFeedback app.AppFeedback
	err := c.ShouldBindQuery(&appFeedback)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappFeedback, err := appFeedbackService.GetAppFeedback(appFeedback.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappFeedback": reappFeedback}, c)
	}
}

// GetAppFeedbackList 分页获取AppFeedback列表
// @Tags AppFeedback
// @Summary 分页获取AppFeedback列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.AppFeedbackSearch true "分页获取AppFeedback列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appFeedback/getAppFeedbackList [get]
func (appFeedbackApi *AppFeedbackApi) GetAppFeedbackList(c *gin.Context) {
	var pageInfo appReq.AppFeedbackSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appFeedbackService.GetAppFeedbackInfoList(pageInfo); err != nil {
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

//提交反馈内容
func (appFeedbackApi *AppFeedbackApi) Create(c *gin.Context) {
	var appFeedback app.AppFeedback
	err := c.ShouldBindJSON(&appFeedback)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	verify := utils.Rules{
		"Contents": {utils.NotEmpty()},
	}
	if err := utils.Verify(appFeedback, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := c.MustGet("id").(uint)
	appFeedback.Uid = uid
	if err := appFeedbackService.CreateAppFeedback(appFeedback); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
