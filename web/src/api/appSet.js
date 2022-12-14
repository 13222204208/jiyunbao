import service from '@/utils/request'

// @Tags AppSet
// @Summary 创建AppSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppSet true "创建AppSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appSet/createAppSet [post]
export const createAppSet = (data) => {
  return service({
    url: '/appSet/createAppSet',
    method: 'post',
    data
  })
}

// @Tags AppSet
// @Summary 删除AppSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppSet true "删除AppSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appSet/deleteAppSet [delete]
export const deleteAppSet = (data) => {
  return service({
    url: '/appSet/deleteAppSet',
    method: 'delete',
    data
  })
}

// @Tags AppSet
// @Summary 删除AppSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appSet/deleteAppSet [delete]
export const deleteAppSetByIds = (data) => {
  return service({
    url: '/appSet/deleteAppSetByIds',
    method: 'delete',
    data
  })
}

// @Tags AppSet
// @Summary 更新AppSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppSet true "更新AppSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appSet/updateAppSet [put]
export const updateAppSet = (data) => {
  return service({
    url: '/appSet/updateAppSet',
    method: 'put',
    data
  })
}

// @Tags AppSet
// @Summary 用id查询AppSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppSet true "用id查询AppSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appSet/findAppSet [get]
export const findAppSet = (params) => {
  return service({
    url: '/appSet/findAppSet',
    method: 'get',
    params
  })
}

// @Tags AppSet
// @Summary 分页获取AppSet列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppSet列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appSet/getAppSetList [get]
export const getAppSetList = (params) => {
  return service({
    url: '/appSet/getAppSetList',
    method: 'get',
    params
  })
}
