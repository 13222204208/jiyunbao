import service from '@/utils/request'

// @Tags AppService
// @Summary 创建AppService
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppService true "创建AppService"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appService/createAppService [post]
export const createAppService = (data) => {
  return service({
    url: '/appService/createAppService',
    method: 'post',
    data
  })
}

// @Tags AppService
// @Summary 删除AppService
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppService true "删除AppService"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appService/deleteAppService [delete]
export const deleteAppService = (data) => {
  return service({
    url: '/appService/deleteAppService',
    method: 'delete',
    data
  })
}

// @Tags AppService
// @Summary 删除AppService
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppService"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appService/deleteAppService [delete]
export const deleteAppServiceByIds = (data) => {
  return service({
    url: '/appService/deleteAppServiceByIds',
    method: 'delete',
    data
  })
}

// @Tags AppService
// @Summary 更新AppService
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppService true "更新AppService"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appService/updateAppService [put]
export const updateAppService = (data) => {
  return service({
    url: '/appService/updateAppService',
    method: 'put',
    data
  })
}

// @Tags AppService
// @Summary 用id查询AppService
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppService true "用id查询AppService"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appService/findAppService [get]
export const findAppService = (params) => {
  return service({
    url: '/appService/findAppService',
    method: 'get',
    params
  })
}

// @Tags AppService
// @Summary 分页获取AppService列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppService列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appService/getAppServiceList [get]
export const getAppServiceList = (params) => {
  return service({
    url: '/appService/getAppServiceList',
    method: 'get',
    params
  })
}
