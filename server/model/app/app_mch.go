// 自动生成模板AppMch
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"

)

// AppMch 结构体
type AppMch struct {
      global.GVA_MODEL
      Uid  uint `json:"uid" form:"uid" gorm:"column:uid;comment:;"`
      FirmName  string `json:"firmName" form:"firmName" gorm:"column:firm_name;comment:企业名称;"`
      StoreName  string `json:"storeName" form:"storeName" gorm:"column:store_name;comment:店铺名称;"`
      MainType  string `json:"mainType" form:"mainType" gorm:"column:main_type;comment:主营品类;"`
      Commission  *int `json:"commission" form:"commission" gorm:"column:commission;comment:;"`
      Final  *int `json:"final" form:"final" gorm:"column:final;comment:T+1结算;"`
      Acquirer  *int `json:"acquirer" form:"acquirer" gorm:"column:acquirer;comment:特约收单;"`
      LegalPerson  string `json:"legalPerson" form:"legalPerson" gorm:"column:legal_person;comment:与营业执照相同的法人名字;"`
      Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:热线电话;"`
      Avatar  string `json:"avatar" form:"avatar" gorm:"column:avatar;comment:负责 人免冠头像;"`
      CardFront  string `json:"cardFront" form:"cardFront" gorm:"column:card_front;comment:身份证正面;"`
      CardReverse  string `json:"cardReverse" form:"cardReverse" gorm:"column:card_reverse;comment:身份证反面;"`
      Entrust  string `json:"entrust" form:"entrust" gorm:"column:entrust;comment:授权委托书;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:;"`
      Status  int `json:"status" form:"status" gorm:"column:status;comment:商户认证状态;default:0"`
}


// TableName AppMch 表名
func (AppMch) TableName() string {
  return "app_mch"
}

