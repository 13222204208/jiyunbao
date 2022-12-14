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

type AppPayApi struct {
}

var appPayService = service.ServiceGroupApp.AppServiceGroup.AppPayService

// CreateAppPay 创建AppPay
// @Tags AppPay
// @Summary 创建AppPay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppPay true "创建AppPay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appPay/createAppPay [post]
func (appPayApi *AppPayApi) CreateAppPay(c *gin.Context) {
	var appPay app.AppPay
	err := c.ShouldBindJSON(&appPay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"BusinessLicense": {utils.NotEmpty()},
		"OpenCard":        {utils.NotEmpty()},
		"MchName":         {utils.NotEmpty()},
		"MchAddress":      {utils.NotEmpty()},
		"LegalPerson":     {utils.NotEmpty()},
		"Phone":           {utils.NotEmpty()},
		"LegalCard":       {utils.NotEmpty()},
		"CardFront":       {utils.NotEmpty()},
		"CardReverse":     {utils.NotEmpty()},
		"CardValidity":    {utils.NotEmpty()},
		"MchType":         {utils.NotEmpty()},
		"Status":          {utils.NotEmpty()},
		"LinkmanName":     {utils.NotEmpty()},
		"LinkmanPhone":    {utils.NotEmpty()},
		"LinkmanMail":     {utils.NotEmpty()},
		"LinkmanAddress":  {utils.NotEmpty()},
		"BankFrontImg":    {utils.NotEmpty()},
		"CloseType":       {utils.NotEmpty()},
		"CloseCardNum":    {utils.NotEmpty()},
		"BankPhone":       {utils.NotEmpty()},
		"OpenBankNum":     {utils.NotEmpty()},
		"OpenBankName":    {utils.NotEmpty()},
		"OpenBankCity":    {utils.NotEmpty()},
		"StartMoney":      {utils.NotEmpty()},
		"StoreHeadImg":    {utils.NotEmpty()},
	}
	if err := utils.Verify(appPay, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appPayService.CreateAppPay(appPay); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppPay 删除AppPay
// @Tags AppPay
// @Summary 删除AppPay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppPay true "删除AppPay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appPay/deleteAppPay [delete]
func (appPayApi *AppPayApi) DeleteAppPay(c *gin.Context) {
	var appPay app.AppPay
	err := c.ShouldBindJSON(&appPay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appPayService.DeleteAppPay(appPay); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppPayByIds 批量删除AppPay
// @Tags AppPay
// @Summary 批量删除AppPay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppPay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appPay/deleteAppPayByIds [delete]
func (appPayApi *AppPayApi) DeleteAppPayByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appPayService.DeleteAppPayByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppPay 更新AppPay
// @Tags AppPay
// @Summary 更新AppPay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppPay true "更新AppPay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appPay/updateAppPay [put]
func (appPayApi *AppPayApi) UpdateAppPay(c *gin.Context) {
	var appPay app.AppPay
	err := c.ShouldBindJSON(&appPay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"BusinessLicense": {utils.NotEmpty()},
		"OpenCard":        {utils.NotEmpty()},
		"MchName":         {utils.NotEmpty()},
		"MchAddress":      {utils.NotEmpty()},
		"LegalPerson":     {utils.NotEmpty()},
		"Phone":           {utils.NotEmpty()},
		"LegalCard":       {utils.NotEmpty()},
		"CardFront":       {utils.NotEmpty()},
		"CardReverse":     {utils.NotEmpty()},
		"CardValidity":    {utils.NotEmpty()},
		"MchType":         {utils.NotEmpty()},
		"Status":          {utils.NotEmpty()},
		"LinkmanName":     {utils.NotEmpty()},
		"LinkmanPhone":    {utils.NotEmpty()},
		"LinkmanMail":     {utils.NotEmpty()},
		"LinkmanAddress":  {utils.NotEmpty()},
		"BankFrontImg":    {utils.NotEmpty()},
		"CloseType":       {utils.NotEmpty()},
		"CloseCardNum":    {utils.NotEmpty()},
		"BankPhone":       {utils.NotEmpty()},
		"OpenBankNum":     {utils.NotEmpty()},
		"OpenBankName":    {utils.NotEmpty()},
		"OpenBankCity":    {utils.NotEmpty()},
		"StartMoney":      {utils.NotEmpty()},
		"StoreHeadImg":    {utils.NotEmpty()},
	}
	if err := utils.Verify(appPay, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appPayService.UpdateAppPay(appPay); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppPay 用id查询AppPay
// @Tags AppPay
// @Summary 用id查询AppPay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppPay true "用id查询AppPay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appPay/findAppPay [get]
func (appPayApi *AppPayApi) FindAppPay(c *gin.Context) {
	var appPay app.AppPay
	err := c.ShouldBindQuery(&appPay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappPay, err := appPayService.GetAppPay(appPay.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappPay": reappPay}, c)
	}
}

// GetAppPayList 分页获取AppPay列表
// @Tags AppPay
// @Summary 分页获取AppPay列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.AppPaySearch true "分页获取AppPay列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appPay/getAppPayList [get]
func (appPayApi *AppPayApi) GetAppPayList(c *gin.Context) {
	var pageInfo appReq.AppPaySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appPayService.GetAppPayInfoList(pageInfo); err != nil {
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

//支付认证
func (appPayApi *AppPayApi) Attestation(c *gin.Context) {
	var appPay app.AppPay
	err := c.ShouldBindJSON(&appPay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"BusinessLicense": {utils.NotEmpty()},
		"OpenCard":        {utils.NotEmpty()},
		"MchName":         {utils.NotEmpty()},
		"MchAddress":      {utils.NotEmpty()},
		"LegalPerson":     {utils.NotEmpty()},
		"Phone":           {utils.NotEmpty()},
		"LegalCard":       {utils.NotEmpty()},
		"CardFront":       {utils.NotEmpty()},
		"CardReverse":     {utils.NotEmpty()},
		"CardValidity":    {utils.NotEmpty()},
		"MchType":         {utils.NotEmpty()},
		"LinkmanName":     {utils.NotEmpty()},
		"LinkmanPhone":    {utils.NotEmpty()},
		"LinkmanMail":     {utils.NotEmpty()},
		"LinkmanAddress":  {utils.NotEmpty()},
		"BankFrontImg":    {utils.NotEmpty()},
		"CloseType":       {utils.NotEmpty()},
		"CloseCardNum":    {utils.NotEmpty()},
		"BankPhone":       {utils.NotEmpty()},
		"OpenBankNum":     {utils.NotEmpty()},
		"OpenBankName":    {utils.NotEmpty()},
		"OpenBankCity":    {utils.NotEmpty()},
		"StartMoney":      {utils.NotEmpty()},
		"StoreHeadImg":    {utils.NotEmpty()},
	}
	if err := utils.Verify(appPay, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := c.MustGet("id").(uint)
	appPay.Uid = uid

	if err := appPayService.Attestation(appPay); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

//修改商户认证信息
func (appPayApi *AppPayApi) Edit(c *gin.Context) {
	var appPay app.AppPay
	err := c.ShouldBindJSON(&appPay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := c.MustGet("id").(uint)
	if err := appPayService.Edit(uid, appPay); err != nil {
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}
