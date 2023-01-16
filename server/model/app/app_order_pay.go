// 自动生成模板AppOrderPay
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AppOrderPay 结构体
type AppOrderPay struct {
	global.GVA_MODEL
	OrderNo     string  `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:;"`
	MercId      string  `json:"mercId" form:"mercId" gorm:"column:merc_id;comment:子商户号;"`
	Subject     string  `json:"subject" form:"subject" gorm:"column:subject;comment:;"`
	PayType     int     `json:"payType" form:"payType" gorm:"column:pay_type;comment:;"`
	TotalAmount float64 `json:"totalAmount" form:"totalAmount" gorm:"column:total_amount;comment:;type:decimal(10,2);"`
	BuyerId     string  `json:"buyerId" form:"buyerId" gorm:"column:buyer_id;comment:;"`
	Appid       string  `json:"appid" form:"appid" gorm:"column:appid;comment:;"`
	PayState    *int    `json:"payState" form:"payState" gorm:"column:pay_state;comment:;default:0"`
}

// TableName AppOrderPay 表名
func (AppOrderPay) TableName() string {
	return "app_order_pay"
}
