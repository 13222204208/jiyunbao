import service from '@/utils/request'

// @Tags AppStore
// @Summary 创建AppStore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppStore true "创建AppStore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appStore/createAppStore [post]
export const createAppStore = (data) => {
  return service({
    url: '/appStore/createAppStore',
    method: 'post',
    data
  })
}

// @Tags AppStore
// @Summary 删除AppStore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppStore true "删除AppStore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appStore/deleteAppStore [delete]
export const deleteAppStore = (data) => {
  return service({
    url: '/appStore/deleteAppStore',
    method: 'delete',
    data
  })
}

// @Tags AppStore
// @Summary 删除AppStore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppStore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appStore/deleteAppStore [delete]
export const deleteAppStoreByIds = (data) => {
  return service({
    url: '/appStore/deleteAppStoreByIds',
    method: 'delete',
    data
  })
}

// @Tags AppStore
// @Summary 更新AppStore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppStore true "更新AppStore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appStore/updateAppStore [put]
export const updateAppStore = (data) => {
  return service({
    url: '/appStore/updateAppStore',
    method: 'put',
    data
  })
}

// @Tags AppStore
// @Summary 用id查询AppStore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppStore true "用id查询AppStore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appStore/findAppStore [get]
export const findAppStore = (params) => {
  return service({
    url: '/appStore/findAppStore',
    method: 'get',
    params
  })
}

// @Tags AppStore
// @Summary 分页获取AppStore列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppStore列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appStore/getAppStoreList [get]
export const getAppStoreList = (params) => {
  return service({
    url: '/appStore/getAppStoreList',
    method: 'get',
    params
  })
}
