import service from '@/utils/request'

// @Tags MiniClassify
// @Summary 创建MiniClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MiniClassify true "创建MiniClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /miniClassify/createMiniClassify [post]
export const createMiniClassify = (data) => {
  return service({
    url: '/miniClassify/createMiniClassify',
    method: 'post',
    data
  })
}

// @Tags MiniClassify
// @Summary 删除MiniClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MiniClassify true "删除MiniClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /miniClassify/deleteMiniClassify [delete]
export const deleteMiniClassify = (data) => {
  return service({
    url: '/miniClassify/deleteMiniClassify',
    method: 'delete',
    data
  })
}

// @Tags MiniClassify
// @Summary 删除MiniClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MiniClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /miniClassify/deleteMiniClassify [delete]
export const deleteMiniClassifyByIds = (data) => {
  return service({
    url: '/miniClassify/deleteMiniClassifyByIds',
    method: 'delete',
    data
  })
}

// @Tags MiniClassify
// @Summary 更新MiniClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MiniClassify true "更新MiniClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /miniClassify/updateMiniClassify [put]
export const updateMiniClassify = (data) => {
  return service({
    url: '/miniClassify/updateMiniClassify',
    method: 'put',
    data
  })
}

// @Tags MiniClassify
// @Summary 用id查询MiniClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MiniClassify true "用id查询MiniClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /miniClassify/findMiniClassify [get]
export const findMiniClassify = (params) => {
  return service({
    url: '/miniClassify/findMiniClassify',
    method: 'get',
    params
  })
}

// @Tags MiniClassify
// @Summary 分页获取MiniClassify列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MiniClassify列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /miniClassify/getMiniClassifyList [get]
export const getMiniClassifyList = (params) => {
  return service({
    url: '/miniClassify/getMiniClassifyList',
    method: 'get',
    params
  })
}
