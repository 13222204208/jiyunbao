package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mini"
	"time"
)

type MiniUserSearch struct {
	mini.MiniUser
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type MiniUserLogin struct {
	Code     string `json:"code" form:"code"`
	Avatar   string `json:"avatar" form:"avatar"`
	Nickname string `json:"nickname" form:"nickname"`
	MiniType int    `json:"miniType" form:"miniType"`
}

type GetPhone struct {
	Code          string `json:"code" form:"code"` //微信
	MiniType      int    `json:"miniType" form:"miniType"`
	EncryptedData string `json:"encryptedData" form:"encryptedData"` //支付宝
}
