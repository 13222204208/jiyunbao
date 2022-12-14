// 自动生成模板AppNotice
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// AppNotice 结构体
type AppNotice struct {
      global.GVA_MODEL
      Contents  string `json:"contents" form:"contents" gorm:"column:contents;comment:;size:5000;"`
      Way  *int `json:"way" form:"way" gorm:"column:way;comment:公告的类型服务公告等;"`
}


// TableName AppNotice 表名
func (AppNotice) TableName() string {
  return "app_notice"
}

