package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AppStoreApi struct {
}

var appStoreService = service.ServiceGroupApp.AppServiceGroup.AppStoreService

// CreateAppStore 创建AppStore
// @Tags AppStore
// @Summary 创建AppStore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppStore true "创建AppStore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appStore/createAppStore [post]
func (appStoreApi *AppStoreApi) CreateAppStore(c *gin.Context) {
	var appStore app.AppStore
	err := c.ShouldBindJSON(&appStore)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appStoreService.CreateAppStore(appStore); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppStore 删除AppStore
// @Tags AppStore
// @Summary 删除AppStore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppStore true "删除AppStore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appStore/deleteAppStore [delete]
func (appStoreApi *AppStoreApi) DeleteAppStore(c *gin.Context) {
	var appStore app.AppStore
	err := c.ShouldBindJSON(&appStore)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appStoreService.DeleteAppStore(appStore); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppStoreByIds 批量删除AppStore
// @Tags AppStore
// @Summary 批量删除AppStore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppStore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appStore/deleteAppStoreByIds [delete]
func (appStoreApi *AppStoreApi) DeleteAppStoreByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appStoreService.DeleteAppStoreByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppStore 更新AppStore
// @Tags AppStore
// @Summary 更新AppStore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppStore true "更新AppStore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appStore/updateAppStore [put]
func (appStoreApi *AppStoreApi) UpdateAppStore(c *gin.Context) {
	var appStore app.AppStore
	err := c.ShouldBindJSON(&appStore)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appStoreService.UpdateAppStore(appStore); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppStore 用id查询AppStore
// @Tags AppStore
// @Summary 用id查询AppStore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppStore true "用id查询AppStore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appStore/findAppStore [get]
func (appStoreApi *AppStoreApi) FindAppStore(c *gin.Context) {
	var appStore app.AppStore
	err := c.ShouldBindQuery(&appStore)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappStore, err := appStoreService.GetAppStore(appStore.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappStore": reappStore}, c)
	}
}

// GetAppStoreList 分页获取AppStore列表
// @Tags AppStore
// @Summary 分页获取AppStore列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.AppStoreSearch true "分页获取AppStore列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appStore/getAppStoreList [get]
func (appStoreApi *AppStoreApi) GetAppStoreList(c *gin.Context) {
	var pageInfo appReq.AppStoreSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appStoreService.GetAppStoreInfoList(pageInfo); err != nil {
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

//门店设置
func (appStoreApi *AppStoreApi) Edit(c *gin.Context) {
	var appStore app.AppStore
	err := c.ShouldBindJSON(&appStore)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := c.MustGet("id").(uint)
	appStore.Uid = uid
	if err := appStoreService.Edit(appStore); err != nil {
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithMessage("成功", c)
	}
}

//门店详情
func (appStoreApi *AppStoreApi) Detail(c *gin.Context) {
	uid := c.MustGet("id").(uint)
	if err, info := appStoreService.Detail(uid); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(gin.H{"info": info}, c)
	}
}
