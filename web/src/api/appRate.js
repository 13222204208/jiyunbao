import service from '@/utils/request'

// @Tags AppRate
// @Summary 创建AppRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppRate true "创建AppRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appRate/createAppRate [post]
export const createAppRate = (data) => {
  return service({
    url: '/appRate/createAppRate',
    method: 'post',
    data
  })
}

// @Tags AppRate
// @Summary 删除AppRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppRate true "删除AppRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appRate/deleteAppRate [delete]
export const deleteAppRate = (data) => {
  return service({
    url: '/appRate/deleteAppRate',
    method: 'delete',
    data
  })
}

// @Tags AppRate
// @Summary 删除AppRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appRate/deleteAppRate [delete]
export const deleteAppRateByIds = (data) => {
  return service({
    url: '/appRate/deleteAppRateByIds',
    method: 'delete',
    data
  })
}

// @Tags AppRate
// @Summary 更新AppRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppRate true "更新AppRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appRate/updateAppRate [put]
export const updateAppRate = (data) => {
  return service({
    url: '/appRate/updateAppRate',
    method: 'put',
    data
  })
}

// @Tags AppRate
// @Summary 用id查询AppRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppRate true "用id查询AppRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appRate/findAppRate [get]
export const findAppRate = (params) => {
  return service({
    url: '/appRate/findAppRate',
    method: 'get',
    params
  })
}

// @Tags AppRate
// @Summary 分页获取AppRate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppRate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appRate/getAppRateList [get]
export const getAppRateList = (params) => {
  return service({
    url: '/appRate/getAppRateList',
    method: 'get',
    params
  })
}
