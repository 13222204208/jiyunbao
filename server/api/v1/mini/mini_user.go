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

type MiniUserApi struct {
}

var miniUserService = service.ServiceGroupApp.MiniServiceGroup.MiniUserService

// CreateMiniUser 创建MiniUser
// @Tags MiniUser
// @Summary 创建MiniUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mini.MiniUser true "创建MiniUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /miniUser/createMiniUser [post]
func (miniUserApi *MiniUserApi) CreateMiniUser(c *gin.Context) {
	var miniUser mini.MiniUser
	err := c.ShouldBindJSON(&miniUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := miniUserService.CreateMiniUser(miniUser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMiniUser 删除MiniUser
// @Tags MiniUser
// @Summary 删除MiniUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mini.MiniUser true "删除MiniUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /miniUser/deleteMiniUser [delete]
func (miniUserApi *MiniUserApi) DeleteMiniUser(c *gin.Context) {
	var miniUser mini.MiniUser
	err := c.ShouldBindJSON(&miniUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := miniUserService.DeleteMiniUser(miniUser); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMiniUserByIds 批量删除MiniUser
// @Tags MiniUser
// @Summary 批量删除MiniUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MiniUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /miniUser/deleteMiniUserByIds [delete]
func (miniUserApi *MiniUserApi) DeleteMiniUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := miniUserService.DeleteMiniUserByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMiniUser 更新MiniUser
// @Tags MiniUser
// @Summary 更新MiniUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mini.MiniUser true "更新MiniUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /miniUser/updateMiniUser [put]
func (miniUserApi *MiniUserApi) UpdateMiniUser(c *gin.Context) {
	var miniUser mini.MiniUser
	err := c.ShouldBindJSON(&miniUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := miniUserService.UpdateMiniUser(miniUser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMiniUser 用id查询MiniUser
// @Tags MiniUser
// @Summary 用id查询MiniUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mini.MiniUser true "用id查询MiniUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /miniUser/findMiniUser [get]
func (miniUserApi *MiniUserApi) FindMiniUser(c *gin.Context) {
	var miniUser mini.MiniUser
	err := c.ShouldBindQuery(&miniUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reminiUser, err := miniUserService.GetMiniUser(miniUser.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reminiUser": reminiUser}, c)
	}
}

// GetMiniUserList 分页获取MiniUser列表
// @Tags MiniUser
// @Summary 分页获取MiniUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query miniReq.MiniUserSearch true "分页获取MiniUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /miniUser/getMiniUserList [get]
func (miniUserApi *MiniUserApi) GetMiniUserList(c *gin.Context) {
	var pageInfo miniReq.MiniUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := miniUserService.GetMiniUserInfoList(pageInfo); err != nil {
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

//小程序用户登陆
func (miniUserApi *MiniUserApi) Login(c *gin.Context) {
	var m miniReq.MiniUserLogin
	err := c.ShouldBindJSON(&m)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := utils.Verify(m, utils.MiniUserLoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//如果是微信小程序登陆
	//if m.MiniType == 1 {
	//
	//}
}
