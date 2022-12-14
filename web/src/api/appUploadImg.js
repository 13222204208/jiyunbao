import service from '@/utils/request'

// @Tags AppUploadImg
// @Summary 创建AppUploadImg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppUploadImg true "创建AppUploadImg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appUploadImg/createAppUploadImg [post]
export const createAppUploadImg = (data) => {
  return service({
    url: '/appUploadImg/createAppUploadImg',
    method: 'post',
    data
  })
}

// @Tags AppUploadImg
// @Summary 删除AppUploadImg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppUploadImg true "删除AppUploadImg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appUploadImg/deleteAppUploadImg [delete]
export const deleteAppUploadImg = (data) => {
  return service({
    url: '/appUploadImg/deleteAppUploadImg',
    method: 'delete',
    data
  })
}

// @Tags AppUploadImg
// @Summary 删除AppUploadImg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppUploadImg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appUploadImg/deleteAppUploadImg [delete]
export const deleteAppUploadImgByIds = (data) => {
  return service({
    url: '/appUploadImg/deleteAppUploadImgByIds',
    method: 'delete',
    data
  })
}

// @Tags AppUploadImg
// @Summary 更新AppUploadImg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppUploadImg true "更新AppUploadImg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appUploadImg/updateAppUploadImg [put]
export const updateAppUploadImg = (data) => {
  return service({
    url: '/appUploadImg/updateAppUploadImg',
    method: 'put',
    data
  })
}

// @Tags AppUploadImg
// @Summary 用id查询AppUploadImg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppUploadImg true "用id查询AppUploadImg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appUploadImg/findAppUploadImg [get]
export const findAppUploadImg = (params) => {
  return service({
    url: '/appUploadImg/findAppUploadImg',
    method: 'get',
    params
  })
}

// @Tags AppUploadImg
// @Summary 分页获取AppUploadImg列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppUploadImg列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appUploadImg/getAppUploadImgList [get]
export const getAppUploadImgList = (params) => {
  return service({
    url: '/appUploadImg/getAppUploadImgList',
    method: 'get',
    params
  })
}
