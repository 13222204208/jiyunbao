import service from '@/utils/request'

// @Tags MiniSet
// @Summary 创建MiniSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MiniSet true "创建MiniSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /miniSet/createMiniSet [post]
export const createMiniSet = (data) => {
  return service({
    url: '/miniSet/createMiniSet',
    method: 'post',
    data
  })
}

// @Tags MiniSet
// @Summary 删除MiniSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MiniSet true "删除MiniSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /miniSet/deleteMiniSet [delete]
export const deleteMiniSet = (data) => {
  return service({
    url: '/miniSet/deleteMiniSet',
    method: 'delete',
    data
  })
}

// @Tags MiniSet
// @Summary 删除MiniSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MiniSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /miniSet/deleteMiniSet [delete]
export const deleteMiniSetByIds = (data) => {
  return service({
    url: '/miniSet/deleteMiniSetByIds',
    method: 'delete',
    data
  })
}

// @Tags MiniSet
// @Summary 更新MiniSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MiniSet true "更新MiniSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /miniSet/updateMiniSet [put]
export const updateMiniSet = (data) => {
  return service({
    url: '/miniSet/updateMiniSet',
    method: 'put',
    data
  })
}

// @Tags MiniSet
// @Summary 用id查询MiniSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MiniSet true "用id查询MiniSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /miniSet/findMiniSet [get]
export const findMiniSet = (params) => {
  return service({
    url: '/miniSet/findMiniSet',
    method: 'get',
    params
  })
}

// @Tags MiniSet
// @Summary 分页获取MiniSet列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MiniSet列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /miniSet/getMiniSetList [get]
export const getMiniSetList = (params) => {
  return service({
    url: '/miniSet/getMiniSetList',
    method: 'get',
    params
  })
}
