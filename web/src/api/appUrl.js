import service from '@/utils/request'

// @Tags AppUrl
// @Summary 创建AppUrl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppUrl true "创建AppUrl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appUrl/createAppUrl [post]
export const createAppUrl = (data) => {
  return service({
    url: '/appUrl/createAppUrl',
    method: 'post',
    data
  })
}

// @Tags AppUrl
// @Summary 删除AppUrl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppUrl true "删除AppUrl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appUrl/deleteAppUrl [delete]
export const deleteAppUrl = (data) => {
  return service({
    url: '/appUrl/deleteAppUrl',
    method: 'delete',
    data
  })
}

// @Tags AppUrl
// @Summary 删除AppUrl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppUrl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appUrl/deleteAppUrl [delete]
export const deleteAppUrlByIds = (data) => {
  return service({
    url: '/appUrl/deleteAppUrlByIds',
    method: 'delete',
    data
  })
}

// @Tags AppUrl
// @Summary 更新AppUrl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppUrl true "更新AppUrl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appUrl/updateAppUrl [put]
export const updateAppUrl = (data) => {
  return service({
    url: '/appUrl/updateAppUrl',
    method: 'put',
    data
  })
}

// @Tags AppUrl
// @Summary 用id查询AppUrl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppUrl true "用id查询AppUrl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appUrl/findAppUrl [get]
export const findAppUrl = (params) => {
  return service({
    url: '/appUrl/findAppUrl',
    method: 'get',
    params
  })
}

// @Tags AppUrl
// @Summary 分页获取AppUrl列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppUrl列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appUrl/getAppUrlList [get]
export const getAppUrlList = (params) => {
  return service({
    url: '/appUrl/getAppUrlList',
    method: 'get',
    params
  })
}
