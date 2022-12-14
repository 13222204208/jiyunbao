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

type AppSignContractApi struct {
}

var appSignContractService = service.ServiceGroupApp.AppServiceGroup.AppSignContractService

// CreateAppSignContract 创建AppSignContract
// @Tags AppSignContract
// @Summary 创建AppSignContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppSignContract true "创建AppSignContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appSignContract/createAppSignContract [post]
func (appSignContractApi *AppSignContractApi) CreateAppSignContract(c *gin.Context) {
	var appSignContract app.AppSignContract
	err := c.ShouldBindJSON(&appSignContract)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"ContractType": {utils.NotEmpty()},
		"IsSendConMsg": {utils.NotEmpty()},
		"CustId":       {utils.NotEmpty()},
	}
	if err := utils.Verify(appSignContract, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appSignContractService.CreateAppSignContract(appSignContract); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppSignContract 删除AppSignContract
// @Tags AppSignContract
// @Summary 删除AppSignContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppSignContract true "删除AppSignContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appSignContract/deleteAppSignContract [delete]
func (appSignContractApi *AppSignContractApi) DeleteAppSignContract(c *gin.Context) {
	var appSignContract app.AppSignContract
	err := c.ShouldBindJSON(&appSignContract)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appSignContractService.DeleteAppSignContract(appSignContract); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppSignContractByIds 批量删除AppSignContract
// @Tags AppSignContract
// @Summary 批量删除AppSignContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppSignContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appSignContract/deleteAppSignContractByIds [delete]
func (appSignContractApi *AppSignContractApi) DeleteAppSignContractByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appSignContractService.DeleteAppSignContractByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppSignContract 更新AppSignContract
// @Tags AppSignContract
// @Summary 更新AppSignContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppSignContract true "更新AppSignContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appSignContract/updateAppSignContract [put]
func (appSignContractApi *AppSignContractApi) UpdateAppSignContract(c *gin.Context) {
	var appSignContract app.AppSignContract
	err := c.ShouldBindJSON(&appSignContract)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"ContractType": {utils.NotEmpty()},
		"IsSendConMsg": {utils.NotEmpty()},
		"CustId":       {utils.NotEmpty()},
	}
	if err := utils.Verify(appSignContract, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appSignContractService.UpdateAppSignContract(appSignContract); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("申请成功", c)
	}
}

// FindAppSignContract 用id查询AppSignContract
// @Tags AppSignContract
// @Summary 用id查询AppSignContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppSignContract true "用id查询AppSignContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appSignContract/findAppSignContract [get]
func (appSignContractApi *AppSignContractApi) FindAppSignContract(c *gin.Context) {
	var appSignContract app.AppSignContract
	err := c.ShouldBindQuery(&appSignContract)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappSignContract, err := appSignContractService.GetAppSignContract(appSignContract.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappSignContract": reappSignContract}, c)
	}
}

// GetAppSignContractList 分页获取AppSignContract列表
// @Tags AppSignContract
// @Summary 分页获取AppSignContract列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.AppSignContractSearch true "分页获取AppSignContract列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appSignContract/getAppSignContractList [get]
func (appSignContractApi *AppSignContractApi) GetAppSignContractList(c *gin.Context) {
	var pageInfo appReq.AppSignContractSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appSignContractService.GetAppSignContractInfoList(pageInfo); err != nil {
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

//签约状态查询
func (appSignContractApi *AppSignContractApi) QueryAuthInfo(c *gin.Context) {
	var appSignContract app.AppSignContract
	err := c.ShouldBindJSON(&appSignContract)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"AuthId": {utils.NotEmpty()},
	}
	if err := utils.Verify(appSignContract, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, res := appSignContractService.QueryAuthInfo(appSignContract); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage(res, c)
	}
}

func (appSignContractApi *AppSignContractApi) BalanceQuery(c *gin.Context) {
	var s app.AppSignContract
	err := c.ShouldBindQuery(&s)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if s.MercId == "" {
		response.FailWithMessage("商户号不能为空", c)
		return
	}
	if err, res := appSignContractService.BalanceQuery(s.MercId); err != nil {
		global.GVA_LOG.Error("查询余额失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage(res, c)
	}
}
