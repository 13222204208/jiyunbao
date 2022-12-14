import service from '@/utils/request'

// @Tags AppCustInfo
// @Summary 创建AppCustInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppCustInfo true "创建AppCustInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appCustInfo/createAppCustInfo [post]
export const createAppCustInfo = (data) => {
  return service({
    url: '/appCustInfo/createAppCustInfo',
    method: 'post',
    data
  })
}

//基本信息变更申请
export const changeMercBaseInfo = (data) => {
  return service({
    url: '/appCustInfo/changeMercBaseInfo',
    method: 'post',
    data
  })
}

//资料上送
export const addCustInfoApply = (data) => {
  return service({
    url: '/appCustInfo/addCustInfoApply',
    method: 'post',
    data
  })
}

//资料确认
export const auditCustInfoApply = (data) => {
  return service({
    url:'/appCustInfo/auditCustInfoApply',
    method: 'post',
    data
  })
}

//商户状态查询
export const queryCustApply = (data) => {
  return service({
    url:'/appCustInfo/queryCustApply',
    method: 'post',
    data
  })
}



// @Tags AppCustInfo
// @Summary 删除AppCustInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppCustInfo true "删除AppCustInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appCustInfo/deleteAppCustInfo [delete]
export const deleteAppCustInfo = (data) => {
  return service({
    url: '/appCustInfo/deleteAppCustInfo',
    method: 'delete',
    data
  })
}

// @Tags AppCustInfo
// @Summary 删除AppCustInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppCustInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appCustInfo/deleteAppCustInfo [delete]
export const deleteAppCustInfoByIds = (data) => {
  return service({
    url: '/appCustInfo/deleteAppCustInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags AppCustInfo
// @Summary 更新AppCustInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppCustInfo true "更新AppCustInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appCustInfo/updateAppCustInfo [put]
export const updateAppCustInfo = (data) => {
  return service({
    url: '/appCustInfo/updateAppCustInfo',
    method: 'put',
    data
  })
}

// @Tags AppCustInfo
// @Summary 用id查询AppCustInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppCustInfo true "用id查询AppCustInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appCustInfo/findAppCustInfo [get]
export const findAppCustInfo = (params) => {
  return service({
    url: '/appCustInfo/findAppCustInfo',
    method: 'get',
    params
  })
}

// @Tags AppCustInfo
// @Summary 分页获取AppCustInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppCustInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appCustInfo/getAppCustInfoList [get]
export const getAppCustInfoList = (params) => {
  return service({
    url: '/appCustInfo/getAppCustInfoList',
    method: 'get',
    params
  })
}
