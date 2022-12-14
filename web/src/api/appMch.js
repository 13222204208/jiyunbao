import service from '@/utils/request'

// @Tags AppMch
// @Summary 创建AppMch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppMch true "创建AppMch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appMch/createAppMch [post]
export const createAppMch = (data) => {
  return service({
    url: '/appMch/createAppMch',
    method: 'post',
    data
  })
}

// @Tags AppMch
// @Summary 删除AppMch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppMch true "删除AppMch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appMch/deleteAppMch [delete]
export const deleteAppMch = (data) => {
  return service({
    url: '/appMch/deleteAppMch',
    method: 'delete',
    data
  })
}

// @Tags AppMch
// @Summary 删除AppMch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppMch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appMch/deleteAppMch [delete]
export const deleteAppMchByIds = (data) => {
  return service({
    url: '/appMch/deleteAppMchByIds',
    method: 'delete',
    data
  })
}

// @Tags AppMch
// @Summary 更新AppMch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppMch true "更新AppMch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appMch/updateAppMch [put]
export const updateAppMch = (data) => {
  return service({
    url: '/appMch/updateAppMch',
    method: 'put',
    data
  })
}

// @Tags AppMch
// @Summary 用id查询AppMch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppMch true "用id查询AppMch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appMch/findAppMch [get]
export const findAppMch = (params) => {
  return service({
    url: '/appMch/findAppMch',
    method: 'get',
    params
  })
}

// @Tags AppMch
// @Summary 分页获取AppMch列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppMch列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appMch/getAppMchList [get]
export const getAppMchList = (params) => {
  return service({
    url: '/appMch/getAppMchList',
    method: 'get',
    params
  })
}
