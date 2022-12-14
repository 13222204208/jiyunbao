import service from '@/utils/request'

// @Tags AppPay
// @Summary 创建AppPay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppPay true "创建AppPay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appPay/createAppPay [post]
export const createAppPay = (data) => {
  return service({
    url: '/appPay/createAppPay',
    method: 'post',
    data
  })
}

// @Tags AppPay
// @Summary 删除AppPay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppPay true "删除AppPay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appPay/deleteAppPay [delete]
export const deleteAppPay = (data) => {
  return service({
    url: '/appPay/deleteAppPay',
    method: 'delete',
    data
  })
}

// @Tags AppPay
// @Summary 删除AppPay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppPay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appPay/deleteAppPay [delete]
export const deleteAppPayByIds = (data) => {
  return service({
    url: '/appPay/deleteAppPayByIds',
    method: 'delete',
    data
  })
}

// @Tags AppPay
// @Summary 更新AppPay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppPay true "更新AppPay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appPay/updateAppPay [put]
export const updateAppPay = (data) => {
  return service({
    url: '/appPay/updateAppPay',
    method: 'put',
    data
  })
}

// @Tags AppPay
// @Summary 用id查询AppPay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppPay true "用id查询AppPay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appPay/findAppPay [get]
export const findAppPay = (params) => {
  return service({
    url: '/appPay/findAppPay',
    method: 'get',
    params
  })
}

// @Tags AppPay
// @Summary 分页获取AppPay列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppPay列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appPay/getAppPayList [get]
export const getAppPayList = (params) => {
  return service({
    url: '/appPay/getAppPayList',
    method: 'get',
    params
  })
}
