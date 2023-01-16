package mini

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mini"
	miniReq "github.com/flipped-aurora/gin-vue-admin/server/model/mini/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MiniClassifyApi struct {
}

var miniClassifyService = service.ServiceGroupApp.MiniServiceGroup.MiniClassifyService

// CreateMiniClassify 创建MiniClassify
// @Tags MiniClassify
// @Summary 创建MiniClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mini.MiniClassify true "创建MiniClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /miniClassify/createMiniClassify [post]
func (miniClassifyApi *MiniClassifyApi) CreateMiniClassify(c *gin.Context) {
	var miniClassify mini.MiniClassify
	err := c.ShouldBindJSON(&miniClassify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Title": {utils.NotEmpty()},
	}
	if err := utils.Verify(miniClassify, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := miniClassifyService.CreateMiniClassify(miniClassify); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMiniClassify 删除MiniClassify
// @Tags MiniClassify
// @Summary 删除MiniClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mini.MiniClassify true "删除MiniClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /miniClassify/deleteMiniClassify [delete]
func (miniClassifyApi *MiniClassifyApi) DeleteMiniClassify(c *gin.Context) {
	var miniClassify mini.MiniClassify
	err := c.ShouldBindJSON(&miniClassify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := miniClassifyService.DeleteMiniClassify(miniClassify); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMiniClassifyByIds 批量删除MiniClassify
// @Tags MiniClassify
// @Summary 批量删除MiniClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MiniClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /miniClassify/deleteMiniClassifyByIds [delete]
func (miniClassifyApi *MiniClassifyApi) DeleteMiniClassifyByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := miniClassifyService.DeleteMiniClassifyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMiniClassify 更新MiniClassify
// @Tags MiniClassify
// @Summary 更新MiniClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mini.MiniClassify true "更新MiniClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /miniClassify/updateMiniClassify [put]
func (miniClassifyApi *MiniClassifyApi) UpdateMiniClassify(c *gin.Context) {
	var miniClassify mini.MiniClassify
	err := c.ShouldBindJSON(&miniClassify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Title": {utils.NotEmpty()},
	}
	if err := utils.Verify(miniClassify, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := miniClassifyService.UpdateMiniClassify(miniClassify); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMiniClassify 用id查询MiniClassify
// @Tags MiniClassify
// @Summary 用id查询MiniClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mini.MiniClassify true "用id查询MiniClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /miniClassify/findMiniClassify [get]
func (miniClassifyApi *MiniClassifyApi) FindMiniClassify(c *gin.Context) {
	var miniClassify mini.MiniClassify
	err := c.ShouldBindQuery(&miniClassify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reminiClassify, err := miniClassifyService.GetMiniClassify(miniClassify.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reminiClassify": reminiClassify}, c)
	}
}

// GetMiniClassifyList 分页获取MiniClassify列表
// @Tags MiniClassify
// @Summary 分页获取MiniClassify列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query miniReq.MiniClassifySearch true "分页获取MiniClassify列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /miniClassify/getMiniClassifyList [get]
func (miniClassifyApi *MiniClassifyApi) GetMiniClassifyList(c *gin.Context) {
	var pageInfo miniReq.MiniClassifySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := miniClassifyService.GetMiniClassifyInfoList(pageInfo); err != nil {
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

//小程序首页分类
func (miniClassifyApi *MiniClassifyApi) List(c *gin.Context) {
	var p miniReq.Classify
	err := c.ShouldBindQuery(&p)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if p.Pid == "" {
		p.Pid = "0"
	}

	if list, err := miniClassifyService.List(p.Pid); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"list": list}, "获取成功", c)
	}
}
