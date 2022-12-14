import service from '@/utils/request'

// @Tags AppNotice
// @Summary 创建AppNotice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppNotice true "创建AppNotice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appNotice/createAppNotice [post]
export const createAppNotice = (data) => {
  return service({
    url: '/appNotice/createAppNotice',
    method: 'post',
    data
  })
}

// @Tags AppNotice
// @Summary 删除AppNotice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppNotice true "删除AppNotice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appNotice/deleteAppNotice [delete]
export const deleteAppNotice = (data) => {
  return service({
    url: '/appNotice/deleteAppNotice',
    method: 'delete',
    data
  })
}

// @Tags AppNotice
// @Summary 删除AppNotice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppNotice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appNotice/deleteAppNotice [delete]
export const deleteAppNoticeByIds = (data) => {
  return service({
    url: '/appNotice/deleteAppNoticeByIds',
    method: 'delete',
    data
  })
}

// @Tags AppNotice
// @Summary 更新AppNotice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppNotice true "更新AppNotice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appNotice/updateAppNotice [put]
export const updateAppNotice = (data) => {
  return service({
    url: '/appNotice/updateAppNotice',
    method: 'put',
    data
  })
}

// @Tags AppNotice
// @Summary 用id查询AppNotice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppNotice true "用id查询AppNotice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appNotice/findAppNotice [get]
export const findAppNotice = (params) => {
  return service({
    url: '/appNotice/findAppNotice',
    method: 'get',
    params
  })
}

// @Tags AppNotice
// @Summary 分页获取AppNotice列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppNotice列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appNotice/getAppNoticeList [get]
export const getAppNoticeList = (params) => {
  return service({
    url: '/appNotice/getAppNoticeList',
    method: 'get',
    params
  })
}
