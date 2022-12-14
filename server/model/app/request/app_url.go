package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type AppUrlSearch struct {
	app.AppUrl
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type BizContent struct {
	NotifyType string `json:"notifyType" form:"notifyType"`
	Status     string `json:"status" form:"status"`
	Auth       Auth
}

type Auth struct {
	AuthId string `json:"authId"` //权限Id
	MercId string `json:"MercId"` //商户号
	Note   string `json:"note"`   //备注
}
