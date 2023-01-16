package mini

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	miniReq "github.com/flipped-aurora/gin-vue-admin/server/model/mini/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
)

type StoreApi struct {
}

var storeService = service.ServiceGroupApp.MiniServiceGroup.StoreService

func (storeApi *StoreApi) List(c *gin.Context) {
	var r miniReq.StoreSearch
	err := c.ShouldBindQuery(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("引入的数据", r)
	if r.Id > 0 {
		if info, err := storeService.Detail(r.Id); err != nil {
			response.FailWithMessage("获取失败", c)
			return
		} else {
			response.OkWithData(gin.H{"info": info}, c)
			return
		}
	}

	if list, err := storeService.List(r.Search); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": list}, c)
	}
}
