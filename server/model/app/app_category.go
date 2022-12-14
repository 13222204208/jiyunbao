// 自动生成模板AppCategory
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// AppCategory 结构体
type AppCategory struct {
      global.GVA_MODEL
      Title  string `json:"title" form:"title" gorm:"column:title;comment:品类名称;"`
      CodeId  string `json:"codeId" form:"codeId" gorm:"column:code_id;comment:银盛行业类别;"`
      Pid  *int `json:"pid" form:"pid" gorm:"column:pid;comment:;"`
}


// TableName AppCategory 表名
func (AppCategory) TableName() string {
  return "app_category"
}

