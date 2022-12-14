import service from '@/utils/request'

// @Tags AppRelieveStore
// @Summary 创建AppRelieveStore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppRelieveStore true "创建AppRelieveStore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appRelieveStore/createAppRelieveStore [post]
export const createAppRelieveStore = (data) => {
  return service({
    url: '/appRelieveStore/createAppRelieveStore',
    method: 'post',
    data
  })
}

// @Tags AppRelieveStore
// @Summary 删除AppRelieveStore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppRelieveStore true "删除AppRelieveStore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appRelieveStore/deleteAppRelieveStore [delete]
export const deleteAppRelieveStore = (data) => {
  return service({
    url: '/appRelieveStore/deleteAppRelieveStore',
    method: 'delete',
    data
  })
}

// @Tags AppRelieveStore
// @Summary 删除AppRelieveStore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppRelieveStore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appRelieveStore/deleteAppRelieveStore [delete]
export const deleteAppRelieveStoreByIds = (data) => {
  return service({
    url: '/appRelieveStore/deleteAppRelieveStoreByIds',
    method: 'delete',
    data
  })
}

// @Tags AppRelieveStore
// @Summary 更新AppRelieveStore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppRelieveStore true "更新AppRelieveStore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appRelieveStore/updateAppRelieveStore [put]
export const updateAppRelieveStore = (data) => {
  return service({
    url: '/appRelieveStore/updateAppRelieveStore',
    method: 'put',
    data
  })
}

// @Tags AppRelieveStore
// @Summary 用id查询AppRelieveStore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppRelieveStore true "用id查询AppRelieveStore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appRelieveStore/findAppRelieveStore [get]
export const findAppRelieveStore = (params) => {
  return service({
    url: '/appRelieveStore/findAppRelieveStore',
    method: 'get',
    params
  })
}

// @Tags AppRelieveStore
// @Summary 分页获取AppRelieveStore列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppRelieveStore列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appRelieveStore/getAppRelieveStoreList [get]
export const getAppRelieveStoreList = (params) => {
  return service({
    url: '/appRelieveStore/getAppRelieveStoreList',
    method: 'get',
    params
  })
}
