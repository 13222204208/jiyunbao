// 自动生成模板AppMch
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AppMch 结构体
type AppMch struct {
	global.GVA_MODEL
	Uid             uint    `json:"uid" form:"uid" gorm:"column:uid;comment:;"`
	MchType         int     `json:"mchType" form:"mchType" gorm:"column:mch_type;comment:1，小微商户1.0  ;   2,小微商户2.0; 3 个体商户；4 企业;default:1"`
	FirmName        string  `json:"firmName" form:"firmName" gorm:"column:firm_name;comment:企业名称;"`
	StoreName       string  `json:"storeName" form:"storeName" gorm:"column:store_name;comment:店铺名称;"`
	MainType        string  `json:"mainType" form:"mainType" gorm:"column:main_type;comment:主营品类;"`
	Commission      *int    `json:"commission" form:"commission" gorm:"column:commission;comment:;"`
	Final           *int    `json:"final" form:"final" gorm:"column:final;comment:T+1结算;"`
	Acquirer        *int    `json:"acquirer" form:"acquirer" gorm:"column:acquirer;comment:特约收单;"`
	LegalPerson     string  `json:"legalPerson" form:"legalPerson" gorm:"column:legal_person;comment:与营业执照相同的法人名字;"`
	Phone           string  `json:"phone" form:"phone" gorm:"column:phone;comment:热线电话;"`
	Avatar          string  `json:"avatar" form:"avatar" gorm:"column:avatar;comment:负责 人免冠头像;"`
	CardFront       string  `json:"cardFront" form:"cardFront" gorm:"column:card_front;comment:身份证正面;"`
	CardReverse     string  `json:"cardReverse" form:"cardReverse" gorm:"column:card_reverse;comment:身份证反面;"`
	ManAndCar       string  `json:"manAndCar" form:"manAndCar" gorm:"column:man_and_car;comment:人车辆合影照"`
	Job             string  `json:"job" form:"job" gorm:"column:job;comment:从业资格证"`
	Transport       string  `json:"transport" form:"transport" gorm:"column:transport;comment:车辆运输证"`
	Cab             string  `json:"cab" form:"cab" gorm:"column:cab;comment:驾驶室照片"`
	Entrust         string  `json:"entrust" form:"entrust" gorm:"column:entrust;comment:授权委托书;"`
	Remark          string  `json:"remark" form:"remark" gorm:"column:remark;comment:;"`
	Status          int     `json:"status" form:"status" gorm:"column:status;comment:商户认证状态;default:0"`
	OperatingStatus int     `json:"operatingStatus" form:"operatingStatus" gorm:"column:operating_status;comment:1暂停，2载客，3空车;default:1"`
	Forbidden       int     `json:"forbidden" form:"forbidden" gorm:"column:forbidden;comment:是否被禁用 1 正常  2 禁用;default:1"`
	AppUser         AppUser `gorm:"foreignKey:Uid"`
}

// TableName AppMch 表名
func (AppMch) TableName() string {
	return "app_mch"
}
