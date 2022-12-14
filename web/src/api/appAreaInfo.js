import service from '@/utils/request'

// @Tags AppAreaInfo
// @Summary 创建AppAreaInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppAreaInfo true "创建AppAreaInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appAreaInfo/createAppAreaInfo [post]
export const createAppAreaInfo = (data) => {
  return service({
    url: '/appAreaInfo/createAppAreaInfo',
    method: 'post',
    data
  })
}

// @Tags AppAreaInfo
// @Summary 删除AppAreaInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppAreaInfo true "删除AppAreaInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appAreaInfo/deleteAppAreaInfo [delete]
export const deleteAppAreaInfo = (data) => {
  return service({
    url: '/appAreaInfo/deleteAppAreaInfo',
    method: 'delete',
    data
  })
}

// @Tags AppAreaInfo
// @Summary 删除AppAreaInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppAreaInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appAreaInfo/deleteAppAreaInfo [delete]
export const deleteAppAreaInfoByIds = (data) => {
  return service({
    url: '/appAreaInfo/deleteAppAreaInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags AppAreaInfo
// @Summary 更新AppAreaInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppAreaInfo true "更新AppAreaInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appAreaInfo/updateAppAreaInfo [put]
export const updateAppAreaInfo = (data) => {
  return service({
    url: '/appAreaInfo/updateAppAreaInfo',
    method: 'put',
    data
  })
}

// @Tags AppAreaInfo
// @Summary 用id查询AppAreaInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppAreaInfo true "用id查询AppAreaInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appAreaInfo/findAppAreaInfo [get]
export const findAppAreaInfo = (params) => {
  return service({
    url: '/appAreaInfo/findAppAreaInfo',
    method: 'get',
    params
  })
}

// @Tags AppAreaInfo
// @Summary 分页获取AppAreaInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppAreaInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appAreaInfo/getAppAreaInfoList [get]
export const getAppAreaInfoList = (params) => {
  return service({
    url: '/appAreaInfo/getAppAreaInfoList',
    method: 'get',
    params
  })
}
