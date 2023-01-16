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

type AppCustInfoApi struct {
}

var appCustInfoService = service.ServiceGroupApp.AppServiceGroup.AppCustInfoService

// CreateAppCustInfo 创建AppCustInfo
// @Tags AppCustInfo
// @Summary 创建AppCustInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppCustInfo true "创建AppCustInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appCustInfo/createAppCustInfo [post]
func (appCustInfoApi *AppCustInfoApi) CreateAppCustInfo(c *gin.Context) {
	var appCustInfo app.AppCustInfo
	err := c.ShouldBindJSON(&appCustInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(appCustInfo, utils.CustInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if len(appCustInfo.CrpCertNo) != 18 {
		response.FailWithMessage("法人证件号不正确", c)
	}
	if err := appCustInfoService.CreateAppCustInfo(appCustInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

//商户端提交的支付认证
func (appCustInfoApi *AppCustInfoApi) Attestation(c *gin.Context) {
	var appCustInfo app.AppCustInfo
	err := c.ShouldBindJSON(&appCustInfo)
	fmt.Println("接收的数据", appCustInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(appCustInfo, utils.AppCustInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if len(appCustInfo.CrpCertNo) != 18 {
		response.FailWithMessage("法人证件号不正确", c)
	}
	appCustInfo.Uid = c.MustGet("id").(uint)
	if err := appCustInfoService.Attestation(appCustInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppCustInfo 删除AppCustInfo
// @Tags AppCustInfo
// @Summary 删除AppCustInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppCustInfo true "删除AppCustInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appCustInfo/deleteAppCustInfo [delete]
func (appCustInfoApi *AppCustInfoApi) DeleteAppCustInfo(c *gin.Context) {
	var appCustInfo app.AppCustInfo
	err := c.ShouldBindJSON(&appCustInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appCustInfoService.DeleteAppCustInfo(appCustInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppCustInfoByIds 批量删除AppCustInfo
// @Tags AppCustInfo
// @Summary 批量删除AppCustInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppCustInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appCustInfo/deleteAppCustInfoByIds [delete]
func (appCustInfoApi *AppCustInfoApi) DeleteAppCustInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appCustInfoService.DeleteAppCustInfoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppCustInfo 更新AppCustInfo
// @Tags AppCustInfo
// @Summary 更新AppCustInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppCustInfo true "更新AppCustInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appCustInfo/updateAppCustInfo [put]
func (appCustInfoApi *AppCustInfoApi) UpdateAppCustInfo(c *gin.Context) {
	var appCustInfo app.AppCustInfo
	err := c.ShouldBindJSON(&appCustInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(appCustInfo, utils.CustInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if len(appCustInfo.CrpCertNo) != 18 {
		response.FailWithMessage("法人证件号不正确", c)
	}

	if err := appCustInfoService.UpdateAppCustInfo(appCustInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("修改资料成功", c)

		//if err := appCustInfoService.ModifyCustInfoApply(appCustInfo.ID); err != nil {
		//	response.FailWithMessage(err.Error(), c)
		//} else {
		//	response.OkWithMessage("修改资料成功", c)
		//}

	}
}

// FindAppCustInfo 用id查询AppCustInfo
// @Tags AppCustInfo
// @Summary 用id查询AppCustInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppCustInfo true "用id查询AppCustInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appCustInfo/findAppCustInfo [get]
func (appCustInfoApi *AppCustInfoApi) FindAppCustInfo(c *gin.Context) {
	var appCustInfo app.AppCustInfo
	err := c.ShouldBindQuery(&appCustInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappCustInfo, err := appCustInfoService.GetAppCustInfo(appCustInfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappCustInfo": reappCustInfo}, c)
	}
}

// GetAppCustInfoList 分页获取AppCustInfo列表
// @Tags AppCustInfo
// @Summary 分页获取AppCustInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.AppCustInfoSearch true "分页获取AppCustInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appCustInfo/getAppCustInfoList [get]
func (appCustInfoApi *AppCustInfoApi) GetAppCustInfoList(c *gin.Context) {
	var pageInfo appReq.AppCustInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appCustInfoService.GetAppCustInfoInfoList(pageInfo); err != nil {
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

//资料上送
func (appCustInfoApi *AppCustInfoApi) AddCustInfoApply(c *gin.Context) {
	var appCustInfo app.AppCustInfo
	err := c.ShouldBindJSON(&appCustInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	if err := utils.Verify(appCustInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := appCustInfoService.AddCustInfoApply(appCustInfo.ID); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("上送成功", c)
	}
}

//资料确认
func (appCustInfoApi *AppCustInfoApi) AuditCustInfoApply(c *gin.Context) {
	var appCustInfo app.AppCustInfo
	err := c.ShouldBindJSON(&appCustInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	if err := utils.Verify(appCustInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := appCustInfoService.AuditCustInfoApply(appCustInfo.ID); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("确认成功", c)
	}
}

//商户状态查询
func (appCustInfoApi *AppCustInfoApi) QueryCustApply(c *gin.Context) {
	var appCustInfo app.AppCustInfo
	err := c.ShouldBindJSON(&appCustInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	if err := utils.Verify(appCustInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := appCustInfoService.QueryCustApply(appCustInfo.ID); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("入网申请状态成功", c)
	}
}

//基本信息变更申请
func (appCustInfoApi *AppCustInfoApi) ChangeMercBaseInfo(c *gin.Context) {
	var appCustInfo app.AppCustInfo
	err := c.ShouldBindJSON(&appCustInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"custId": {utils.NotEmpty()},
	}
	if err := utils.Verify(appCustInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := appCustInfoService.ChangeMercBaseInfo(appCustInfo.CustId); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("提交成功", c)
	}
}
