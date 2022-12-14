package mini

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/mini"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    miniReq "github.com/flipped-aurora/gin-vue-admin/server/model/mini/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type MiniSetApi struct {
}

var miniSetService = service.ServiceGroupApp.MiniServiceGroup.MiniSetService


// CreateMiniSet 创建MiniSet
// @Tags MiniSet
// @Summary 创建MiniSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mini.MiniSet true "创建MiniSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /miniSet/createMiniSet [post]
func (miniSetApi *MiniSetApi) CreateMiniSet(c *gin.Context) {
	var miniSet mini.MiniSet
	err := c.ShouldBindJSON(&miniSet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "MiniType":{utils.NotEmpty()},
    }
	if err := utils.Verify(miniSet, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := miniSetService.CreateMiniSet(miniSet); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMiniSet 删除MiniSet
// @Tags MiniSet
// @Summary 删除MiniSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mini.MiniSet true "删除MiniSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /miniSet/deleteMiniSet [delete]
func (miniSetApi *MiniSetApi) DeleteMiniSet(c *gin.Context) {
	var miniSet mini.MiniSet
	err := c.ShouldBindJSON(&miniSet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := miniSetService.DeleteMiniSet(miniSet); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMiniSetByIds 批量删除MiniSet
// @Tags MiniSet
// @Summary 批量删除MiniSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MiniSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /miniSet/deleteMiniSetByIds [delete]
func (miniSetApi *MiniSetApi) DeleteMiniSetByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := miniSetService.DeleteMiniSetByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMiniSet 更新MiniSet
// @Tags MiniSet
// @Summary 更新MiniSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mini.MiniSet true "更新MiniSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /miniSet/updateMiniSet [put]
func (miniSetApi *MiniSetApi) UpdateMiniSet(c *gin.Context) {
	var miniSet mini.MiniSet
	err := c.ShouldBindJSON(&miniSet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "MiniType":{utils.NotEmpty()},
      }
    if err := utils.Verify(miniSet, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := miniSetService.UpdateMiniSet(miniSet); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMiniSet 用id查询MiniSet
// @Tags MiniSet
// @Summary 用id查询MiniSet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mini.MiniSet true "用id查询MiniSet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /miniSet/findMiniSet [get]
func (miniSetApi *MiniSetApi) FindMiniSet(c *gin.Context) {
	var miniSet mini.MiniSet
	err := c.ShouldBindQuery(&miniSet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reminiSet, err := miniSetService.GetMiniSet(miniSet.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reminiSet": reminiSet}, c)
	}
}

// GetMiniSetList 分页获取MiniSet列表
// @Tags MiniSet
// @Summary 分页获取MiniSet列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query miniReq.MiniSetSearch true "分页获取MiniSet列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /miniSet/getMiniSetList [get]
func (miniSetApi *MiniSetApi) GetMiniSetList(c *gin.Context) {
	var pageInfo miniReq.MiniSetSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := miniSetService.GetMiniSetInfoList(pageInfo); err != nil {
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
