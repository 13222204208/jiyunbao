import service from '@/utils/request'

// @Tags AppFeedback
// @Summary 创建AppFeedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppFeedback true "创建AppFeedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appFeedback/createAppFeedback [post]
export const createAppFeedback = (data) => {
  return service({
    url: '/appFeedback/createAppFeedback',
    method: 'post',
    data
  })
}

// @Tags AppFeedback
// @Summary 删除AppFeedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppFeedback true "删除AppFeedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appFeedback/deleteAppFeedback [delete]
export const deleteAppFeedback = (data) => {
  return service({
    url: '/appFeedback/deleteAppFeedback',
    method: 'delete',
    data
  })
}

// @Tags AppFeedback
// @Summary 删除AppFeedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppFeedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appFeedback/deleteAppFeedback [delete]
export const deleteAppFeedbackByIds = (data) => {
  return service({
    url: '/appFeedback/deleteAppFeedbackByIds',
    method: 'delete',
    data
  })
}

// @Tags AppFeedback
// @Summary 更新AppFeedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppFeedback true "更新AppFeedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appFeedback/updateAppFeedback [put]
export const updateAppFeedback = (data) => {
  return service({
    url: '/appFeedback/updateAppFeedback',
    method: 'put',
    data
  })
}

// @Tags AppFeedback
// @Summary 用id查询AppFeedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppFeedback true "用id查询AppFeedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appFeedback/findAppFeedback [get]
export const findAppFeedback = (params) => {
  return service({
    url: '/appFeedback/findAppFeedback',
    method: 'get',
    params
  })
}

// @Tags AppFeedback
// @Summary 分页获取AppFeedback列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppFeedback列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appFeedback/getAppFeedbackList [get]
export const getAppFeedbackList = (params) => {
  return service({
    url: '/appFeedback/getAppFeedbackList',
    method: 'get',
    params
  })
}
