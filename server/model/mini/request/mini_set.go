package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/mini"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type MiniSetSearch struct{
    mini.MiniSet
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
