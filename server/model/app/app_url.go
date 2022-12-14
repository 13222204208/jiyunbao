// 自动生成模板AppUrl
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// AppUrl 结构体
type AppUrl struct {
      global.GVA_MODEL
      Url  string `json:"url" form:"url" gorm:"column:url;comment:请求第三方的地址;"`
      Way  *int `json:"way" form:"way" gorm:"column:way;comment:;"`
}


// TableName AppUrl 表名
func (AppUrl) TableName() string {
  return "app_url"
}

