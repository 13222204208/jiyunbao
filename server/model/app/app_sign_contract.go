// 自动生成模板AppSignContract
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AppSignContract 结构体
type AppSignContract struct {
	global.GVA_MODEL
	ReqId        string `json:"reqId" form:"reqId" gorm:"column:req_id;comment:流水号;"`
	CertId       string `json:"certId" form:"certId" gorm:"column:cert_id;comment:;"`
	BusOpenType  string `json:"busOpenType" form:"busOpenType" gorm:"column:bus_open_type;comment:到账方式,00-扫码工作日到账01-扫码实时到账10-刷卡工作日到账11-刷卡实时到账20-D1到账 允许多选用“|”分隔 T1工作日到账必选;"`
	ContractType *int   `json:"contractType" form:"contractType" gorm:"column:contract_type;comment:合同类型,1-纸质合同 2-电子合同;"`
	IsSendConMsg *int   `json:"isSendConMsg" form:"isSendConMsg" gorm:"column:is_send_con_msg;comment:;"`
	NotifyUrl    string `json:"notifyUrl" form:"notifyUrl" gorm:"column:notify_url;comment:;"`
	CustId       string `json:"custId" form:"custId" gorm:"column:cust_id;comment:;"`
	SignUrl      string `json:"signUrl" form:"signUrl" gorm:"column:sign_url;comment:签约地址"`
	SignId       string `json:"signId" form:"signId" gorm:"column:sign_id;comment:签约ID;"`
	AuthId       string `json:"authId" form:"authId" gorm:"column:auth_id;comment:权限id"`
	MercId       string `json:"mercId" form:"mercId" gorm:"column:merc_id;comment:商户号"`
}

// TableName AppSignContract 表名
func (AppSignContract) TableName() string {
	return "app_sign_contract"
}
