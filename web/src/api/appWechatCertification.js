import service from '@/utils/request'

// @Tags AppWechatCertification
// @Summary 创建AppWechatCertification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppWechatCertification true "创建AppWechatCertification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appWechatCertification/createAppWechatCertification [post]
export const createAppWechatCertification = (data) => {
  return service({
    url: '/appWechatCertification/createAppWechatCertification',
    method: 'post',
    data
  })
}

//实名认证
export const wechatCertification = (data) => {
  return service({
    url: '/appWechatCertification/wechatCertification',
    method: 'post',
    data
  })
}

//明细信息查询
export const getAuthMessagesByMercId = (params) => {
  return service({
    url:'/appWechatCertification/getAuthMessagesByMercId',
    method: 'get',
    params
  })
}

//授权状态
export const getAuthState = (params) => {
  return service({
    url:'/appWechatCertification/getAuthState',
    method: 'get',
    params
  })
}

// @Tags AppWechatCertification
// @Summary 删除AppWechatCertification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppWechatCertification true "删除AppWechatCertification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appWechatCertification/deleteAppWechatCertification [delete]
export const deleteAppWechatCertification = (data) => {
  return service({
    url: '/appWechatCertification/deleteAppWechatCertification',
    method: 'delete',
    data
  })
}

// @Tags AppWechatCertification
// @Summary 删除AppWechatCertification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppWechatCertification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appWechatCertification/deleteAppWechatCertification [delete]
export const deleteAppWechatCertificationByIds = (data) => {
  return service({
    url: '/appWechatCertification/deleteAppWechatCertificationByIds',
    method: 'delete',
    data
  })
}

// @Tags AppWechatCertification
// @Summary 更新AppWechatCertification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppWechatCertification true "更新AppWechatCertification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appWechatCertification/updateAppWechatCertification [put]
export const updateAppWechatCertification = (data) => {
  return service({
    url: '/appWechatCertification/updateAppWechatCertification',
    method: 'put',
    data
  })
}

// @Tags AppWechatCertification
// @Summary 用id查询AppWechatCertification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppWechatCertification true "用id查询AppWechatCertification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appWechatCertification/findAppWechatCertification [get]
export const findAppWechatCertification = (params) => {
  return service({
    url: '/appWechatCertification/findAppWechatCertification',
    method: 'get',
    params
  })
}


//微信实名认证申请状态
export const certificationState = (params) => {
  return service({
    url: '/appWechatCertification/certificationState',
    method: 'get',
    params
  })
}

// @Tags AppWechatCertification
// @Summary 分页获取AppWechatCertification列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppWechatCertification列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appWechatCertification/getAppWechatCertificationList [get]
export const getAppWechatCertificationList = (params) => {
  return service({
    url: '/appWechatCertification/getAppWechatCertificationList',
    method: 'get',
    params
  })
}
