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
	Deduct      float64 `json:"deduct" form:"deduct" gorm:"column:deduct;comment:扣除的金额;type:decimal(10,2);default:0"`
	IsMember    int     `json:"isMember" form:"isMember" gorm:"column:is_member;comment:是否是会员 0 非会员，1 会员也就是老铁;default:0"`
	BuyerId     string  `json:"buyerId" form:"buyerId" gorm:"column:buyer_id;comment:;"`
	Appid       string  `json:"appid" form:"appid" gorm:"column:appid;comment:;"`
	PayTime     string  `json:"payTime" form:"payTime" gorm:"column:pay_time;comment:支付时间;default:2023-01-16"`
	PayState    *int    `json:"payState" form:"payState" gorm:"column:pay_state;comment:;default:0"`
}

// TableName AppOrderPay 表名
func (AppOrderPay) TableName() string {
	return "app_order_pay"
}
