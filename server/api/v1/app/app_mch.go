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

type AppMchApi struct {
}

var appMchService = service.ServiceGroupApp.AppServiceGroup.AppMchService

// CreateAppMch 创建AppMch
// @Tags AppMch
// @Summary 创建AppMch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppMch true "创建AppMch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appMch/createAppMch [post]
func (appMchApi *AppMchApi) CreateAppMch(c *gin.Context) {
	var appMch app.AppMch
	err := c.ShouldBindJSON(&appMch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"FirmName":    {utils.NotEmpty()},
		"MainType":    {utils.NotEmpty()},
		"Commission":  {utils.NotEmpty()},
		"LegalPerson": {utils.NotEmpty()},
		"Avatar":      {utils.NotEmpty()},
		"CardFront":   {utils.NotEmpty()},
		"CardReverse": {utils.NotEmpty()},
		"Entrust":     {utils.NotEmpty()},
		"Status":      {utils.NotEmpty()},
	}
	if err := utils.Verify(appMch, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appMchService.CreateAppMch(appMch); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppMch 删除AppMch
// @Tags AppMch
// @Summary 删除AppMch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppMch true "删除AppMch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appMch/deleteAppMch [delete]
func (appMchApi *AppMchApi) DeleteAppMch(c *gin.Context) {
	var appMch app.AppMch
	err := c.ShouldBindJSON(&appMch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appMchService.DeleteAppMch(appMch); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppMchByIds 批量删除AppMch
// @Tags AppMch
// @Summary 批量删除AppMch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppMch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appMch/deleteAppMchByIds [delete]
func (appMchApi *AppMchApi) DeleteAppMchByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appMchService.DeleteAppMchByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppMch 更新AppMch
// @Tags AppMch
// @Summary 更新AppMch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppMch true "更新AppMch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appMch/updateAppMch [put]
func (appMchApi *AppMchApi) UpdateAppMch(c *gin.Context) {
	var appMch app.AppMch
	err := c.ShouldBindJSON(&appMch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"FirmName":    {utils.NotEmpty()},
		"MainType":    {utils.NotEmpty()},
		"Commission":  {utils.NotEmpty()},
		"LegalPerson": {utils.NotEmpty()},
		"Avatar":      {utils.NotEmpty()},
		"CardFront":   {utils.NotEmpty()},
		"CardReverse": {utils.NotEmpty()},
		"Entrust":     {utils.NotEmpty()},
		"Status":      {utils.NotEmpty()},
	}
	if err := utils.Verify(appMch, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appMchService.UpdateAppMch(appMch); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppMch 用id查询AppMch
// @Tags AppMch
// @Summary 用id查询AppMch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppMch true "用id查询AppMch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appMch/findAppMch [get]
func (appMchApi *AppMchApi) FindAppMch(c *gin.Context) {
	var appMch app.AppMch
	err := c.ShouldBindQuery(&appMch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappMch, err := appMchService.GetAppMch(appMch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappMch": reappMch}, c)
	}
}

// GetAppMchList 分页获取AppMch列表
// @Tags AppMch
// @Summary 分页获取AppMch列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.AppMchSearch true "分页获取AppMch列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appMch/getAppMchList [get]
func (appMchApi *AppMchApi) GetAppMchList(c *gin.Context) {
	var pageInfo appReq.AppMchSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appMchService.GetAppMchInfoList(pageInfo); err != nil {
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

//商户认证
func (appMchApi *AppMchApi) Attestation(c *gin.Context) {
	var appMch app.AppMch
	err := c.ShouldBindJSON(&appMch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"FirmName":    {utils.NotEmpty()},
		"MainType":    {utils.NotEmpty()},
		"Commission":  {utils.NotEmpty()},
		"LegalPerson": {utils.NotEmpty()},
		"Avatar":      {utils.NotEmpty()},
		"CardFront":   {utils.NotEmpty()},
		"CardReverse": {utils.NotEmpty()},
		"Entrust":     {utils.NotEmpty()},
	}
	if err := utils.Verify(appMch, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := c.MustGet("id").(uint)
	appMch.Uid = uid
	if err := appMchService.Attestation(appMch); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

//修改商户认证信息
func (appMchApi *AppMchApi) Edit(c *gin.Context) {
	var appMch app.AppMch
	err := c.ShouldBindJSON(&appMch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := c.MustGet("id").(uint)
	if err := appMchService.Edit(uid, appMch); err != nil {
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}
