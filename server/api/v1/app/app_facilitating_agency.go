package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/app"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type AppFacilitatingAgencyApi struct {
}

var appFacilitatingAgencyService = service.ServiceGroupApp.AppServiceGroup.AppFacilitatingAgencyService


// CreateAppFacilitatingAgency 创建AppFacilitatingAgency
// @Tags AppFacilitatingAgency
// @Summary 创建AppFacilitatingAgency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppFacilitatingAgency true "创建AppFacilitatingAgency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appFacilitatingAgency/createAppFacilitatingAgency [post]
func (appFacilitatingAgencyApi *AppFacilitatingAgencyApi) CreateAppFacilitatingAgency(c *gin.Context) {
	var appFacilitatingAgency app.AppFacilitatingAgency
	err := c.ShouldBindJSON(&appFacilitatingAgency)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Phone":{utils.NotEmpty()},
        "Name":{utils.NotEmpty()},
        "Way":{utils.NotEmpty()},
        "Principal":{utils.NotEmpty()},
        "Status":{utils.NotEmpty()},
        "Certification":{utils.NotEmpty()},
    }
	if err := utils.Verify(appFacilitatingAgency, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := appFacilitatingAgencyService.CreateAppFacilitatingAgency(appFacilitatingAgency); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppFacilitatingAgency 删除AppFacilitatingAgency
// @Tags AppFacilitatingAgency
// @Summary 删除AppFacilitatingAgency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppFacilitatingAgency true "删除AppFacilitatingAgency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appFacilitatingAgency/deleteAppFacilitatingAgency [delete]
func (appFacilitatingAgencyApi *AppFacilitatingAgencyApi) DeleteAppFacilitatingAgency(c *gin.Context) {
	var appFacilitatingAgency app.AppFacilitatingAgency
	err := c.ShouldBindJSON(&appFacilitatingAgency)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appFacilitatingAgencyService.DeleteAppFacilitatingAgency(appFacilitatingAgency); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppFacilitatingAgencyByIds 批量删除AppFacilitatingAgency
// @Tags AppFacilitatingAgency
// @Summary 批量删除AppFacilitatingAgency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppFacilitatingAgency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appFacilitatingAgency/deleteAppFacilitatingAgencyByIds [delete]
func (appFacilitatingAgencyApi *AppFacilitatingAgencyApi) DeleteAppFacilitatingAgencyByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appFacilitatingAgencyService.DeleteAppFacilitatingAgencyByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppFacilitatingAgency 更新AppFacilitatingAgency
// @Tags AppFacilitatingAgency
// @Summary 更新AppFacilitatingAgency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppFacilitatingAgency true "更新AppFacilitatingAgency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appFacilitatingAgency/updateAppFacilitatingAgency [put]
func (appFacilitatingAgencyApi *AppFacilitatingAgencyApi) UpdateAppFacilitatingAgency(c *gin.Context) {
	var appFacilitatingAgency app.AppFacilitatingAgency
	err := c.ShouldBindJSON(&appFacilitatingAgency)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Phone":{utils.NotEmpty()},
          "Name":{utils.NotEmpty()},
          "Way":{utils.NotEmpty()},
          "Principal":{utils.NotEmpty()},
          "Status":{utils.NotEmpty()},
          "Certification":{utils.NotEmpty()},
      }
    if err := utils.Verify(appFacilitatingAgency, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := appFacilitatingAgencyService.UpdateAppFacilitatingAgency(appFacilitatingAgency); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppFacilitatingAgency 用id查询AppFacilitatingAgency
// @Tags AppFacilitatingAgency
// @Summary 用id查询AppFacilitatingAgency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppFacilitatingAgency true "用id查询AppFacilitatingAgency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appFacilitatingAgency/findAppFacilitatingAgency [get]
func (appFacilitatingAgencyApi *AppFacilitatingAgencyApi) FindAppFacilitatingAgency(c *gin.Context) {
	var appFacilitatingAgency app.AppFacilitatingAgency
	err := c.ShouldBindQuery(&appFacilitatingAgency)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappFacilitatingAgency, err := appFacilitatingAgencyService.GetAppFacilitatingAgency(appFacilitatingAgency.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappFacilitatingAgency": reappFacilitatingAgency}, c)
	}
}

// GetAppFacilitatingAgencyList 分页获取AppFacilitatingAgency列表
// @Tags AppFacilitatingAgency
// @Summary 分页获取AppFacilitatingAgency列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.AppFacilitatingAgencySearch true "分页获取AppFacilitatingAgency列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appFacilitatingAgency/getAppFacilitatingAgencyList [get]
func (appFacilitatingAgencyApi *AppFacilitatingAgencyApi) GetAppFacilitatingAgencyList(c *gin.Context) {
	var pageInfo appReq.AppFacilitatingAgencySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appFacilitatingAgencyService.GetAppFacilitatingAgencyInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
