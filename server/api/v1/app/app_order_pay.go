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

type AppOrderPayApi struct {
}

var appOrderPayService = service.ServiceGroupApp.AppServiceGroup.AppOrderPayService

// CreateAppOrderPay 创建AppOrderPay
// @Tags AppOrderPay
// @Summary 创建AppOrderPay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppOrderPay true "创建AppOrderPay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appOrderPay/createAppOrderPay [post]
func (appOrderPayApi *AppOrderPayApi) CreateAppOrderPay(c *gin.Context) {
	var appOrderPay app.AppOrderPay
	err := c.ShouldBindJSON(&appOrderPay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"PayType": {utils.NotEmpty()},
	}
	if err := utils.Verify(appOrderPay, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appOrderPayService.CreateAppOrderPay(appOrderPay); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppOrderPay 删除AppOrderPay
// @Tags AppOrderPay
// @Summary 删除AppOrderPay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppOrderPay true "删除AppOrderPay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appOrderPay/deleteAppOrderPay [delete]
func (appOrderPayApi *AppOrderPayApi) DeleteAppOrderPay(c *gin.Context) {
	var appOrderPay app.AppOrderPay
	err := c.ShouldBindJSON(&appOrderPay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appOrderPayService.DeleteAppOrderPay(appOrderPay); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppOrderPayByIds 批量删除AppOrderPay
// @Tags AppOrderPay
// @Summary 批量删除AppOrderPay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppOrderPay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appOrderPay/deleteAppOrderPayByIds [delete]
func (appOrderPayApi *AppOrderPayApi) DeleteAppOrderPayByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appOrderPayService.DeleteAppOrderPayByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppOrderPay 更新AppOrderPay
// @Tags AppOrderPay
// @Summary 更新AppOrderPay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppOrderPay true "更新AppOrderPay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appOrderPay/updateAppOrderPay [put]
func (appOrderPayApi *AppOrderPayApi) UpdateAppOrderPay(c *gin.Context) {
	var appOrderPay app.AppOrderPay
	err := c.ShouldBindJSON(&appOrderPay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"PayType": {utils.NotEmpty()},
	}
	if err := utils.Verify(appOrderPay, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appOrderPayService.UpdateAppOrderPay(appOrderPay); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppOrderPay 用id查询AppOrderPay
// @Tags AppOrderPay
// @Summary 用id查询AppOrderPay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppOrderPay true "用id查询AppOrderPay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appOrderPay/findAppOrderPay [get]
func (appOrderPayApi *AppOrderPayApi) FindAppOrderPay(c *gin.Context) {
	var appOrderPay app.AppOrderPay
	err := c.ShouldBindQuery(&appOrderPay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappOrderPay, err := appOrderPayService.GetAppOrderPay(appOrderPay.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappOrderPay": reappOrderPay}, c)
	}
}

// GetAppOrderPayList 分页获取AppOrderPay列表
// @Tags AppOrderPay
// @Summary 分页获取AppOrderPay列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.AppOrderPaySearch true "分页获取AppOrderPay列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appOrderPay/getAppOrderPayList [get]
func (appOrderPayApi *AppOrderPayApi) GetAppOrderPayList(c *gin.Context) {
	var pageInfo appReq.AppOrderPaySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appOrderPayService.GetAppOrderPayInfoList(pageInfo); err != nil {
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

//聚合支付
func (appOrderPayApi *AppOrderPayApi) Pay(c *gin.Context) {
	var p appReq.OrderPay
	err := c.ShouldBindJSON(&p)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"PayType":     {utils.NotEmpty()},
		"Code":        {utils.NotEmpty()},
		"TotalAmount": {utils.NotEmpty()},
		"MercId":      {utils.NotEmpty()},
	}
	if err := utils.Verify(p, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err, info := appOrderPayService.Pay(p); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(gin.H{"info": info}, c)
	}
}

func (appOrderPayApi *AppOrderPayApi) Notify(c *gin.Context) {
	var r appReq.OrderPayNotify
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("回调的数据", r)

	_ = c.ShouldBindQuery(&r)
	fmt.Println("回调的表单", r)
	if r.TradeStatus == "TRADE_SUCCESS" {
		if err := appOrderPayService.Notify(r.OutTradeNo); err != nil {
			global.GVA_LOG.Error("修改失败!", zap.Error(err))
			//response.FailWithMessage(err.Error(), c)
		}
	}
}
