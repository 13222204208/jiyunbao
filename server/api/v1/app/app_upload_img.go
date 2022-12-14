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

type AppUploadImgApi struct {
}

var appUploadImgService = service.ServiceGroupApp.AppServiceGroup.AppUploadImgService

// CreateAppUploadImg 创建AppUploadImg
// @Tags AppUploadImg
// @Summary 创建AppUploadImg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppUploadImg true "创建AppUploadImg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appUploadImg/createAppUploadImg [post]
func (appUploadImgApi *AppUploadImgApi) CreateAppUploadImg(c *gin.Context) {
	var appUploadImg app.AppUploadImg
	err := c.ShouldBindJSON(&appUploadImg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"PicType": {utils.NotEmpty()},
		"ImgUrl":  {utils.NotEmpty()},
	}
	if err := utils.Verify(appUploadImg, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appUploadImgService.CreateAppUploadImg(appUploadImg); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppUploadImg 删除AppUploadImg
// @Tags AppUploadImg
// @Summary 删除AppUploadImg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppUploadImg true "删除AppUploadImg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appUploadImg/deleteAppUploadImg [delete]
func (appUploadImgApi *AppUploadImgApi) DeleteAppUploadImg(c *gin.Context) {
	var appUploadImg app.AppUploadImg
	err := c.ShouldBindJSON(&appUploadImg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appUploadImgService.DeleteAppUploadImg(appUploadImg); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppUploadImgByIds 批量删除AppUploadImg
// @Tags AppUploadImg
// @Summary 批量删除AppUploadImg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppUploadImg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appUploadImg/deleteAppUploadImgByIds [delete]
func (appUploadImgApi *AppUploadImgApi) DeleteAppUploadImgByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appUploadImgService.DeleteAppUploadImgByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppUploadImg 更新AppUploadImg
// @Tags AppUploadImg
// @Summary 更新AppUploadImg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppUploadImg true "更新AppUploadImg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appUploadImg/updateAppUploadImg [put]
func (appUploadImgApi *AppUploadImgApi) UpdateAppUploadImg(c *gin.Context) {
	var appUploadImg app.AppUploadImg
	err := c.ShouldBindJSON(&appUploadImg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"PicType": {utils.NotEmpty()},
		"ImgUrl":  {utils.NotEmpty()},
	}
	if err := utils.Verify(appUploadImg, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appUploadImgService.UpdateAppUploadImg(appUploadImg); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppUploadImg 用id查询AppUploadImg
// @Tags AppUploadImg
// @Summary 用id查询AppUploadImg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppUploadImg true "用id查询AppUploadImg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appUploadImg/findAppUploadImg [get]
func (appUploadImgApi *AppUploadImgApi) FindAppUploadImg(c *gin.Context) {
	var appUploadImg app.AppUploadImg
	err := c.ShouldBindQuery(&appUploadImg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappUploadImg, err := appUploadImgService.GetAppUploadImg(appUploadImg.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappUploadImg": reappUploadImg}, c)
	}
}

// GetAppUploadImgList 分页获取AppUploadImg列表
// @Tags AppUploadImg
// @Summary 分页获取AppUploadImg列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.AppUploadImgSearch true "分页获取AppUploadImg列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appUploadImg/getAppUploadImgList [get]
func (appUploadImgApi *AppUploadImgApi) GetAppUploadImgList(c *gin.Context) {
	var pageInfo appReq.AppUploadImgSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appUploadImgService.GetAppUploadImgInfoList(pageInfo); err != nil {
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
