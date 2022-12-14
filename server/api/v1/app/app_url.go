package app

import (
	"fmt"
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

type AppUrlApi struct {
}

var appUrlService = service.ServiceGroupApp.AppServiceGroup.AppUrlService

// CreateAppUrl 创建AppUrl
// @Tags AppUrl
// @Summary 创建AppUrl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppUrl true "创建AppUrl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appUrl/createAppUrl [post]
func (appUrlApi *AppUrlApi) CreateAppUrl(c *gin.Context) {
	var appUrl app.AppUrl
	err := c.ShouldBindJSON(&appUrl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Url": {utils.NotEmpty()},
		"Way": {utils.NotEmpty()},
	}
	if err := utils.Verify(appUrl, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appUrlService.CreateAppUrl(appUrl); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppUrl 删除AppUrl
// @Tags AppUrl
// @Summary 删除AppUrl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppUrl true "删除AppUrl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appUrl/deleteAppUrl [delete]
func (appUrlApi *AppUrlApi) DeleteAppUrl(c *gin.Context) {
	var appUrl app.AppUrl
	err := c.ShouldBindJSON(&appUrl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appUrlService.DeleteAppUrl(appUrl); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppUrlByIds 批量删除AppUrl
// @Tags AppUrl
// @Summary 批量删除AppUrl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppUrl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appUrl/deleteAppUrlByIds [delete]
func (appUrlApi *AppUrlApi) DeleteAppUrlByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appUrlService.DeleteAppUrlByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppUrl 更新AppUrl
// @Tags AppUrl
// @Summary 更新AppUrl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppUrl true "更新AppUrl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appUrl/updateAppUrl [put]
func (appUrlApi *AppUrlApi) UpdateAppUrl(c *gin.Context) {
	var appUrl app.AppUrl
	err := c.ShouldBindJSON(&appUrl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Url": {utils.NotEmpty()},
		"Way": {utils.NotEmpty()},
	}
	if err := utils.Verify(appUrl, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appUrlService.UpdateAppUrl(appUrl); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppUrl 用id查询AppUrl
// @Tags AppUrl
// @Summary 用id查询AppUrl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppUrl true "用id查询AppUrl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appUrl/findAppUrl [get]
func (appUrlApi *AppUrlApi) FindAppUrl(c *gin.Context) {
	var appUrl app.AppUrl
	err := c.ShouldBindQuery(&appUrl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappUrl, err := appUrlService.GetAppUrl(appUrl.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappUrl": reappUrl}, c)
	}
}

// GetAppUrlList 分页获取AppUrl列表
// @Tags AppUrl
// @Summary 分页获取AppUrl列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.AppUrlSearch true "分页获取AppUrl列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appUrl/getAppUrlList [get]
func (appUrlApi *AppUrlApi) GetAppUrlList(c *gin.Context) {
	var pageInfo appReq.AppUrlSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appUrlService.GetAppUrlInfoList(pageInfo); err != nil {
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

func (appUrlApi *AppUrlApi) Notify(c *gin.Context) {
	var b appReq.BizContent
	err := c.ShouldBindQuery(&b)
	if err != nil {
		fmt.Println("错误", err.Error())
	}
	fmt.Println("返回的结果", err)
}
