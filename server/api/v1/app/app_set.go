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

type AppSetApi struct {
}

var appSetService = service.ServiceGroupApp.AppServiceGroup.AppSetService

// CreateAppSet 创建AppSet
// @Tags AppSet
// @Summary 创建AppSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppSet true "创建AppSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appSet/createAppSet [post]
func (appSetApi *AppSetApi) CreateAppSet(c *gin.Context) {
	var appSet app.AppSet
	err := c.ShouldBindJSON(&appSet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appSetService.CreateAppSet(appSet); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppSet 删除AppSet
// @Tags AppSet
// @Summary 删除AppSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppSet true "删除AppSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appSet/deleteAppSet [delete]
func (appSetApi *AppSetApi) DeleteAppSet(c *gin.Context) {
	var appSet app.AppSet
	err := c.ShouldBindJSON(&appSet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appSetService.DeleteAppSet(appSet); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppSetByIds 批量删除AppSet
// @Tags AppSet
// @Summary 批量删除AppSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appSet/deleteAppSetByIds [delete]
func (appSetApi *AppSetApi) DeleteAppSetByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appSetService.DeleteAppSetByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppSet 更新AppSet
// @Tags AppSet
// @Summary 更新AppSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppSet true "更新AppSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appSet/updateAppSet [put]
func (appSetApi *AppSetApi) UpdateAppSet(c *gin.Context) {
	var appSet app.AppSet
	err := c.ShouldBindJSON(&appSet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appSetService.UpdateAppSet(appSet); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppSet 用id查询AppSet
// @Tags AppSet
// @Summary 用id查询AppSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppSet true "用id查询AppSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appSet/findAppSet [get]
func (appSetApi *AppSetApi) FindAppSet(c *gin.Context) {
	var appSet app.AppSet
	err := c.ShouldBindQuery(&appSet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappSet, err := appSetService.GetAppSet(appSet.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappSet": reappSet}, c)
	}
}

// GetAppSetList 分页获取AppSet列表
// @Tags AppSet
// @Summary 分页获取AppSet列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.AppSetSearch true "分页获取AppSet列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appSet/getAppSetList [get]
func (appSetApi *AppSetApi) GetAppSetList(c *gin.Context) {
	var pageInfo appReq.AppSetSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appSetService.GetAppSetInfoList(pageInfo); err != nil {
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

//url生成二维码
func (appSetApi *AppSetApi) Qrcode(c *gin.Context) {

	uid := c.MustGet("id").(uint)

	if err, imgUrl := appSetService.Qrcode(uid); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(gin.H{"info": imgUrl}, c)
	}
}
