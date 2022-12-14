package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type AppSetSearch struct {
	app.AppSet
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

//url生成二维码
type Qrcode struct {
	Url string `json:"url" form:"url"`
}
