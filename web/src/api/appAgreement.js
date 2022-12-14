import service from '@/utils/request'

// @Tags AppAgreement
// @Summary 创建AppAgreement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppAgreement true "创建AppAgreement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appAgreement/createAppAgreement [post]
export const createAppAgreement = (data) => {
  return service({
    url: '/appAgreement/createAppAgreement',
    method: 'post',
    data
  })
}

// @Tags AppAgreement
// @Summary 删除AppAgreement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppAgreement true "删除AppAgreement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appAgreement/deleteAppAgreement [delete]
export const deleteAppAgreement = (data) => {
  return service({
    url: '/appAgreement/deleteAppAgreement',
    method: 'delete',
    data
  })
}

// @Tags AppAgreement
// @Summary 删除AppAgreement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppAgreement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appAgreement/deleteAppAgreement [delete]
export const deleteAppAgreementByIds = (data) => {
  return service({
    url: '/appAgreement/deleteAppAgreementByIds',
    method: 'delete',
    data
  })
}

// @Tags AppAgreement
// @Summary 更新AppAgreement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppAgreement true "更新AppAgreement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appAgreement/updateAppAgreement [put]
export const updateAppAgreement = (data) => {
  return service({
    url: '/appAgreement/updateAppAgreement',
    method: 'put',
    data
  })
}

// @Tags AppAgreement
// @Summary 用id查询AppAgreement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppAgreement true "用id查询AppAgreement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appAgreement/findAppAgreement [get]
export const findAppAgreement = (params) => {
  return service({
    url: '/appAgreement/findAppAgreement',
    method: 'get',
    params
  })
}

// @Tags AppAgreement
// @Summary 分页获取AppAgreement列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppAgreement列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appAgreement/getAppAgreementList [get]
export const getAppAgreementList = (params) => {
  return service({
    url: '/appAgreement/getAppAgreementList',
    method: 'get',
    params
  })
}
