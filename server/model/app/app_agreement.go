// 自动生成模板AppAgreement
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AppAgreement 结构体
type AppAgreement struct {
	global.GVA_MODEL
	Way      *int   `json:"way" form:"way" gorm:"column:way;comment:协议类型;"`
	Title    string `json:"title" form:"title" gorm:"column:title;comment:协议标题;"`
	Contents string `json:"contents" form:"contents" gorm:"column:contents;comment:协议内容;type:text;size:60000;"`
}

// TableName AppAgreement 表名
func (AppAgreement) TableName() string {
	return "app_agreement"
}
