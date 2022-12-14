// 自动生成模板AppSet
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AppSet 结构体
type AppSet struct {
	global.GVA_MODEL
	PhoneKeyId        string `json:"phoneKeyId" form:"phoneKeyId" gorm:"column:phone_key_id;comment:;"`
	PhoneKeySecret    string `json:"phoneKeySecret" form:"phoneKeySecret" gorm:"column:phone_key_secret;comment:;"`
	ThirdUrl          string `json:"thirdUrl" form:"thirdUrl" gorm:"column:third_url;comment:;"`
	NotifyUrl         string `json:"notifyUrl" form:"notifyUrl" gorm:"column:notify_url;comment:异步通知地址;"`
	PhoneSignName     string `json:"phoneSignName" form:"phoneSignName" gorm:"column:phone_sign_name;comment:;"`
	PhoneTemplateCode string `json:"phoneTemplateCode" form:"phoneTemplateCode" gorm:"column:phone_template_code;comment:;"`
}

// TableName AppSet 表名
func (AppSet) TableName() string {
	return "app_set"
}
