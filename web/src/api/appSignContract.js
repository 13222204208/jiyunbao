import service from '@/utils/request'

// @Tags AppSignContract
// @Summary 创建AppSignContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppSignContract true "创建AppSignContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appSignContract/createAppSignContract [post]
export const createAppSignContract = (data) => {
  return service({
    url: '/appSignContract/createAppSignContract',
    method: 'post',
    data
  })
}

// @Tags AppSignContract
// @Summary 删除AppSignContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppSignContract true "删除AppSignContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appSignContract/deleteAppSignContract [delete]
export const deleteAppSignContract = (data) => {
  return service({
    url: '/appSignContract/deleteAppSignContract',
    method: 'delete',
    data
  })
}

// @Tags AppSignContract
// @Summary 删除AppSignContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppSignContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appSignContract/deleteAppSignContract [delete]
export const deleteAppSignContractByIds = (data) => {
  return service({
    url: '/appSignContract/deleteAppSignContractByIds',
    method: 'delete',
    data
  })
}

// @Tags AppSignContract
// @Summary 更新AppSignContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppSignContract true "更新AppSignContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appSignContract/updateAppSignContract [put]
export const updateAppSignContract = (data) => {
  return service({
    url: '/appSignContract/updateAppSignContract',
    method: 'put',
    data
  })
}

//签约状态查询
export const queryAuthInfo = (data) => {
  return service({
    url: '/appSignContract/queryAuthInfo',
    method: 'post',
    data
  })
}


// @Tags AppSignContract
// @Summary 用id查询AppSignContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppSignContract true "用id查询AppSignContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appSignContract/findAppSignContract [get]
export const findAppSignContract = (params) => {
  return service({
    url: '/appSignContract/findAppSignContract',
    method: 'get',
    params
  })
}

// @Tags AppSignContract
// @Summary 分页获取AppSignContract列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppSignContract列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appSignContract/getAppSignContractList [get]
export const getAppSignContractList = (params) => {
  return service({
    url: '/appSignContract/getAppSignContractList',
    method: 'get',
    params
  })
}

//余额查询
export const balanceQuery = (params) => {
  return service({
    url: '/appSignContract/balanceQuery',
    method: 'get',
    params
  })
}