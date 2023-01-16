package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type AppOrderPaySearch struct {
	app.AppOrderPay
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

//聚合支付请求参数
type OrderPay struct {
	Code        string  `json:"code" form:"code"`
	MercId      string  `json:"MercId" form:"MercId"`
	TotalAmount float64 `json:"totalAmount" form:"totalAmount"`
	PayType     int     `json:"payType" form:"payType"`
	Subject     string  `json:"subject" form:"subject"`
}

//聚合支付回调参数
type OrderPayNotify struct {
	OutTradeNo  string `json:"out_trade_no" form:"out_trade_no"`
	TradeStatus string `json:"trade_status" form:"trade_status"`
	PayeeFee    string `json:"pyee_fee" form:"payere_fee"`
}
