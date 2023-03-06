// 自动生成模板AppRate
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// AppRate 结构体
type AppRate struct {
      global.GVA_MODEL
      Title  string `json:"title" form:"title" gorm:"column:title;comment:;"`
      Money  *float64 `json:"money" form:"money" gorm:"column:money;comment:;"`
      Rate  string `json:"rate" form:"rate" gorm:"column:rate;comment:;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:;"`
}


// TableName AppRate 表名
func (AppRate) TableName() string {
  return "app_rate"
}

