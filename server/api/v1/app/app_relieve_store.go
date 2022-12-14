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

type AppRelieveStoreApi struct {
}

var appRelieveStoreService = service.ServiceGroupApp.AppServiceGroup.AppRelieveStoreService

// CreateAppRelieveStore 创建AppRelieveStore
// @Tags AppRelieveStore
// @Summary 创建AppRelieveStore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppRelieveStore true "创建AppRelieveStore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appRelieveStore/createAppRelieveStore [post]
func (appRelieveStoreApi *AppRelieveStoreApi) CreateAppRelieveStore(c *gin.Context) {
	var appRelieveStore app.AppRelieveStore
	err := c.ShouldBindJSON(&appRelieveStore)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Contents": {utils.NotEmpty()},
	}
	if err := utils.Verify(appRelieveStore, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appRelieveStoreService.CreateAppRelieveStore(appRelieveStore); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppRelieveStore 删除AppRelieveStore
// @Tags AppRelieveStore
// @Summary 删除AppRelieveStore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppRelieveStore true "删除AppRelieveStore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appRelieveStore/deleteAppRelieveStore [delete]
func (appRelieveStoreApi *AppRelieveStoreApi) DeleteAppRelieveStore(c *gin.Context) {
	var appRelieveStore app.AppRelieveStore
	err := c.ShouldBindJSON(&appRelieveStore)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appRelieveStoreService.DeleteAppRelieveStore(appRelieveStore); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppRelieveStoreByIds 批量删除AppRelieveStore
// @Tags AppRelieveStore
// @Summary 批量删除AppRelieveStore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppRelieveStore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appRelieveStore/deleteAppRelieveStoreByIds [delete]
func (appRelieveStoreApi *AppRelieveStoreApi) DeleteAppRelieveStoreByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appRelieveStoreService.DeleteAppRelieveStoreByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppRelieveStore 更新AppRelieveStore
// @Tags AppRelieveStore
// @Summary 更新AppRelieveStore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppRelieveStore true "更新AppRelieveStore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appRelieveStore/updateAppRelieveStore [put]
func (appRelieveStoreApi *AppRelieveStoreApi) UpdateAppRelieveStore(c *gin.Context) {
	var appRelieveStore app.AppRelieveStore
	err := c.ShouldBindJSON(&appRelieveStore)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Contents": {utils.NotEmpty()},
	}
	if err := utils.Verify(appRelieveStore, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appRelieveStoreService.UpdateAppRelieveStore(appRelieveStore); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppRelieveStore 用id查询AppRelieveStore
// @Tags AppRelieveStore
// @Summary 用id查询AppRelieveStore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppRelieveStore true "用id查询AppRelieveStore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appRelieveStore/findAppRelieveStore [get]
func (appRelieveStoreApi *AppRelieveStoreApi) FindAppRelieveStore(c *gin.Context) {
	var appRelieveStore app.AppRelieveStore
	err := c.ShouldBindQuery(&appRelieveStore)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappRelieveStore, err := appRelieveStoreService.GetAppRelieveStore(appRelieveStore.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappRelieveStore": reappRelieveStore}, c)
	}
}

// GetAppRelieveStoreList 分页获取AppRelieveStore列表
// @Tags AppRelieveStore
// @Summary 分页获取AppRelieveStore列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.AppRelieveStoreSearch true "分页获取AppRelieveStore列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appRelieveStore/getAppRelieveStoreList [get]
func (appRelieveStoreApi *AppRelieveStoreApi) GetAppRelieveStoreList(c *gin.Context) {
	var pageInfo appReq.AppRelieveStoreSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appRelieveStoreService.GetAppRelieveStoreInfoList(pageInfo); err != nil {
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

//申请解除合作
func (appRelieveStoreApi *AppRelieveStoreApi) Apply(c *gin.Context) {
	var appRelieveStore app.AppRelieveStore
	err := c.ShouldBindJSON(&appRelieveStore)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Contents": {utils.NotEmpty()},
		"StoreId":  {utils.NotEmpty()},
	}
	if err := utils.Verify(appRelieveStore, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := c.MustGet("id").(uint)
	appRelieveStore.Uid = uid
	if err := appRelieveStoreService.Apply(appRelieveStore); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
