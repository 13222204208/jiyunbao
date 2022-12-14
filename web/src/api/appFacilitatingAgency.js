import service from '@/utils/request'

// @Tags AppFacilitatingAgency
// @Summary 创建AppFacilitatingAgency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppFacilitatingAgency true "创建AppFacilitatingAgency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appFacilitatingAgency/createAppFacilitatingAgency [post]
export const createAppFacilitatingAgency = (data) => {
  return service({
    url: '/appFacilitatingAgency/createAppFacilitatingAgency',
    method: 'post',
    data
  })
}

// @Tags AppFacilitatingAgency
// @Summary 删除AppFacilitatingAgency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppFacilitatingAgency true "删除AppFacilitatingAgency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appFacilitatingAgency/deleteAppFacilitatingAgency [delete]
export const deleteAppFacilitatingAgency = (data) => {
  return service({
    url: '/appFacilitatingAgency/deleteAppFacilitatingAgency',
    method: 'delete',
    data
  })
}

// @Tags AppFacilitatingAgency
// @Summary 删除AppFacilitatingAgency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppFacilitatingAgency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appFacilitatingAgency/deleteAppFacilitatingAgency [delete]
export const deleteAppFacilitatingAgencyByIds = (data) => {
  return service({
    url: '/appFacilitatingAgency/deleteAppFacilitatingAgencyByIds',
    method: 'delete',
    data
  })
}

// @Tags AppFacilitatingAgency
// @Summary 更新AppFacilitatingAgency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppFacilitatingAgency true "更新AppFacilitatingAgency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appFacilitatingAgency/updateAppFacilitatingAgency [put]
export const updateAppFacilitatingAgency = (data) => {
  return service({
    url: '/appFacilitatingAgency/updateAppFacilitatingAgency',
    method: 'put',
    data
  })
}

// @Tags AppFacilitatingAgency
// @Summary 用id查询AppFacilitatingAgency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppFacilitatingAgency true "用id查询AppFacilitatingAgency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appFacilitatingAgency/findAppFacilitatingAgency [get]
export const findAppFacilitatingAgency = (params) => {
  return service({
    url: '/appFacilitatingAgency/findAppFacilitatingAgency',
    method: 'get',
    params
  })
}

// @Tags AppFacilitatingAgency
// @Summary 分页获取AppFacilitatingAgency列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppFacilitatingAgency列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appFacilitatingAgency/getAppFacilitatingAgencyList [get]
export const getAppFacilitatingAgencyList = (params) => {
  return service({
    url: '/appFacilitatingAgency/getAppFacilitatingAgencyList',
    method: 'get',
    params
  })
}
