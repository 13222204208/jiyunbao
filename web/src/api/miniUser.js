import service from '@/utils/request'

// @Tags MiniUser
// @Summary 创建MiniUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MiniUser true "创建MiniUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /miniUser/createMiniUser [post]
export const createMiniUser = (data) => {
  return service({
    url: '/miniUser/createMiniUser',
    method: 'post',
    data
  })
}

// @Tags MiniUser
// @Summary 删除MiniUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MiniUser true "删除MiniUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /miniUser/deleteMiniUser [delete]
export const deleteMiniUser = (data) => {
  return service({
    url: '/miniUser/deleteMiniUser',
    method: 'delete',
    data
  })
}

// @Tags MiniUser
// @Summary 删除MiniUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MiniUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /miniUser/deleteMiniUser [delete]
export const deleteMiniUserByIds = (data) => {
  return service({
    url: '/miniUser/deleteMiniUserByIds',
    method: 'delete',
    data
  })
}

// @Tags MiniUser
// @Summary 更新MiniUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MiniUser true "更新MiniUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /miniUser/updateMiniUser [put]
export const updateMiniUser = (data) => {
  return service({
    url: '/miniUser/updateMiniUser',
    method: 'put',
    data
  })
}

// @Tags MiniUser
// @Summary 用id查询MiniUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MiniUser true "用id查询MiniUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /miniUser/findMiniUser [get]
export const findMiniUser = (params) => {
  return service({
    url: '/miniUser/findMiniUser',
    method: 'get',
    params
  })
}

// @Tags MiniUser
// @Summary 分页获取MiniUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MiniUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /miniUser/getMiniUserList [get]
export const getMiniUserList = (params) => {
  return service({
    url: '/miniUser/getMiniUserList',
    method: 'get',
    params
  })
}
