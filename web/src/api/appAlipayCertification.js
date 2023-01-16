import service from '@/utils/request'

// @Tags AppAlipayCertification
// @Summary 创建AppAlipayCertification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppAlipayCertification true "创建AppAlipayCertification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appAlipayCertification/createAppAlipayCertification [post]
export const createAppAlipayCertification = (data) => {
  return service({
    url: '/appAlipayCertification/createAppAlipayCertification',
    method: 'post',
    data
  })
}

//实名认证
export const alipayCertification = (data) => {
  return service({
    url: '/appAlipayCertification/alipayCertification',
    method: 'post',
    data
  })
}

// @Tags AppAlipayCertification
// @Summary 删除AppAlipayCertification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppAlipayCertification true "删除AppAlipayCertification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appAlipayCertification/deleteAppAlipayCertification [delete]
export const deleteAppAlipayCertification = (data) => {
  return service({
    url: '/appAlipayCertification/deleteAppAlipayCertification',
    method: 'delete',
    data
  })
}

// @Tags AppAlipayCertification
// @Summary 删除AppAlipayCertification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppAlipayCertification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appAlipayCertification/deleteAppAlipayCertification [delete]
export const deleteAppAlipayCertificationByIds = (data) => {
  return service({
    url: '/appAlipayCertification/deleteAppAlipayCertificationByIds',
    method: 'delete',
    data
  })
}

// @Tags AppAlipayCertification
// @Summary 更新AppAlipayCertification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppAlipayCertification true "更新AppAlipayCertification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appAlipayCertification/updateAppAlipayCertification [put]
export const updateAppAlipayCertification = (data) => {
  return service({
    url: '/appAlipayCertification/updateAppAlipayCertification',
    method: 'put',
    data
  })
}

// @Tags AppAlipayCertification
// @Summary 用id查询AppAlipayCertification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppAlipayCertification true "用id查询AppAlipayCertification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appAlipayCertification/findAppAlipayCertification [get]
export const findAppAlipayCertification = (params) => {
  return service({
    url: '/appAlipayCertification/findAppAlipayCertification',
    method: 'get',
    params
  })
}

//支付宝实名状态查询
export const alipayCertificationState = (params) => {
  return service({
    url: '/appAlipayCertification/alipayCertificationState',
    method: 'get',
    params
  })
}


//授权查询
export const getAlipayAuthState = (params) => {
  return service({
    url: '/appAlipayCertification/getAlipayAuthState',
    method: 'get',
    params
  })
}

// @Tags AppAlipayCertification
// @Summary 分页获取AppAlipayCertification列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppAlipayCertification列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appAlipayCertification/getAppAlipayCertificationList [get]
export const getAppAlipayCertificationList = (params) => {
  return service({
    url: '/appAlipayCertification/getAppAlipayCertificationList',
    method: 'get',
    params
  })
}
