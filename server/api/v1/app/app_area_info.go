package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/app"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type AppAreaInfoApi struct {
}

var appAreaInfoService = service.ServiceGroupApp.AppServiceGroup.AppAreaInfoService


// CreateAppAreaInfo 创建AppAreaInfo
// @Tags AppAreaInfo
// @Summary 创建AppAreaInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppAreaInfo true "创建AppAreaInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appAreaInfo/createAppAreaInfo [post]
func (appAreaInfoApi *AppAreaInfoApi) CreateAppAreaInfo(c *gin.Context) {
	var appAreaInfo app.AppAreaInfo
	err := c.ShouldBindJSON(&appAreaInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appAreaInfoService.CreateAppAreaInfo(appAreaInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppAreaInfo 删除AppAreaInfo
// @Tags AppAreaInfo
// @Summary 删除AppAreaInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppAreaInfo true "删除AppAreaInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appAreaInfo/deleteAppAreaInfo [delete]
func (appAreaInfoApi *AppAreaInfoApi) DeleteAppAreaInfo(c *gin.Context) {
	var appAreaInfo app.AppAreaInfo
	err := c.ShouldBindJSON(&appAreaInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appAreaInfoService.DeleteAppAreaInfo(appAreaInfo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppAreaInfoByIds 批量删除AppAreaInfo
// @Tags AppAreaInfo
// @Summary 批量删除AppAreaInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppAreaInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appAreaInfo/deleteAppAreaInfoByIds [delete]
func (appAreaInfoApi *AppAreaInfoApi) DeleteAppAreaInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appAreaInfoService.DeleteAppAreaInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppAreaInfo 更新AppAreaInfo
// @Tags AppAreaInfo
// @Summary 更新AppAreaInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppAreaInfo true "更新AppAreaInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appAreaInfo/updateAppAreaInfo [put]
func (appAreaInfoApi *AppAreaInfoApi) UpdateAppAreaInfo(c *gin.Context) {
	var appAreaInfo app.AppAreaInfo
	err := c.ShouldBindJSON(&appAreaInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appAreaInfoService.UpdateAppAreaInfo(appAreaInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppAreaInfo 用id查询AppAreaInfo
// @Tags AppAreaInfo
// @Summary 用id查询AppAreaInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppAreaInfo true "用id查询AppAreaInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appAreaInfo/findAppAreaInfo [get]
func (appAreaInfoApi *AppAreaInfoApi) FindAppAreaInfo(c *gin.Context) {
	var appAreaInfo app.AppAreaInfo
	err := c.ShouldBindQuery(&appAreaInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappAreaInfo, err := appAreaInfoService.GetAppAreaInfo(appAreaInfo.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappAreaInfo": reappAreaInfo}, c)
	}
}

// GetAppAreaInfoList 分页获取AppAreaInfo列表
// @Tags AppAreaInfo
// @Summary 分页获取AppAreaInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.AppAreaInfoSearch true "分页获取AppAreaInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appAreaInfo/getAppAreaInfoList [get]
func (appAreaInfoApi *AppAreaInfoApi) GetAppAreaInfoList(c *gin.Context) {
	var pageInfo appReq.AppAreaInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appAreaInfoService.GetAppAreaInfoInfoList(pageInfo); err != nil {
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
