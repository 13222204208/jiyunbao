import service from '@/utils/request'

// @Tags AppOrderPay
// @Summary 创建AppOrderPay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppOrderPay true "创建AppOrderPay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appOrderPay/createAppOrderPay [post]
export const createAppOrderPay = (data) => {
  return service({
    url: '/appOrderPay/createAppOrderPay',
    method: 'post',
    data
  })
}

// @Tags AppOrderPay
// @Summary 删除AppOrderPay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppOrderPay true "删除AppOrderPay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appOrderPay/deleteAppOrderPay [delete]
export const deleteAppOrderPay = (data) => {
  return service({
    url: '/appOrderPay/deleteAppOrderPay',
    method: 'delete',
    data
  })
}

// @Tags AppOrderPay
// @Summary 删除AppOrderPay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppOrderPay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appOrderPay/deleteAppOrderPay [delete]
export const deleteAppOrderPayByIds = (data) => {
  return service({
    url: '/appOrderPay/deleteAppOrderPayByIds',
    method: 'delete',
    data
  })
}

// @Tags AppOrderPay
// @Summary 更新AppOrderPay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppOrderPay true "更新AppOrderPay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appOrderPay/updateAppOrderPay [put]
export const updateAppOrderPay = (data) => {
  return service({
    url: '/appOrderPay/updateAppOrderPay',
    method: 'put',
    data
  })
}

// @Tags AppOrderPay
// @Summary 用id查询AppOrderPay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppOrderPay true "用id查询AppOrderPay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appOrderPay/findAppOrderPay [get]
export const findAppOrderPay = (params) => {
  return service({
    url: '/appOrderPay/findAppOrderPay',
    method: 'get',
    params
  })
}

// @Tags AppOrderPay
// @Summary 分页获取AppOrderPay列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppOrderPay列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appOrderPay/getAppOrderPayList [get]
export const getAppOrderPayList = (params) => {
  return service({
    url: '/appOrderPay/getAppOrderPayList',
    method: 'get',
    params
  })
}
